const express=require('express');
const app=express();
const mongoose=require('mongoose');
const PORT=process.env.PORT || 3000;

const {MONGO_USER,MONGO_PASS,MONGO_IP,MONGO_PORT} = require('./config/config');

const mongoURL = `mongodb://${MONGO_USER}:${MONGO_PASS}@${MONGO_IP}:${MONGO_PORT}/?authSource=admin`;

// routes
const postRouter = require("./routes/postRoutes");
const userRouter = require("./routes/userRoutes");


// mongoose has inbuilt 30 second rule where it tries to connect to mongo
// though if you have an application dependent on each other like mongo and node this a rough way to actually make sure that our mongo connects first so as to avoid errors.
const connectWithRetry = () =>{
  // can get the IP address of mongo container via 'docker inspect container_name' under networks section
  mongoose.connect(
    mongoURL,{useNewUrlParser: true,useUnifiedTopology: true,useFindAndModify:false,useCreateIndex:true}
  )// IP address:Port(on which mongo runs default is 27017)
  .then(()=>{console.log('Connected to MongoDB 🎉via Docker Container')})
  .catch((err)=>{
    console.log(err);
    // however this is not a best practice
    setTimeout(connectWithRetry,5000);
  });

}

connectWithRetry();

app.use(express.json());

app.get('/',(req,res)=>{
  res.send("<h1>Inside Docker Development✨✨</h1>");
});

app.use("/api/v1/posts",postRouter);
app.use("/api/v1/users",userRouter);


app.listen(PORT,process.env.IP,()=>{console.log(`Server Started at ${PORT}`)});

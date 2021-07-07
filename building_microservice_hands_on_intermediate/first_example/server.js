const express=require('express');
const app=express();
const mongoose=require('mongoose');
const PORT=process.env.PORT || 3000;

const {MONGO_USER,MONGO_PASS,MONGO_IP,MONGO_PORT} = require('./config/config');

const mongoURL = `mongodb://${MONGO_USER}:${MONGO_PASS}@${MONGO_IP}:${MONGO_PORT}/?authSource=admin`;

// mongoose has inbuilt 30 second rule where it tries to connect to mongo
// though if you have an application dependent on each other like mongo and node this a rough way to actually make sure that our mongo connects first so as to avoid errors.
const connectWithRetry = () =>{
  // can get the IP address of mongo container via 'docker inspect container_name' under networks section
  mongoose.connect(
    mongoURL,{useNewUrlParser: true,useUnifiedTopology: true,useFindAndModify:false}
  )// IP address:Port(on which mongo runs default is 27017)
  .then(()=>{console.log('Connected to MongoDB 🎉via Docker Container')})
  .catch((err)=>{
    console.log(err);
    // however this is not a best practice
    setTimeout(connectWithRetry,5000);
  });

}

connectWithRetry();


app.get('/',(req,res)=>{
  res.send("<h1>docker-compose.prod.yml works fine!!!</h1>");
});

app.listen(PORT,process.env.IP,()=>{console.log(`Server Started at ${PORT}`)});

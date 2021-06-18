const express=require('express');
const app=express();
const mongoose=require('mongoose');
const PORT=process.env.PORT || 3000;


// can get the IP address of mongo container via 'docker inspect container_name' under networks section
mongoose.connect(
  "mongodb://jasmeet:mypassword@172.24.0.2:27017/?authSource=admin"
)// IP address:Port(on which mongo runs default is 27017)
.then(()=>{console.log('Connected to MongoDB via Docker Container')})
.catch((err)=>{console.log(err);});

app.get('/',(req,res)=>{
  res.send("<h1>docker-compose.prod.yml works fine!!!</h1>");
});

app.listen(PORT,process.env.IP,()=>{console.log(`Server Started at ${PORT}`)});

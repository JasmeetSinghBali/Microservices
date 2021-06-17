const express=require('express'),
      app=express(),
      PORT=process.env.PORT || 3000;

app.get('/',(req,res)=>{
  res.send("<h1>docker-compose.prod.yml works fine!!</h1>");
});

app.listen(PORT,process.env.IP,()=>{console.log(`Server Started at ${PORT}`)});

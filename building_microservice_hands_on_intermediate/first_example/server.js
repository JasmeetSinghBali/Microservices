const express=require('express'),
      app=express(),
      PORT=process.env.PORT || 5000;

app.get('/',(req,res)=>{
  res.send("<h1>This is a docker compose up!</h1>");
});

app.listen(PORT,process.env.IP,()=>{console.log(`Server Started at ${PORT}`)});

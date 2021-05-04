const express=require('express'),
      app=express(),
      PORT= 5000;

app.get('/',(req,res)=>{
  res.send('Node Application Running');
})

app.listen(PORT,process.env.IP,()=>{
  console.log(`Server running at PORT : ${PORT}`);
})

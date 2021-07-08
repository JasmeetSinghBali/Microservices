const User = require('../models/Users');

const bcrypt = require('bcryptjs');

exports.signUp = async(req,res)=>{
  const {username,password} = req.body;
  try{
    const hashpassword = await bcrypt.hash(password,12);
    const newUser = await User.create({
      username,
      password : hashpassword
    });
    res.status(201).json({
      status:'Success',
      data:{
        new_user: newUser
      }
    })
  }catch(err){
    res.status(400).json({
      status: 'Something Went Wrong',
      error: err
    });
    console.log(err);
  }
}

exports.login = async(req,res)=>{
  const {username,password} = req.body;

  try{
    const user = await User.findOne({username});

    if(!user){
      res.status(404).json({
        status:'failed',
        message: 'user not found'
      })
    }

    const isAuth = await bcrypt.compare(password,user.password);
    if(isAuth){
      res.status(200).json({
        message:'logged in!!'
      });
    }else{
      res.json({
        message: 'Username/Password is Wrong!!'
      });
    }
  }catch(err){
    res.status(400).json({
      status: 'Something Went Wrong',
      error: err
    })
  }
}

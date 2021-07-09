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
    // so this attachs a user object and creates a session on the server side so as long as this user object is attached with the session the session is valid.
    req.session.user = newUser;
    return res.status(201).json({
      status:'Success',
      data:{
        new_user: newUser
      }
    })
  }catch(err){
    console.log(err);
    return res.status(400).json({
      status: 'Something Went Wrong',
      error: err
    });

  }
}

exports.login = async(req,res)=>{
  const {username,password} = req.body;

  try{
    const user = await User.findOne({username});

    if(!user){
      return res.status(404).json({
        status:'failed',
        message: 'user not found'
      })
    }

    const isAuth = await bcrypt.compare(password,user.password);
    if(isAuth){
      // we are assigning the user object to req.session object with key user if password is correct
      req.session.user = user;
      return res.status(200).json({
        message:'logged in!!'
      });
    }else{
      return res.json({
        message: 'Username/Password is Wrong!!'
      });
    }
  }catch(err){
    return res.status(400).json({
      status: 'Something Went Wrong',
      error: err
    })
  }
}

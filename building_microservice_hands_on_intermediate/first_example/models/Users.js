const mongoose= require('mongoose');

const UserSchema = new mongoose.Schema({
  username:{
    type:String,
    require: [true,'Username required'],
    unique: true
  },
  password:{
    type:String,
    require:[true,'Password required']
  }
});

const User = mongoose.model('User',UserSchema);
module.exports = User;

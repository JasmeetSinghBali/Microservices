const mongoose= require('mongoose');

const PostSchema = new mongoose.Schema({
  title:{
    type:String,
    require: [true,'Title for the post required']
  },
  body:{
    type:String,
    require: [true,'Body for the post required']
  }
});

const Post = mongoose.model('Post',PostSchema);
module.exports = Post;

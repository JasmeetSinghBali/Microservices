const Post = require('../models/Posts');

exports.getAllPosts = async (req,res,next)=>{
  try{
    const posts = await Post.find();
    res.status(200).json({
      status :'Success',
      message :'All posts retrieved successfully',
      result : posts.length,
      data : {
        posts
      }
    });
  }catch(err){
    res.status(400).json({
      error: err,
      message: 'Failed! Something went wrong!!'
    });
  }
};

exports.getOnePosts = async (req,res,next)=>{
  try{
    const post = await Post.findById(req.params.id);
    res.status(200).json({
      status :'Success',
      message :'Single post retrieved successfully',
      data : {
        post
      }
    });
  }catch(err){
    res.status(400).json({
      error: err,
      message: 'Failed! Something went wrong!!'
    });
  }
};

exports.createPost = async (req,res,next)=>{
  try{
    const post = await Post.create(req.body);
    res.status(200).json({
      status :'Success',
      message :'Post created successfully',
      data : {
        post
      }
    });
  }catch(err){
    res.status(400).json({
      error: err,
      message: 'Failed! Something went wrong!!'
    });
  }
};

exports.updatePost = async (req,res,next)=>{
  try{
    const post = await Post.findByIdAndUpdate(req.params.id,req.body,{new:true,runValidators: true});
    res.status(200).json({
      status :'Success',
      message :'Post updated successfully',
      data : {
        post
      }
    });
  }catch(err){
    res.status(400).json({
      error: err,
      message: 'Failed! Something went wrong!!'
    });
  }
};

exports.deletePost = async (req,res,next)=>{
  try{
    const post = await Post.findByIdAndDelete(req.params.id);
    res.status(200).json({
      status :'Success',
      message :'Post deleted successfully'
    });
  }catch(err){
    res.status(400).json({
      error: err,
      message: 'Failed! Something went wrong!!'
    });
  }
};

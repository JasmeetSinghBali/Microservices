const express = require('express');

const postController = require('../controllers/postController');

const router = express.Router();

// get and post route chaining as they are going to the same route just the difference of get and post
router.route("/")
    .get(postController.getAllPosts)
    .post(postController.createPost);

// get,patch,delete by id routes
router.route("/:id")
    .get(postController.getOnePosts)
    .patch(postController.updatePost)
    .delete(postController.deletePost);

module.exports = router;

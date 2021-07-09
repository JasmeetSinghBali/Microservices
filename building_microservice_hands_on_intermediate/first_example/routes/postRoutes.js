const express = require('express');

const postController = require('../controllers/postController');

const protect = require('../middleware/authMiddleware');

const router = express.Router();

// get and post route chaining as they are going to the same route just the difference of get and post
router.route("/")
    .get(protect,postController.getAllPosts)
    .post(protect, postController.createPost);

// get,patch,delete by id routes
router.route("/:id")
    .get(protect,postController.getOnePosts)
    .patch(protect,postController.updatePost)
    .delete(protect,postController.deletePost);

module.exports = router;

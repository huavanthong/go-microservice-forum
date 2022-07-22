package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"github.com/huavanthong/microservice-golang/user-api-v3/payload"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) PostController {
	return PostController{postService}
}

// CreatePost godoc
// @Summary Creates a new post
// @Description User create a new post for topic
// @Tags posts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param post body models.CreatePostRequest true "New Post"
// @Failure 400 {object} payload.Response
// @Failure 409 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 201 {object} payload.CreatePostSuccess
// @Router /posts [post]
func (pc *PostController) CreatePost(ctx *gin.Context) {
	// prepare a post request from ctx
	var post *models.CreatePostRequest

	// from context, bind a new post info to json
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// call post service to create the post
	newPost, err := pc.postService.CreatePost(post)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusConflict,
					Message: err.Error(),
				})
			return
		}

		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated,
		payload.CreatePostSuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Create a new post success",
			Data:    newPost,
		})
}

// UpdatePost godoc
// @Summary Update a post
// @Description User update the exist post
// @Tags posts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param post body models.UpdatePost true "Update post"
// @Param postId path string true "Post ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.Response
// @Router /posts/{postId} [patch]
func (pc *PostController) UpdatePost(ctx *gin.Context) {

	// get post ID from URL path
	postId := ctx.Param("postId")

	// from context, bind a new post info to json
	var post *models.UpdatePost
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// call post service to update info
	updatedPost, err := pc.postService.UpdatePost(postId, post)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.UpdatePostSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Update a exist post success",
			Data:    updatedPost,
		})
}

// FindPostById godoc
// @Summary Retrieve a single post
// @Description User find the post by postId
// @Tags posts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param postId path string true "Post ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.Response
// @Router /posts/{postId} [get]
func (pc *PostController) FindPostById(ctx *gin.Context) {

	// get post ID from URL path
	postId := ctx.Param("postId")

	// call post service to find post by ID
	post, err := pc.postService.FindPostById(postId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.GetPostSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get the post success",
			Data:    post,
		})
}

// FindPosts godoc
// @Summary Retrieve all posts
// @Description User delete the post by postId
// @Tags posts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param page path string true "Post ID"
// @Param limit path string true "Post ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 204 {object} payload.Response
// @Router /posts/ [get]
func (pc *PostController) FindPosts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	posts, err := pc.postService.FindPosts(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

// DeletePost godoc
// @Summary Delete a post
// @Description User delete the post by postId
// @Tags posts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param postId path string true "Post ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 204 {object} payload.Response
// @Router /posts/{postId} [delete]
func (pc *PostController) DeletePost(ctx *gin.Context) {
	// get post ID from URL path
	postId := ctx.Param("postId")

	// call post service to delete post by ID
	err := pc.postService.DeletePost(postId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusNoContent,
		payload.Response{
			Status:  "success",
			Code:    http.StatusNoContent,
			Message: "Delete post success",
		})
}

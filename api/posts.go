package api

import (
	"Chatting/config"
	"Chatting/model"
	"Chatting/repository"
	"Chatting/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

type PostController struct{
	postRepository repository.PostRepository
}

// GetAllPost godoc
// @Summary Get all posts
// @Description Get all user items
// @Tags posts
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Success 200 {array} model.Post
// @Failure 500 {object} handler.APIError
// @Router /posts [get]
func (PostController *PostController) GetAllPost(c echo.Context) error{
	mdb, _ := config.MongoConnection()
	posts, _ := PostController.postRepository.GetAllPost(mdb, bson.M{})
	return utils.Negotiate(c,http.StatusOK, posts)
}

// SavePost godoc
// @Summary Create a post
// @Description Create a new post item
// @Tags posts
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param user body model.PostInput true "New Post"
// @Success 201 {object} model.User
// @Failure 400 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /posts [post]
func (postController *PostController) SavePost(c echo.Context) error {
	payload := new(model.PostInput)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return err
	}
	var post model.Post
	post.PostInput = payload
	mdb, _ := config.MongoConnection()
	createdPost, err := postController.postRepository.SavePost(mdb, post)
	if err != nil {
		return err
	}
	return utils.Negotiate(c, http.StatusCreated, createdPost)
}

// GetPost godoc
// @Summary Get a post
// @Description Get a post item
// @Tags posts
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Post ID"
// @Success 200 {object} model.Post
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /posts/{id} [get]
func (postController *PostController) GetPost(c echo.Context) error {

	mdb, _ := config.MongoConnection()
	id := c.Param("id")
	post, err := postController.postRepository.GetOnePost(mdb, bson.M{"id":id})
	if err != nil {
		return err
	}
	return utils.Negotiate(c, http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update a post
// @Description Update a post item
// @Tags posts
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Post ID"
// @Param user body model.PostInput true "Post Info"
// @Success 200 {object} model.Post
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /posts/{id} [put]
func (postController *PostController) UpdatePost(c echo.Context) error {
	post := new(model.Post)
	if err := utils.BindAndValidate(c, post) ; err != nil{
		log.Println("post BindAndValidate Failed" , err)
		return err
	}
	mdb, _ := config.MongoConnection()
	id := c.Param("id")
	user, err := postController.postRepository.UpdatePost(mdb,post,bson.M{"id":id})
	if err != nil {
		return err
	}
	return utils.Negotiate(c, http.StatusOK, user)
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a new post item
// @Tags posts
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Post ID"
// @Success 204 {object} model.Post
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /posts/{id} [delete]
func (postController *PostController) DeletePost(c echo.Context) error {
	id := c.Param("id")
	mdb, _ := config.MongoConnection()

	err := postController.postRepository.DeletePost(mdb,bson.M{"id":id})
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}



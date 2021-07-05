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
	"strconv"
)

type PostController struct{
	postRepository repository.PostRepository
}

// GetAllPost godoc
// @Summary Get all Posts
// @Description post get All
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {array}	model.Post
// @Failure 500 {object} api.APIError
// @Router /posts [get]
func (postController *PostController) GetAllPost(c echo.Context) error{
	mdb, _ := config.MongoConnection()
	posts, _ := repository.GetAllPost(mdb, bson.M{})
	return c.JSON(http.StatusOK,posts)
}
// SavePost godoc
// @Summary Save Post
// @Description save post
// @Tags posts
// @Param post body model.Post true "post Info"
// @Accept json
// @Produce json
// @Success 201 {object} model.Post
// @Failure 500 {object} api.APIError
// @Router /posts [post]
func (postController *PostController) SavePost(c echo.Context) error {
	payload := new(model.Post)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return err
	}
	var post model.Post
	post = *payload
	mdb, _ := config.MongoConnection()
	createdPost, err := repository.SavePost(mdb, post)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated,createdPost)
}

// GetPost godoc
// @Summary Get One Post
// @Description Get one post
// @Tags posts
// @Param id path int true "Post ID"
// @Accept json
// @Produce json
// @Success 200 {object} model.Post
// @Failure 400 {object} api.APIError
// @Failure 404 {object} api.APIError
// @Failure 500 {object} api.APIError
// @Router /posts/{id} [get]
func (postController *PostController) GetPost(c echo.Context) error {

	mdb, _ := config.MongoConnection()
	id,_ := strconv.Atoi(c.Param("id"))
	post, err := repository.GetOnePost(mdb, bson.M{"id":id})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,post)
}

// UpdatePost godoc
// @Summary Update Post
// @Description Update post
// @Tags posts
// @Param id path int true "Post ID"
// @Param post body model.Post true "post Info"
// @Accept json
// @Produce json
// @Success 200 {object} model.Post
// @Failure 500 {object} api.APIError
// @Router /posts/{id} [put]
func (postController *PostController) UpdatePost(c echo.Context) error {
	post := new(model.Post)
	if err := utils.BindAndValidate(c, post) ; err != nil{
		log.Println("post BindAndValidate Failed" , err)
		return err
	}
	mdb, _ := config.MongoConnection()
	id := c.Param("id")
	updateCnt, err := repository.UpdatePost(mdb,post,bson.M{"id":id})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,updateCnt)
}

// DeletePost godoc
// @Summary delete post
// @Description Delete post
// @Tags posts
// @Param id path int true "Post ID"
// @Accept json
// @Produce json
// @Success 200 {object} model.Post
// @Failure 500 {object} api.APIError
// @Router /posts/{id} [delete]
func (postController *PostController) DeletePost(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	mdb, _ := config.MongoConnection()

	_,err := repository.DeletePost(mdb,bson.M{"id":id})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,true)
}



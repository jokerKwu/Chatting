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

type UserController struct {
	userRepository repository.UserRepository
}
func NewUserController(userRepository repository.UserRepository) *UserController {
	return &UserController{userRepository: userRepository}
}
// PostLoginUser godoc
// @Summary User Login
// @Description User login
// @Tags users
// @Accept json
// @Produce json
// @Param name body model.User true "User Info"
// @Success 200 {object} model.User
// @Failure 500 {object} api.APIError
// @Router /login [post]
func (UserController *UserController) PostLoginUser(c echo.Context) error{
	ctx := utils.CtxGenerate(c.Request(),"","")
	payload := new(model.User)
	if err := utils.BindAndValidate(c, payload); err != nil{
		return err
	}
	mdb, _ := config.MongoConnection()
	dbUser, err := repository.GetOneUser(mdb,bson.M{"name":payload.Name})

	//User 정보가 일치하는지 체크
	if dbUser.Name != payload.Name || dbUser.Password != payload.Password{
		log.Println("User Info different")
		return err
	}
	accessToken, refreshToken, _, err := utils.GenerateTokenPair(ctx,utils.TokenData{Name:payload.Name})
	if err != nil{
		return err
	}

	log.Println("test ======================")
	log.Println(c.Cookie("access_token"))
	log.Println(c.Cookie("refresh_token"))
	return c.JSON(http.StatusOK,bson.M{"access_token":accessToken,"refresh_token":refreshToken})
}
// PostUser godoc
// @Summary Post User
// @Description user Save
// @Param user body model.User true "user info"
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 500 {object} api.APIError
// @Router /join [post]
func(UserController *UserController) PostUser(c echo.Context) error{
	payload := new(model.User)
	if err := utils.BindAndValidate(c, payload) ; err != nil{
		return err
	}
	var user model.User
	user = *payload
	mdb, _ := config.MongoConnection()
	log.Println(user.Name, user.Password)
	createdUser, err := repository.PostUser(mdb,&user)
	if err != nil{
		return err
	}
	return c.JSON(http.StatusCreated,createdUser)
}
package api

import (
	"Chatting/model"
	"Chatting/repository"
	"Chatting/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	return &UserController{userRepository: userRepository}
}

func (userController *UserController) GetAllUser(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

	pagedUser, _ := userController.userRepository.GetAllUser(page, limit)
	return utils.Negotiate(c, http.StatusOK, pagedUser)
}
func (userController *UserController) SaveUser(c echo.Context) error {
	payload := new(model.UserInput)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return err
	}

	createdUser, err := userController.userRepository.SaveUser(&model.User{UserInput: payload})
	if err != nil {
		return err
	}

	return utils.Negotiate(c, http.StatusCreated, createdUser)
}

func (userController *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := userController.userRepository.GetUser(id)
	if err != nil {
		return err
	}

	return utils.Negotiate(c, http.StatusOK, user)
}

func (userController *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	payload := new(model.UserInput)

	if err := utils.BindAndValidate(c, payload); err != nil {
		return err
	}

	user, err := userController.userRepository.UpdateUser(id, &model.User{UserInput: payload})
	if err != nil {
		return err
	}
	return utils.Negotiate(c, http.StatusOK, user)
}

func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := userController.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

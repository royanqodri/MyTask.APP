package handler

import (
	"net/http"
	"royan/cleanarch/features/user"
	"royan/cleanarch/helpers"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c *gin.Context) {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "error bind data. data not valid", nil)
		return
	}

	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			helpers.WebResponse(c, http.StatusBadRequest, err.Error(), nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "error login", nil)
		}
		return
	}

	response := map[string]interface{}{
		"token":   token,
		"user_id": dataLogin.ID,
		"name":    dataLogin.Name,
	}

	helpers.WebResponse(c, http.StatusOK, "success login", response)
}

func (handler *UserHandler) GetAllUser(c *gin.Context) {
	result, err := handler.userService.GetAll()
	if err != nil {
		helpers.WebResponse(c, http.StatusInternalServerError, "error read data", nil)
		return
	}

	var usersResponse []UserResponse
	for _, value := range result {
		usersResponse = append(usersResponse, UserResponse{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			CreatedAt:   value.CreatedAt,
		})
	}

	helpers.WebResponse(c, http.StatusOK, "success read data", usersResponse)
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "error bind data. data not valid", nil)
		return
	}

	userCore := RequestToCore(*userInput)
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			helpers.WebResponse(c, http.StatusBadRequest, err.Error(), nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "error insert data", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusCreated, "success insert data", nil)
}

func (handler *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("user_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "wrong id", nil)
		return
	}

	result, err := handler.userService.GetById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			helpers.WebResponse(c, http.StatusBadRequest, err.Error(), nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "error read data", nil)
		}
		return
	}

	resultResponse := UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		Address:     result.Address,
		PhoneNumber: result.PhoneNumber,
		CreatedAt:   result.CreatedAt,
	}

	helpers.WebResponse(c, http.StatusOK, "success read data", resultResponse)
}

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "Error binding data", nil)
		return
	}

	userCore := RequestToCore(*userInput)
	err := handler.userService.Update(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			helpers.WebResponse(c, http.StatusNotFound, "User not found", nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "error updating data user", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusOK, "User Updated successfully", nil)
}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		helpers.WebResponse(c, http.StatusBadRequest, "Error binding data", nil)
		return
	}

	var userCore user.Core
	err := handler.userService.Delete(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			helpers.WebResponse(c, http.StatusNotFound, "User not found", nil)
		} else {
			helpers.WebResponse(c, http.StatusInternalServerError, "error Deleted data user", nil)
		}
		return
	}

	helpers.WebResponse(c, http.StatusOK, "User Deleted successfully", nil)
}

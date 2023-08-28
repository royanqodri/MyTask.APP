package handler

import (
	"net/http"
	"royan/cleanarch/features/user"
	"royan/cleanarch/helpers"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	response := map[string]any{
		"token":   token,
		"user_id": dataLogin.ID,
		"name":    dataLogin.Name,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success login", response))
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	result, err := handler.userService.GetAll()
	if err != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]any{
		// 	"code": 500,
		// 	"message": "hello world",
		// })
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
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
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", usersResponse))
	// return c.JSON(http.StatusOK, map[string]any{
	// 	"code":    200,
	// 	"message": "success read data",
	// 	"data":    usersResponse,
	// })

}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	userCore := RequestToCore(*userInput)
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("user_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}
	result, err := handler.userService.GetById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	resultResponse := UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		Address:     result.Address,
		PhoneNumber: result.PhoneNumber,
		CreatedAt:   result.CreatedAt,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", resultResponse))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {

	// Ambil data pengguna yang akan diperbarui dari permintaan JSON
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error binding data", nil))
	}

	// Mapping data dari UserRequest ke struct Core
	userCore := RequestToCore(*userInput)
	err := handler.userService.Update(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "User not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating data user", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User Updateed successfully", nil))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error binding data", nil))
	}

	// Mapping data dari UserRequest ke struct Core
	var userCore user.Core
	err := handler.userService.Delete(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "User not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error Deleted data user", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User Deleted successfully", nil))
}

// Panggil fungsi service untuk menghapus pengguna berdasarkan `user_id`
// 	err = handler.userService.Delete(uint(userID))
// 	if err != nil {
// 		if strings.Contains(err.Error(), "not found") {
// 			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "User not found", nil))
// 		}
// 		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error deleting user", nil))

// 	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User deleted successfully", nil))
// }

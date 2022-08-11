package delivery

import (
	"log"
	"net/http"
	"petadopter/domain"
	common "petadopter/features/common"
	auth "petadopter/utils/google"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var (
	oauthStateString = "pseudo-random"
)

type userHandler struct {
	userUsecase domain.UserUseCase
	oauth       *oauth2.Config
}

func New(us domain.UserUseCase, o *oauth2.Config) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
		oauth:       o,
	}
}

// LoginGoogle implements domain.UserHandler
func (us *userHandler) SignUpGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := us.oauth.AuthCodeURL(oauthStateString)
		return c.Redirect(http.StatusFound, url)
	}
}

// LoginGoogle implements domain.UserHandler
func (us *userHandler) CallbackGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		cost := 10

		data, err, token := auth.GetUserInfo(us.oauth, c.FormValue("state"), c.FormValue("code"), oauthStateString)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		newuser.Email = data.Email
		newuser.Fullname = data.Fullname
		newuser.Photoprofile = data.Photoprofile
		newuser.Username = data.Fullname

		status := us.userUsecase.RegisterUser(newuser.ToModel(), cost, token)

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == http.StatusConflict {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"code":    status,
				"message": "Cant input existing data",
			})
		}

		if status == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Register success",
		})
	}
}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		var arrmap = map[string]interface{}{}

		errLog := c.Bind(&userLogin)

		if errLog != nil {
			log.Println("invalid input")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		data, err := us.userUsecase.Login(userLogin.ToModelLogin())

		if err != nil {
			log.Println("Login failed", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong username or password",
			})
		}

		token := common.GenerateToken(data)

		arrmap["token"] = token
		arrmap["username"] = data.Username
		arrmap["role"] = data.Role

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    arrmap,
			"code":    http.StatusOK,
			"message": "Login success",
		})
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := common.ExtractData(c)

		status, err := us.userUsecase.Delete(data.ID)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if !status {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success delete user",
		})
	}
}

// Register implements domain.UserHandler
func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		bind := c.Bind(&newuser)
		cost := 10

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		status := us.userUsecase.RegisterUser(newuser.ToModel(), cost, nil)

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == http.StatusConflict {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"code":    status,
				"message": "Cant input existing data",
			})
		}

		if status == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Register success",
		})
	}
}

// Update implements domain.UserHandler
func (us *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UpdateFormat
		cost := 10
		param := common.ExtractData(c)
		bind := c.Bind(&newuser)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		status := us.userUsecase.UpdateUser(newuser.ToModelUpdate(), param.ID, cost)

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == http.StatusConflict {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"code":    status,
				"message": "Cant input existing data",
			})
		}

		if status == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Update success",
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := common.ExtractData(c)
		var res = map[string]interface{}{}
		data, err := uh.userUsecase.GetProfile(usr.ID)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}
		res["username"] = data.Username
		res["fullname"] = data.Fullname
		res["phonenumber"] = data.Phonenumber
		res["email"] = data.Email
		res["address"] = data.Address
		res["photoprofile"] = data.PhotoProfile
		res["city"] = data.City

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"code":    http.StatusOK,
			"message": "get data success",
		})
	}
}

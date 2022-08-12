package delivery

import (
	"log"
	"net/http"
	"petadopter/config"
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

func (us *userHandler) LoginGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		us.oauth.RedirectURL = "http://localhost:8000/callback/login"
		url := us.oauth.AuthCodeURL(oauthStateString)

		return c.Redirect(http.StatusFound, url)
	}
}

func (us *userHandler) SignUpGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		us.oauth.RedirectURL = "http://localhost:8000/callback/signup"
		url := us.oauth.AuthCodeURL(oauthStateString)

		return c.Redirect(http.StatusFound, url)
	}
}

// CallbackGoogleLogin implements domain.UserHandler
func (us *userHandler) CallbackGoogleLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataLogin UserInfoFormat

		dataInfo, err, token := auth.GetUserInfo(us.oauth, c.FormValue("state"), c.FormValue("code"), oauthStateString)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		dataLogin = UserInfoFormat(dataInfo)

		res, status := us.userUsecase.Login(dataLogin.ToModelUserInfoFormat(), token)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"code":    http.StatusOK,
			"message": "Register success",
		})
	}
}

// CallbackGoogleSignUp implements domain.UserHandler
func (us *userHandler) CallbackGoogleSignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat

		data, err, token := auth.GetUserInfo(us.oauth, c.FormValue("state"), c.FormValue("code"), oauthStateString)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}
		status := us.userUsecase.RegisterUser(newuser.ToModel(), config.COST, token, data)

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if status == http.StatusConflict {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"code":    http.StatusConflict,
				"message": "Cant input existing data",
			})
		}

		if status == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
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

		errLog := c.Bind(&userLogin)
		if errLog != nil {
			log.Println("invalid input")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "There is an error in internal server",
			})
		}

		data, status := us.userUsecase.Login(userLogin.ToModelLogin(), nil)
		if status == 400 {
			log.Println("Login failed")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong username or password",
			})
		}

		if status == 404 {
			log.Println("Login failed")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Wrong username or password",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    http.StatusOK,
			"message": "Login success",
		})
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := common.ExtractData(c)
		status := us.userUsecase.Delete(data.ID)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if status == 500 {
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
		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		status := us.userUsecase.RegisterUser(newuser.ToModel(), config.COST, nil, domain.UserInfo{})

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if status == http.StatusConflict {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"code":    http.StatusConflict,
				"message": "Cant input existing data",
			})
		}

		if status == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Register success",
		})
	}
}

// Update implements domain.UserHandler
func (us *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UpdateFormat
		param := common.ExtractData(c)

		bind := c.Bind(&newuser)
		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		status := us.userUsecase.UpdateUser(newuser.ToModelUpdate(), param.ID, config.COST)

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
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
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Update success",
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := common.ExtractData(c)

		arrmap, status := uh.userUsecase.GetProfile(usr.ID)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    arrmap,
			"code":    http.StatusOK,
			"message": "get data success",
		})
	}
}

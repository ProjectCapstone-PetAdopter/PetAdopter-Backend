package delivery

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"petadopter/config"
	"petadopter/domain"
	common "petadopter/features/common"
	"petadopter/utils/google"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var (
	oauthStateString = "pseudo-random"
)

type userHandler struct {
	userUsecase domain.UserUseCase
	oauth       *oauth2.Config
	client      *google.ClientUploader
}

func New(us domain.UserUseCase, o *oauth2.Config, cl *google.ClientUploader) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
		oauth:       o,
		client:      cl,
	}
}

// GetbyID implements domain.UserHandler
func (us *userHandler) GetbyID() echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := c.Param("id")

		cnv, err := strconv.Atoi(usr)
		if err != nil {
			log.Println("cant convert to int")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		arrmap, status := us.userUsecase.GetProfile(cnv)

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

func (us *userHandler) LoginGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		us.oauth.RedirectURL = "https://golangprojectku.site/callback/login"
		url := us.oauth.AuthCodeURL(oauthStateString)

		return c.Redirect(http.StatusFound, url)
	}
}

func (us *userHandler) SignUpGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		us.oauth.RedirectURL = "https://golangprojectku.site/callback/signup"
		url := us.oauth.AuthCodeURL(oauthStateString)

		return c.Redirect(http.StatusFound, url)
	}
}

// CallbackGoogleLogin implements domain.UserHandler
func (us *userHandler) CallbackGoogleLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataLogin UserInfoFormat

		dataInfo, err, token := google.GetUserInfo(us.oauth, c.FormValue("state"), c.FormValue("code"), oauthStateString)
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

		urlstr := fmt.Sprintf("http://localhost:3000/auth/redirect?token=%s&role=%s&tokenoauth=%s", res["token"], res["role"], res["tokenoauth"])

		var buf bytes.Buffer
		buf.WriteString(urlstr)
		v := url.Values{}
		buf.WriteString(v.Encode())

		return c.Redirect(http.StatusFound, buf.String())
		// return c.JSON(http.StatusOK, map[string]interface{}{
		// 	"data":    string(buf.String()),
		// 	"code":    http.StatusOK,
		// 	"message": "Register success",
		// })
	}
}

// CallbackGoogleSignUp implements domain.UserHandler
func (us *userHandler) CallbackGoogleSignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat

		data, err, token := google.GetUserInfo(us.oauth, c.FormValue("state"), c.FormValue("code"), oauthStateString)
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

		urlstr := "http://localhost:3000/auth/redirect"

		var buf bytes.Buffer
		buf.WriteString(urlstr)
		v := url.Values{}
		buf.WriteString(v.Encode())

		return c.Redirect(http.StatusFound, buf.String())
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
		if newuser.Photoprofile != "" {
			form, err := c.FormFile("photoprofile")
			if err != nil {
				log.Println(err, "cant get form")
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "There is an error in internal server",
				})
			}

			file, err := form.Open()
			if err != nil {
				log.Println(err, "cant open file")
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "There is an error in internal server",
				})
			}

			id := uuid.New()
			filename := fmt.Sprintf("%sPP-%s.jpg", newuser.Username, id.String())
			config.UPLOADPATH = "profile/"

			link, err := us.client.UploadFile(file, config.UPLOADPATH, filename)
			if err != nil {
				log.Println(err, "cant upload file")
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "There is an error in internal server",
				})
			}
			newuser.Photoprofile = link
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

func (us *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := common.ExtractData(c)

		arrmap, status := us.userUsecase.GetProfile(usr.ID)

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

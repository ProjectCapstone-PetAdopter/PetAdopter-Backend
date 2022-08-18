package delivery

import (
	"fmt"
	"log"
	"net/http"
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/common"
	"petadopter/utils/google"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type petsHandler struct {
	petsUsecase domain.PetsUseCase
	client      *google.ClientUploader
}

func New(cu domain.PetsUseCase, cl *google.ClientUploader) domain.PetsHandler {
	return &petsHandler{
		petsUsecase: cu,
		client:      cl,
	}
}

func (ph *petsHandler) InsertPets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp PetsInsertRequest

		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input data",
			})
		}

		token := common.ExtractData(c)
		if token.ID == 0 {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		form, err := c.FormFile("petphoto")
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
		filename := fmt.Sprintf("%dPost-%s.jpg", tmp.Userid, id.String())
		config.UPLOADPATH = "post/"

		link, err := ph.client.UploadFile(file, config.UPLOADPATH, filename)
		if err != nil {
			log.Println(err, "cant upload file")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}
		tmp.Petphoto = link

		_, errs := ph.petsUsecase.AddPets(tmp.ToDomain(), token.ID)
		if errs != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Post pet success",
		})

	}
}

func (ph *petsHandler) UpdatePets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp PetsInsertRequest

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		if cnv <= 0 {
			log.Println("Cant get id")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		token := common.ExtractData(c)
		if token.ID == 0 {
			log.Println("Cant get token data")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		res := c.Bind(&tmp)
		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Please enter data correctly",
			})
		}
		if tmp.Petphoto != "" {
			form, err := c.FormFile("petphoto")
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
			filename := fmt.Sprintf("%dPost-%s.jpg", tmp.Userid, id.String())
			config.UPLOADPATH = "post/"

			link, err := ph.client.UploadFile(file, config.UPLOADPATH, filename)
			if err != nil {
				log.Println(err, "cant upload file")
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "There is an error in internal server",
				})
			}
			tmp.Petphoto = link
		}
		data, errs := ph.petsUsecase.UpPets(cnv, tmp.ToDomain(), token.ID)
		if errs != nil {
			log.Println("Cannot update data", errs)
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data Not Found",
			})
		}

		if data.ID == 0 {
			log.Println("Cannot update data", err)
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data Not Found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"code":    http.StatusOK,
		})
	}
}

func (ph *petsHandler) DeletePets() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		if cnv <= 0 {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		data, err := ph.petsUsecase.DelPets(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success delete pet data",
		})
	}
}

func (ph *petsHandler) GetAllPets() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ph.petsUsecase.GetAllP()

		if err != nil {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if data == nil {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    http.StatusOK,
			"message": "success update data",
		})
	}
}

func (ph *petsHandler) GetPetsID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idNews := c.Param("id")

		id, errs := strconv.Atoi(idNews)
		if errs != nil {
			log.Println("Cannot convert to int", errs.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		data, err := ph.petsUsecase.GetSpecificPets(id)
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get Data",
			"data":    data,
		})
	}
}

func (ph *petsHandler) GetmyPets() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := common.ExtractData(c)

		if token.ID == 0 {
			log.Println("User not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		data, err := ph.petsUsecase.GetmyPets(token.ID)
		if err != nil {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get my pets",
			"data":    data,
		})
	}
}

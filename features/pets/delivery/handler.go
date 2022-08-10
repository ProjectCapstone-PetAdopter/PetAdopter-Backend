package delivery

import (
	"log"
	"net/http"
	"petadopter/domain"
	"petadopter/features/common"
	"strconv"

	"github.com/labstack/echo/v4"
)

type petsHandler struct {
	petsUsecase domain.PetsUseCase
}

func New(cu domain.PetsUseCase) domain.PetsHandler {
	return &petsHandler{
		petsUsecase: cu,
	}
}

func (ph *petsHandler) InsertPets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp PetsInsertRequest

		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong input data",
			})
		}

		token := common.ExtractData(c)
		if token.ID == 0 {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		tmp.Userid = token.ID

		_, errs := ph.petsUsecase.AddPets(tmp.ToDomain())
		if errs != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Post pet success",
		})

	}
}

func (ph *petsHandler) UpdatePets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp PetsInsertRequest
		qry := map[string]interface{}{}

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		if cnv <= 0 {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		token := common.ExtractData(c)
		if token.ID == 0 {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		tmp.Userid = token.ID

		res := c.Bind(&tmp)
		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Please enter data correctly",
			})
		}

		if tmp.Petname != "" {
			qry["petname"] = tmp.Petname
		}
		if tmp.Gender != 0 {
			qry["gender"] = tmp.Gender
		}
		if tmp.Age != 0 {
			qry["age"] = tmp.Age
		}
		if tmp.Color != "" {
			qry["color"] = tmp.Color
		}
		if tmp.Description != "" {
			qry["description"] = tmp.Description
		}
		if tmp.Petphoto != "" {
			qry["petphoto"] = tmp.Petphoto
		}

		data, errs := ph.petsUsecase.UpPets(cnv, tmp.ToDomain())
		if errs != nil {
			log.Println("Cannot update data", err)
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data Not Found",
			})
		}

		if data.ID == 0 {
			log.Println("Cannot update data", err)
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data Not Found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"code":    200,
		})
	}
}

func (ph *petsHandler) DeletePets() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		if cnv <= 0 {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		data, err := ph.petsUsecase.DelPets(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Success delete pet data",
		})
	}
}

func (ph *petsHandler) GetAllPets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var arrmap = []map[string]interface{}{}

		data, err := ph.petsUsecase.GetAllP()
		if len(data) == 0 {
			log.Println("Data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		for i := 0; i < len(data); i++ {
			var res = map[string]interface{}{}
			res["id"] = data[i].ID
			res["petname"] = data[i].Petname
			res["petphoto"] = data[i].Petphoto
			res["species"] = data[i].Speciesid
			res["gender"] = data[i].Gender
			res["age"] = data[i].Age
			res["color"] = data[i].Color
			res["description"] = data[i].Description

			arrmap = append(arrmap, res)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    arrmap,
			"code":    200,
			"message": "success update data",
		})
	}
}

func (ph *petsHandler) GetPetsID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res = map[string]interface{}{}
		idNews := c.Param("id")

		id, errs := strconv.Atoi(idNews)
		if errs != nil {
			log.Println("Cannot convert to int", errs.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		data, datapetuser, err := ph.petsUsecase.GetSpecificPets(id)
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		res["petname"] = data[0].Petname
		res["petphoto"] = data[0].Petphoto
		res["species"] = data[0].Speciesid
		res["gender"] = data[0].Gender
		res["age"] = data[0].Age
		res["color"] = data[0].Color
		res["description"] = data[0].Description
		res["ownername"] = datapetuser.Fullname
		res["city"] = datapetuser.City

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get Data",
			"data":    res,
		})
	}
}

func (ph *petsHandler) GetmyPets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var arrmap = []map[string]interface{}{}
		token := common.ExtractData(c)

		if token.ID == 0 {
			log.Println("User not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		data, err := ph.petsUsecase.GetmyPets(token.ID)
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		for i := 0; i < len(data); i++ {
			var res = map[string]interface{}{}
			res["petname"] = data[i].Petname
			res["petphoto"] = data[i].Petphoto
			res["species"] = data[i].Speciesid
			res["gender"] = data[i].Gender
			res["age"] = data[i].Age
			res["color"] = data[i].Color
			res["description"] = data[i].Description

			arrmap = append(arrmap, res)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get my pets",
			"data":    arrmap,
		})
	}
}

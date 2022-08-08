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
			c.JSON(http.StatusBadRequest, "error read input")
		}

		userid, _ := common.ExtractData2(c)
		data, err := ph.petsUsecase.AddPets(userid, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}

func (ph *petsHandler) UpdatePets() echo.HandlerFunc {
	return func(c echo.Context) error {

		qry := map[string]interface{}{}
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp PetsInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		if tmp.Name != "" {
			qry["name"] = tmp.Name
		}
		if tmp.Gender != "" {
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
		if tmp.Images != "" {
			qry["images"] = tmp.Images
		}
		data, err := ph.petsUsecase.UpPets(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "success update data",
			"id":          data.ID,
			"Name":        data.Name,
			"Gender":      data.Gender,
			"Age":         data.Age,
			"Color":       data.Color,
			"Description": data.Description,
			"Images":      data.Images,
		})
	}
}

func (ph *petsHandler) DeletePets() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := ph.petsUsecase.DelPets(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete Data")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete Data",
		})
	}
}

func (ph *petsHandler) GetAllPets() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ph.petsUsecase.GetAllP()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all Data",
			"data":    data,
		})
	}
}

func (ph *petsHandler) GetPetsID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idNews := c.Param("id")
		id, _ := strconv.Atoi(idNews)
		data, err := ph.petsUsecase.GetSpecificPets(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get Data",
			"data":    data,
		})
	}
}

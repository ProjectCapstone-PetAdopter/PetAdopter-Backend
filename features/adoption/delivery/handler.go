package delivery

import (
	"log"
	"net/http"
	"petadopter/domain"
	"petadopter/features/common"
	"strconv"

	"github.com/labstack/echo/v4"
)

type adoptionHandler struct {
	adoptionUsecase domain.AdoptionUseCase
}

func New(pu domain.AdoptionUseCase) domain.AdoptionHandler {
	return &adoptionHandler{
		adoptionUsecase: pu,
	}
}
func (ad *adoptionHandler) UpdateAdoption() echo.HandlerFunc {
	return func(c echo.Context) error {

		qry := map[string]interface{}{}
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp AdoptionInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		if tmp.PetsID != 0 {
			qry["pets_id"] = tmp.PetsID
		}
		if tmp.UserID != 0 {
			qry["user_id"] = tmp.UserID
		}
		if tmp.Petphoto != "" {
			qry["petphoto"] = tmp.Petphoto
		}
		if tmp.Status != "" {
			qry["status"] = tmp.Status
		}

		data, err := ad.adoptionUsecase.UpAdoption(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success Update",
			"data":    FromDomain(data),
		})
	}
}

func (ad *adoptionHandler) InsertAdoption() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp AdoptionInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		userid, _ := common.ExtractData2(c)

		data, err := ad.adoptionUsecase.AddAdoption(userid, tmp.ToDomain())

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

func (ad *adoptionHandler) DeleteAdoption() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := ad.adoptionUsecase.DelAdoption(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete",
		})
	}
}

func (ad *adoptionHandler) GetAllAdoption() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ad.adoptionUsecase.GetAllAP()
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

func (ad *adoptionHandler) GetAdoptionID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idOrder := c.Param("id")
		id, _ := strconv.Atoi(idOrder)
		data, err := ad.adoptionUsecase.GetSpecificAdoption(id)
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data",
			"data":    data,
		})
	}
}

func (ad *adoptionHandler) GetMYAdopt() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, _ := common.ExtractData2(c)
		data, err := ad.adoptionUsecase.GetmyAdoption(userid)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my Data",
			"data":    data,
		})
	}
}

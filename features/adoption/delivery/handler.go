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
		var tmp AdoptionInsertRequest

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}

		if cnv <= 0 {
			log.Println("data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		res := c.Bind(&tmp)
		if res != nil {
			log.Println("Cannot bind data", res)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input data",
			})
		}

		_, errs := ad.adoptionUsecase.UpAdoption(cnv, tmp.ToDomain())
		if errs != nil {
			log.Println("Cannot update data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input data",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success update data",
		})
	}
}

func (ad *adoptionHandler) InsertAdoption() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp AdoptionInsertRequest

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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}

		_, errs := ad.adoptionUsecase.AddAdoption(token.ID, tmp.ToDomain())
		if errs != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input data",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Post pets success",
		})
	}
}

func (ad *adoptionHandler) DeleteAdoption() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}

		data, err := ad.adoptionUsecase.DelAdoption(cnv)
		if err != nil {
			log.Println("cant delete data")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if !data {
			log.Println("cant delete data")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success delete adoption data",
		})
	}
}

//get owner applier data
func (ad *adoptionHandler) GetAllAdoption() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := common.ExtractData(c)
		if token.ID == 0 {
			log.Println("Cannot get token")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}

		data, err := ad.adoptionUsecase.GetAllAP(token.ID)
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get all Data",
			"data":    data,
		})
	}
}

func (ad *adoptionHandler) GetAdoptionID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idOrder := c.Param("id")

		id, errs := strconv.Atoi(idOrder)
		if errs != nil {
			log.Println("Cannot convert to int", errs.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		data, err := ad.adoptionUsecase.GetSpecificAdoption(id)
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
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    data,
		})
	}
}

//user history adoption
func (ad *adoptionHandler) GetMYAdopt() echo.HandlerFunc {
	return func(c echo.Context) error {

		token := common.ExtractData(c)

		data, err := ad.adoptionUsecase.GetmyAdoption(token.ID)
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
			"code":    http.StatusOK,
			"message": "success get my Data",
			"data":    data,
		})
	}
}

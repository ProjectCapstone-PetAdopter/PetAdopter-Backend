package delivery

import (
	"fmt"
	"log"
	"net/http"
	"petadopter/domain"
	"petadopter/features/common"
	"strconv"

	"github.com/labstack/echo/v4"
)

type speciesHandler struct {
	speciesUseCase domain.SpeciesUsecase
}

func New(su domain.SpeciesUsecase) domain.SpeciesHandler {
	return &speciesHandler{
		speciesUseCase: su,
	}
}

func (sh *speciesHandler) AddSpecies() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("ok")
		token := common.ExtractData(c)
		if token.Role == "user" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "unauthorized",
			})
		}

		// errGet := sh.speciesUseCase.GetUser(uint(token.ID))
		// if errGet != nil {
		// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		// 		"code":http.StatusUnauthorized,
		// 		"message": errGet.Error(),
		// 	})
		// }

		var tmp InserFormat
		errBind := c.Bind(&tmp)
		if errBind != nil {
			log.Println("cannot parse data", errBind)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}
		_, err := sh.speciesUseCase.AddSpecies(tmp.ToModel())
		if err != nil {
			log.Println("cannot proces data", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "succes operation",
		})
	}
}

func (sh *speciesHandler) GetSpecies() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := sh.speciesUseCase.GetAllSpecies()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get species data",
			"data":    FromModelToList(data),
		})
	}
}

func (sh *speciesHandler) UpdateDataSpecies() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := common.ExtractData(c)
		if token.Role == "user" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "unauthorized",
			})
		}

		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Please enter data correctly",
			})
		}

		var updatedData InserFormat
		err = c.Bind(&updatedData)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		data, err := sh.speciesUseCase.UpdateSpecies(id, updatedData.ToModel())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success update species " + param,
			"data":    data,
		})
	}
}

func (sh *speciesHandler) DeleteDataSpecies() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := common.ExtractData(c)
		if token.Role == "user" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "unauthorized",
			})
		}

		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Please enter data correctly",
			})
		}

		_, errDel := sh.speciesUseCase.DeleteSpecies(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success Operation",
		})
	}
}

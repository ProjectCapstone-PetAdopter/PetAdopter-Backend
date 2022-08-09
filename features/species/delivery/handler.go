package delivery

import (
	"log"
	"net/http"
	"petadopter/config"
	"petadopter/domain"
	"petadopter/features/common"
	"petadopter/features/species/delivery/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type speciesHandler struct {
	speciesUseCase domain.SpeciesUsecase
}

func New(e *echo.Echo, su domain.SpeciesUsecase) {
	handler := &speciesHandler{
		speciesUseCase: su,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/species", handler.AddSpeciesHandler(), useJWT)
}

func (sh *speciesHandler) AddSpeciesHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		errGet := sh.speciesUseCase.GetUser(uint(common.ExtractData3(c)))
		if errGet != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": errGet.Error(),
			})
		}

		var tmp InserFormat
		errBind := c.Bind(&tmp)
		if errBind != nil {
			log.Println("cannot parse data", errBind)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error",
			})
		}
		_, err := sh.speciesUseCase.AddSpeciesUseCase(tmp.ToModel())
		if err != nil {
			log.Println("cannot proces data", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "succes operation",
		})
	}
}

func (sh *speciesHandler) GetSpecies() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := sh.speciesUseCase.GetAllSpecies()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all species data",
			"data":    data,
		})
	}
}

func (sh *speciesHandler) UpdateDataSpecies() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := sh.speciesUseCase.GetUser(uint(common.ExtractData3(c)))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": err.Error(),
			})
		}

		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		var updatedData InserFormat
		err = c.Bind(&updatedData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		data, err := sh.speciesUseCase.UpdateSpecies(id, updatedData.ToModel())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update species " + param,
			"data":    data,
		})
	}
}

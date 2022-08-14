package delivery

import (
	"log"
	"net/http"
	"petadopter/domain"
	"petadopter/features/common"
	"strconv"

	"github.com/labstack/echo/v4"
)

type meetingHandler struct {
	meetingUsecase domain.MeetingUsecase
}

func New(mu domain.MeetingUsecase) domain.MeetingHandler {
	return &meetingHandler{
		meetingUsecase: mu,
	}
}

func (mh *meetingHandler) InsertMeeting() echo.HandlerFunc {
	return func(c echo.Context) error {
		var insertMeeting InsertMeeting
		err := c.Bind((&insertMeeting))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong input data",
			})
		}
		token := common.ExtractData(c)
		insertMeeting.Userid = token.ID
		if token.ID == 0 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		_, err = mh.meetingUsecase.AddMeeting(insertMeeting.ToModel())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "post meeting success",
		})
	}
}

func (mh *meetingHandler) UpdateDataMeeting() echo.HandlerFunc {
	return func(c echo.Context) error {

		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Please enter data correctly",
			})
		}

		var updatedData InsertMeeting
		err = c.Bind(&updatedData)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		_, errMeet := mh.meetingUsecase.UpdateMeeting(updatedData.ToModel(), id)
		if errMeet != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success update data",
		})
	}
}

func (mh *meetingHandler) DeleteDataMeeting() echo.HandlerFunc {
	return func(c echo.Context) error {

		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Please enter data correctly",
			})
		}

		errDel := mh.meetingUsecase.DeleteMeeting(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Success Operation",
		})
	}
}

func (mh *meetingHandler) GetAdopt() echo.HandlerFunc {
	return func(c echo.Context) error {
		var meetingid int
		data, err := mh.meetingUsecase.GetMyMeeting(meetingid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}
		if data == nil {
			log.Println("data not found")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "data not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get meeting",
			"data":    data,
		})
	}
}

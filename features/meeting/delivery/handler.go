package delivery

import (
	"fmt"
	"log"
	"net/http"
	"petadopter/domain"
	"petadopter/features/common"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type meetingHandler struct {
	meetingUsecase domain.MeetingUsecase
	oauth          *oauth2.Config
}

func New(mu domain.MeetingUsecase, o *oauth2.Config) domain.MeetingHandler {
	return &meetingHandler{
		meetingUsecase: mu,
		oauth:          o,
	}
}

func (mh *meetingHandler) InsertMeeting() echo.HandlerFunc {
	return func(c echo.Context) error {
		var insertMeeting InsertMeeting

		token := common.ExtractData(c)
		insertMeeting.Userid = token.ID

		if token.ID == 0 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		err := c.Bind((&insertMeeting))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong input data",
			})
		}

		idMeet, err := mh.meetingUsecase.AddMeeting(insertMeeting.ToModel())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
			})
		}

		if insertMeeting.Token != "" {
			owner, seeker := mh.meetingUsecase.GetEmail(token.ID, idMeet)
			if owner.Email == "" {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    500,
					"message": "Internal Server Error",
				})
			}

			dateTime := fmt.Sprintf("%sT%s+07:00", insertMeeting.Date, insertMeeting.Time)
			location := fmt.Sprintf("%s,%s", owner.Address, owner.City)

			events := &calendar.Event{
				Summary:     "Meeting for adoption",
				Description: "Meeting with owner",
				Start: &calendar.EventDateTime{
					DateTime: dateTime,
					TimeZone: "Asia/Jakarta",
				},
				End: &calendar.EventDateTime{
					DateTime: dateTime,
					TimeZone: "Asia/Jakarta",
				},
				Attendees: []*calendar.EventAttendee{
					{Email: owner.Email},
					{Email: seeker.Email},
				},
				Location: location,
			}

			tokenOauth := &oauth2.Token{AccessToken: insertMeeting.Token}

			client := mh.oauth.Client(c.Request().Context(), tokenOauth)

			srv, err := calendar.NewService(c.Request().Context(), option.WithHTTPClient(client))
			if err != nil {
				log.Printf("Unable to retrieve Calendar client: %v", err)
			}

			_, err = srv.Events.Insert("primary", events).SendUpdates("all").Do()
			if err != nil {
				log.Printf("Unable to create event. %v\n", err)
			}
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

func (mh *meetingHandler) GetMeeting() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := common.ExtractData(c)

		data, err := mh.meetingUsecase.GetOwnerMeeting(token.ID)
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

// GetMyMeeting implements domain.MeetingHandler
func (mh *meetingHandler) GetMyMeeting() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := common.ExtractData(c)

		data, err := mh.meetingUsecase.GetSeekMeeting(token.ID)
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

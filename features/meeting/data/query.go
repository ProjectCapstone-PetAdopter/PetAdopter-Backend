package data

import (
	"errors"
	"fmt"
	"log"
	"petadopter/domain"

	"gorm.io/gorm"
)

type meetingData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.MeetingData {
	return &meetingData{
		db: DB,
	}
}

// GetMyMeetingID implements domain.MeetingData
func (md *meetingData) GetMyMeetingID(id int) []domain.MeetingOwner {
	var data []MeetingOwner
	var seeker []Seekerdata

	err := md.db.Model(&Meeting{}).Select("users.fullname, adoptions.user_id").
		Joins("join adoptions on meetings.adoption_id = adoptions.id").Joins("join users on adoptions.user_id = users.id").
		Where("adoptions.user_id = ?", id).Scan(&seeker)
	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	err = md.db.Model(&Meeting{}).Select("meetings.id, meetings.adoption_id, meetings.user_id, meetings.time, meetings.date, pets.petname, pets.petphoto, users.fullname, users.photo_profile, users.address").
		Joins("join adoptions on meetings.adoption_id = adoptions.id").Joins("join pets on adoptions.pets_id = pets.id").Joins("join users on pets.userid = users.id").
		Where("adoptions.user_id = ?", id).Scan(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	for i := 0; i < len(seeker); i++ {
		data[i].Seekername = seeker[i].Fullname
		data[i].Seekerid = seeker[i].UserID
	}

	return ParseToArrMeeting(data)
}

// GetEmailData implements domain.MeetingData
func (md *meetingData) GetEmailData(userID, meetingID int) (domain.Ownerdata, domain.Seekerdata, int) {
	var owner Ownerdata
	var seeker Seekerdata

	getOwner := md.db.Model(&Meeting{}).Select("users.email, users.address, users.city").Joins("join users on meetings.user_id = users.id").
		Where("meetings.user_id = ?", userID).Limit(1).Scan(&owner)

	if getOwner.Error != nil {
		log.Println("query error", getOwner.Error)
		return domain.Ownerdata{}, domain.Seekerdata{}, 500
	}

	if getOwner.RowsAffected == 0 {
		log.Println("data not found in db")
		return domain.Ownerdata{}, domain.Seekerdata{}, 404
	}

	getSeeker := md.db.Model(&Meeting{}).Select("users.email").Joins("join adoptions on meetings.adoption_id = adoptions.id").
		Joins("join users on adoptions.user_id = users.id").Where("meetings.id = ?", meetingID).Scan(&seeker)

	if getSeeker.Error != nil {
		log.Println("query error", getSeeker.Error)
		return domain.Ownerdata{}, domain.Seekerdata{}, 500
	}

	if getSeeker.RowsAffected == 0 {
		log.Println("data not found in db")
		return domain.Ownerdata{}, domain.Seekerdata{}, 404
	}

	return owner.ToModelOwnerdata(), seeker.ToModelSeekerdata(), 200
}

func (md *meetingData) Insert(data domain.Meeting) (idMeet int, err error) {

	meetingData := FromModel(data)
	var seekerID int

	getOwnerID := md.db.Table("adoptions").Select("adoptions.user_id").Where("id = ?", data.AdoptionID).Scan(&seekerID)

	if getOwnerID.Error != nil {
		log.Println("Cannot get adopt id", getOwnerID.Error.Error())
		return 0, errors.New("cannot insert meeting")
	}

	if data.UserID == seekerID {
		log.Println("error db")
		return -1, errors.New("only owner can add meeting")
	}

	result := md.db.Create(&meetingData)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("cannot insert meeting")
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}

	getMeetingId := md.db.Last(&Meeting{}, "meetings.user_id = ?", data.UserID).Select("meetings.id").Scan(&idMeet)

	if getMeetingId.Error != nil {
		log.Println("cant get meeting id", getMeetingId.Error)
		return -1, errors.New("meeting id query error")
	}

	if result.RowsAffected == 0 {
		log.Println("meeting id not foun", getMeetingId.Error)
		return 0, errors.New("failed get meeting id")
	}

	return idMeet, nil
}

func (md *meetingData) Update(updatedData domain.Meeting, id int) (idMeet int, err error) {
	var cnv = FromModel(updatedData)
	var seekerid int

	getownerid := md.db.Table("adoptions").Select("adoptions.user_id").Where("id = ?", updatedData.AdoptionID).Scan(&seekerid)

	if getownerid.Error != nil {
		log.Println("Cannot get adopt id", getownerid.Error.Error())
		return 0, errors.New("cannot update meeting")
	}

	if updatedData.UserID == seekerid {
		log.Println("error db")
		return -1, errors.New("only owner can update meeting")
	}

	cnv.ID = uint(id)
	fmt.Println(cnv.ID)
	result := md.db.Model(&Meeting{}).Where("ID = ?", id).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("cannot update meeting")
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (md *meetingData) Delete(id int) error {

	res := md.db.Delete(&Meeting{}, id)
	if res.Error != nil {
		log.Println("cannot delete data", res.Error.Error())
		return res.Error
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted", res.Error.Error())
		return fmt.Errorf("failed to delete species")
	}
	return nil
}

func (md *meetingData) GetMeetingID(id int) []domain.MeetingOwner {
	var data []MeetingOwner
	var seeker []Seekerdata
	var getData domain.Meeting
	var getMeeting []domain.MeetingOwner
	var seekerid int

	getownerid := md.db.Table("adoptions").Select("adoptions.user_id").Where("id = ?", getData.AdoptionID).Scan(&seekerid)

	if getownerid.Error != nil {
		log.Println("Cannot get adopt id", getownerid.Error.Error())
		return getMeeting
	}

	var userMeeting int
	if userMeeting != getData.UserID || userMeeting != seekerid {
		log.Println("error db")
		return getMeeting
	}

	err := md.db.Model(&Meeting{}).Select("users.fullname, adoptions.user_id").
		Joins("join adoptions on meetings.adoption_id = adoptions.id").Joins("join users on adoptions.user_id = users.id").
		Where("meetings.user_id = ?", id).Scan(&seeker)
	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	err = md.db.Model(&Meeting{}).Select("meetings.id, meetings.adoption_id, meetings.user_id, meetings.time, meetings.date, pets.petname, pets.petphoto, users.fullname, users.photo_profile, users.address").
		Joins("join adoptions on meetings.adoption_id = adoptions.id").Joins("join pets on adoptions.pets_id = pets.id").Joins("join users on pets.userid = users.id").
		Where("meetings.user_id = ?", id).Scan(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	for i := 0; i < len(seeker); i++ {
		data[i].Seekername = seeker[i].Fullname
		data[i].Seekerid = seeker[i].UserID
	}

	return ParseToArrMeeting(data)
}

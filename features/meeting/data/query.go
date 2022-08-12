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

func (md *meetingData) Insert(data domain.Meeting) (row int, err error) {

	meetingData := FromModel(data)
	var seekerid int

	// getownerid := md.db.Model(&Meeting{}).Select("adoptions.user_id").Joins("join adoptions on meetings.adoption_id = adoptions.id").
	// Where("meetings.adoption_id = ?", data.AdoptionID).Scan(&seekerid)
	getownerid := md.db.Table("adoptions").Select("adoptions.user_id").Where("id = ?", data.AdoptionID).Scan(&seekerid)

	if getownerid.Error != nil {
		log.Println("Cannot get adopt id", getownerid.Error.Error())
		return 0, errors.New("cannot insert meeting")
	}
	//memverifikasi user yang mengakses
	fmt.Println(seekerid)
	fmt.Println(data.UserID)
	if data.UserID == seekerid {
		log.Println("error db")
		return -1, errors.New("Cant adopt owned pet")
	}

	result := md.db.Create(&meetingData)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("cannot insert meeting")
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}

	return int(result.RowsAffected), nil
}

func (md *meetingData) Update(updatedData domain.Meeting, id int) error {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(id)
	err := md.db.Model(&Meeting{}).Where("ID = ?", id).Updates(cnv).Error
	if err != nil {
		return err
	}
	return nil
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

func (md *meetingData) GetMeetingID(meetingID int) []domain.MeetingOwner {
	var data []MeetingOwner

	err := md.db.Model(&Meeting{}).Select("meetings.id, pets.petname, pets.petphoto, users.fullname, users.photo_profile, users.address, adoptions.status").
		Joins("join pets on adoptions.pets_id = pets.id").Joins("join users on pets.userid = users.id").Where("meetings.id = ?", meetingID).Scan(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}

	return ParseToArrMeeting(data)
}

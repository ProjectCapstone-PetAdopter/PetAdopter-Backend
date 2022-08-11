package data

import (
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

func (md *meetingData) Insert(data domain.Meeting) error {
	meetingData := FromModel(data)
	err := md.db.Create(&meetingData).Error
	if err != nil {
		return err
	}
	return nil
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

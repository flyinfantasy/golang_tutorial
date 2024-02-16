package ptstatus

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Ptstatus struct {
	gorm.Model
	Pigeon_club         string `json:"pigeon_club" db:"pigeon_club" gorm:"varchar(200);not null"`
	Pigeon_loft         string `json:"pigeon_loft" db:"pigeon_loft" gorm:"varchar(200);not null"`
	Longitude           string `json:"longitude" db:"longitude" gorm:"varchar(200);not null"`
	Latitude            string `json:"latitude" db:"latitude" gorm:"varchar(200);not null"`
	Electricity         string `json:"electricity" db:"electricity" gorm:"varchar(200);not null;"`
	Plugged_in          string `json:"plugged_in" db:"plugged_in" gorm:"varchar(200);not null;"`
	Out_of_range        string `json:"out_of_range" db:"out_of_range" gorm:"varchar(200);not null;"`
	UID_Table           string `json:"uID_Table" db:"uID_Table" gorm:"varchar(200);not null;unique;primary_key;"`
	Status_code         string `json:"status_code" db:"status_code" gorm:"varchar(200);not null"`
	Signal_strength     string `json:"signal_strength" db:"signal_strength" gorm:"varchar(200);not null"`
	Server_receive_time string `json:"server_receive_time" db:"server_receive_time" gorm:"varchar(200);not null"`
	Server_receive_date string `json:"server_receive_date" db:"server_receive_date" gorm:"varchar(200);not null"`
	Pigeon_loft_tel     string `json:"pigeon_loft_tel" db:"pigeon_loft_tel" gorm:"varchar(200)"`
}

func (Ptstatus) TableName() string {
	return "ptstatuss"
}

func (u *Ptstatus) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}

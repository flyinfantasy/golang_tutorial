package pigeon

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Pigeon struct {
	gorm.Model
	Uid                    string `json:"uid" db:"uid" gorm:"varchar(200);not null;unique;primary_key;"`
	IMEI                   string `json:"imei" db:"imei" gorm:"varchar(200)"`
	Pigeon_club            string `json:"pigeon_club" db:"pigeon_club" gorm:"varchar(200)"`
	Pigeon_loft            string `json:"pigeon_loft" db:"pigeon_loft" gorm:"varchar(200)"`
	Pigeon_number          string `json:"pigeon_number" db:"pigeon_number" gorm:"varchar(200)"`
	Pigeon_aluminum        string `json:"pigeon_aluminum" db:"pigeon_aluminum" gorm:"varchar(200)"`
	Aluminum_UID           string `json:"aluminum_UID" db:"aluminum_UID" gorm:"varchar(200)"`
	Longitude              string `json:"longitude" db:"longitude" gorm:"varchar(200)"`
	Latitude               string `json:"latitude" db:"latitude" gorm:"varchar(200)"`
	Induction_date         string `json:"induction_date" db:"induction_date" gorm:"varchar(200)"`
	Induction_time         string `json:"induction_time" db:"induction_time" gorm:"varchar(200)"`
	Server_receive_date    string `json:"server_receive_date" db:"server_receive_date" gorm:"varchar(200)"`
	Server_receive_time    string `json:"server_receive_time" db:"server_receive_time" gorm:"varchar(200)"`
	Signal_strength        string `json:"signal_strength" db:"signal_strength" gorm:"varchar(200)"`
	Electricity            string `json:"electricity" db:"electricity" gorm:"varchar(200)"`
	Plugged_in             string `json:"plugged_in" db:"plugged_in" gorm:"varchar(200)"`
	UDP_receive_time       string `json:"UDP_receive_time" db:"UDP_receive_time" gorm:"varchar(200)"`
	UDP_receive_date       string `json:"UDP_receive_date" db:"UDP_receive_date" gorm:"varchar(200)"`
	Status_code            string `json:"status_code" db:"status_code" gorm:"varchar(200)"`
	Get_web_receive_date   string `json:"get_web_receive_date" db:"get_web_receive_date" gorm:"varchar(200)"`
	Get_web_receive_time   string `json:"get_web_receive_time" db:"get_web_receive_time" gorm:"varchar(200)"`
	Get_web_longitude      string `json:"get_web_longitude" db:"get_web_longitude" gorm:"varchar(200)"`
	Get_web_latitude       string `json:"get_web_latitude" db:"get_web_latitude" gorm:"varchar(200)"`
	Get_web_number         string `json:"get_web_number" db:"get_web_number" gorm:"varchar(200)"`
	Get_web_Induction_date string `json:"get_web_Induction_date" db:"get_web_Induction_date" gorm:"varchar(200)"`
	Get_web_Induction_time string `json:"get_web_Induction_time" db:"get_web_Induction_time" gorm:"varchar(200)"`
	Get_web_check          string `json:"get_web_check" db:"get_web_check" gorm:"varchar(200);default:'0'"`
	Get_pt_check           string `json:"get_pt_check" db:"get_pt_check" gorm:"varchar(200);default:'0'"`
}

func (Pigeon) TableName() string {
	return "pigeons"
}

func (u *Pigeon) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}

package login

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Login struct {
	gorm.Model
	Pigeon_tel  string `json:"pigeon_tel" db:"pigeon_tel" gorm:"varchar(200);unique;primary_key;not null"`
	Pigeon_loft string `json:"pigeon_loft" db:"pigeon_loft" gorm:"varchar(200);not null"`
	Pigeon_club string `json:"pigeon_club" db:"pigeon_club" gorm:"varchar(200);not null"`
}

type Loginp struct {
	gorm.Model
	Pigeon_club    string `json:"pigeon_club" db:"pigeon_club"`
	Pigeon_account string `json:"pigeon_account" db:"pigeon_account"`
}

func (Login) TableName() string {
	return "logins"
}

func (u *Login) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}

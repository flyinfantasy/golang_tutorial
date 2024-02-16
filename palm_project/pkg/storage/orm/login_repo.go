package orm

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/login"
)

type loginRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewLoginRepo(db *gorm.DB, log logger.LogInfoFormat) login.Repository {
	return &loginRepo{db, log}
}

func (l *loginRepo) Signin(pigeon_tel string, pigeon_loft string) string {
	l.log.Debugf("Signin by pigeon_tel : %s", pigeon_tel)
	login := &login.Login{}
	err := l.db.Where("pigeon_tel = ?", pigeon_tel).First(&login).Error
	if err != nil {
		l.log.Errorf("pigeon not found with pigeon_tel : %s, reason : %v", pigeon_tel, err)
	}
	if login.Pigeon_loft == pigeon_loft {
		return login.Pigeon_club
	} else {
		return "error"
	}
}

func (l *loginRepo) Delete(Pigeon_account string) error {
	l.log.Debugf("deleting the pigeon with Pigeon_account : %s", Pigeon_account)

	if l.db.Delete(&login.Login{}, "Pigeon_account = ?", Pigeon_account).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the pigeon with Pigeon_account : %s", Pigeon_account)
		l.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (l *loginRepo) GetAll() ([]*login.Login, error) {
	l.log.Debug("get all the acc")

	login := make([]*login.Login, 0)
	err := l.db.Select([]string{"pigeon_account", "pigeon_club"}).Find(&login).Error
	if err != nil {
		l.log.Debug("no single login found")
		return nil, err
	}
	return login, nil
}

func (l *loginRepo) Update(logi *login.Login) error {
	l.log.Debugf("updating the Pigeon_account : %v", logi.Pigeon_tel)
	fmt.Printf("updating the Pigeon_tel : %v", logi.Pigeon_tel)
	err := l.db.Model(&logi).Where("pigeon_account = ?", logi.Pigeon_tel).Updates(login.Login{
		Pigeon_club: logi.Pigeon_club,
		Pigeon_loft: logi.Pigeon_loft}).Error
	if err != nil {
		l.log.Errorf("error while updating the Pigeon_account, reason : %v", err)
		return err
	}
	return nil
}

func (l *loginRepo) Store(logi *login.Login) error {
	l.log.Debugf("creating the login with Pigeon_account : %v", logi.Pigeon_tel)
	err := l.db.Create(&logi).Error
	if err != nil {
		l.log.Errorf("error while creating the Pigeon_account, reason : %v", err)
		return err
	}
	return nil
}

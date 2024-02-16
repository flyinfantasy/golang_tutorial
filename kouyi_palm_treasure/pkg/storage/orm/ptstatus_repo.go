package orm

import (
	"errors"
	"fmt"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/ptstatus"

	"github.com/jinzhu/gorm"
)

type ptstatusRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewPtstatusRepo(db *gorm.DB, log logger.LogInfoFormat) ptstatus.Repository {
	return &ptstatusRepo{db, log}
}

func (p *ptstatusRepo) Delete(pigeon_club_pigeon_loft string) error {
	p.log.Debugf("deleting the ptstatus with Pigeon_club_pigeon_loft : %s", pigeon_club_pigeon_loft)

	if p.db.Delete(&ptstatus.Ptstatus{}, "Pigeon_club_pigeon_loft = ?", pigeon_club_pigeon_loft).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the ptstatus with id : %s", pigeon_club_pigeon_loft)
		p.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (p *ptstatusRepo) GetAll() ([]*ptstatus.Ptstatus, error) {
	p.log.Debug("get all the ptstatuss")

	ptstatuss := make([]*ptstatus.Ptstatus, 0)
	err := p.db.Find(&ptstatuss).Error
	if err != nil {
		p.log.Debug("no single ptstatus found")
		return nil, err
	}
	return ptstatuss, nil
}

func (p *ptstatusRepo) GetByID(uid string) (*ptstatus.Ptstatus, error) {
	p.log.Debugf("get user details by Uid : %s", uid)

	ptstatus := &ptstatus.Ptstatus{}
	err := p.db.Where("Uid = ?", uid).First(&ptstatus).Error
	if err != nil {
		p.log.Errorf("user not found with Uid : %s, reason : %v", uid, err)
		return nil, err
	}
	return ptstatus, nil
}

func (p *ptstatusRepo) Store(ptstatu *ptstatus.Ptstatus) error {
	ptstatuss := &ptstatus.Ptstatus{}
	results := p.db.Where("UID_Table = ?", ptstatu.UID_Table).First(&ptstatuss)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			p.log.Debugf("creating the ptstatu with Uid number : %v", ptstatu.UID_Table)
			err := p.db.Create(&ptstatu).Error
			if err != nil {
				p.log.Errorf("error while creating the ptstatu, reason : %v", err)
				return err
			}
		}
	} else {
		p.Update(ptstatu)
	}
	return nil
}

func (p *ptstatusRepo) GetByDate(palm_treasure_date string) ([]*ptstatus.Ptstatus, error) {
	p.log.Debugf("get ptstatus details by Palm_treasure_date : %s", palm_treasure_date)

	ptstatu := make([]*ptstatus.Ptstatus, 0)
	err := p.db.Where("Palm_treasure_date = ?", palm_treasure_date).Find(&ptstatu).Error
	if err != nil {
		p.log.Debug("no single Palm_treasure_date found")
		return nil, err
	}
	return ptstatu, nil
}

func (p *ptstatusRepo) Update(ptstatu *ptstatus.Ptstatus) error {
	p.log.Debugf("updating the user, Uid number : %v", ptstatu.UID_Table)

	err := p.db.Model(&ptstatu).Where("UID_Table = ?", ptstatu.UID_Table).Updates(ptstatus.Ptstatus{
		Pigeon_club: ptstatu.Pigeon_club, Pigeon_loft: ptstatu.Pigeon_loft, Longitude: ptstatu.Longitude,
		Latitude: ptstatu.Latitude, Server_receive_date: ptstatu.Server_receive_date,
		Server_receive_time: ptstatu.Server_receive_time, Signal_strength: ptstatu.Signal_strength,
		Electricity: ptstatu.Electricity, Plugged_in: ptstatu.Plugged_in,
		Status_code: ptstatu.Status_code, Pigeon_loft_tel: ptstatu.Pigeon_loft_tel, Out_of_range: ptstatu.Out_of_range}).Error
	if err != nil {
		p.log.Errorf("error while updating the user, reason : %v", err)
		return err
	}
	return nil
}

func (p *ptstatusRepo) GetByClub(pigeon_club string) ([]*ptstatus.Ptstatus, error) {
	p.log.Debugf("get ptstatus details by pigeon_club : %s", pigeon_club)
	//fmt.Println(pigeon_club)
	ptstatus := make([]*ptstatus.Ptstatus, 0)
	err := p.db.Where("pigeon_club = ?", pigeon_club).Find(&ptstatus).Error
	if err != nil {
		p.log.Debug("no single pigeon_club found")
		return nil, err
	}
	return ptstatus, nil
}

func (p *ptstatusRepo) GetByClubandDate(pigeon_club string, date string) ([]*ptstatus.Ptstatus, error) {
	p.log.Debugf("get ptstatus details by pigeon_club & server_receive_date : %s & %s", pigeon_club, date)

	ptstatus := make([]*ptstatus.Ptstatus, 0)
	err := p.db.Where("pigeon_club = ? AND server_receive_date = ? ", pigeon_club, date).Find(&ptstatus).Error
	//err := p.db.Where("induction_date = ? ", date).Or("get_web_Induction_date = ? ", date).Find(&pigeon, "pigeon_club = ?", pigeon_club).Error
	if err != nil {
		p.log.Debug("no single pigeon_club & server_receive_date found")
		return nil, err
	}
	return ptstatus, nil
}

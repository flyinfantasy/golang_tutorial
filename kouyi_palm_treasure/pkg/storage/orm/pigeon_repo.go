package orm

import (
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/pigeon"

	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type pigeonRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewPigeonRepo(db *gorm.DB, log logger.LogInfoFormat) pigeon.Repository {
	return &pigeonRepo{db, log}
}

func (p *pigeonRepo) Delete(Uid string) error {
	p.log.Debugf("deleting the pigeon with Uid : %s", Uid)

	if p.db.Delete(&pigeon.Pigeon{}, "Uid = ?", Uid).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the pigeon with Uid : %s", Uid)
		p.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (p *pigeonRepo) GetAll() ([]*pigeon.Pigeon, error) {
	p.log.Debug("get all the users")

	pigeon := make([]*pigeon.Pigeon, 0)
	err := p.db.Find(&pigeon).Error
	if err != nil {
		p.log.Debug("no single pigeon found")
		return nil, err
	}
	return pigeon, nil
}

func (p *pigeonRepo) GetByID(Uid string) (*pigeon.Pigeon, error) {
	p.log.Debugf("get pigeon details by Uid : %s", Uid)

	pigeon := &pigeon.Pigeon{}
	err := p.db.Where("Uid = ?", Uid).First(&pigeon).Error
	if err != nil {
		p.log.Errorf("pigeon not found with Uid : %s, reason : %v", Uid, err)
		return nil, err
	}
	return pigeon, nil
}

func (p *pigeonRepo) GetByDate(induction_date string) ([]*pigeon.Pigeon, error) {
	p.log.Debugf("get pigeon details by induction_date : %s", induction_date)

	pigeon := make([]*pigeon.Pigeon, 0)
	err := p.db.Where("induction_date = ?", induction_date).Find(&pigeon).Error
	if err != nil {
		p.log.Debug("no single induction_date found")
		return nil, err
	}
	return pigeon, nil
}

// TODO 雙向寫入
func (p *pigeonRepo) WebStore(pigeo *pigeon.Pigeon) error {
	pigeons := &pigeon.Pigeon{}
	results := p.db.Where("uid = ?", pigeo.Uid).First(&pigeons)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			p.log.Debugf("creating the pigeo with Pigeon number : %v", pigeo.Uid)
			err := p.db.Create(&pigeo).Error
			if err != nil {
				p.log.Errorf("error while creating the pigeo, reason : %v", err)
				return err
			}
		}
	} else {
		resultss := p.db.Where("uid = ?", pigeo.Uid).Where("Get_pt_check =?", pigeo.Get_pt_check).First(&pigeons)
		if resultss.Error != nil {
			if resultss.Error == gorm.ErrRecordNotFound {
				p.log.Debugf("WebUpdate the pigeo with Pigeon number : %v", pigeo.Uid)
				p.WebUpdate(pigeo)
			}
		}
	}
	return nil
}

func (p *pigeonRepo) WebUpdate(pigeo *pigeon.Pigeon) error {
	p.log.Debugf("WebUpdating the pigeon, Uid : %v", pigeo.Uid)
	err := p.db.Model(&pigeo).Where("uid = ?", pigeo.Uid).Updates(
		pigeon.Pigeon{
			Pigeon_loft:            pigeo.Pigeon_loft,
			Pigeon_club:            pigeo.Pigeon_club,
			Pigeon_number:          pigeo.Pigeon_number,
			Get_web_receive_date:   pigeo.Get_web_receive_date,
			Get_web_receive_time:   pigeo.Get_web_receive_time,
			Get_web_longitude:      pigeo.Get_web_longitude,
			Get_web_latitude:       pigeo.Get_web_latitude,
			Get_web_number:         pigeo.Get_web_number,
			Get_web_Induction_date: pigeo.Get_web_Induction_date,
			Get_web_Induction_time: pigeo.Get_web_Induction_time,
			Get_web_check:          pigeo.Get_web_check}).Error
	if err != nil {
		p.log.Errorf("error while updating the pigeon, reason : %v", err)
		return err
	}
	return nil
}

func (p *pigeonRepo) PtStore(pigeo *pigeon.Pigeon) error {
	pigeons := &pigeon.Pigeon{}
	results := p.db.Where("uid = ?", pigeo.Uid).First(&pigeons)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			p.log.Debugf("creating the pigeo with Pigeon number : %v", pigeo.Uid)
			err := p.db.Create(&pigeo).Error
			if err != nil {
				p.log.Errorf("error while creating the pigeo, reason : %v", err)
				return err
			}
		}
	} else {
		resultss := p.db.Where("uid = ?", pigeo.Uid).Where("Get_pt_check =?", pigeo.Get_web_check).First(&pigeons)
		if resultss.Error != nil {
			if resultss.Error == gorm.ErrRecordNotFound {
				p.log.Debugf("PtUpdate the pigeo with Pigeon number : %v", pigeo.Uid)
				fmt.Printf("PtUpdate the pigeo with Pigeon number : %v", pigeo.Uid)
				p.PtUpdate(pigeo)
			}
		}
	}
	return nil
}

/*
	func (p *ptstatusRepo) Update(ptstatu *ptstatus.Ptstatus) error {
		p.log.Debugf("updating the user, Pigeon_club_pigeon_loft number : %v", ptstatu.Pigeon_club_pigeon_loft)

		err := p.db.Model(&ptstatu).Where("Pigeon_club_pigeon_loft = ?", ptstatu.Pigeon_club_pigeon_loft).Updates(ptstatus.Ptstatus{
			Pigeon_club: ptstatu.Pigeon_club, Pigeon_loft: ptstatu.Pigeon_loft, Palm_treasure_date: ptstatu.Palm_treasure_date,
			Palm_treasure_time: ptstatu.Palm_treasure_time, Longitude: ptstatu.Longitude,
			Latitude: ptstatu.Latitude, Server_receive_date: ptstatu.Server_receive_date,
			Server_receive_time: ptstatu.Server_receive_time, Signal_strength: ptstatu.Signal_strength,
			Electricity: ptstatu.Electricity, Plugged_in: ptstatu.Plugged_in,
			Status_code: ptstatu.Status_code, IMEI: ptstatu.IMEI}).Error
		if err != nil {
			p.log.Errorf("error while updating the user, reason : %v", err)
			return err
		}
		return nil
	}
*/
func (p *pigeonRepo) PtUpdate(pigeo *pigeon.Pigeon) error {
	p.log.Debugf("Ptupdating the pigeon, uid : %v", pigeo.Uid)
	fmt.Printf("Ptupdating the pigeon, uid : %v", pigeo.Uid)
	err := p.db.Model(&pigeo).Where("uid = ?", pigeo.Uid).Updates(pigeon.Pigeon{
		Pigeon_club:         pigeo.Pigeon_club,
		IMEI:                pigeo.IMEI,
		Pigeon_loft:         pigeo.Pigeon_loft,
		Pigeon_aluminum:     pigeo.Pigeon_aluminum,
		Pigeon_number:       pigeo.Pigeon_number,
		Aluminum_UID:        pigeo.Aluminum_UID,
		Longitude:           pigeo.Longitude,
		Latitude:            pigeo.Latitude,
		Induction_date:      pigeo.Induction_date,
		Induction_time:      pigeo.Induction_time,
		Server_receive_date: pigeo.Server_receive_date,
		Server_receive_time: pigeo.Server_receive_time,
		UDP_receive_time:    pigeo.UDP_receive_time,
		UDP_receive_date:    pigeo.UDP_receive_date,
		Signal_strength:     pigeo.Signal_strength,
		Electricity:         pigeo.Electricity,
		Plugged_in:          pigeo.Plugged_in,
		Status_code:         pigeo.Status_code}).Error
	if err != nil {
		p.log.Errorf("error while Ptupdating the pigeon, reason : %v", err)
		return err
	}
	return nil
}

func (p *pigeonRepo) Store(pigeo *pigeon.Pigeon) error {
	p.log.Debugf("creating the pigeo with Pigeon number : %v", pigeo.Uid)

	err := p.db.Create(&pigeo).Error
	if err != nil {
		p.log.Errorf("error while creating the pigeo, reason : %v", err)
		return err
	}
	return nil
}

func (p *pigeonRepo) Update(pigeo *pigeon.Pigeon) error {
	p.log.Debugf("updating the pigeon, Uid : %v", pigeo.Uid)

	err := p.db.Model(&pigeo).Updates(pigeon.Pigeon{Pigeon_club: pigeo.Pigeon_club,
		Pigeon_loft: pigeo.Pigeon_loft, Pigeon_aluminum: pigeo.Pigeon_aluminum, Aluminum_UID: pigeo.Aluminum_UID,
		Longitude: pigeo.Longitude, Latitude: pigeo.Latitude, Induction_date: pigeo.Induction_date,
		Induction_time: pigeo.Induction_time, Server_receive_date: pigeo.Server_receive_date,
		Server_receive_time: pigeo.Server_receive_time, UDP_receive_time: pigeo.UDP_receive_time, UDP_receive_date: pigeo.UDP_receive_date,
		Signal_strength: pigeo.Signal_strength, Electricity: pigeo.Electricity, Plugged_in: pigeo.Plugged_in,
		Status_code: pigeo.Status_code, Uid: pigeo.Uid}).Error
	if err != nil {
		p.log.Errorf("error while updating the pigeon, reason : %v", err)
		return err
	}
	return nil
}

//GetBYCLUB

/*

func (p *pigeonRepo) GetByDate(induction_date string) ([]*pigeon.Pigeon, error) {
	p.log.Debugf("get pigeon details by induction_date : %s", induction_date)

	pigeon := make([]*pigeon.Pigeon, 0)
	err := p.db.Where("induction_date = ?", induction_date).Find(&pigeon).Error
	if err != nil {
		p.log.Debug("no single induction_date found")
		return nil, err
	}
	return pigeon, nil
}
*/

func (p *pigeonRepo) GetByClub(pigeon_club string) ([]*pigeon.Pigeon, error) {
	p.log.Debugf("get pigeon details by pigeon_club : %s", pigeon_club)
	//fmt.Println(pigeon_club)
	pigeon := make([]*pigeon.Pigeon, 0)
	err := p.db.Where("pigeon_club = ?", pigeon_club).Find(&pigeon).Error
	if err != nil {
		p.log.Debug("no single pigeon_club found")
		return nil, err
	}
	return pigeon, nil
}

func (p *pigeonRepo) GetByClubandDate(pigeon_club string, date string) ([]*pigeon.Pigeon, error) {
	p.log.Debugf("get pigeon details by pigeon_club & get_web_Induction_date : %s & %s", pigeon_club, date)

	pigeon := make([]*pigeon.Pigeon, 0)
	err := p.db.Where("pigeon_club = ? AND get_web_Induction_date = ? ", pigeon_club, date).Or("pigeon_club = ? AND induction_date = ? ", pigeon_club, date).Find(&pigeon).Error
	//err := p.db.Where("induction_date = ? ", date).Or("get_web_Induction_date = ? ", date).Find(&pigeon, "pigeon_club = ?", pigeon_club).Error
	if err != nil {
		p.log.Debug("no single pigeon_club & get_web_Induction_date found")
		return nil, err
	}
	return pigeon, nil
}

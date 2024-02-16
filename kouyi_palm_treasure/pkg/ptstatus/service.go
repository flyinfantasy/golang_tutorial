package ptstatus

type Service interface {
	Delete(id string) error
	GetAll() ([]*Ptstatus, error)
	GetByID(id string) (*Ptstatus, error)
	GetByDate(palm_treasure_date string) ([]*Ptstatus, error)
	Store(u *Ptstatus) error
	Update(u *Ptstatus) error
	GetByClub(pigeon_club string) ([]*Ptstatus, error)
	GetByClubandDate(pigeon_club string, date string) ([]*Ptstatus, error)
}

type ptstatusService struct {
	repo Repository
}

func NewpigeonService(repo Repository) Service {
	return &ptstatusService{
		repo: repo,
	}
}

func (svc *ptstatusService) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *ptstatusService) GetAll() ([]*Ptstatus, error) { return svc.repo.GetAll() }

func (svc *ptstatusService) GetByID(id string) (*Ptstatus, error) { return svc.repo.GetByID(id) }

func (svc *ptstatusService) GetByDate(palm_treasure_date string) ([]*Ptstatus, error) {
	return svc.repo.GetByDate(palm_treasure_date)
}

func (svc *ptstatusService) Store(u *Ptstatus) error { return svc.repo.Store(u) }

func (svc *ptstatusService) Update(u *Ptstatus) error { return svc.repo.Update(u) }

func (svc *ptstatusService) GetByClub(pigeon_club string) ([]*Ptstatus, error) {
	return svc.repo.GetByClub(pigeon_club)
}

func (svc *ptstatusService) GetByClubandDate(pigeon_club string, date string) ([]*Ptstatus, error) {
	return svc.repo.GetByClubandDate(pigeon_club, date)
}

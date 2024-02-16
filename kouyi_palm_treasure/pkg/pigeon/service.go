package pigeon

type Service interface {
	Delete(id string) error
	GetAll() ([]*Pigeon, error)
	GetByID(id string) (*Pigeon, error)
	GetByDate(induction_date string) ([]*Pigeon, error)
	Store(u *Pigeon) error
	Update(u *Pigeon) error
	WebStore(u *Pigeon) error
	WebUpdate(u *Pigeon) error
	PtStore(u *Pigeon) error
	PtUpdate(u *Pigeon) error
	GetByClub(pigeon_club string) ([]*Pigeon, error)
	GetByClubandDate(pigeon_club string, date string) ([]*Pigeon, error)
}

type pigeonService struct {
	repo Repository
}

func NewpigeonService(repo Repository) Service {
	return &pigeonService{
		repo: repo,
	}
}

func (svc *pigeonService) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *pigeonService) GetAll() ([]*Pigeon, error) { return svc.repo.GetAll() }

func (svc *pigeonService) GetByID(id string) (*Pigeon, error) { return svc.repo.GetByID(id) }

func (svc *pigeonService) GetByDate(induction_date string) ([]*Pigeon, error) {
	return svc.repo.GetByDate(induction_date)
}

func (svc *pigeonService) Store(u *Pigeon) error { return svc.repo.Store(u) }

func (svc *pigeonService) Update(u *Pigeon) error { return svc.repo.Update(u) }

func (svc *pigeonService) WebStore(u *Pigeon) error { return svc.repo.WebStore(u) }

func (svc *pigeonService) WebUpdate(u *Pigeon) error { return svc.repo.WebUpdate(u) }

func (svc *pigeonService) PtStore(u *Pigeon) error { return svc.repo.PtStore(u) }

func (svc *pigeonService) PtUpdate(u *Pigeon) error { return svc.repo.PtUpdate(u) }

func (svc *pigeonService) GetByClub(pigeon_club string) ([]*Pigeon, error) {
	return svc.repo.GetByClub(pigeon_club)
}

func (svc *pigeonService) GetByClubandDate(pigeon_club string, date string) ([]*Pigeon, error) {
	return svc.repo.GetByClubandDate(pigeon_club, date)
}

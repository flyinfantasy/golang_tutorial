package pigeon

type Repository interface {
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

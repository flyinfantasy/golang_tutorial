package ptstatus

type Repository interface {
	Delete(id string) error
	GetAll() ([]*Ptstatus, error)
	GetByID(id string) (*Ptstatus, error)
	GetByDate(palm_treasure_date string) ([]*Ptstatus, error)
	Store(u *Ptstatus) error
	Update(u *Ptstatus) error
	GetByClub(pigeon_club string) ([]*Ptstatus, error)
	GetByClubandDate(pigeon_club string, date string) ([]*Ptstatus, error)
}

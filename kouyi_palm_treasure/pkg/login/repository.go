package login

type Repository interface {
	Signin(pigeon_account string, pigeon_password string) string
	Store(u *Login) error
	Delete(pigeon_account string) error
	Update(u *Login) error
	GetAll() ([]*Login, error)
}

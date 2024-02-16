package login

type Service interface {
	Signin(pigeon_account string, pigeon_password string) string
	Store(u *Login) error
	Delete(pigeon_account string) error
	Update(u *Login) error
	GetAll() ([]*Login, error)
}

type loginService struct {
	repo Repository
}

func NewloginService(repo Repository) Service {
	return &loginService{
		repo: repo,
	}
}

func (svc *loginService) Signin(pigeon_account string, pigeon_password string) string {
	return svc.repo.Signin(pigeon_account, pigeon_password)
}

func (svc *loginService) Store(u *Login) error {
	return svc.repo.Store(u)
}

func (svc *loginService) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *loginService) Update(u *Login) error { return svc.repo.Update(u) }

func (svc *loginService) GetAll() ([]*Login, error) { return svc.repo.GetAll() }

package repo

import "bitbucket.org/service-ekspedisi/models"

type AboutUsRepoInterface interface {
	AddAbout(v models.AboutUsDb) (models.AboutUsDb, error)
	GetAboutUs() (models.AboutUsDb, error)
	GetById(id int) (models.AboutUsDb, error)
	UpdateData(id int, v models.AboutUsRequest) (models.AboutUsRequest, error)
	DeleteData(id []string) error
}

type ExpeditionRepoInterface interface {
	AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	GetById(id int) (models.ExpeditionSchedule, error)
	GetAll(imit, offset int) (models.ResponseDataPagination, error)
	Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	DeleteData(id []string) error
	GetByRoute(rf, rd string) ([]models.ExpeditionSchedule, error)
	GetKotaById(id int) (models.KotaKab, error)
}

type UserRepoInterface interface {
	AddUser(v models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	UpdateData(id int, v models.User) (models.User, error)
	DeleteData(id []string) error
}

type LoginRepoInterface interface {
	LoginUser(email string) (models.User, error)
	CreateAuth(authFix models.Auth) (models.Auth, error)
	DeleteAuthData(givenUuid string) (int, error)
	GetAuthByEmailAndAuthID(email string, authUUID string) (*models.Auth, error)
}

type BlogRepoInterface interface {
	AddBlog(v models.Blog) (models.Blog, error)
	GetAll() ([]models.Blog, error)
	GetById(id int) (models.Blog, error)
	UpdateData(id int, v models.Blog) (models.Blog, error)
	DeleteData(id []string) error
}

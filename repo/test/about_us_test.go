package test

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/repo/about_us"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository repo.AboutUsRepoInterface
}

var AboutUs = models.AboutUsDb{
	Model:             gorm.Model{},
	Profil:            "1",
	Visi:              "2",
	Misi:              `[{"item":"1"},{"item":"2"}]`,
	Motto:             "4",
	PerusahaanRekanan: `[{"nama_perusahaan":"b"}]`,
	SocialMedia:       `{"instagram":"instagram user","facebook":"fb user","twitter":"twitter user"}`,
	Email:             "abc@eemail.com",
	NoTelp:            "08123456",
	Office:            "abc office",
	Warehouse:         "abc warehouse",
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	logg := s.DB.Logger
	s.DB.Debug()
	logg.LogMode(logger.Info)

	s.repository = about_us.NewAboutUsRepo(s.DB, nil)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Get() {

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "about_us_dbs"`)).
		WillReturnRows(sqlmock.NewRows([]string{"profil", "visi", "no_telp", "misi", "motto", "perusahaan_rekanan", "social_media", "email", "office", "warehouse"}).
			AddRow(AboutUs.Profil, AboutUs.Visi, AboutUs.NoTelp, AboutUs.Misi, AboutUs.Motto, AboutUs.PerusahaanRekanan,
				AboutUs.SocialMedia, AboutUs.Email, AboutUs.Office,
				AboutUs.Warehouse))

	res, err := s.repository.GetAboutUs()

	if res.ID == 0 {
		s.Suite.Fail("ID is not set")
	}
	require.NoError(s.T(), err)
	assert.NotEmpty(s.T(), res)
	require.Nil(s.T(), deep.Equal(AboutUs, res))
}

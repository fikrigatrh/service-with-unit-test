package mocks

import (
	"bitbucket.org/service-ekspedisi/models"
	"github.com/stretchr/testify/mock"
)

type BlogMockRepository struct {
	mock.Mock
}

func (b *BlogMockRepository) GetAllBlog() ([]models.Blog, error) {
	args := b.Called()
	users := []models.Blog{
		{
			Title:    "judulll",
			UrlImage: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
			Content:  "ini content",
		},
	}
	return users, args.Error(1)
}

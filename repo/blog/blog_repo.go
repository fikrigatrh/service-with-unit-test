package blog

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"gorm.io/gorm"
)

type BlogRepoStruct struct {
	db  *gorm.DB
	log *log.LogCustom
}

func NewBlogRepo(db *gorm.DB, log *log.LogCustom) repo.BlogRepoInterface {
	return &BlogRepoStruct{db: db, log: log}
}

func (b BlogRepoStruct) AddBlog(v models.Blog) (models.Blog, error) {
	tx := b.db.Begin()
	err := tx.Debug().Create(&v).Error
	if err != nil {
		b.log.Error(err, "error add blog", "", nil, v, nil)
		tx.Rollback()
		return v, err
	}
	tx.Commit()
	return v, nil
}

func (b BlogRepoStruct) GetAll() ([]models.Blog, error) {
	var v []models.Blog
	err := b.db.Debug().Model(&models.Blog{}).Find(&v).Error
	if err != nil {
		b.log.Error(err, "error get all blog", "", nil, nil, nil)
		return v, err
	}
	return v, nil
}

func (b BlogRepoStruct) GetById(id int) (models.Blog, error) {
	var v models.Blog
	err := b.db.Debug().Model(&models.Blog{}).Where("id = ?", id).First(&v).Error
	if err != nil {
		b.log.Error(err, "error get blog by id", "", nil, nil, nil)
		return v, err
	}
	return v, nil
}

func (b BlogRepoStruct) UpdateData(id int, v models.Blog) (models.Blog, error) {
	tx := b.db.Begin()
	err := tx.Debug().Model(&models.Blog{}).Where("id = ?", id).Updates(v).Error
	if err != nil {
		b.log.Error(err, "error update blog", "", nil, v, nil)
		tx.Rollback()
		return v, err
	}
	tx.Commit()
	return v, nil
}

func (b BlogRepoStruct) DeleteData(id []string) error {
	tx := b.db.Begin()
	err := tx.Debug().Model(&models.Blog{}).Where("id IN (?)", id).Delete(&models.Blog{}).Error
	if err != nil {
		b.log.Error(err, "error delete blog", "", nil, nil, nil)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

package services

import (
	"github.com/yz124/superstar/dao"
	"github.com/yz124/superstar/models"
	"github.com/yz124/superstar/datasource"
)

type SuperstarService interface {
	GetAll() []models.StarInfo
	Search(country string) []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() SuperstarService {
	return &superstarService{
		dao: dao.NewSuperstarDao(datasource.InstanceMaster()),
	}
}

func (s *superstarService)GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarService)Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}

func (s *superstarService)Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *superstarService)Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *superstarService)Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *superstarService)Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}

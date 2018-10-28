package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/yz124/superstar/models"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{
		engine:engine,
	}
}

func (d *SuperstarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *SuperstarDao) GetAll() []models.StarInfo {
	datalist := make([]models.StarInfo, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperstarDao) Search(country string) []models.StarInfo {
	datalist := make([]models.StarInfo, 0)
	err := d.engine.Where("country=?", country).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperstarDao) Delete(id int) error {
	data := &models.StarInfo{Id:id, SysStatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *SuperstarDao) Create(data *models.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}

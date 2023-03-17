package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Catalog struct {
	CatId       sql.NullInt32  `gorm:"column:cat_id;primary_key"`
	CatName     sql.NullString `gorm:"column:cat_name"`
	ItemId      sql.NullInt32  `gorm:"column:item_id"`
	SNumber     int32          `gorm:"column:s_number;default:99"`
	Addtime     int32          `gorm:"column:addtime;default:0"`
	ParentCatId int32          `gorm:"column:parent_cat_id;default:0;NOT NULL"`
	Level       int32          `gorm:"column:level;default:2;NOT NULL"`
	Conn        *gorm.DB       `gorm:"-"`
}

func NewCatalog(conn *gorm.DB) *Catalog {
	return &Catalog{Conn: conn}
}

func (c *Catalog) TableName() string {
	return "catalog"
}

func (c *Catalog) TakeCatalogs(ItemId sql.NullInt32) (catalogs []*Catalog, err error) {
	catalog := &Catalog{}
	err = c.Conn.Model(catalog).Where("item_id = ?", ItemId).Find(&catalogs).Order("s_number ASC").Error
	return catalogs, err
}

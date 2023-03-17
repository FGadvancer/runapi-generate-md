package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Page struct {
	PageId         sql.NullInt32  `gorm:"column:page_id;primary_key"`
	AuthorUid      sql.NullInt32  `gorm:"column:author_uid"`
	AuthorUsername sql.NullString `gorm:"column:author_username"`
	ItemId         sql.NullInt32  `gorm:"column:item_id"`
	CatId          sql.NullInt32  `gorm:"column:cat_id"`
	PageTitle      sql.NullString `gorm:"column:page_title"`
	PageContent    sql.NullString `gorm:"column:page_content"`
	SNumber        int32          `gorm:"column:s_number;default:99"`
	Addtime        int32          `gorm:"column:addtime;default:0"`
	PageComments   string         `gorm:"column:page_comments;NOT NULL"`
	IsDel          int32          `gorm:"column:is_del;default:0;NOT NULL"`
	PageAddtime    int32          `gorm:"column:page_addtime;default:0;NOT NULL"`
	Conn           *gorm.DB       `gorm:"-"`
}

func NewPage(conn *gorm.DB) *Page {
	return &Page{Conn: conn}
}

func (p *Page) TableName() string {
	return "page"
}

func (p *Page) TakePages(catId sql.NullInt32, ItemId sql.NullInt32) (pages []*Page, err error) {
	page := &Page{}
	err = p.Conn.Model(page).Where("cat_id = ? And item_id = ?", catId, ItemId).Find(&pages).Order("s_number ASC").Error
	return pages, err
}

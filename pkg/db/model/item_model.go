package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Item struct {
	ItemId          sql.NullInt32  `gorm:"column:item_id;primary_key"`
	ItemName        sql.NullString `gorm:"column:item_name"`
	ItemDescription sql.NullString `gorm:"column:item_description"`
	Uid             sql.NullInt32  `gorm:"column:uid"`
	Username        sql.NullString `gorm:"column:username"`
	Password        sql.NullString `gorm:"column:password"`
	Addtime         sql.NullInt32  `gorm:"column:addtime"`
	LastUpdateTime  int32          `gorm:"column:last_update_time;default:0"`
	ItemDomain      string         `gorm:"column:item_domain;NOT NULL"`
	ItemType        int32          `gorm:"column:item_type;default:1;NOT NULL"`
	IsArchived      int32          `gorm:"column:is_archived;default:0;NOT NULL"`
	IsDel           int32          `gorm:"column:is_del;default:0;NOT NULL"`
	conn            *gorm.DB       `gorm:"-"`
}

func NewItem(conn *gorm.DB) *Item {
	return &Item{conn: conn}
}

func (i *Item) TableName() string {
	return "item"
}

func (i *Item) TakeItem(ItemName string) (*Item, error) {
	item := &Item{}
	err := i.conn.Model(item).Where("item_name = ? and item_type!=1", ItemName).Take(&item).Error
	return item, err
}

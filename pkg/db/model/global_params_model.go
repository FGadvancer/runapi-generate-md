package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type RunapiGlobalParam struct {
	Id             sql.NullInt32 `gorm:"column:id;primary_key"`
	ItemId         int32         `gorm:"column:item_id;default:0;NOT NULL"`
	ParamType      string        `gorm:"column:param_type;NOT NULL"`
	ContentJsonStr string        `gorm:"column:content_json_str;NOT NULL"`
	Addtime        string        `gorm:"column:addtime;NOT NULL"`
	LastUpdateTime string        `gorm:"column:last_update_time;NOT NULL"`
	Conn           *gorm.DB      `gorm:"-"`
}

func NewRunapiGlobalParam(conn *gorm.DB) *RunapiGlobalParam {
	return &RunapiGlobalParam{Conn: conn}
}

func (r *RunapiGlobalParam) TableName() string {
	return "runapi_global_param"
}
func (r *RunapiGlobalParam) TakeRunapiGlobalHeaderParam(ItemId sql.NullInt32) (runapiGlobalParams *RunapiGlobalParam, err error) {
	runapiGlobalParam := &RunapiGlobalParam{}
	err = r.Conn.Model(runapiGlobalParam).Where("item_id = ? And param_type= ? ", ItemId, "header").Find(&runapiGlobalParams).Error
	return runapiGlobalParams, err
}

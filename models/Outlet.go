package models

type Outlet struct {
	id         int64  `gorm:"column:_id;primaryKey"`
	outletid   string `gorm:"column:outletid"`
	status     string `gorm:"column:status"`
	outlet     string `gorm:"column:outlet"`
	keterangan string `gorm:"column:keterangan"`
}

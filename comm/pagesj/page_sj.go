package pagesj

import "gorm.io/gorm"

type PageFunc func(db *gorm.DB) *gorm.DB

// Paginate 分页公用组件
func Paginate(pageNo, pageSize int32) PageFunc {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo <= 0 {
			pageNo = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (pageNo - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

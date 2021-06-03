package mysql

import "github.com/jinzhu/gorm"

// Collections is mysql's collection of function
type Collections interface {
	Table(name string) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Last(out interface{}, where ...interface{}) *gorm.DB
}

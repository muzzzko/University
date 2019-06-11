package utils

import "github.com/jinzhu/gorm"

func NotFound(err error) bool {
	if nil == err {
		return false
	}

	if gorm.ErrRecordNotFound == err {
		return true
	}

	if err.Error() == "redis: nil" {
		return true
	}

	panic(err)
}

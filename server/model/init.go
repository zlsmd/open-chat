/**
 * @Author: li.zhang
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/12/11 下午4:49
 */
package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const TablePre = "bg_"

var db *gorm.DB

func InitOrm() {
	var err error
	if db == nil {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "127.0.0.1", 3306, "blog"))
		if err != nil {
			panic(err)
		}
	}
}

package helper

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var db *gorm.DB
var err error

func GetDb() *gorm.DB {
	godotenv.Load()
	if db == nil {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DB"),
		))
	}
	if err != nil {
		fmt.Print(err)
	}
	db.SingularTable(true)
	return db
}

type InTransaction func(tx *gorm.DB) error

func DoInTransaction(fn InTransaction) error {
	tx := GetDb().Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := fn(tx)
	if err != nil {
		xerr := tx.Rollback().Error
		if xerr != nil {
			return xerr
		}
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func DoInTransactionBatch(fn []InTransaction) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	for _, v := range fn {
		err := v(tx)
		if err != nil {
			xerr := tx.Rollback().Error
			if xerr != nil {
				return xerr
			}
			return err
		}
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

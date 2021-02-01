package db

import (
	"fmt"
	"log"
	"math"
	"net"
	"time"

	"github.com/jccatrinck/go-libs/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

// Load connect to database
func Load() (err error) {
	mysqlDSN := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		env.Get("MYSQL_USER", "root"),
		env.Get("MYSQL_PASSWORD", ""),
		env.Get("MYSQL_HOST", "localhost"),
		env.Get("MYSQL_PORT", "3306"),
		env.Get("MYSQL_DATABASE", "user_api"),
	)

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
			Logger:                                   logger.Discard,
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		if _, ok := err.(*net.OpError); ok {
			t := math.Pow(2, float64(i))
			log.Println("Trying to connect to MySQL again in", uint(t), "seconds")
			time.Sleep(time.Duration(t) * time.Second)
			continue
		}

		if err != nil {
			return
		}

		break
	}

	return
}

// DB from gorm lib
func DB() *gorm.DB {
	if db == nil {
		log.Panic("db is not set, verify database connection")
	}

	return db
}

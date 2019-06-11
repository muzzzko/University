package store

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"univers/pkg/config"
)

type Store struct {
	db *gorm.DB
	rds *redis.Client
}

var store *Store

func GetStore() *Store {
	if nil == store {
		store = &Store{}
		if err := store.init(); nil != err {
			panic(err)
		}
	}

	return store
}

func (s *Store) init() error {
	cfg := config.GetConfig()

	db, err := gorm.Open("mysql", cfg.MySQLDBO)
	if nil != err {
		return errors.Wrap(err, "fail connect mysql")
	}

	if err := db.DB().Ping(); nil != err {
		return errors.Wrap(err, "fail ping mysql")
	}

	s.db = db

	rds := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
		Password: "",
		DB: 0,
	})
	if _, err := rds.Ping().Result(); nil != err {
		return errors.Wrap(err, "fail ping redis")
	}
	s.rds = rds
	return nil
}

func (s *Store) close() {
	_ = s.db.Close()
}

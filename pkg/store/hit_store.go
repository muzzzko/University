package store

import (
	"strconv"
	"strings"
	"time"
	"univers/pkg/utils"
)

func (s *Store) GetLimit(IP string) (int64, bool) {
	idx := strings.LastIndex(IP, ":")
	if 0 != idx {
		IP = IP[:idx]
	}

	valueS, err := s.rds.Get(IP).Result()
	if utils.NotFound(err) {
		return 0, false
	}

	value, err := strconv.Atoi(valueS)
	if nil != err {
		panic(err)
	}
	return int64(value), true
}

func (s *Store) IncrLimit(IP string) {
	idx := strings.LastIndex(IP, ":")
	if 0 != idx {
		IP = IP[:idx]
	}

	if _, err := s.rds.Get(IP).Result(); utils.NotFound(err) {
		if _, err := s.rds.Set(IP, 1, time.Minute).Result(); nil != err {
			panic(err)
		}
	} else {
		if _, err := s.rds.Incr(IP).Result(); nil != err {
			panic(err)
		}
	}
}

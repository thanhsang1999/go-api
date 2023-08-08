package singleton

import (
	"log"
	"sync"
	"time"
)

type Singleton interface {
	AddOne() int
}
type singleton struct {
	count int
}

var (
	instance *singleton
	once     sync.Once
)

func init() {
	log.Printf("Initializing Singleton")
}

func GetInstance() *singleton {
	once.Do(func() {
		time.Sleep(time.Second)
		instance = &singleton{count: 100}
	})
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

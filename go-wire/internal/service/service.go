package service

import (
	"fmt"
	"go-wire/internal/dao"

	"github.com/google/wire"
)

// wire 定义Service需要的依赖
var Provider = wire.NewSet(New)

type Service struct {
	dao dao.Dao
}

func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = s.Close
	return
}

func (s *Service) Test() {
	s.dao.Ping()
	fmt.Println("test...")
}

func (s *Service) Close() {
	fmt.Println("service close...")
}

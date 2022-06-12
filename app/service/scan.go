package service

import (
	"example/hello/app/model"
	"fmt"
)

func (s *Server) AddScan(c model.Scan) {

	fmt.Println(c)

}

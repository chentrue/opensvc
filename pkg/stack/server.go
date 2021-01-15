package stack

import (
	"fmt"
	"opensvc/models"
)

type Server1 struct {
	Client models.Vms
}

func (s *Server1) Getvms(region string){
	aa,_ := s.Client.List(region)
	fmt.Println(aa)

}



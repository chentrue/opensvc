package models

import (
	"gorm.io/datatypes"
)


type Vm struct {

	ID      string `gorm:"primary_key;column:id"`
	Name	string `gorm:"column:name"`
	Status  string `gorm:"column:status"`
	Image   string `gorm:"column:image"`
	Size  string `gorm:"column:size"`
	Nic    	datatypes.JSON  `gorm:"column:nic"`
	Project string `gorm:"column:project"`
	Created string `gorm:"column:created"`
	Update  string `gorm:"column:update"`
}

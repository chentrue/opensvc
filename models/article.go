package models

import "github.com/jinzhu/gorm"




func GetALLCloud()  ([]*Model, error){
	var clouds []*Model

	err := Db.Where("deleted = ?", "0").Find(&clouds).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return  nil, err
	}
	return clouds, nil
}

func GetClouds(region string) (*Model, error) {
	var cloud Model
	err := Db.Where("region = ? AND deleted = ?", region, "0").First(&cloud).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return  nil, err
	}
	return &cloud, nil
}


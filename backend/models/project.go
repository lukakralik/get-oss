package models

import (
	_ "encoding/json"
	_ "gorm.io/gorm"
	 "github.com/golang-module/carbon/v2"
)

type Project struct {
	ID		int				`json:"id"`	
	Name	string			`json:"name"`
	Author	string			`json:"author"`
	Rating	float32			`json:"rating"`
	Time	carbon.Carbon	`json:"time"`
}

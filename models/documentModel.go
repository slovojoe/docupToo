package models

import "gorm.io/gorm"

type Document struct{
    gorm.Model 

    Name string
    Body string
    UserID int
}
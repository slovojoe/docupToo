package models

import "gorm.io/gorm"

type Document struct{
    gorm.Model 

    Name string
    URL string
    FileKey string
    UserID int
}
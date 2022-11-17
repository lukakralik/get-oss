package models

import (
  "gorm.io/gorm" //just the ORM
  "gorm.io/driver/sqlite" //support for SQLite from the ORM
)

var DB *gorm.DB //DB initialization

func ConnectDatabase() {
        database, err := gorm.Open(sqlite.Open("projects.db"), &gorm.Config{}) //create new DB and configurate it
        if err != nil {
                panic("Failed to connect to database!") //handle them errors
        }
        err = database.AutoMigrate(&Project{}) //this migrates the database model schema
        if err != nil {
                return
        }
        DB = database
}

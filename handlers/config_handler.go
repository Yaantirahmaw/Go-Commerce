package handlers

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

const (
    host     = "localhost"
    port     = 5432
    user     = ""
    password = ""
    dbname   = "go-commerce"
)

var db *gorm.DB

func init() {
    var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    var err error
    db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database: " + err.Error())
    }

    fmt.Println("Connected!")
}

func GetDB() *gorm.DB {
    return db
}

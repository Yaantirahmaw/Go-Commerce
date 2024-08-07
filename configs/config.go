package configs

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "" // your username
    password = "" // your pass
    dbname   = "go-commerce"
)

func ConnectDB() *gorm.DB {
    var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database: " + err.Error())
    }

    fmt.Println("Connected!")
    return db
}
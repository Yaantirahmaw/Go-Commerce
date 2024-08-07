# Go-Commerce API

Go-Commerce is a JSON-based API project for managing products and transactions in an e-commerce application. This API is built using Go, Gin, and GORM with PostgreSQL as the database.

## Features

- CRUD Operations for Products
- CRUD Operations for Product Categories
- Transaction Management
- User Authentication

## Prerequisites

Make sure you have installed the following tools:

- Go 1.16 or later
- PostgreSQL
- Git

## Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/yaantirahmaw/go-commerce.git
    cd go-commerce
    ```

2. Install dependencies using `go mod`:

    ```bash
    go mod tidy
    ```

3. Create the database configuration file `config/database.go`:

    ```go
    package configs

    import (
        "gorm.io/driver/postgres"
        "gorm.io/gorm"
    )

    var DB *gorm.DB

    func ConnectDB() (*gorm.DB, error) {
        dsn := "user=postgres password=Password dbname=go-commerce port=5432 sslmode=disable"
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            return nil, err
        }
        DB = db
        return DB, nil
    }
    ```

4. Run the migrations to create the necessary database tables:

    ```go
    package migrations

    import (
        "gorm.io/gorm"
    )

    func Migrate(db *gorm.DB) error {
        err := db.AutoMigrate(&Product{}, &ProductCategory{})
        if err != nil {
            return err
        }
        return nil
    }

    type Product struct {
        ID        uint    `gorm:"primaryKey"`
        Name      string  `gorm:"size:255"`
        Price     float64
        CategoryID uint
    }

    type ProductCategory struct {
        ID   uint   `gorm:"primaryKey"`
        Name string `gorm:"size:255"`
    }
    ```

5. Seed the database (optional):

    ```go
    package seeders

    import (
        "gorm.io/gorm"
    )

    func Seed(db *gorm.DB) error {
        // Add initial data seeding logic here
        return nil
    }
    ```

6. Run the server:

    ```bash
    go run main.go
    ```

## Usage

### List Products

- **URL:** `/products`
- **Method:** `GET`
- **Auth required:** Yes

### Get Product by ID

- **URL:** `/products/:id`
- **Method:** `GET`
- **Auth required:** No

### Create a Product

- **URL:** `/products`
- **Method:** `POST`
- **Auth required:** Yes

### Update a Product

- **URL:** `/products/:id`
- **Method:** `PUT`
- **Auth required:** Yes

### Delete a Product

- **URL:** `/products/:id`
- **Method:** `DELETE`
- **Auth required:** Yes

### User Authentication

- **Login URL:** `/login`
- **Register URL:** `/register`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Contact

Yanti Rahmawati - [yaantirahmaw@gmail.com](mailto:yaantirahmaw@gmail.com)

Project Link: [yaantirahmaw/go-commerce](https://github.com/yaantirahmaw/go-commerce)

package main

import (
	"fmt" // "context"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	port = ":9091"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Docker need time to startup the database for the app to connect
	// this needs to be hanlded in a better way.
	// https://docs.docker.com/compose/startup-order/
	time.Sleep(7 * time.Second)

	config := readENV()

	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.PG_HOST, config.PG_PORT, config.PG_USER, config.PG_DBNAME, config.PG_PASSWORD, config.PG_SSLMODE)

	db, err := gorm.Open("postgres", pgConnStr)
	if err != nil {
		print(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}

func readENV() *Config {
	return &Config{
		PG_HOST:     "db",
		PG_PORT:     "5432",
		PG_USER:     "in_user",
		PG_DBNAME:   "in_db",
		PG_PASSWORD: "in_password",
		PG_SSLMODE:  "disable",
	}
}

type Config struct {
	PG_HOST     string
	PG_PORT     string
	PG_USER     string
	PG_DBNAME   string
	PG_PASSWORD string
	PG_SSLMODE  string
}

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
	// product := Product{Name: "Fone", Price: 152.0}
	// db.Create(&product)

	// products := []Product{
	// 	{Name: "Caneca", Price: 20.0},
	// 	{Name: "Lampada", Price: 33.99},
	// }

	// db.Create(&products)

	// select one
	var product Product
	db.First(&product, 3)

	fmt.Printf("%v\n", product)

	// select all
	var products []Product
	db.Find(&products)

	for _, p := range products {
		fmt.Printf("O produto %s custa %.2f\n", p.Name, p.Price)
	}

}

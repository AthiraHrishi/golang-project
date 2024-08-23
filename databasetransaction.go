/*Demonstrate Database Transaction of at least 10 different tables with proper error handling using
a single api(use GIN).*/

package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"uniqueIndex;size:255"`
}

type Order struct {
	ID     uint   `gorm:"primaryKey"`
	Amount float64
	UserID uint
	User   User
}

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Price float64
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

type OrderItem struct {
	ID       uint    `gorm:"primaryKey"`
	OrderID  uint
	Product  Product
	ProductID uint
	Quantity uint
}

type Payment struct {
	ID        uint   `gorm:"primaryKey"`
	OrderID   uint
	Order     Order
	Amount    float64
	PaymentType string `gorm:"size:255"`
}

type Shipment struct {
	ID      uint   `gorm:"primaryKey"`
	OrderID uint
	Address string `gorm:"size:255"`
	Status  string `gorm:"size:255"`
}

type Review struct {
	ID       uint   `gorm:"primaryKey"`
	ProductID uint
	Product   Product
	Rating    uint
	Comment   string `gorm:"size:255"`
}

type Inventory struct {
	ID       uint `gorm:"primaryKey"`
	ProductID uint
	Quantity  uint
}

type Supplier struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func setupDatabase() (*gorm.DB, error) {
	// Update this connection string with your database credentials
	dsn := "host=localhost user=postgres password=yourpassword dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{}, &Order{}, &Product{}, &Category{}, &OrderItem{}, &Payment{},
		&Shipment{}, &Review{}, &Inventory{}, &Supplier{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

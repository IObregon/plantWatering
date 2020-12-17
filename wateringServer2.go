package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Readings struct {
	Origin      string  `gorm:"not null" form:"origin" json:"origin"`
	ReadingType string  `gorm:"not null" form:"reading_type" json:"reading_type"`
	Plant       string  `gorm:"null" form:"plant" json:"plant"`
	Measurement float32 `gorm:"not null" form:"measurement" json:"measurement"`
	DateTime    string  `gorm:"not null" form:"datetime" json:"datetime"`
}

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./watering.db")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&Readings{}) {
		db.CreateTable(&Readings{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Readings{})
	}
	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/readings", PostReading)
		v1.GET("/readings", GetReadings)
	}

	r.Run(":8080")
}

func PostReading(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var reading Readings
	c.Bind(&reading)

	db.Create(&reading)
	c.JSON(201, gin.H{"success": reading})
}

func GetReadings(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var readings []Readings
	db.Find(&readings)

	c.JSON(200, readings)
}

package db

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

//Bucket the db struct
type Bucket struct {
	Number    int    `json:"id" db:"id"`
	Placename string `json:"Placename" db:"Placename"`
	Latitude  string `json:"Lat" db:"Latitude"`
	Longitude string `json:"Long" db:"Longitude"`
	Visited   bool   `json:"Visited" db:"Visited"`
}

//GetVisitedLocations select all the visited locations from the database
func GetVisitedLocations(ctx *fiber.Ctx) error {
	db := openDB()
	defer db.Close()
	rows, error := selectFromBucket(db, 1)
	if error != nil {
		return error
	}
	buckets := []Bucket{}
	defer rows.Close()
	for rows.Next() {
		var row Bucket
		error = rows.Scan(&row.Number, &row.Placename, &row.Latitude, &row.Longitude, &row.Visited)
		if error != nil {
			log.Infoln("0")
			return error
		}
		buckets = append(buckets, row)
	}
	log.Infoln(buckets)
	ctx.JSON(buckets)
	return nil
}

//GetNotVisitedLocations select all the not visited locations from the database
func GetNotVisitedLocations(ctx *fiber.Ctx) error {
	db := openDB()
	defer db.Close()
	rows, error := selectFromBucket(db, 0)
	if error != nil {
		return error
	}
	buckets := []Bucket{}
	defer rows.Close()
	for rows.Next() {
		var row Bucket
		error = rows.Scan(&row.Number, &row.Placename, &row.Latitude, &row.Longitude, &row.Visited)
		if error != nil {
			log.Infoln("0")
			return error
		}
		buckets = append(buckets, row)
	}
	log.Infoln(buckets)
	ctx.JSON(buckets)
	return nil
}

//UpdateLocations updates the one or more locations
func UpdateLocations(ctx *fiber.Ctx) error {
	ctx.SendString("Update locations")
	return nil
}

//InsertLocations insert one ore more locations into the database
func InsertLocations(ctx *fiber.Ctx) error {
	ctx.SendString("Insert locations")
	return nil
}

//DeleteLocations delete one or more locations from the database
func DeleteLocations(ctx *fiber.Ctx) error {
	ctx.SendString("Delete locations")
	return nil
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		log.Fatalf("Error with open DB %v", err)
	}
	return db
}

func selectFromBucket(db *sql.DB, visited int) (*sql.Rows, error) {
	rows, error := db.Query("SELECT * FROM Buckets WHERE Visited = ?", visited)
	if error != nil {
		return rows, error
	}
	return rows, nil
}

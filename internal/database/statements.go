package database

import (
	"context"
	"database/sql"

	scan "github.com/blockloop/scan/v2"
	_ "modernc.org/sqlite"
)

// Bucket the db struct
type Bucket struct {
	Number    int     `json:"id" db:"ID"`
	Placename string  `json:"Placename" db:"Placename"`
	Latitude  float64 `json:"Lat" db:"Latitude"`
	Longitude float64 `json:"Long" db:"Longitude"`
	Visited   bool    `json:"Visited" db:"Visited"`
}

var ctx = context.Background()

func Open() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite", "./bucketlist.db")
	if err != nil {
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, err
}

func Close(db *sql.DB) error {
	err := db.Close()
	return err
}

func Initialize(db *sql.DB) error {
	sqlQuarry := "Create Table IF NOT EXISTS Buckets (ID INTEGER PRIMARY KEY AUTOINCREMENT, Placename TEXT(255), Latitude TEXT(255),Longitude TEXT(255), Visited int(1));"
	_, err := db.ExecContext(ctx, sqlQuarry)
	if err != nil {
		return err
	}
	return nil
}

func GetLocations(db *sql.DB) ([]Bucket, error) {
	var locations []Bucket
	rows, err := db.QueryContext(ctx, "SELECT * FROM Buckets")
	if err != nil {
		return locations, err
	}
	err = scan.Rows(&locations, rows)
	if err != nil {
		return locations, err
	}

	return locations, nil
}

func TempData(db *sql.DB) error {
	demo := []string{
		"Insert into Buckets (Placename, Latitude, Longitude, Visited ) Values (\"Kyoto palace\", \"35.02509\", \"135.76193\", 1);",
		"Insert into Buckets (Placename, Latitude, Longitude, Visited ) Values (\"Osaka train-station\", \"34.7332\", \"135.49928\", 0);",
	}

	for _, stmt := range demo {
		_, err := db.ExecContext(ctx, stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

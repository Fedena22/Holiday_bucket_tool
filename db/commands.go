package db

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
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
func GetVisitedLocations(ctx *gin.Context) {
	db := openDB()
	defer db.Close()
	rows, error := selectFromBucket(db, 1)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": error.Error()})
		return
	}
	buckets := []Bucket{}
	defer rows.Close()
	for rows.Next() {
		var row Bucket
		error = rows.Scan(&row.Number, &row.Placename, &row.Latitude, &row.Longitude, &row.Visited)
		if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": error.Error()})
			return
		}
		buckets = append(buckets, row)
	}
	log.Infoln(buckets)
	ctx.JSON(http.StatusOK, buckets)
}

//GetNotVisitedLocations select all the not visited locations from the database
func GetNotVisitedLocations(ctx *gin.Context) {
	db := openDB()
	defer db.Close()
	rows, error := selectFromBucket(db, 0)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": error.Error()})
		return
	}
	buckets := []Bucket{}
	defer rows.Close()
	for rows.Next() {
		var row Bucket
		error = rows.Scan(&row.Number, &row.Placename, &row.Latitude, &row.Longitude, &row.Visited)
		if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": error.Error()})
			return
		}
		buckets = append(buckets, row)
	}

	ctx.JSON(http.StatusOK, &buckets)
}

//UpdateLocations updates the one or more locations
func UpdateLocations(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Update locations"})
}

//InsertLocations insert one ore more locations into the database
func InsertLocations(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Insert locations"})

}

//DeleteLocations delete one or more locations from the database
func DeleteLocations(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Delete locations"})

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

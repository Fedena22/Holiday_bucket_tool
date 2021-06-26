package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"

	packr "github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

//Migrate Creates and updates the db schema
func Migrate() {
	// migrations := &migrate.FileMigrationSource{
	// 	Dir: "./db/migrations",
	// }
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}
	log.Printf("%s", migrations)
	db, err := sql.Open("sqlite3", "./bucketlist.db")
	if err != nil {
		log.Fatalf("Error with open DB %v", err)
	}
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Error with migration %v", err)
	}
	log.Printf("Applied %d migrations!", n)
	db.Close()
}

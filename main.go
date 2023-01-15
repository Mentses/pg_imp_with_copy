package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func check_err(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
func runCopyFrom(filename string, copy_cmd string) {
	log.Println(filename)
	ioReader, err := os.Open(filename)
	check_err(err)
	defer func() { _ = ioReader.Close() }()
	dbconn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	check_err(err)
	defer func() { _ = dbconn.Close(context.Background()) }()
	// copy_cmd := fmt.Sprintf("COPY %s FROM STDIN DELIMITER '\t' CSV HEADER ENCODING 'UTF8'", db_aim)
	log.Println(copy_cmd)
	RowsAffected, err := dbconn.PgConn().CopyFrom(context.Background(), ioReader, copy_cmd)
	check_err(err)
	log.Println(RowsAffected)
}

func main() {
	if len(os.Args) == 3 {
		filename := os.Args[1]
		copy_cmd := os.Args[2]
		runCopyFrom(filename, copy_cmd)

	} else {
		log.Println("parameter count error")
		os.Exit(1)
	}
}

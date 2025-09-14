package db

// db module connection.
// uses env file to access keys
// has fallbacks credentials so be careful.
// use db.Init to start the connection.
// use db.DB.Query("") to make sql queries, store in rows, err.
// defer -> rows.Close()

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	
	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver database/sql
)

var DB *sql.DB

func Init() {
	user := getEnv("PGUSER","postgres")
	pass := getEnv("PGPASSWORD","postgres")
	host := getEnv("PGHOST","localhost")
	port := getEnv("PGPORT","5432")
	dbname := getEnv("PGDATABASE","postgres")
	
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",user,pass,host,port,dbname)
	DB, err = sql.Open("pgx",dsn)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(30*time.Minute)
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := DB.PingContext(ctx); err != nil{
		log.Fatalf("[1] Not connected: %v", err)
	}
	log.Println("Connected to DB")
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

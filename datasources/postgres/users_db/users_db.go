package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	pq_host     = "localhost"
	pq_port     = 5432
	pq_user     = "howiewang"
	pq_password = "fef"
	pq_dbname   = "users_db"
)

var (
	Client *sql.DB
	// 	host     = os.Getenv("pq_host")
	// 	port     = os.Getenv("pq_port")
	// 	user     = os.Getenv("pq_user")
	// 	password = os.Getenv("pq_password")
	// 	dbname   = os.Getenv("pq_dbname")
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pq_host, pq_port, pq_user, pq_password, pq_dbname)
	var err error
	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = Client.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

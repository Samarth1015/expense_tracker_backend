// package postgres

// import (
// 	"fmt"
// 	"log"

// 	"github.com/Samarth1015/expense/model"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var Db *gorm.DB
// var err error

// func init() {
// 	fmt.Print("this is also bein called")
// 	dsn := "psql -h postgres-service.go-app.svc.cluster.local -p 5432 -U postgres -d expense"
// 	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Println("Error in connecting DB")

// 	}
// 	err = Db.AutoMigrate(&model.User{}, &model.Expense{})

// 	if err != nil {
// 		fmt.Print("--->err", err)

//		}
//	}
package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/Samarth1015/expense/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	// Read settings from env vars so you can change them in Kubernetes
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5050")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "example")
	dbname := getEnv("DB_NAME", "expense")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port,
	)

	fmt.Println("connecting to:", dsn) // optional debug

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting DB: %v", err)
	}

	if err = Db.AutoMigrate(&model.User{}, &model.Expense{}); err != nil {
		log.Printf("AutoMigrate error: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

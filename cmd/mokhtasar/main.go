package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	handler "github.com/davidcopperfield1991/mokhtasar/handlers"
	"github.com/davidcopperfield1991/mokhtasar/pkg"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var DB *gorm.DB

func randString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var rootCmd = &cobra.Command{
	Use: "",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "serves an http server",
	Long:  "serves an http server gardash",
	Run: func(cmd *cobra.Command, args []string) {

		dsn := "host=127.0.0.1 user=postgres password=admin dbname=mokhtasar port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// fmt.Println("hi from main")
		// fmt.Println(db)

		// fmt.Println("bye from main")
		// rows := db.First(&DB).Model(db.Table("urls"))
		// fmt.Println(rows)

		// db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode =%s",
		// 	// db, err := sql.Open("postgres", fmt.Sprintf("DATABASE_HOST=%s user=%s password=%s database=%s sslmode =%s",
		// 	// db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
		// 	config.DatabaseHost,
		// 	config.DatabaseUser,
		// 	config.DatabasePass,
		// 	config.DatabaseName,
		// 	config.DatabaseSSLMode,
		// ))
		if err != nil {
			panic(err)
		}
		// err = db.Ping()
		// if err != nil {
		// 	// fmt.Println("ping nashood")
		// 	panic(err)
		// }
		mokhtasar := &pkg.PostgresStore{
			DB:              db,
			RandomGenerator: randString,
		}
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		sl := logger.Sugar()
		handler := &handler.HTTPHandler{Mokhtasar: mokhtasar, Logger: sl}
		http.HandleFunc("/short", handler.Shorten)
		http.HandleFunc("/long", handler.Long)
		http.ListenAndServe(":8011", nil)

	},
}

var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "give digili url",
	Long:  "giv chokh digili url gardash",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := "host=127.0.0.1 user=postgres password=admin dbname=mokhtasar port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode =%s",
		// 	config.DatabaseHost,
		// 	config.DatabaseUser,
		// 	config.DatabasePass,
		// 	config.DatabaseName,
		// 	config.DatabaseSSLMode,
		// ))
		if err != nil {
			panic(err)
		}
		// err = db.Ping()
		// if err != nil {
		// 	// fmt.Println("ping nashood")
		// 	panic(err)
		// }
		mokhtasar := &pkg.PostgresStore{
			DB:              db,
			RandomGenerator: randString,
		}
		if len(args) < 1 {
			panic("you need to pass url")
		}
		key, err := mokhtasar.Shorten(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(key)
	},
}

var longCmd = &cobra.Command{
	Use:   "long",
	Short: "give orginal url",
	Long:  "giv chokh orginal url gardash",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := "host=127.0.0.1 user=postgres password=admin dbname=mokhtasar port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode =%s",
		// 	config.DatabaseHost,
		// 	config.DatabaseUser,
		// 	config.DatabasePass,
		// 	config.DatabaseName,
		// 	config.DatabaseSSLMode,
		// ))
		if err != nil {
			panic(err)
		}
		// err = db.Ping()
		// if err != nil {
		// 	// fmt.Println("ping nashood")
		// 	panic(err)
		// }
		mokhtasar := &pkg.PostgresStore{
			DB:              db,
			RandomGenerator: randString,
		}
		if len(args) < 1 {
			panic("you need pass key")
		}
		url, err := mokhtasar.GetOrginalurl(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(url)

	},
}

func main() {
	// rootCmd.AddCommand(shortCmd, longCmd, serverCmd)
	rootCmd.AddCommand(shortCmd, longCmd, serverCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

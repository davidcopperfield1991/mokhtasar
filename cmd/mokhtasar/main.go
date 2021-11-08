package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/davidcopperfield1991/mokhtasar/config"
	handler "github.com/davidcopperfield1991/mokhtasar/handlers"
	"github.com/davidcopperfield1991/mokhtasar/pkg"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var DB *sql.DB

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
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode =%s",
			config.DatabaseHost,
			config.DatabaseUser,
			config.DatabasePass,
			config.DatabaseName,
			config.DatabaseSSLMode,
		))
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			// fmt.Println("ping nashood")
			panic(err)
		}
		mokhtasar := &pkg.Mokhtasar{
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
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode =%s",
			config.DatabaseHost,
			config.DatabaseUser,
			config.DatabasePass,
			config.DatabaseName,
			config.DatabaseSSLMode,
		))
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			// fmt.Println("ping nashood")
			panic(err)
		}
		mokhtasar := &pkg.Mokhtasar{
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
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode =%s",
			config.DatabaseHost,
			config.DatabaseUser,
			config.DatabasePass,
			config.DatabaseName,
			config.DatabaseSSLMode,
		))
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			// fmt.Println("ping nashood")
			panic(err)
		}
		mokhtasar := &pkg.Mokhtasar{
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
	rootCmd.AddCommand(shortCmd, longCmd, serverCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

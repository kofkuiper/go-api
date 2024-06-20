package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kofkuiper/echo-api/handlers"
	"github.com/kofkuiper/echo-api/repositories"
	"github.com/kofkuiper/echo-api/services"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initConfig()
	db := initDB()

	accountRepo := repositories.NewAccountRepository(db)
	accountSrv := services.NewAccountService(accountRepo)
	accountHlr := handlers.NewAccountHandler(accountSrv)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go Echo API")
	})
	e.POST("/signup", accountHlr.SignUp)
	e.Logger.Fatal(e.Start(":3000"))
}

func initConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	// support replace env
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
		viper.GetInt("db.port"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to Postgres database")
	return db
}

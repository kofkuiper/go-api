package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kofkuiper/echo-api/config"
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
	cfg := config.ReadConfig()
	db := initDB(cfg.Db)

	accountRepo := repositories.NewAccountRepository(db)
	accountSrv := services.NewAccountService(cfg, accountRepo)
	accountHlr := handlers.NewAccountHandler(accountSrv)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go Echo API")
	})
	e.POST("/signup", accountHlr.SignUp)
	e.POST("/login", accountHlr.Login)
	e.POST("/validate", accountHlr.Validate)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.App.Port)))
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

func initDB(cfg config.DB) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DataBase,
		cfg.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to Postgres database")
	return db
}

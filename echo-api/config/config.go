package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App        App
		Db         DB
		JwtCfg     JwtConfig
		BlockChain BlockChain
	}

	App struct {
		Port int16
	}

	DB struct {
		Driver   string
		Host     string
		Port     int16
		Username string
		Password string
		DataBase string
	}

	JwtConfig struct {
		Secret  string
		Timeout time.Duration
	}

	BlockChain struct {
		ChainID       int
		RpcUrl        string
		BlockExplorer string
	}
)

func ReadConfig() Config {
	return Config{
		App: App{
			Port: int16(viper.GetInt("app.port")),
		},
		Db: DB{
			Driver:   viper.GetString("db.driver"),
			Host:     viper.GetString("db.host"),
			Port:     int16(viper.GetInt("db.port")),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			DataBase: viper.GetString("db.database"),
		},
		JwtCfg: JwtConfig{
			Secret:  viper.GetString("jwt.secret"),
			Timeout: readJwtTimeout(),
		},
		BlockChain: BlockChain{
			ChainID:       viper.GetInt("chain.id"),
			RpcUrl:        viper.GetString("chain.rpc"),
			BlockExplorer: viper.GetString("chain.blockExplorer"),
		},
	}
}

func readJwtTimeout() time.Duration {
	td, err := time.ParseDuration(viper.GetString("jwt.timeout"))
	if err != nil {
		panic(err)
	}
	return td
}

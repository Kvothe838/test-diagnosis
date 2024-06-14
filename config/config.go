package config

import (
	"context"
	"github.com/spf13/viper"
	"io"
	"log"
	"strings"
)

type State interface {
	GetString(string) string
}

type Source interface {
	Name() string
	Load(context.Context, State) (io.Reader, error)
}

type Configuration struct {
	Port     string       `mapstructure:"port"`
	Postgres PostgresConf `mapstructure:"postgres"`
}

type PostgresConf struct {
	User     string `mapstructure:"user"`
	DbName   string `mapstructure:"db_name"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}

func New(ctx context.Context, filePath string, sources ...Source) *Configuration {
	conf := &Configuration{}

	vip := viper.New()
	vip.SetConfigType("yaml")

	bindEnvironmentVariables(vip)
	readConfigFile(vip, filePath)
	for _, s := range sources {
		r, err2 := s.Load(ctx, vip)
		if err2 != nil {
			log.Fatalf("could not load config from source %q: %s", s.Name(), err2)
		}

		err2 = vip.MergeConfig(r)
		if err2 != nil {
			log.Fatalf("could not merge config from source %q: %s", s.Name(), err2)
		}
	}

	err := vip.Unmarshal(conf)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return conf
}

func bindEnvironmentVariables(vip *viper.Viper) {
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	vip.MustBindEnv("port")              // PORT
	vip.MustBindEnv("postgres.user")     // POSTGRES_USER
	vip.MustBindEnv("postgres.db_name")  // POSTGRES_DB_NAME
	vip.MustBindEnv("postgres.password") // POSTGRES_PASSWORD
	vip.MustBindEnv("postgres.host")     // POSTGRES_HOST
}

func readConfigFile(vip *viper.Viper, configFilePath string) {
	if len(configFilePath) == 0 {
		return
	}
	vip.SetConfigFile(configFilePath)
	if err := vip.ReadInConfig(); err != nil {
		log.Fatalf("error reading configuration file, %s", err)
	}
}

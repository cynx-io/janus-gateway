package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

var Config *AppConfig

type AppConfig struct {
	Hermes struct {
		Url string `json:"url"`
	} `json:"hermes"`
	Mercury struct {
		Url string `json:"url"`
	} `json:"mercury"`
	Plato struct {
		Url string `json:"url"`
	} `json:"plato"`
	Elastic struct {
		Url   string `json:"url"`
		Level string `json:"level"`
	} `json:"elastic"`
	Cookie struct {
		Name     string `json:"name"`
		Domain   string `json:"domain"`
		Path     string `json:"path"`
		Secure   bool   `json:"secure"`
		HttpOnly bool   `json:"http_only"`
	} `json:"cookie"`
	JWT struct {
		Secret    string `json:"secret"`
		ExpiresIn int    `json:"expires_in"`
	} `json:"jwt"`
	App struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Key     string `json:"key"`
		Port    int    `json:"port"`
		Debug   bool   `json:"debug"`
	} `json:"app"`
	CORS struct {
		Domain  string   `json:"domain"`
		Origins []string `json:"origins"`
		Enabled bool     `json:"enabled"`
	} `json:"cors"`
}

func Init() {
	// Load .env file into environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	// Set environment variable prefix for nested configs
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	Config = &AppConfig{}
	bindEnvs(Config, "")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(Config); err != nil {
		panic("failed to unmarshal config: " + err.Error())
	}
}

func bindEnvs(iface interface{}, parentKey string) {
	t := reflect.TypeOf(iface)
	v := reflect.ValueOf(iface)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		tag := field.Tag.Get("mapstructure")
		if tag == "" {
			continue
		}

		fullKey := tag
		if parentKey != "" {
			fullKey = parentKey + "." + tag
		}

		// Handle nested structs
		if fieldVal.Kind() == reflect.Struct {
			bindEnvs(fieldVal.Addr().Interface(), fullKey)
			continue
		}

		// Bind environment variable
		err := viper.BindEnv(fullKey)
		if err != nil {
			panic("error binding env: " + err.Error())
		}
	}
}

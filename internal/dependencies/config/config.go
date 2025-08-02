package config

import "github.com/cynx-io/cynx-core/src/configuration"

var Config *AppConfig

type AppConfig struct {
	Hermes struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"hermes"`
	Mercury struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"mercury"`
	Plato struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"plato"`
	Philyra struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"philyra"`
	Elastic struct {
		Url   string `mapstructure:"url"`
		Level string `mapstructure:"level"`
	} `mapstructure:"elastic"`
	Cookie struct {
		Name     string `mapstructure:"name"`
		Domain   string `mapstructure:"domain"`
		Path     string `mapstructure:"path"`
		Secure   bool   `mapstructure:"secure"`
		HttpOnly bool   `mapstructure:"http_only"`
	} `mapstructure:"cookie"`
	JWT struct {
		Secret         string `mapstructure:"secret"`
		ExpiresInHours int    `mapstructure:"expiresInHours"`
	} `mapstructure:"jwt"`
	App struct {
		Address string `mapstructure:"address"`
		Name    string `mapstructure:"name"`
		Key     string `mapstructure:"key"`
		Port    int    `mapstructure:"port"`
		Debug   bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
	CORS struct {
		Domain  string   `mapstructure:"domain"`
		Origins []string `mapstructure:"origins"`
		Enabled bool     `mapstructure:"enabled"`
	} `mapstructure:"cors"`
	Auth0 struct {
		Domain        string `mapstructure:"domain"`
		ClientId      string `mapstructure:"client_id"`
		ClientSecret  string `mapstructure:"client_secret"`
		CallbackUrl   string `mapstructure:"callback_url"`
		FrontendUrl   string `mapstructure:"frontend_url"`
		SessionSecret string `mapstructure:"session_secret"`
	} `mapstructure:"auth0"`
}

func Init() {

	Config = &AppConfig{}
	err := configuration.InitConfig("config.json", Config)
	if err != nil {
		panic("failed to initialize config: " + err.Error())
	}
}

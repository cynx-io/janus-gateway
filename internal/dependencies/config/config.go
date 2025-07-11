package config

import "github.com/cynxees/cynx-core/src/configuration"

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
		Secret         string `json:"secret"`
		ExpiresInHours int    `json:"expiresInHours"`
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

	Config = &AppConfig{}
	err := configuration.InitConfig("config.json", Config)
	if err != nil {
		panic("failed to initialize config: " + err.Error())
	}
}

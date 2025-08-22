package config

import (
	"github.com/cynx-io/cynx-core/src/configuration"
	"github.com/cynx-io/janus-gateway/internal/constant"
)

var Config *AppConfig

type AppConfig struct {
	Sites   SitesConfig `mapstructure:"sites"`
	Elastic struct {
		Url   string `mapstructure:"url"`
		Level string `mapstructure:"level"`
	} `mapstructure:"elastic"`
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
	Plutus struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"plutus"`
	Ananke struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"Ananke"`
	Auth0 struct {
		Domain string `mapstructure:"domain"`
	} `mapstructure:"auth0"`
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
		Enabled bool `mapstructure:"enabled"`
	} `mapstructure:"cors"`
}

type SitesConfig struct {
	Makeadle SiteConfig `mapstructure:"makeadle"`
	Rizzume  SiteConfig `mapstructure:"rizzume"`
}

type SiteConfig struct {
	Auth0 struct {
		ClientId      string `mapstructure:"client_id"`
		ClientSecret  string `mapstructure:"client_secret"`
		CallbackUrl   string `mapstructure:"callback_url"`
		FrontendUrl   string `mapstructure:"frontend_url"`
		SessionSecret string `mapstructure:"session_secret"`
	} `mapstructure:"auth0"`
	ApiUrl string   `mapstructure:"api_url"`
	Domain string   `mapstructure:"domain"`
	Urls   []string `mapstructure:"urls"`
}

func (s SitesConfig) Iterate(fn func(constant.SiteKey, SiteConfig)) {
	fn(constant.SiteMakeadle, s.Makeadle)
	fn(constant.SiteRizzume, s.Rizzume)
}

func (s SitesConfig) Get(siteKey constant.SiteKey) SiteConfig {
	switch siteKey {
	case constant.SiteMakeadle:
		return s.Makeadle
	case constant.SiteRizzume:
		return s.Rizzume
	default:
		return SiteConfig{}
	}
}

func Init() {

	Config = &AppConfig{}
	err := configuration.InitConfig("config.json", Config)
	if err != nil {
		panic("failed to initialize config: " + err.Error())
	}
}

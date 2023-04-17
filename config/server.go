package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pedidosya/@project_name@/models"
	"github.com/pedidosya/peya-go/logs"
	"github.com/pedidosya/peya-go/vault"
)

const (
	newRelicKey = "nrkey"
)

func LoadConfigurations(env string) *models.Configuration {
	dest := &models.Configuration{}

	file, err := os.Open(fmt.Sprintf("env_%s.json", env))

	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(dest)
	}

	if err != nil {
		panic(
			fmt.Sprintf("can't read config for environment %s: %s", env, err.Error()))
	}

	//Read app config
	dest.App.AppVersion = os.Getenv("VERSION")
	dest.App.AppEnv = env

	//Read vault config
	if vault, err := vault.Read(dest.Vault.URL, env); err != nil {
		logs.Error(fmt.Sprintf("can't read vault config %s: %s", env, err.Error()))
	} else {

		//New relic config
		if dest.NewRelic.AppName == "" {
			dest.NewRelic.AppName = dest.App.AppName
		}
		if dest.NewRelic.LicenseKey == "" {
			dest.NewRelic.LicenseKey = fmt.Sprintf("%s", vault[newRelicKey])
		}
	}

	return dest
}

package config

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"strings"
)

type NoctiLucaAuthConf struct {
	JWT NoctiLucaAuthJwtConf `koanf:"jwt"`
}

type NoctiLucaAuthJwtConf struct {
	Enabled bool   `koanf:"enabled"`
	Key     string `koanf:"key"`
	Issuer  string `koanf:"issuer"`
	// Timeout of the JWT in seconds
	Timeout int    `koanf:"timeout"`
	// Refresh timeout of the JWT in seconds
	Refresh int    `koanf:"refresh"`
}

type NoctiLucaServeConf struct {
	Host string `koanf:"host"`
	Port string `koanf:"port"`
}

type NoctiLucaConf struct {
	DSN   string             `koanf:"dsn"`
	Serve NoctiLucaServeConf `koanf:"serve"`
	Auth  NoctiLucaAuthConf  `koanf:"auth"`
}

var NlConfig *NoctiLucaConf

func LoadConfigs(configPath string) {
	defaultConfig := NoctiLucaConf{
		Serve: NoctiLucaServeConf{
			Host: "0.0.0.0",
			Port: "8000",
		},
		DSN: "",
		Auth: NoctiLucaAuthConf{JWT: NoctiLucaAuthJwtConf{
			Enabled: true,
			Key:     "somereallysecretivekey1234!@#*!",
			Issuer:  "NoctiLuca",
			Timeout: 604800,
			Refresh: 604800,
		}},
	}

	k := koanf.New(".")

	if configPath != "" {
		if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
			panic(fmt.Errorf("error loading config: %v", err))
		}
	} else {
		// Set defaults
		if err := k.Load(structs.Provider(defaultConfig, "koanf"), nil); err != nil {
			panic(fmt.Errorf("error setting default config: %v", err))
		}

		// Get set NOC_ env variables. This will overwrite any settings that have been set
		if err := k.Load(env.ProviderWithValue("NOC_", ".", func(s string, v string) (string, interface{}) {
			key := strings.Replace(strings.ToLower(strings.TrimPrefix(s, "NOC_")), "_", ".", -1)
			// Check to exist if we have a configuration option already and see if it's a slice
			switch k.Get(key).(type) {
			case []interface{}, []string:
				// Convert our environment variable to a slice by splitting on space
				return key, strings.Split(v, " ")
			}

			return key, s
		}), nil); err != nil {
			panic(fmt.Errorf("error parsing environment config: %v", err))
		}
	}

	err := k.Unmarshal("", &NlConfig)

	if err != nil {
		panic(err)
	}
}

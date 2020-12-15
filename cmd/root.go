package cmd

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type NoctiLucaConf struct {
	Host string `koanf:"host"`
	Port string `koanf:"port"`
	DNS  string `koanf:"dns"`
}

var (
	rootCmd = &cobra.Command{
		Use:              "app",
		TraverseChildren: true,
	}

	configPath = "/conf/noctiluca.yaml"

	defaultConfig = NoctiLucaConf{
		Host: "0.0.0.0",
		Port: "8000",
		DNS:  "",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "-c /path/to/file.yaml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readConf() (*NoctiLucaConf, error) {
	k := koanf.New(".")

	if configPath != "" {
		if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
			return nil, fmt.Errorf("error loading config: %v", err)
		}
	} else {
		// Set defaults
		if err := k.Load(structs.Provider(defaultConfig, "koanf"), nil); err != nil {
			return nil, fmt.Errorf("error setting default config: %v", err)
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
			return nil, fmt.Errorf("error parsing environment config: %v", err)
		}
	}
	var noctilucaConf *NoctiLucaConf

	err := k.Unmarshal("", &noctilucaConf)

	if err != nil {
		return nil, err
	}

	return noctilucaConf, nil
}

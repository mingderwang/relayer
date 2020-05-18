/*
Copyright © 2020 Jack Zampolin jack.zampolin@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package graph

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	codecstd "github.com/cosmos/cosmos-sdk/std"
	"github.com/gogo/protobuf/codec"
	"github.com/iqlusioninc/relayer/relayer"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	cfgPath     string
	homePath    string
	debug       bool
	config      *Config
	defaultHome = os.ExpandEnv("$HOME/.relayer")
	cdc         *codec.Codec
	appCodec    *codecstd.Codec

	// Default identifiers for dummy usage
	dcli = "defaultclientid"
	dcon = "defaultconnectionid"
	dcha = "defaultchannelid"
	dpor = "defaultportid"
	dord = "ordered"
)

// Config represents the config file for the relayer
type Config struct {
	Global GlobalConfig   `yaml:"global" json:"global"`
	Chains relayer.Chains `yaml:"chains" json:"chains"`
	Paths  relayer.Paths  `yaml:"paths" json:"paths"`
}

func defaultConfig() []byte {
	return Config{
		Global: newDefaultGlobalConfig(),
		Chains: relayer.Chains{},
		Paths:  relayer.Paths{},
	}.MustYAML()
}

// GlobalConfig describes any global relayer settings
type GlobalConfig struct {
	Timeout       string `yaml:"timeout" json:"timeout"`
	LiteCacheSize int    `yaml:"lite-cache-size" json:"lite-cache-size"`
}

// newDefaultGlobalConfig returns a global config with defaults set
func newDefaultGlobalConfig() GlobalConfig {
	return GlobalConfig{
		Timeout:       "10s",
		LiteCacheSize: 20,
	}
}

// MustYAML returns the yaml string representation of the Paths
func (c Config) MustYAML() []byte {
	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	return out
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {

	config = &Config{}
	home := os.Getenv("HOME")
	if home == "" {
		return fmt.Errorf("$HOME(%s) not set", home)
	}
	cfgPath := path.Join(home, ".relayer/config", "config.yaml")
	if _, err := os.Stat(cfgPath); err == nil {
		viper.SetConfigFile(cfgPath)
		if err := viper.ReadInConfig(); err == nil {
			// read the config file bytes
			file, err := ioutil.ReadFile(viper.ConfigFileUsed())
			if err != nil {
				fmt.Println("Error reading file:", err)
				os.Exit(1)
			}

			// unmarshall them into the struct
			err = yaml.Unmarshal(file, config)
			if err != nil {
				fmt.Println("Error unmarshalling config:", err)
				os.Exit(1)
			}
		}
	}
	return nil
}

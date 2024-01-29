package config

import (
	"os"

	"github.com/spf13/viper"
)

func Load() error {
	configFilename := "config.yaml"
	// ファイルが存在しない場合、config.yamlを生成する。
	if !fileExists(configFilename) {
		if err := initConfigFile(configFilename); err != nil {
			return err
		}
	}

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func initConfigFile(configFilename string) error {
	viper.SetConfigType("yaml")
	viper.Set(LINE_NOTIFY_API_TOKEN, "")

	err := viper.WriteConfigAs(configFilename)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

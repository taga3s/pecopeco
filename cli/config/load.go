package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func Load() {
	configFilename := "config.yaml"
	// ファイルが存在しない場合、config.yamlを生成する。
	if !fileExists(configFilename) {
		if err := initConfigFile(configFilename); err != nil {
			log.Fatal(err)
		}
	}

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
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

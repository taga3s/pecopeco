package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Load() {
	configFilename := "config.yaml"
	// ファイルが存在しない場合、config.yamlを生成する。
	if !FileExists(configFilename) {
		viper.SetConfigType("yaml")
		viper.Set("line_notify_api_token", "")

		if err := viper.WriteConfigAs(configFilename); err != nil {
			fmt.Println("Failed to write the file:", err)
			return
		}
	}

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

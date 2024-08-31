package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Get(property string) string {
	value := viper.GetString(property)
	return value
}

// ローカルに存在しない場合のみ、保存
func Save(property string, value string) error {
	if viper.GetString(property) != "" {
		return nil
	}
	viper.Set(property, value)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config file: %s", err)
	}
	return nil
}

func Revoke(property string) error {
	viper.Set(property, "")
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config file: %s", err)
	}
	return nil
}

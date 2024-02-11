package config

import (
	"fmt"
	"unicode/utf8"

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
		return fmt.Errorf("Error writing config file: %s\n", err)
	}
	return nil
}

func Revoke(property string) error {
	viper.Set(property, "")
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("Error writing config file: %s\n", err)
	}
	return nil
}

func IsLogin() bool {
	value := viper.GetString(PECOPECO_API_TOKEN)
	if utf8.RuneCountInString(value) != 0 {
		return true
	}
	return false
}

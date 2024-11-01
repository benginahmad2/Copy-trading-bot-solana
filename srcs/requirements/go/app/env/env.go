package env

import (
	"bot/data"
	"errors"
	"os"
)

func GetEnv(env *data.Env) error {
	api := os.Getenv("API")
	if api == "" {
		return errors.New("Api not found")
	}
	api_secret := os.Getenv("API_SECRET")
	if api_secret == "" {
		return errors.New("Api_secret not found")
	}
	env.Api_telegram = os.Getenv("API_TELEGRAM")
	if env.Api_telegram == "" {
		return errors.New("Api_telegram not found")
	}
	env.Url = os.Getenv("URL")
	if env.Url == "" {
		return errors.New("Url not found")
	}
	env.BotName = os.Getenv("ID_CHANNEL")
	if env.BotName == "" {
		return errors.New("Bot name not found")
	}
	env.AddApi(api, api_secret)
	return nil
}

func LoadEnv(env *data.Env) error {
	err := GetEnv(env)
	if err != nil {
		return err
	}
	return nil
}

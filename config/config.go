package config

type Config struct {
	BOT bot `toml:"bot"`
	REDIS  redis `toml:"redis"`
}

type bot struct {
	BotName  string `toml:"bot_name"`
	ApiKey string `toml:"api_key"`
}

type redis struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
	Db int64 `toml: db`
}
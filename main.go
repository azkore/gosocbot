package main

import (
	"github.com/azkore/gosocbot/config"

	"github.com/zhulik/margelet"
	"github.com/BurntSushi/toml"

	"fmt"
)

func main() {
	var conf config.Config
	_, err := toml.DecodeFile("config.toml", &conf)
	checkErr(err)

	fmt.Println("Bot started.")

	bot, err := margelet.NewMargelet(conf.BOT.BotName, conf.REDIS.Address, conf.REDIS.Password,
		conf.REDIS.Db, conf.BOT.ApiKey, false)
    checkErr(err)

	bot.AddCommandHandler("cat", CatHandler{})
	bot.AddCommandHandler("котейка", CatHandler{})
	bot.AddSessionHandler("start", ConfigSessionHandler{})

	go randomCatSender(bot)

	bot.Run()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

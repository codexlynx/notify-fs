package main

import (
	"github.com/codexlynx/notify-fs/pkg/services"
	"github.com/codexlynx/notify-fs/pkg/watcher"
	"github.com/h2non/filetype"
	"log"
	"os"
	"strconv"
)

const version = "0.1.0"

type Notifier struct {
	Telegram   *services.Telegram
	ChatId     int
	OnlyImages bool
}

func (notifier *Notifier) File(filePath string) {
	log.Printf("New file: %s", filePath)
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	if filetype.IsImage(fileData) {
		err := notifier.Telegram.SendImage(notifier.ChatId, filePath)
		if err != nil {
			log.Println(err)
		}
	} else {
		if !notifier.OnlyImages {
			err := notifier.Telegram.SendFile(notifier.ChatId, filePath)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Printf("Ignoring file: %s", filePath)
		}
	}
}

type Config struct {
	ChaiId     int
	Directory  string
	Token      string
	OnlyImages bool
}

func GetConfig() (*Config, error) {
	var exit = false
	chatIdString, ok := os.LookupEnv("TARGET_CHAT_ID")
	if !ok {
		log.Println("Missing TARGET_CHAT_ID environment variable")
		exit = true
	}
	directory, ok := os.LookupEnv("TARGET_DIRECTORY")
	if !ok {
		log.Println("Missing TARGET_DIRECTORY environment variable")
		exit = true
	}
	token, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Println("Missing TELEGRAM_TOKEN environment variable")
		exit = true
	}

	var onlyImages = false
	_, ok = os.LookupEnv("ONLY_IMAGES")
	if ok {
		onlyImages = true
	}
	if exit {
		log.Fatalln("Invalid service configuration")
	}

	chatId, err := strconv.Atoi(chatIdString)
	if err != nil {
		log.Fatalln("Invalid TARGET_CHAT_ID environment variable")
	}
	return &Config{
		ChaiId:     chatId,
		Directory:  directory,
		Token:      token,
		OnlyImages: onlyImages,
	}, nil
}

func main() {
	log.Printf("notify-fs %s / Connecting filesystem events with instant messaging", version)

	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	notifier := &Notifier{
		Telegram:   services.NewTelegram(config.Token),
		ChatId:     config.ChaiId,
		OnlyImages: config.OnlyImages,
	}
	log.Printf("Telegram target Chat id: %d", config.ChaiId)
	watcherFs, err := watcher.NewFs(config.Directory, notifier.File)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Watching: %s directory", config.Directory)

	defer watcherFs.Close()
	watcherFs.Watch()
}

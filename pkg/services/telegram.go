package services

import (
	"bytes"
	"context"
	"fmt"
	telegram "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"path/filepath"
)

type Telegram struct {
	client *telegram.Bot
}

func (notifier *Telegram) SendImage(chatId int, filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	_, err = notifier.client.SendPhoto(context.TODO(), &telegram.SendPhotoParams{
		ChatID: chatId,
		Photo:  &models.InputFileUpload{Data: bytes.NewReader(fileData)},
		//Caption: fileName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (notifier *Telegram) SendFile(chatId int, filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileName := filepath.Base(filePath)
	_, err = notifier.client.SendMediaGroup(context.TODO(), &telegram.SendMediaGroupParams{
		ChatID: chatId,
		Media: []models.InputMedia{
			&models.InputMediaDocument{
				Media:           fmt.Sprintf("attach://%s", fileName),
				MediaAttachment: bytes.NewReader(fileData),
				//Caption:         fileName,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func NewTelegram(token string) *Telegram {
	bot := telegram.New(token)
	return &Telegram{client: bot}
}

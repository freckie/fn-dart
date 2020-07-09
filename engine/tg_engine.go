package engine

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"fn-dart/models"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const MaxPrevMessageQueueSize = 5

type TGEngine struct {
	Bot          *telegram.BotAPI
	Cfg          *models.Config
	PrevMessages []string
}

func (tg *TGEngine) GenerateBot() error {
	bot, err := telegram.NewBotAPI(tg.Cfg.Telegram.BotToken)
	if err != nil {
		return err
	}

	tg.Bot = bot
	return nil
}

func (tg *TGEngine) SendMessage(item models.Report) error {
	if tg.IsDuplicated(item) {
		return errors.New("Message Duplicated.")
	}

	var targetConfig models.ConfigItemTargetItem
	if item.CrawlerID == "1" {
		targetConfig = tg.Cfg.Targets.CID1
	} else if item.CrawlerID == "2" {
		targetConfig = tg.Cfg.Targets.CID2
	} else if item.CrawlerID == "3" {
		targetConfig = tg.Cfg.Targets.CID3
	} else if item.CrawlerID == "4" {
		targetConfig = tg.Cfg.Targets.CID4
	} else {
		return fmt.Errorf("invalid crawler id")
	}

	msgStr := targetConfig.MessageFormat
	msgStr = strings.Replace(msgStr, "%(url)", item.ReportURL, -1)
	msgStr = strings.Replace(msgStr, "%(corp_name)", "<b>"+item.CorpName+"</b>", -1)
	msgStr = strings.Replace(msgStr, "%(description)", "<b>"+targetConfig.Description+"</b>", -1)

	for i := 0; i < targetConfig.ValuesCount; i++ {
		param := fmt.Sprintf("%%(%d)", i)
		msgStr = strings.Replace(msgStr, param, item.Values[i], -1)
	}

	for _, channel := range tg.Cfg.Telegram.Channels {
		msgType := telegram.MessageConfig{
			BaseChat: telegram.BaseChat{
				ChatID:           channel,
				ReplyToMessageID: 0,
			},
			Text:                  msgStr,
			ParseMode:             "html",
			DisableWebPagePreview: false,
		}
		_, err := tg.Bot.Send(msgType)
		if err != nil {
			log.Println("[ERROR] 메세지 전송 실패 : ", err)
		}
		log.Printf("\n======== 채널(%d)에 메세지 전송 ======== \n보고서 제목: %v\n기업 이름: %v", channel, item.Title, item.CorpName)
	}

	tg.AddMessage(item)

	return nil
}

func (tg TGEngine) TestMessage() error {
	msgStr := "<b>테스트 메세지입니다.</b>"

	for _, channel := range tg.Cfg.Telegram.Channels {
		msgType := telegram.MessageConfig{
			BaseChat: telegram.BaseChat{
				ChatID:           channel,
				ReplyToMessageID: 0,
			},
			Text:                  msgStr,
			ParseMode:             "html",
			DisableWebPagePreview: false,
		}
		sentMsg, err := tg.Bot.Send(msgType)
		if err != nil {
			log.Println("[ERROR] 테스트 메세지 전송 실패 : ", err)
		}
		log.Printf("채널(%d)에 테스트 메세지 전송 : %v", channel, sentMsg.Text)
	}

	return nil
}

func (tg TGEngine) IsDuplicated(item models.Report) bool {
	for _, prevMsg := range tg.PrevMessages {
		if item.RceptNo == prevMsg {
			return true
		}
	}
	return false
}

func (tg *TGEngine) AddMessage(item models.Report) {
	if len(tg.PrevMessages) < 1 {
		tg.PrevMessages = append(tg.PrevMessages, item.RceptNo)
	} else {
		tg.PrevMessages = append(tg.PrevMessages, item.RceptNo)
		if len(tg.PrevMessages) > MaxPrevMessageQueueSize {
			tg.PrevMessages = tg.PrevMessages[1:]
		}
	}
}

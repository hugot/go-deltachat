package deltabot

import (
	"log"
	"os"

	"github.com/hugot/go-deltachat/deltachat"
)

type Bot struct {
	commands []Command
	logger   deltachat.Logger
}

func NewBot(logger deltachat.Logger) *Bot {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	return &Bot{
		logger: logger,
	}
}

func (b *Bot) AddCommand(c Command) {
	b.commands = append(b.commands, c)
}

func (b *Bot) GetCommandForMessage(chat *deltachat.Chat, message *deltachat.Message) Command {
	for _, command := range b.commands {
		if command.Accepts(chat, message) {
			return command
		}
	}

	return nil
}

func (b *Bot) HandleMessage(c *deltachat.Context, e *deltachat.Event) {
	chatID, err := e.Data1.Int()

	if err != nil {
		b.logger.Println(err)
		return
	}

	messageID, err := e.Data2.Int()

	if err != nil {
		b.logger.Println(err)
		return
	}

	chat := c.GetChat(uint32(*chatID))
	defer chat.Unref()

	message := c.GetMessage(uint32(*messageID))
	defer message.Unref()

	command := b.GetCommandForMessage(chat, message)

	if command != nil {
		command.Execute(c, chat, message)
	}
}

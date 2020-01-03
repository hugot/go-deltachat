package deltabot

import "github.com/hugot/go-deltachat/deltachat"

// Bot commands  should implement this interface
type Command interface {
	// Determines whether a command can be executed for a message.
	Accepts(chat *deltachat.Chat, message *deltachat.Message) bool

	// Executes the bot command
	Execute(c *deltachat.Context, chat *deltachat.Chat, message *deltachat.Message)
}

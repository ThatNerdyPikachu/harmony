package harmony

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)


func NewHandler(prefix string) (CommandHandler) {
	return CommandHandler{prefix, []Command{}}
}

func (h *CommandHandler) AddCommand(name string, function func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)) {
	h.commands = append(h.commands, Command{name, function})
}

func (h *CommandHandler) MainEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	for _, command := range h.commands {
		if strings.HasPrefix(strings.ToLower(m.Content), h.prefix + command.name) {
			args := strings.Split(m.Content, " ")
			command.function(s, m, args)
		}
	}
}
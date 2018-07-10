package harmony

import "github.com/bwmarrin/discordgo"

type CommandHandler struct {
	prefix string
	commands []Command
}

type Command struct {
	name string
	function func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

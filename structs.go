package harmony

import "github.com/bwmarrin/discordgo"

// CommandHandler is a structure that contains data required for the command handler to function
type CommandHandler struct {
	Prefix    string
	Commands  map[string]*Command
	OnMessage func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

// Command is a structure that contains data that helps the CommandHandler execute commands
type Command struct {
	Run       func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
	SingleUse bool
}

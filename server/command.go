package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          botUsername,
		DisplayName:      botDisplayName,
		Description:      botDescription,
		AutoComplete:     true,
		AutoCompleteDesc: "[command] [type] command is list, add, pop and type is project, category, doc",
		AutoCompleteHint: "[command] [type]",
	}
}

func (p *Plugin) postCommandResponse(args *model.CommandArgs, text string) {
	post := &model.Post{
		UserId:    p.botUserId,
		ChannelId: args.ChannelId,
		Message:   text,
	}
	_ = p.API.SendEphemeralPost(args.UserId, post)
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split := strings.Fields(args.Command)

	command := split[0]

	// parameters := []string{}

	if command != "/knowbot" {
		return &model.CommandResponse{}, nil
	}

	p.postCommandResponse(args, "Command executed")
	return &model.CommandResponse{}, nil
}

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func getHelp() string {
	return `Available commands:
	[command] [type]
	command:
		list -> List all of given type
		add -> Add of the given type
		remove -> Remove of the given typr
		help -> Display help
		ques -> Question
	type:
		project -> Command on all projects
		category -> Command on all categories
		doc -> Command on all docs
	`
}

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          botUsername,
		DisplayName:      botDisplayName,
		Description:      botDescription,
		AutoComplete:     true,
		AutoCompleteDesc: "[command] [actionType] command is list, add, pop and actionType is project, category, doc",
		AutoCompleteHint: "[command] [actionType]",
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

	// parameters := []string{}

	if split[0] != "/knowbot" {
		return &model.CommandResponse{}, nil
	}

	command := split[1]

	actionType := split[2]

	switch command {
	case "list":
		switch actionType {
		case "project":
			p.handleListProjects(Projects, args)
		case "category":
			p.handleListCategories(Categories, args)
		}
	case "add":
		switch actionType {
		case "project":
			p.handleAddProject(args, split[2:])
		case "category":
			p.handleAddCategory(args, split[2:])
		}
	case "remove":
		switch actionType {
		case "project":
			p.handleRemoveProject(args, split[2:])
		case "category":
			p.handleRemoveCategory(args, split[2:])
		}
	}

	p.postCommandResponse(args, "Command executed")
	return &model.CommandResponse{}, nil
}

func (p *Plugin) handleListProjects(projs []Project, args *model.CommandArgs) {
	str := convertProjectsToStr(projs)
	p.postCommandResponse(args, str)
}

func (p *Plugin) handleListCategories(cats []Category, args *model.CommandArgs) {
	str := convertCategoriesToStr(cats)
	p.postCommandResponse(args, str)
}

func (p *Plugin) handleListDocs(docs []Doc, args *model.CommandArgs) {
	str := convertDocsToStr(docs)
	p.postCommandResponse(args, str)
}

func (p *Plugin) handleAddProject(args *model.CommandArgs, nameContent []string) {
	var proj Project

	nme := ""

	for _, s := range nameContent {
		nme = nme + " " + s
	}
	proj.Id = model.NewId()
	proj.Name = nme

	Projects = append(Projects, proj)
	str := convertProjectsToStr(Projects)
	p.postCommandResponse(args, "Updated Projects:\n"+str)
}

func (p *Plugin) handleAddDoc(args *model.CommandArgs, checkContent []string) {
	var doc Doc

	projName := ""
	docContent := ""
	catName := ""
	curAddition := "doc"

	for _, s := range checkContent {
		if s == "'project'" {
			curAddition = "project"
			continue
		} else if s == "'category'" {
			curAddition = "category"
			continue
		}
		if curAddition == "doc" {
			docContent = docContent + " " + s
		} else if curAddition == "project" {
			projName = projName + " " + s
		} else if curAddition == "category" {
			catName = catName + " " + s
		}
	}

	p.postCommandResponse(args, "Successfully added doc"+docContent+" for project "+projName)

	var reqProj Project
	var reqCategory Category

	for _, p := range Projects {
		if p.Name == projName {
			reqProj = p
			break
		}
	}
	for _, c := range Categories {
		if c.Name == catName {
			reqCategory = c
			break
		}
	}
	doc.Id = model.NewId()
	doc.Content = docContent
	doc.ProjectId = reqProj.Id
	doc.ProjectName = reqProj.Name
	doc.CategoryId = reqCategory.Id
	doc.CategoryName = reqCategory.Name

	Docs = append(Docs, doc)
	// str := "Successfully added"
	// p.postCommandResponse(args, str)
}

func (p *Plugin) handleAddCategory(args *model.CommandArgs, nameContent []string) {
	var cat Category

	nme := ""

	for _, s := range nameContent {
		nme = nme + " " + s
	}
	cat.Id = model.NewId()
	cat.Name = nme

	Categories = append(Categories, cat)
	str := convertCategoriesToStr(Categories)
	p.postCommandResponse(args, "Updated Categories:\n"+str)
}

func (p *Plugin) handleRemoveProject(args *model.CommandArgs, nameContent []string) {
	// var proj Project
	nme := ""

	for _, s := range nameContent {
		nme = nme + " " + s
	}

	for index, p := range Projects {
		if p.Name == nme {
			Projects = append(Projects[:index], Projects[index+1:]...)
		}
	}

	str := convertProjectsToStr(Projects)
	p.postCommandResponse(args, "Updated Projects: \n"+str)
}

func (p *Plugin) handleRemoveCategory(args *model.CommandArgs, nameContent []string) {
	// var cat Category
	nme := ""

	for _, s := range nameContent {
		nme = nme + " " + s
	}

	for index, c := range Categories {
		if c.Name == nme {
			Categories = append(Categories[:index], Categories[index+1:]...)
		}
	}

	str := convertCategoriesToStr(Categories)
	p.postCommandResponse(args, "Updated Categories: \n"+str)
}

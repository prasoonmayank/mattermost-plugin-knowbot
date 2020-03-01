package main

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var Projects []Project

func convertProjectsToStr(projs []Project) string {
	str := "\nRequired list of projects are:\n"

	for _, p := range projs {
		str += "* " + p.Name + "\n"
	}
	return str
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var Categories []Category

func convertCategoriesToStr(cats []Category) string {
	str := "\nRequired list of categories are:\n"

	for _, cat := range cats {
		str += "* " + cat.Name + "\n"
	}
	return str
}

type Doc struct {
	Id           string `json:"id"`
	Content      string `json:"content"`
	ProjectId    string `json:"project_id"`
	ProjectName  string `json:"project_name"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

var Docs []Doc

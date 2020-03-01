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

type CatDoc struct {
	Cat   Category
	CatId string
	Docs  []Doc
}

type ProjCatDoc struct {
	Proj    Project
	ProjId  string
	CatDocs []CatDoc
}

func convertDocsToStr(docs []Doc) string {
	str := "\nRequired list of docs are:\n"

	// var projCatDocs []ProjCatDoc
	// var projectNames []Project
	// var categoryNames []Category
	str = str + "ProjectName\tCategoryName\tDoc\n"

	for _, dc := range docs {
		str += "* " + dc.ProjectName + "\t" + dc.CategoryName + "\t" + dc.Content
	}

	return str

	// for _, dc := range docs {
	// 	inProjectName := false
	// 	for _, p in range projectNames {
	// 		if dc.ProjectName == p.Name {
	// 			inProjectName = true
	// 			var tempProjCatDoc projCatDocs
	// 			for _, pcd := range projCatDocs {
	// 				if pcd.ProjId == dc.ProjectId {
	// 					tempProjCatDoc = pcd
	// 					break
	// 				}
	// 			}
	// 			inCategoryName := false
	// 			for _, c in range categoryNames {
	// 				for _, cd := range tempProjCatDoc.CatDocs {
	// 					if cd.CatId == c.Id {
	// 						inCategoryName = true
	// 						cd.Docs = append(cd.Docs, dc)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	if !inProjectName {
	// 		var tempCatDoc CatDoc
	// 		vat tempProjCatDoc ProjCatDoc
	// 		var cat Category
	// 		var docs []Doc
	// 		var tempCatDocs []CatDoc
	// 		for _, c := range Categories {
	// 			if c.Id == dc.CategoryId {
	// 				cat = c
	// 				break
	// 			}
	// 		}
	// 		var proj Project
	// 		for _, p := range Projects {
	// 			if p.Id == dc.ProjectId {
	// 				proj = p
	// 				break
	// 			}
	// 		}
	// 		tempCatDoc = CatDoc{
	// 			Cat: cat,
	// 			CatId: dc.CategoryId,
	// 			docs: append(docs, dc)
	// 		}
	// 		tempProjCatDoc = ProjCatDoc{
	// 			Proj: proj,
	// 			ProjId: proj.Id,
	// 			CatDocs: append(tempCatDocs, tempCatDoc),
	// 		}
	// 		projCatDocs = append(projCatDocs, tempProjCatDoc)
	// 	}
	// }
}

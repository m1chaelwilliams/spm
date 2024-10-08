package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"spm/shared"
)

type Project struct {
	Name     string         `json:"name"`
	Path     string         `json:"path"`
	MetaData map[string]any `json:"metadata"`
}

func NewProject(name string, path string, metaData map[string]any) *Project {
	return &Project{
		Name:     name,
		Path:     path,
		MetaData: metaData,
	}
}

func (p *Project) ToString() string {
	return p.Name
}

func (p *Project) ToStringDetailed() string {
	return fmt.Sprintf(
		"Name: %-30s Path: %-30s",
		color.GreenString(p.Name),
		color.GreenString(p.Path),
	)
}

type ProjectData struct {
	Projects []*Project
	ExePath  string
}

func (p *ProjectData) CheckDuplicates(newProj *Project) *Project {
	for _, project := range p.Projects {
		if project.Name == newProj.Name {
			return project
		}
	}

	return nil
}

func (p *ProjectData) FindProject(name string) (*Project, bool) {
	for _, project := range p.Projects {
		if name == project.Name {
			return project, true
		}
	}
	return nil, false
}

func (p *ProjectData) ReplaceProject(newProj *Project) {
	for _, project := range p.Projects {
		if project.Name == newProj.Name {
			project.Path = newProj.Path
			project.MetaData = newProj.MetaData
			break
		}
	}
}

func (p *ProjectData) RemoveProject(name string) error {
	if len(p.Projects) < 1 {
		return errors.New("no projects in database")
	}

	target := -1

	for index, project := range p.Projects {
		if project.Name == name {
			target = index
		}
	}

	if target > -1 {
		lastElemIndex := len(p.Projects) - 1

		p.Projects[target] = p.Projects[lastElemIndex]
		p.Projects = p.Projects[:lastElemIndex]

		return nil
	}

	return fmt.Errorf("project with name %s not found", name)
}

func (p *ProjectData) Serialize() error {
	projDataString, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return os.WriteFile(
		filepath.Join(p.ExePath, shared.PROJECT_DATA_FILEPATH),
		projDataString,
		shared.FILEMODE_WRITE,
	)
}

func (p *ProjectData) UpdateProject(proj *Project) error {
	target, exists := p.FindProject(proj.Name)
	if !exists {
		return errors.New("project does not exist")
	}

	target.Name = proj.Name
	target.Path = proj.Path
	target.MetaData = proj.MetaData

	return nil
}

// sorts by the metadata target key
// returns the sorted list plus items that do not have that data
// func (p ProjectData) SortBy(metaTarget string) ([]*Project, []*Project, error) {
//
// 	if !utils.IsSupportedSortStrategy(metaTarget) {
// 		return nil, nil, errors.New("sorting strategy is unsupported")
// 	}
//
// 	projects := make([]*Project, 0)
// 	outliers := make([]*Project, 0)
//
// 	for i := range p.Projects {
//
// 	inner:
// 		for j := range p.Projects {
// 			switch metaTarget {
// 			case "date_added":
// 				date1, exists := p.Projects[i].MetaData["date_added"].(string)
// 				if exists == false {
// 					break inner
// 				}
// 				date2, exists := p.Projects[j].MetaData["date_added"].(string)
// 				if exists == false {
// 					continue
// 				}
// 				if utils.IsDateStrLess(date1, date2) {
// 					// wip
// 				}
// 			}
// 		}
// 	}
//
// 	return projects, outliers, nil
// }

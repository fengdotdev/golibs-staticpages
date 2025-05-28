package webwriter

import (
	"os"

	container "github.com/fengdotdev/golibs-datacontainer"
)

type IOTask struct {
	filePath string
	content  container.DataContainer
}

type paths string

type IOTasker struct {
	tasks     []IOTask
	doneTasks []paths
}

func NewIOTasker() *IOTasker {
	return &IOTasker{
		tasks:     []IOTask{},
		doneTasks: []paths{},
	}
}

func (i *IOTasker) AddTask(filePath string, content container.DataContainer) {
	i.tasks = append(i.tasks, IOTask{
		filePath: filePath,
		content:  content,
	})
}
func (i *IOTasker) Do() error {
	for _, task := range i.tasks {
		data := task.content.Get().(string)
		err := WriteToFileInWorkingDir(task.filePath, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveAllInWorkingDir(filePath string) error {
	wd, err := GetWorkingDir()
	if err != nil {
		return err
	}

	fullPath := wd + "/" + filePath
	return os.RemoveAll(fullPath)
}

func GetWorkingDir() (string, error) {
	return os.Getwd()
}

func MkdirAllInWorkingDir(filePath string) error {
	wd, err := GetWorkingDir()
	if err != nil {
		return err
	}
	fullPath := wd + "/" + filePath
	return os.MkdirAll(fullPath, os.ModePerm)
}

func WriteToFileInWorkingDir(filePath string, content string) error {

	wd, err := GetWorkingDir()
	if err != nil {
		return err
	}
	fullPath := wd + "/" + filePath

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}


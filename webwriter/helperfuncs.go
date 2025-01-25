package webwriter

import "os"

func RemoveAll(filePath string) error {
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

func MkdirAll(filePath string) error {
	wd, err := GetWorkingDir()
	if err != nil {
		return err
	}
	fullPath := wd + "/" + filePath
	return os.MkdirAll(fullPath, os.ModePerm)
}

func WriteToFile(filePath string, content string) error {

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

// filepath without extension
func WriteCSS(filePath string, content string) error {
	extension := ".css"
	return WriteToFile(filePath+extension, content)
}

// filepath without extension
func WriteHTML(filePath string, content string) error {
	extension := ".html"
	return WriteToFile(filePath+extension, content)
}

// filepath without extension
func WriteJS(filePath string, content string) error {
	extension := ".js"
	return WriteToFile(filePath+extension, content)
}

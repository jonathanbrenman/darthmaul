package commands

import (
	"darthmaul/templates"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type generateCMD struct {
	Name string
	Entity string
}

func NewGenerateCMD(entity string, name string) Command {
	return &generateCMD{
		Name: name,
		Entity: entity,
	}
}

func (c generateCMD) Execute() (err error){
	fmt.Println("Executing command generate for entity", c.Entity, "with name", c.Name+"...")

	fistLetterLowerCase := c.lcFirst(c.Name)
	fistLetterUpperCase := c.ucFirst(c.Name)

	entityUpperCase := c.ucFirst(c.Entity)

	templateFile := templates.Boilerplate

	path, _ := c.CreateDir()

	newContents := strings.Replace(string(templateFile), "darthImpl", fistLetterLowerCase + entityUpperCase + "Impl", -1)
	newContents = strings.Replace(string(newContents), "Darth", fistLetterUpperCase + entityUpperCase, -1)
	newContents = strings.Replace(string(newContents), "package template", "package " + path, -1)

	f, err := os.Create(path+"/"+fistLetterLowerCase+"_"+c.Entity+".go")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(newContents)
	return err
}

func (c generateCMD) ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func (c generateCMD) lcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func (c generateCMD) CreateDir() (path string, err error) {
	entityMap := make(map[string]string)
	entityMap["controller"] = "controllers"
	entityMap["service"] = "services"
	entityMap["repository"] = "repositories"
	// Create directory if it doesn't exist yet
	if _, err := os.Stat(entityMap[c.Entity]); os.IsNotExist(err) {
		os.Mkdir(entityMap[c.Entity], 0775)
	}
	return entityMap[c.Entity], err
}
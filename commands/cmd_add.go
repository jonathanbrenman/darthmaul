package commands

import (
	"darthmaul/templates"
	"fmt"
	"os"
)

type addCMD struct {
	Entity string
}

func NewAddCMD(entity string) Command {
	return &addCMD{
		Entity: entity,
	}
}

func (c addCMD) Execute() (err error){
	fmt.Println("Executing command add for entity", c.Entity,"...")
	entityMap := make(map[string]error)
	entityMap["cache"] = c.addCache()

	return entityMap[c.Entity]
}

func (c addCMD) addCache() (err error) {
	c.createDir()

	if err := c.addFile("cache/cache_provider.go", templates.CacheProvider); err != nil {
		return err
	}
	if err := c.addFile("cache/redis_provider.go", templates.RedisProvider); err != nil {
		return err
	}
	fmt.Println(`Run go get "github.com/go-redis/redis/v8"`)
	return nil
}

func (c addCMD) createDir() {
	// Create directory if it doesn't exist yet
	if _, err := os.Stat(c.Entity); os.IsNotExist(err) {
		os.Mkdir(c.Entity, 0775)
	}
	return
}

func (c addCMD) addFile(path, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error WriteString for file %s", path), err)
	}
	return nil
}
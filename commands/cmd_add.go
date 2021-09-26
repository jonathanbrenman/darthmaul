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
	entityMap := make(map[string]func()error)
	entityMap["redis"] = func() error {
		return c.addCache()
	}
	entityMap["postgres"] =  func() error {
		return c.addDBProviderPostgres()
	}
	entityMap["mysql"] =  func() error {
		return c.addDBProviderMysql()
	}
	entityMap["sqlite"] =  func() error {
		return c.addDBProviderSqlite()
	}
	return entityMap[c.Entity]()
}

func (c addCMD) addCache() (err error) {
	path = "cache"
	c.createDir(path)

	if err := c.addFile(path+"/cache_provider.go", templates.CacheProvider); err != nil {
		return err
	}
	if err := c.addFile(path+"/redis_provider.go", templates.RedisProvider); err != nil {
		return err
	}
	fmt.Println(`Run go get "github.com/go-redis/redis/v8"`)
	return nil
}

func (c addCMD) addDbBase(path string) (err error) {
	c.createDir(path)
	if err := c.addFile(path+"/db_provider.go", templates.DatabaseProvider); err != nil {
		return err
	}
	fmt.Println(`Run go get gorm.io/gorm`)
	return nil
}

func (c addCMD) addDBProviderMysql() (err error) {
	path = "database"
	if err := c.addDbBase(path); err != nil {
		return err
	}
	if err := c.addFile(path+"/mysql_provider.go", templates.MysqlProvider); err != nil {
		return err
	}
	fmt.Println(`Run go get gorm.io/driver/mysql`)
	return nil
}

func (c addCMD) addDBProviderPostgres() (err error) {
	path = "database"
	if err := c.addDbBase(path); err != nil {
		return err
	}
	if err := c.addFile(path+"/postgres_provider.go", templates.PostgresProvider); err != nil {
		return err
	}
	fmt.Println(`Run go get gorm.io/driver/postgres`)
	return nil
}

func (c addCMD) addDBProviderSqlite() (err error) {
	path = "database"
	if err := c.addDbBase(path); err != nil {
		return err
	}
	if err := c.addFile(path+"/sqlite_provider.go", templates.SqliteProvider); err != nil {
		return err
	}
	fmt.Println(`Run go get gorm.io/driver/sqlite`)
	return nil
}

func (c addCMD) createDir(dirName string) {
	// Create directory if it doesn't exist yet
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, 0775)
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
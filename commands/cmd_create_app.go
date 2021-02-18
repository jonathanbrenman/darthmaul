package commands

import (
	"darthmaul/templates"
	"fmt"
	"os"
	"sync"
)

var (
	path string
)

type createAppCMD struct {
	AppName string
}

func NewCreateAppCMD(appName string) Command {
	return &createAppCMD{
		AppName: appName,
	}
}

func (c createAppCMD) Execute() (err error){
	fmt.Println("Executing command create-app " + c.AppName)
	// Create main folder (AppName)
	if err := c.CreateDir(c.AppName); err != nil {
		fmt.Println("Error creating folder", c.AppName)
		return err
	}

	// Create files and folders.
	if err := c.GenerateBoilerPlate(); err != nil {
		return err
	}

	return nil
}

func (c createAppCMD) GenerateBoilerPlate() (err error){
	path = c.AppName
	fmt.Println("Generating boilerplate please wait...")
	// Create Files
	var wg sync.WaitGroup

	// Readme.md
	wg.Add(1)
	go c.CreateFile(path+"/Readme.md", templates.Readme, &wg)

	// go.mod
	wg.Add(1)
	go c.CreateFile(path+"/go.mod", fmt.Sprintf(templates.GoModule, c.AppName), &wg)

	// docker-compose.yml
	wg.Add(1)
	go c.CreateFile(path+"/docker-compose.yml", fmt.Sprintf(templates.DockerCompose, c.AppName, c.AppName), &wg)

	// docker-compose.yml
	wg.Add(1)
	go c.CreateFile(path+"/dev.Dockerfile", fmt.Sprintf(templates.DockerDev, c.AppName, c.AppName, c.AppName), &wg)

	// .gitignore
	wg.Add(1)
	go c.CreateFile(path+"/.gitignore", templates.GitIgnore, &wg)

	// Create folder cmd/api
	path = fmt.Sprintf("%s/cmd/api", c.AppName)
	fmt.Println("Changing path to", fmt.Sprintf("%s/cmd/api", c.AppName))
	if err := os.MkdirAll(path, 0775); err != nil {
		fmt.Errorf("Error creating folder " + path, err)
		return err
	}

	// main.go
	wg.Add(1)
	go c.CreateFile(path+"/main.go", fmt.Sprintf(templates.MainTemplate, c.AppName), &wg)

	// Create App folder and file
	if err := c.CreateDir(path+"/app"); err != nil {
		fmt.Errorf("Error creating folder " + path+"/app", err)
		return err
	}

	wg.Add(1)
	go c.CreateFile(path+"/app/app.go", fmt.Sprintf(templates.AppFile, c.AppName), &wg)

	// Create Config folder and file
	if err := c.CreateDir(path+"/config"); err != nil {
		fmt.Errorf("Error creating folder " + path+"/config", err)
		return err
	}

	wg.Add(1)
	go c.CreateFile(path+"/config/settings.go", templates.Settings, &wg)

	// Create ping controller
	if err := c.CreateDir(path+"/controllers"); err != nil {
		fmt.Errorf("Error creating folder " + path+"/controllers", err)
		return err
	}

	wg.Add(1)
	go c.CreateFile(path+"/controllers/ping_controller.go", templates.PingControllerTemplate, &wg)

	wg.Add(1)
	go c.CreateFile(path+"/controllers/ping_controller_test.go", templates.PingControllerTest, &wg)

	// Create middlewares (cors.go)
	if err := c.CreateDir(path+"/middlewares"); err != nil {
		fmt.Errorf("Error creating folder " + path+"/middlewares", err)
		return err
	}

	wg.Add(1)
	go c.CreateFile(path+"/middlewares/cors.go", templates.CorsMiddleware, &wg)

	// Create folder cmd/api/router
	path = fmt.Sprintf("%s/cmd/api/router", c.AppName)
	fmt.Println("Changing path to", fmt.Sprintf("%s/cmd/api/router", c.AppName))
	if err := os.MkdirAll(path+"/factory", 0775); err != nil {
		fmt.Errorf("Error creating folder " + path+"/factory", err)
		return err
	}

	// urls.go
	wg.Add(1)
	go c.CreateFile(path+"/urls.go", fmt.Sprintf(templates.UrlsMapping, c.AppName), &wg)

	// router.go
	wg.Add(1)
	go c.CreateFile(path+"/router.go", fmt.Sprintf(templates.RouterTemplate, c.AppName), &wg)

	// router_test.go
	wg.Add(1)
	go c.CreateFile(path+"/router_test.go", templates.RouterTest, &wg)

	// Change path to router/factory
	path = fmt.Sprintf("%s/cmd/api/router/factory", c.AppName)
	fmt.Println("Changing path to", fmt.Sprintf("%s/cmd/api/router/factory", c.AppName))

	// controller_factory.go
	wg.Add(1)
	go c.CreateFile(path+"/controller_factory.go", fmt.Sprintf(templates.ControllerFactory, c.AppName), &wg)

	// controller_factory_test.go
	wg.Add(1)
	go c.CreateFile(path+"/controller_factory_test.go", templates.FactoryTest, &wg)

	wg.Wait()
	return err
}

func (c createAppCMD) CreateDir(path string) (err error) {
	// Create directory if it doesn't exist yet
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0775)
	}
	return err
}

func (c createAppCMD) CreateFile(path string, content string, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := os.Create(path)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error creating file %s", path), err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error WriteString for file %s", path), err)
	}
	return
}
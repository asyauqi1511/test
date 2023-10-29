package main

import (
	"fmt"
	_ "github.com/asyauqi1511/test/docs"
	"github.com/asyauqi1511/test/internal/controller"
	empMod "github.com/asyauqi1511/test/internal/model/employees"
	"github.com/asyauqi1511/test/internal/pkg"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	// Load config.
	appConfig, err := pkg.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed load config: %v", err))
	}

	// Init log.
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Connect database.
	db, err := pkg.ConnectDB(appConfig.DB)
	if err != nil {
		panic(fmt.Sprintf("Failed connect database: %v", err))
	}

	// Initialize models.
	employeeModel, err := empMod.New(db)
	if err != nil {
		panic(fmt.Sprintf("Failed initialize employee model: %v", err))
	}

	// Initialize controllers.
	control := controller.New(employeeModel)

	// Route.
	r := gin.Default()
	r.GET("/employees", pkg.Wrap(control.EmployeeGetAll))
	r.GET("/employees/:id", pkg.Wrap(control.EmployeeGet))
	r.POST("/employees", pkg.Wrap(control.EmployeeInsert))
	r.PUT("/employees/:id", pkg.Wrap(control.EmployeeUpdate))
	r.DELETE("/employees/:id", pkg.Wrap(control.EmployeeDelete))

	// Documentation.
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8000")
}

package main

import(
	"fmt"
	"github.com/Midnight-Kittens/CRM/lead"
	"github.com/Midnight-Kittens/CRM/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setup_routes(app *fiber.App){
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}




func init_database(){
	var err error 
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil{
		panic("failed to connect to database")
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.lead{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	init_database()
	setup_routes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
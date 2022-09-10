package database 

import(
	"github.com/jinzhu/gorm"
	"github.com/Midnight-Kittens/CRM/database"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


var(
	DBConn *gorm.DB
)
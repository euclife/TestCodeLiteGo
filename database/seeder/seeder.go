package seeder

import (
"TestCodelite/models"
"github.com/jinzhu/gorm"
"log"
)

var users = []models.User{
	models.User{
		Name	: "Chandra",
		Email	: "euclife@outlook.co.id",
		Password: "password",
	},
	models.User{
		Name: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var articles = []models.Article{
	models.Article{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Article{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Article{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Article{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Article{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		articles[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Article{}).Create(&articles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed article table: %v", err)
		}
	}
}
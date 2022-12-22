package main

import (
	"fmt"
	"log"
	"orm/route"

	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	// db := repository.DB()
	// sql := `insert into users (ID,NAME,EMAIL,PASSWORD)
	// values (4, 'kiran the trainer', 'test@test.com', 'xyz@password')`
	// db.Exec(sql)
	// var users []model.User
	// db.Raw("Select ID,NAME,EMAIL,PASSWORD  from users").Scan(&users)
	// for _, result := range users {
	// 	fmt.Println(result)
	// }
	fmt.Println("Main application starts")
	loadenv()
	log.Fatal(route.RunAPI(":8090"))
}

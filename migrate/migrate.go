package migrate

import (
	"fmt"
	"log"

	"go-project/initializers"
	"go-project/model"
)

func RunMigration() {
	config, err := initializers.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	initializers.ConnectDB(&config)

	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&model.User{})
	fmt.Println("👍 Migration complete")
}

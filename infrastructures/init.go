package infrastructures

import (
	"fmt"
	"touchedFlowed/infrastructures/database"
)

func Init() {
	fmt.Println("Init infrastructures")
	database.RedisConnection()
	database.PgConnection()
}

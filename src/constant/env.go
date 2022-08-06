package constant

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Stage string
}

func (e *Env) Check() {
}

func (e *Env) Init() {
	if "" == os.Getenv("STAGE") {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("get local env err:", err)
		}
	}
}

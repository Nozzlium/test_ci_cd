package lib

import "github.com/joho/godotenv"

func ReadEnv() {
	godotenv.Load("development.env")
}

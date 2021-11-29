package main

import (
	"github.com/barqus/fillq_backend/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}

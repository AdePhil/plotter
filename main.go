package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AdePhil/plotter/cmd"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
			log.Fatal(err)
	}
	fmt.Print(dir)
	err = godotenv.Load(filepath.Join("./", ".env"))
	
  if err != nil {
    log.Fatal("Error loading .env file here", err)
  }

	cmd.RootCmd.Execute()
}

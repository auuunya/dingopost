// Package Dingo is the main entry point to your application.
package app

import (
	"dingo/app/handler"
	"dingo/app/model"
	"fmt"
	"os"

	"github.com/dinever/golf"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// Init loads a public and private key pair used to create and validate JSON
// web tokens, or creates a new pair if they don't exist. It also initializes
// the database connection.
func Init( /*dbPath,*/ privKey, pubKey string) {
	model.InitializeKey(privKey, pubKey)
	if err := model.Initialize(); err != nil {
		err = fmt.Errorf("failed to intialize db: %v", err)
		panic(err)
	}
	//	if err := model.Initialize(dbPath, fileExists(dbPath)); err != nil {
	//		err = fmt.Errorf("failed to intialize db: %v", err)
	//		panic(err)
	//	}
	fmt.Printf("Database is used at ")
}

// Run starts our HTTP server on the given port.
func Run(portNumber string) {
	app := golf.New()
	app = handler.Initialize(app)
	fmt.Printf("Application Started on port %s\n", portNumber)
	app.Run(":" + portNumber)
}

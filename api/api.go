package api

import (
	"fmt"

	"github.com/labstack/echo"
)

// Requesting the API with a curl -X POST command :
// Function to add a block

// Function to print block in the blockchain database based on the ID

// Function to get Hash of a block in the blockchain based on the ID

func Test() {
	echoInstance := echo.New()
	fmt.Println(echoInstance.Start("localhost:8080"))
}

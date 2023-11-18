package cmd

import (
	"ElliotBrookes/halo-platform/internal/startup"
	"fmt"
)

func main() {

	fmt.Println("Starting Server...")
	// get a configured server
	s, err := startup.InitServer()
	if err != nil {
		return
	}

	// Init and configure the routers for the application routes as well as adapters.

	s.ListenAndServe()
}

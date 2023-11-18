package startup

import "net/http"

func InitServer() (*http.Server, error) {
	// Initialise the Server
	s := http.Server{}
	return &s, nil

}

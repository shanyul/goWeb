package main

import "designer-api/server"

func main() {
	// start server
	server.StartApp()
	// close server
	defer server.ExitApp()
}

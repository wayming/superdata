package main

import "github.com/wayming/superdata/internal/loader"

func main() {
	sunLoader := loader.Loader{TableName: "SUNSUPER"}
	sunLoader.Connect()
	sunLoader.Create()
	sunLoader.Disconnect()
}

package main

import "github.com/wayming/superdata/internal/loader"

func main() {
	sunLoader := loader.Loader{TableName: "SUNSUPER", DateFormat: "02 Jan 2006"}
	sunLoader.Connect()
	sunLoader.Create()
	sunLoader.Load("datafiles/SUNSUPER/balanced_pool_prices.csv")
	sunLoader.Disconnect()
}

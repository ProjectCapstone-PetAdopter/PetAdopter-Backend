package config

import "os"

var (
	SECRET     string = os.Getenv("SECRET")
	SERVERPORT int    = 8000
	COST       int    = 10
	UPLOADPATH string
)

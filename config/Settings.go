package config

import "os"

var ModelicaKeywords = map[string]bool{"der": true, "and": true, "or": true, "not": true, "constant": true}

var USERNAME = os.Getenv("USERNAME")

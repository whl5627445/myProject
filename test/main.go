package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"log"
)

func main() {
	var result map[string]string
	jsonData := `{"name": "john", "email": "john@example.com"}`
	err := sonic.Unmarshal([]byte(jsonData), &result)
	fmt.Println(err)
	log.Println(result)
}

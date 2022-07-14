package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)
	reverse(sd)

	whatIsIt = string(sd)
	fmt.Println(whatIsIt)
}

func reverse(items []byte) {
	for i, j := 0, len(items)-1; i < len(items)/2; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

 // CodeCrafters — Operation Gopher Protocol
       // Module: File Pipeline
       // Author: [Godwin Emaikwu]
       // Squad:  [The Benchmarkers]


package main

import (
	"fmt"
	"regexp"
	"strings"
)

func RemToDo (s string) string {
	words := strings.Fields(s)
	for i = 0; i < len(words); i++ {
		if words[i] = "TODO" {
			words[i] = "ACTION"
		}
	}
	return strings.Join(words," ")
}
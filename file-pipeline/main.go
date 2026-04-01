// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Godwin Emaikwu]
// Squad:  [The Benchmarkers]

package main

import (
	"fmt"
	"strings"
)

func RemToDo(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "TODO" {
			words[i] = "ACTION"
		}
	}
	return strings.Join(words, " ")
}

func allCaps(line string) string {
    if line == strings.ToUpper(line) && line != "" {
        words := strings.Fields(line)
        for i, w := range words {
            words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
        }
        return strings.Join(words, " ")
    }
    return line
}

func lowerToUpper(line string) string {
    if line == strings.ToLower(line) && line != "" {
        return strings.ToUpper(line)
    }
    return line
}

func trim(line string) string {
    return strings.TrimSpace(line)
}

func reverse(line string) string {
    if strings.Contains(line, "REVERSE") {
        words := strings.Fields(line)
        for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
            words[i], words[j] = words[j], words[i]
        }
        return strings.Join(words, " ")
    }
    return line
}

func ifLong(line string) string {
    if len(line) > 80 {
        return line + " [TRUNCATED]"
    }
    return line
}

func main() {

    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . <input.txt> <output.txt>")
        return
    }

    input := os.Args[1]
    output := os.Args[2]

    if input == output {
        fmt.Println("Input and output cannot be the same file.")
        return
    }

    file, err := os.Open(input)
    if err != nil {
        fmt.Printf("File not found: %s\n", input)
        return
    }
    defer file.Close()

    if info, err := os.Stat(output); err == nil && info.IsDir() {
        fmt.Println("Cannot write to output: path is a directory, not a file.")
        return
    }
func main() {
	fmt.Println(RemToDo("Director X TODO why this evil"))
}

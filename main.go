package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {

	// if piping to stdin
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Printf("%s\n", hashit(scanner.Text())[:8])
		}
	} else {
		fmt.Printf("%s\n", hashit(os.Args[1])[:8])
	}
}
func hashit(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

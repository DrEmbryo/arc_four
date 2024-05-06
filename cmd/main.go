package main

import (
	"flag"
	"fmt"
	"os"

	ArcFour "github.com/DrEmbryo/arc_four/lib"
)

func main() {
	input := os.Args[1]
	typeFlagGroup := flag.NewFlagSet("", flag.ExitOnError)
	key := typeFlagGroup.String("key", "secret", "Encription key")
	inputType := typeFlagGroup.String("type", "text", "Input type")
	outputPath := typeFlagGroup.String("output", input, "Output file path")
	typeFlagGroup.Parse(os.Args[2:])

	var source string
	enc := &ArcFour.RC4{}
	enc.Init(*key)
	
	switch *inputType {
		case "text": 
			source = input
			encrypted := enc.Encrypt(source)
			fmt.Printf("Encrypted message: %s", encrypted)
		case "file":
			_, err := os.Stat(input)
			if err != nil {
				panic(fmt.Sprintf("file does not exist at path: %s", input))
			}
			source, err := os.ReadFile(input)
			if err != nil {
				panic(err)
			}

			encrypted := enc.Encrypt(string(source))

			err = os.WriteFile(*outputPath, []byte(encrypted), 0666)
			if err != nil {
				panic(err)
			}
	}
}
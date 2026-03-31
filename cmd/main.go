package main

import (
	"flag"

	"github.com/fatih/color"
	"github.com/myntdeveloper/passgen/internal/generator"
)

func main() {
	length := flag.Int("l", 12, "Password length")
	symbols := flag.Bool("s", false, "Include symbols")
	numbers := flag.Bool("n", false, "Include numbers")
	upper := flag.Bool("u", false, "Include uppercase letters")
	count := flag.Int("c", 1, "Number of passwords to generate")
	flag.Parse()

	cfg := generator.Config{
		Length:  *length,
		Symbols: *symbols,
		Numbers: *numbers,
		Upper:   *upper,
		Count:   *count,
	}

	for i := 0; i < cfg.Count; i++ {
		password, err := generator.Generate(cfg)
		if err != nil {
			color.Red("Error:", err)
			return
		}
		color.Blue(password)
	}

}

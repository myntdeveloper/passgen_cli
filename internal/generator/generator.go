package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type Config struct {
	Length  int
	Symbols bool
	Numbers bool
	Upper   bool
	Count   int
}

func Generate(cfg Config) (string, error) {
	if cfg.Length <= 0 {
		return "", errors.New("password length must be more than 0")
	}

	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>/?"

	charTypes := []struct {
		enabled bool
		set     string
		name    string
	}{
		{true, lower, "lower"},
		{cfg.Upper, upper, "upper"},
		{cfg.Numbers, numbers, "numbers"},
		{cfg.Symbols, symbols, "symbols"},
	}

	var enabledSets []string

	for _, ct := range charTypes {
		if ct.enabled {
			enabledSets = append(enabledSets, ct.set)
		}
	}

	if len(enabledSets) == 0 {
		return "", errors.New("no character sets selected")
	}

	if cfg.Length < len(enabledSets) {
		return "", errors.New("password length too short for selected sets")
	}

	baseQuota := cfg.Length / len(enabledSets)
	remainder := cfg.Length % len(enabledSets)

	setCounts := make([]int, len(enabledSets))
	for i := range setCounts {
		setCounts[i] = baseQuota
	}

	for left := 0; left < remainder; left++ {
		i, err := cryptoRandInt(len(enabledSets))
		if err != nil {
			return "", err
		}
		setCounts[i]++
	}

	var passRunes []rune

	for i, set := range enabledSets {
		for n := 0; n < setCounts[i]; n++ {
			idx, err := cryptoRandInt(len(set))
			if err != nil {
				return "", err
			}
			passRunes = append(passRunes, rune(set[idx]))
		}
	}

	secureShuffle(passRunes)

	return string(passRunes), nil
}

func cryptoRandInt(max int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
}

func secureShuffle(a []rune) {
	for i := len(a) - 1; i > 0; i-- {
		jRand, err := cryptoRandInt(i + 1)
		if err != nil {
			return
		}
		a[i], a[jRand] = a[jRand], a[i]
	}
}

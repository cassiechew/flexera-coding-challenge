package checker

import (
	"strings"

	"github.com/ryanchew3/flexera-coding-challenge/internal/model"
)

// getKeys gets the keys of in incoming mapstructure
func getKeys(in map[int64][]*model.Computer) []int64 {
	keys := make([]int64, len(in))
	i := 0
	for k := range in {
		keys[i] = k
		i++
	}
	return keys
}

// ProcessChecks processes the list of user owned machines and calculates how many licences are needed
func ProcessChecks(in map[int64][]*model.Computer) int {
	keys := getKeys(in)

	totalLicences := 0

	for _, k := range keys {
		numLaptops := 0
		numDesktops := 0
		id := make(map[int64]rune)

		for _, c := range in[k] {
			_, ok := id[c.Id]
			if !ok {
				id[c.Id] = ' '
				switch strings.ToLower(c.Type) {
				case "desktop":
					numDesktops++
				case "laptop":
					numLaptops++
				}
			}
		}

		numDesktops = numDesktops - numLaptops
		if numDesktops < 0 {
			numDesktops = 0
		}
		totalLicences = totalLicences + numDesktops + numLaptops
	}
	return totalLicences
}

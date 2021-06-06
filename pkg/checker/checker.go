package checker

import (
	"strings"

	"github.com/ryanchew3/flexera-coding-challenge/internal/model"
)

func processChecks(in map[int64][]*model.Computer) int {

	keys := make([]int64, len(in))
	i := 0
	for k := range in {
		keys[i] = k
		i++
	}

	totalLicences := 0

	for _, k := range keys {
		numLaptops := 0
		numDesktops := 0
		var id map[int64]rune

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

		numDesktops -= numLaptops
		if numDesktops < 0 {
			numDesktops = 0
		}
		totalLicences = numDesktops + numLaptops
	}

	return totalLicences
}

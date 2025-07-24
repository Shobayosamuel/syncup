package utils

import (
	"fmt"
)

const EarthRadiusKm = 6371.0

func HaversineSQL(lat, lon float64) string {
	return fmt.Sprintf(`
		%[1]f * acos(
			cos(radians(%[2]f)) * cos(radians(latitude)) *
			cos(radians(longitude) - radians(%[3]f)) +
			sin(radians(%[2]f)) * sin(radians(latitude))
		)
	`, EarthRadiusKm, lat, lon)
}

func CountMutualInterests(a, b []string) int {
	set := make(map[string]bool)
	for _, interest := range a {
		set[interest] = true
	}
	count := 0
	for _, interest := range b {
		if set[interest] {
			count++
		}
	}
	return count
}
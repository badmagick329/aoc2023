package day2

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func linesFromFile() chan string {
	lines := make(chan string)
	dat, err := os.ReadFile(FILE)
	if err != nil {
		log.Fatal("Error reading file")
	}
	readLines := strings.Split(string(dat), "\n")
	go func() {
		defer close(lines)
		for _, rl := range readLines {
			if rl == "" {
				break
			}
			lines <- rl
		}
	}()
	return lines
}

func parseColors(line string) []Colors {
	cubes := []Colors{}
	pulls := strings.Split(strings.Split(line, ": ")[1], "; ")
	for _, pull := range pulls {
		colors := toColors(pull)
		cubes = append(cubes, colors)
	}
	return cubes
}

func toColors(pull string) Colors {
	colorsStr := strings.Split(pull, ", ")
	colors := Colors{}
	for _, c := range colorsStr {
		spl := strings.Split(c, " ")
		num, err := strconv.Atoi(spl[0])
		if err != nil {
			log.Fatalf("Error parsing %s\n", c)
		}
		switch spl[1] {
		case "red":
			colors.r += num
		case "green":
			colors.g += num
		case "blue":
			colors.b += num
		default:
			log.Fatalf("Unrecognised color %s", c)
		}
	}
	return colors
}

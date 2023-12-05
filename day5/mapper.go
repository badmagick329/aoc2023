package day5

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type RangeInfo struct {
	start  int
	end    int
	offset int
}

type Map struct {
	name string
	// sourceToDestMap map[int]int
	rangeInfos []RangeInfo
}

func NewMap(name string, lines []string) Map {
	mp := Map{
		name: name,
		// sourceToDestMap: make(map[int]int),
		rangeInfos: []RangeInfo{},
	}
	for _, line := range lines {
		// mp.updateMap(line)
		mp.updateRangeInfo(line)
	}
	return mp
}

func (m *Map) updateRangeInfo(line string) {
	pattern := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
	nums := pattern.FindStringSubmatch(line)
	if len(nums) != 4 {
		log.Fatalf("Num length %d found in %s\n", len(nums), line)
	}
	dst, _ := strconv.Atoi(nums[1])
	src, _ := strconv.Atoi(nums[2])
	rangeLen, _ := strconv.Atoi(nums[3])
	diff := dst - src
	end := src + rangeLen - 1
	m.rangeInfos = append(m.rangeInfos, RangeInfo{
		start:  src,
		end:    end,
		offset: diff,
	})
}

func (m *Map) Dest(src int) int {
	for _, ri := range m.rangeInfos {
		if src >= ri.start && src <= ri.end {
			return src + ri.offset
		}
	}
	return src
}

type SeedRange struct {
	start    int
	rangeLen int
}

type Mapper struct {
	seeds      []int
	seedRanges []SeedRange
	maps       []Map
}

func NewMapper(lines []string) Mapper {
	mapper := Mapper{}
	idx := 0
	if !strings.Contains(lines[0], "seeds: ") {
		log.Fatalf("Seeds not found in line: %s", lines[0])
	} else {
		re, _ := regexp.Compile(`\d+`)
		numsStr := re.FindAllStringSubmatch(lines[0], -1)
		for _, ns := range numsStr {
			conv, _ := strconv.Atoi(ns[0])
			mapper.seeds = append(mapper.seeds, conv)
		}
	}
	for ; idx < len(lines); idx++ {
		if strings.Contains(lines[idx], "map:") {
			break
		}
	}
	mapLines := []string{}
	var src, dst string
	for ; idx < len(lines); idx++ {
		if lines[idx] == "" {
			continue
		}
		if strings.Contains(lines[idx], "map:") {
			if len(mapLines) > 0 {
				mapper.maps = append(mapper.maps, NewMap(src, mapLines))
				mapLines = []string{}
			}
			src, dst = lineToSourceAndDest(lines[idx])
			continue
		}
		mapLines = append(mapLines, lines[idx])
	}
	if len(mapLines) > 0 {
		mapper.maps = append(mapper.maps, NewMap(src, mapLines))
	}
	mapper.maps = append(mapper.maps, NewMap(dst, []string{}))
	return mapper
}

func NewMapper2(lines []string) Mapper {
	mapper := Mapper{}
	idx := 0
	if !strings.Contains(lines[0], "seeds: ") {
		log.Fatalf("Seeds not found in line: %s", lines[0])
	} else {
		re, _ := regexp.Compile(`\d+`)
		numsStr := re.FindAllStringSubmatch(lines[0], -1)
		seedNums := []int{}
		for _, ns := range numsStr {
			conv, _ := strconv.Atoi(ns[0])
			seedNums = append(seedNums, conv)
		}
		for i := 0; i < len(seedNums); i += 2 {
			mapper.seedRanges = append(mapper.seedRanges, SeedRange{
				seedNums[i],
				seedNums[i+1],
			})
		}
	}
	for ; idx < len(lines); idx++ {
		if strings.Contains(lines[idx], "map:") {
			break
		}
	}
	mapLines := []string{}
	var src, dst string
	for ; idx < len(lines); idx++ {
		if lines[idx] == "" {
			continue
		}
		if strings.Contains(lines[idx], "map:") {
			if len(mapLines) > 0 {
				mapper.maps = append(mapper.maps, NewMap(src, mapLines))
				mapLines = []string{}
			}
			src, dst = lineToSourceAndDest(lines[idx])
			continue
		}
		mapLines = append(mapLines, lines[idx])
	}
	if len(mapLines) > 0 {
		mapper.maps = append(mapper.maps, NewMap(src, mapLines))
	}
	mapper.maps = append(mapper.maps, NewMap(dst, []string{}))
	return mapper
}

func lineToSourceAndDest(line string) (string, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error parsing line: %s", line)
		}
	}()
	spl := strings.Split(line, "-to-")
	return spl[0], strings.Split(spl[1], " ")[0]
}

func (m *Mapper) Transform(src, dst string, srcNum int) (int, error) {
	curr := -1
	for _, mp := range m.maps {
		if mp.name == src {
			curr = mp.Dest(srcNum)
			continue
		}
		if mp.name == dst && curr != -1 {
			break
		}
		if curr != -1 {
			curr = mp.Dest(curr)
		}
	}
	if curr != -1 {
		return curr, nil
	}
	return curr, fmt.Errorf("Dest not found\nsrc: %s, dst: %s, srcNum: %d\n", src, dst, srcNum)
}

func (m *Mapper) sourceToDest(from, to string) *[2]int {
	idx := [2]int{-1, -1}
	for i, s := range m.maps {
		if s.name == from {
			idx[0] = i
		} else if s.name == to {
			idx[1] = i
		}
	}
	if idx[0] == -1 || idx[1] == -1 {
		return nil
	}
	return &idx
}

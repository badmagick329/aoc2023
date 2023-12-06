package lib

import (
	"regexp"
	"strconv"
)

func ParseNumsFromText(text string) []int {
    nums := []int{}
    re, _ := regexp.Compile(`\d+`)
    numsStr := re.FindAllStringSubmatch(text, -1)
    for _, ns := range numsStr {
        conv, _ := strconv.Atoi(ns[0])
        nums = append(nums, conv)
    }
    return nums
}

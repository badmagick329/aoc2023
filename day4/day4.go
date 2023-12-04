package day4

import (
	"runtime"
)

const FILE = "day4/file.txt"

var WorkerCount = runtime.NumCPU()

func Run() {
	RunOne()
	RunTwo()
}

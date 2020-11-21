package helpers

import (
	"runtime"
	"strconv"
)

var threads int
var arch string
var goos string
var root string

func New() {
	threads = GetNumThreads()
	arch = runtime.GOARCH
	goos = runtime.GOOS
	root = runtime.GOROOT()
}

func GetNumThreads() int {
	return runtime.GOMAXPROCS(0)
}

func SetNumThreads(num int) int {
	if num == 0 {
		return threads
	}
	if num < 0 {
		Print("Thread number cannot be smaller than zero. Current value not changed")

		return threads
	}
	runtime.GOMAXPROCS(num)
	threads = num

	return num
}

func GetArch() string {
	if EmptyString(arch) {
		Print("Architecture not defined, call the funtion New()")

		return "unknown"
	}
	return arch
}

func GetOS() string {
	if EmptyString(arch) {
		Print("OS not defined, call the funtion New()")

		return "unknown"
	}
	return goos
}

func GetRoot() string {
	if EmptyString(arch) {
		Print("ROOT not defined, call the funtion New()")

		return "unknown"
	}

	return root
}

func LoadPerThread(work int) int {
	var threads int
	threads = GetNumThreads()

	if work%threads == 0 {
		return work / threads
	}
	return work/threads + 1
}

func ThreadPagination(work int, threadID int) (int, int) {
	var threads int
	threads = GetNumThreads() - 1
	if threadID > threads {
		Print("Exceeded the number of threads available")

		return work, work
	}
	if work <= threads {
		return 0, work
	}

	var start, end, each int

	each = LoadPerThread(work)

	start = each * threadID
	if threadID == threads {
		return start, work
	}
	end = start + each - 1

	return start, end
}

func IsInt(s string) bool {
	_, e := strconv.Atoi(s)
	if e != nil {
		return false
	}
	return true
}

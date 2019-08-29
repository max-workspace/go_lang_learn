package main

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/max_workspace/golang_learn/stringutil"
)

type syncInfo struct {
	mu    sync.Mutex
	index int
}

var (
	searchIn    string = "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat         string = "[0-9]+.[0-9]+"
	syncInfoObj syncInfo
)

func main() {
	// 测试regexp
	testPackageRegexp()

	// 测试math/big
	testPackageMathBig()

	// 测试自定义包以及加载过程
	testPackageCustomPackage()

	// 测试sync
	testPackageSync()

	return
}

func testPackageRegexp() {
	fmt.Println("test package regexp")

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}

	re, _ := regexp.Compile(pat)
	// 将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	// 参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
	fmt.Println("\n")
}

func testPackageMathBig() {
	fmt.Println("test package math/big")
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(3)
	ip := big.NewInt(5)
	result := big.NewInt(1)
	fmt.Printf("type: %T|value: %v\n", im, im)
	fmt.Printf("type: %T|value: %v\n", in, in)
	fmt.Printf("type: %T|value: %v\n", io, io)
	fmt.Printf("type: %T|value: %v\n", ip, ip)
	result.Mul(io, ip)
	fmt.Printf("%v * %v = %v\n", io, ip, result)
	result.Add(io, ip)
	fmt.Printf("%v + %v = %v\n", io, ip, result)
	fmt.Println("\n")
}

func testPackageCustomPackage() {
	fmt.Println("test package custom package")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("\n")
}

func testPackageSync() {
	fmt.Println("test package sync")
	// 非同步状态下执行
	syncInfoObj.index = 0
	go updateInfoObj(1000, "first")
	go updateInfoObj(1500, "second")

	// 等待10s保证goroutine的执行
	time.Sleep(10 * time.Second)

	// 同步状态下执行
	go updateSyncInfoObj(1000, "first")
	go updateSyncInfoObj(1500, "second")

	// 等待10s保证goroutine的执行
	time.Sleep(10 * time.Second)
	fmt.Println("\n")
}

func updateInfoObj(sleepMicroseconds int, tag string) {
	for i := 0; i < 5; i++ {
		// 增加等待时间，模拟处理不同数据量时出现的执行时间的细微差别
		time.Sleep(time.Duration(sleepMicroseconds) * time.Microsecond)
		syncInfoObj.index += 1
		fmt.Printf(
			"unsync|action: %v|index: %d|time: %v|value: %d\n",
			tag,
			i,
			time.Now().Format("2006-01-02 15:04:05"),
			syncInfoObj.index)
	}
}

func updateSyncInfoObj(sleepMicroseconds int, tag string) {
	syncInfoObj.mu.Lock()
	for i := 0; i < 5; i++ {
		// 增加等待时间，模拟处理不同数据量时出现的执行时间的细微差别
		time.Sleep(time.Duration(sleepMicroseconds) * time.Microsecond)
		syncInfoObj.index += 1
		fmt.Printf("sync|action: %v|index: %d|time: %v|value: %d\n", tag, i, time.Now().Format("2006-01-02 15:04:05"), syncInfoObj.index)
	}
	syncInfoObj.mu.Unlock()
}

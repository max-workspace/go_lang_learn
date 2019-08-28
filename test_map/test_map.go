package main

import (
	"fmt"
	"sort"
)

var (
	unsortValMap = map[string]int{
		"alpha":   34,
		"bravo":   56,
		"charlie": 23,
		"delta":   87,
		"echo":    56,
		"foxtrot": 12,
		"golf":    34,
		"hotel":   16,
		"indio":   87,
		"juliet":  65,
		"kili":    43,
		"lima":    98,
	}

	capitalMap = map[string]int{
		"New Delhi": 55,
		"Beijing":   20,
		"Washiton":  25,
	}

	capitalAndCountryMap = map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
	}
)

func main() {
	// 测试map是否包含指定元素
	testMapHasElement()

	// 测试map删除指定元素
	testMapDeleteElement()

	// 测试map是否按顺序进行输出
	testMapOrder()

	// 测试map类型的切片
	testMapSlice()

	// 测试map排序
	testMapSort()

	// 测试map反转
	testMapInvert()
}

func testMapHasElement() {
	fmt.Println("test map has element")

	value, isPresent := capitalMap["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in capitalMap is: %d\n", value)
	} else {
		fmt.Printf("capitalMap does not contain Beijing")
	}

	value, isPresent = capitalMap["Paris"]
	fmt.Printf("Is \"Paris\" in capitalMap ?: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)
	fmt.Println("")
}

func testMapDeleteElement() {
	fmt.Println("test map delete element")

	delete(capitalMap, "Washington")
	value, isPresent := capitalMap["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in capitalMap is: %d\n", value)
	} else {
		fmt.Printf("capitalMap does not contain Washington\n")
	}
	fmt.Println("")
}

func testMapOrder() {
	fmt.Println("test map order")
	for key := range capitalAndCountryMap {
		fmt.Println("Map item: capitalAndCountryMap of", key, "is", capitalAndCountryMap[key])
	}
	fmt.Println("")
}

func testMapSlice() {
	fmt.Println("test map slice")
	// Version A
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2 * i
	}
	fmt.Printf("Version A: Value of items: %v\n", items[1:])
	fmt.Printf("Version A: len of items: %v\n", len(items[1:]))
	fmt.Printf("Version A: cap of items: %v\n", cap(items[1:]))

	// Version B
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
	fmt.Println("")
}

func testMapSort() {
	fmt.Println("test map sort")
	fmt.Println("unsorted:")
	for k, v := range unsortValMap {
		fmt.Printf("key: %v, Value: %v / ", k, v)
	}

	keys := make([]string, len((unsortValMap)))
	i := 0
	for k, _ := range unsortValMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, unsortValMap[k])
	}
	fmt.Println("")
}

func testMapInvert() {
	fmt.Println("test map invert")
	invMap := make(map[int]string, len(unsortValMap))
	for k, v := range unsortValMap {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	fmt.Println("")
}

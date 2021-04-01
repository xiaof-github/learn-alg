package main

import (
    "fmt"
    "math/rand"//伪随机
    "time"
)

func main() {
    // 随机种子
    rand.Seed(time.Now().Unix())	
    // 生成 10000 个 [0, 100000) 范围的伪随机数。
	arr := make([]int, 0, 50000)
    for i := 0; i < 10000; i++ {
        result := rand.Intn(100000)
        arr = append(arr, result)
    }
	fmt.Println(arr)
}
package main

import (
	"fmt"
)

var rst []string

func main() {
    str := "623"   
    // rst = append(rst, "abc")
    // fmt.Println(rst[0])
    fmt.Println(letterCombinations(str))
}


func letterCombinations(digits string) []string {
    if digits == "" {
        return make([]string, 0)
    }
    rst = make([]string, 0)
    mapd := map[string][]string{"2":{"a","b","c"}, "3":{"d","e","f"},"4":{"g","h","i"},"5":{"j","k","l"}, "6":{"m","n","o"}, "7":{"p","q","r","s"}, "8":{"t","u","v"},"9":{"w","x","y","z"}}      
    // fmt.Println(mapd)

    slice := make([][]string, 0)
    for i:=0;i<len(digits);i++{
        a := mapd[string(digits[i])]
        slice = append(slice, a)
    }
    // fmt.Println(slice)
    // fmt.Println(digits[2:3])
    // fmt.Println(string(digits[2]))
    // fmt.Println([]rune(digits))
    backTrace(slice, 0, "")
    return rst
}

/*
 * go语言字符串操作： https://www.jianshu.com/p/61dc40f5090d
 * 回溯法，回溯所有组合
*/
func backTrace(slice [][]string, layer int, base string) {

    if layer == len(slice) - 1 {        
        for i:=0;i<len(slice[layer]);i++{
            tmp := base + string(slice[layer][i])
            rst = append(rst, tmp)
            // fmt.Println(rst)
        }
    } else {
        for i:=0;i<len(slice[layer]);i++ {
            backTrace(slice, layer+1, base+string(slice[layer][i]))
        }
    }
}
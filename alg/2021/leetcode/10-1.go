package main

import (
    "fmt"
)

func main() {
    fmt.Println(isMatch("aab", "c*a*b"))
}

func isMatch(s string, p string) bool {
    arr1 := []byte(s)
    arr2 := []byte(p)

    dfs := make([][]int, len(arr1)+1)

    for i := 0; i < len(arr1)+1; i++ {
        tmp := make([]int, len(arr2)+1)        
        dfs[i] = tmp
    }

    return compare(arr1, arr2, dfs, 0, 0)
}

func compare(s []byte, p []byte, dfs [][]int, i int, j int) bool {
    if dfs[i][j] != 0 {
        return dfs[i][j] == 1
    }
    len1 := len(s)
    len2 := len(p)
    flag := false
    if j == len2 {
        if i == len1 {
            flag = true
        }
    }

    if j == len2-1 {
        if i < len1 {
            if s[i] == p[j] || p[j] == byte('.') {
                flag = compare(s, p, dfs, i+1, j+1)
            }
        }
    }

    if j < len2-1 {
        if i < len1 && (s[i] == p[j] || p[j] == byte('.')) {            
            if p[j+1] == byte('*') {
                flag = compare(s, p, dfs, i, j+2) || compare(s, p, dfs, i+1, j) || compare(s, p, dfs, i+1, j+2)
            } else {
                flag = compare(s, p, dfs, i+1, j+1)
            }
        } else {
            if p[j+1] == byte('*') {
                flag = compare(s, p, dfs, i, j+2)
            }
        }        
    }

    if flag == true {
        dfs[i][j] = 1
    } else {
        dfs[i][j] = -1

    }
    return flag
}

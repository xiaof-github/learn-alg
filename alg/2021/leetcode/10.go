package main

import (
	"fmt"
)

func main() {
    fmt.Println(isMatch("fs", ".*.."))
}

func isMatch(s string, p string) bool {
    arr1 := []byte(s)
    arr2 := []byte(p)    

    return compare(arr1, arr2)
}

func compare(s []byte , p []byte) bool {
    len1 := len(s)
    len2 := len(p)

    if len2 == 0 {
        return len1 == 0
    }

    if len1 >= 1 {
        matched := (s[0] == p[0] || p[0] == byte('.'))
                    
        if matched {
            if len2 >= 2 && p[1] == byte('*') {
                return compare(s, p[2:]) || compare(s[1:], p)
            } else {
                return compare(s[1:], p[1:])
            }
        } else {            
            if len2 >= 2 && p[1] == byte('*') {
                return compare(s, p[2:])
            } else {
                return false
            }
        }        
        
    } else if len2 >= 2{
        if p[1] == byte('*') {
            return compare(s, p[2:])
        }        
    }    

    return false
        
}

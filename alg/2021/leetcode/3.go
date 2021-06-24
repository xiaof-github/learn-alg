
package main

import (
    "fmt"
)

func main() {
    fmt.Println(lengthOfLongestSubstring("pewwke"))
}

// start end 滑动窗口
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0  {
        return 0
    }

    start := 0
    end:=start+1
    max := 1
    len := len(s)        
    for ;end < len; {
        // 移动end到出现重复
        
        for i:=start;i < end;i++ {
            if s[end] == s[i] {                
                if max < (end - start) {
                    max = end-start
                }
                start = i+1
                break
            }
        }
        // 移动start到出现重复的字符后一个
        end++
    }
    if max < (end - start) {
        max = end-start
    }
    return max
}

func repeat(sub string, c byte) bool {    
    for i:=0 ; i< len(sub);i++ {        
        if sub[i] == c {
            return true
        }
    }
    return false
}

// 全遍历
func lengthOfLongestSubstring1(s string) int {
    if len(s) == 0  {
        return 0
    }

    start := 0
    max := 1
    len := len(s)
    for ;len - start > max;start++ {
        for end:=start+1;end < len; end++{
            // repeat
            if repeat(s[start:end],s[end]) {
                break
            }
            tmp := end -start + 1
            if ( tmp > max) {
                max = tmp
            }
        }        
    }
    return max
}
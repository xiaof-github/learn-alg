package main

import (
	"fmt"
)

func main() {
    fmt.Println(longestPalindrome("asddddddddfffffffdsfasdcdsddssssdfccsdfsddsdfassdfsdfsdsdfsdfs"))
}

/**
 * 解题思路：
 * 
 * 1.回文字符串，首尾字母相同，字符串长度为1时，为true；字符串长度为2时，s[0] == s[1]为true
 * 2.f(i,j)表示字符串s[i,j]是否为回文字串，转移方程：  f(i,j)  => f(i+1,j-1)  && s[i] == s[j] 为 true
 *                      => 不满足上面条件为  false
 * 3.返回最大的 s[i,j]子串
**/

func longestPalindrome(s string) string {
    len := len(s)
    if len < 2 {
        return s;
    }

    // 长度为1的字串都是回文串
    var dp [][]bool    
    for i:=0;i<len;i++{
        v := make([]bool, len)
        dp = append(dp, v)
        dp[i][i] =true
    }

    maxLen := 1
    begin := 0
    arr := []byte(s)
    // 子串长度
    for L := 2;L<=len;L++ {
        // i左边界
        for i:=0;i<len;i++ {
            // j右边界
            j := L + i - 1
            if j>=len {
                break
            }
            
            if arr[i] != arr[j] {
                dp[i][j] = false
            } else {
                // 3个数以内
                if j-i < 3 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }

            // 记录最长回文字串
            if dp[i][j] && j-i+1 > maxLen {
                maxLen = j-i+1
                begin = i
            }

        }
    }
    // fmt.Println(begin, maxLen)
    // fmt.Println(arr[begin:begin+maxLen])
    return string(arr[begin:begin+maxLen])
}
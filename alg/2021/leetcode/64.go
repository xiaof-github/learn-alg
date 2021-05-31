package main

import "fmt"


func main(){
    grid := [][]int{{1,3,1},{1,5,1},{4,2,1},{5,6,8}}
    a := minPathSum(grid)
    fmt.Println("a: ", a)
}

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])	
    if m == 0 {
        return 0
    }

    tmp := make([][]int, m)
    for i:=0;i<m;i++ {
        tmp[i] = make([]int, n)
        copy(tmp[i], grid[i])
    }
    
    for i:=1;i<m;i++ {
        tmp[i][0] += tmp[i-1][0]
    }
    for i:=1;i<n;i++ {
        tmp[0][i] += tmp[0][i-1]
    }

    // 计算每一个网格的最小值，以动态规划方式计算最后的最小值

    for i:=1;i<m;i++ {
        for j:=1;j<n;j++ {
            tmp[i][j] += min(tmp[i][j-1],tmp[i-1][j])
        }
    }
    return tmp[m-1][n-1]
}

func min(a int,b int) int {
    if a <= b {
        return a
    }
    return b	
}
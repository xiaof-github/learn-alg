package main

import "log"
func main(){
    var a = []int{45,145,245,32,5,2,69,239,12,40}
    log.Print(a)
    quicksort(a)
    log.Print(a)
}

func quicksort(array []int){    
    left := 1
    right := len(array)-1
    swapedL := false
    swapedR := false
    // 轴心
    zhou := array[0]
    if(len(array) == 1){
        log.Println(array)
        return        
    }
    if (len(array) == 2) {
        if (array[0] > array[1]){
            swap(array,0,1)
        }
        log.Println(array)
        return
    }

    for ;left<right;{
        if (zhou < array[left]){            
            swapedL = true
        }else {
            left++
        }
        if (zhou > array[right]){            
            swapedR = true
        } else {
            right--
        }
        
        if (swapedL && swapedR){
            swap(array, left, right)
            left++
            right--
            swapedL = false
            swapedR = false
        }        
    }
    log.Println(left, right, array)
    if (zhou > array[right]){        
        swap(array, 0, right)
    }    
    log.Println(array)
    
    quicksort(array[0:right])
    quicksort(array[right:len(array)])
}

func swap(array []int, i int, j int){
    var tmp int
    tmp = array[i]
    array[i] = array[j]
    array[j] = tmp
}
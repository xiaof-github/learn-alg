package main

import "log"
func main(){
    var a = []int{45,145,245,32,5,2,69,239,12,40}
    quicksort(a)
    log.Print(a)
}

func quicksort(array []int){
    i, j:= 0, len(array)-1
    left := -1
    right := -1
    swaped := 0
    zhou := array[0]
    for ;i<j;{
        if (left == -1 && zhou < array[i]){
            left = i
        }
        if (right == -1 && zhou > array[j]){
            right = j
        }        
        if (left == -1){
            i++
        }
        if (right == -1){
            j--
        }
        if (left != -1 && right != -1){
            swap(array, left, right)
            left = -1
            right = -1
            swaped ++
            j--
            i++
        }        
    }
    if (swaped == 0 && right == -1){
        return
    } else {
        quicksort(array[0:j+1])
        quicksort(array[j+2:len(array)])
    }
}

func swap(array []int, i int, j int){
    var tmp int
    tmp = array[i]
    array[i] = array[j]
    array[j] = tmp
}
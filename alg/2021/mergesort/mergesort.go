package main

import (
	"log"
	"math"
) 
func main(){
    var a = []int{45,145,245,32,5,2,69,239,12,40}
    log.Print("before: ", a)
    mergesort(a)
    log.Print("after: ", a)
}
/**
 * 思路: 逐层分割为两个元素的子序列，对子序列进行排序，对排序完的子序列从下到上进行合并
 * 切分子序列：确定每层的分割序列个数，直到子序列都为1个元素为止
 * 10个元素：[[5,5]],[[2,3],[2,3]],[[1,1],[1,2],[1,1],[1,2]],[[1,1],[1],[1,1],[1,1],[1],[1,1]]
 * 19个元素：[[9,10]],[[4,5],[5,5]],[[2,2],[2,3],[2,3],[2,3]],[[1,1],[1,1],[1,1],[1,2],[1,1],[1,2],[1,1],[1,2]],[[1,1],[1,1],[1,1],[1],[1,1],[1,1],[1],[1,1],[1,1],[1],[1,1]]
 * 排序和归并：两个游标移动，对两个切片中的数据排序后，合为一个切片 
 */
func divide(arr []int) map[int][][]int {

	subPos := make(map[int][][]int)
	seqUnit := make([]int, 2)
	seqUnit[0] = len(arr)/2
	seqUnit[1] = len(arr) - len(arr)/2
	subPos[0] = append(subPos[0], seqUnit)
	lastLayerLen := 1
	layerLen := lastLayerLen*2
	layerNum := int(math.Ceil(math.Log2(float64(len(arr)))))
	log.Print("layerNum: ", layerNum)

	for i:=1;i<layerNum;i++ {						
		seq := subPos[i]
		for j:=0;j<lastLayerLen;j++{
			lastSeqUnit := subPos[i-1][j]
			var seqUnit1,seqUnit2 []int
			if (lastSeqUnit[0] == 1 && lastSeqUnit[1] == 2){
				seqUnit1 = make([]int, 1)
				seqUnit1[0] = 1
				seqUnit2 = make([]int, 2)
				seqUnit2[0] = 1
				seqUnit2[1] = 1
			} else if (lastSeqUnit[0] == 1 && lastSeqUnit[1] == 1){
				seqUnit1 = make([]int, 2)
				seqUnit1[0] = 1
				seqUnit1[1] = 1
				seqUnit2 = nil
			} else {
				seqUnit1 = make([]int, 2)
				seqUnit1[0] = lastSeqUnit[0]/2
				seqUnit1[1] = lastSeqUnit[0]-lastSeqUnit[0]/2
				seqUnit2 = make([]int, 2)
				seqUnit2[0] = lastSeqUnit[1]/2
				seqUnit2[1] = lastSeqUnit[1]-lastSeqUnit[1]/2				
			}
			seq = append(seq, seqUnit1)
			if (seqUnit2 != nil){
				seq = append(seq, seqUnit2)
			}			
		}			
		lastLayerLen = lastLayerLen*2
		layerLen = layerLen*2
		subPos[i] = seq
		log.Print("divide: ", subPos[i])
	}
	return subPos
}

func mergesort(arr []int) {
	layer := divide(arr)
	lenM := len(layer)	
	for i:=lenM-1;i>=0;i-- {
		seq := layer[i]
		base := 0
		for j:=0;j<len(seq);j++ {
			if (len(seq[j]) == 2) {				
				arrL := arr[base : base + seq[j][0]]
				arrR := arr[base + seq[j][0] : base + seq[j][0] + seq[j][1]]
				sort(arrL,arrR)
				base += seq[j][0] + seq[j][1]
			}else{
				base += 1
			}
		}
	}
	return
}

func sort(arrL []int, arrR []int) {
	tempSlice := make([]int, 0, len(arrL)+len(arrR))
	for i,j:=0,0;;{
		if ( i == len(arrL)) {
			for ;j<len(arrR);j++ {
				tempSlice = append(tempSlice, arrR[j])
			}			
			break
		}
		if ( j == len(arrR)) {			
			for ;i<len(arrL);i++ {
				tempSlice = append(tempSlice, arrL[i])
			}		
			break
		}

		if (arrL[i] <= arrR[j]){
			tempSlice = append(tempSlice, arrL[i])
			i++
		} else {
			tempSlice = append(tempSlice, arrR[j])
			j++
		}		
	}
	copy(arrL, tempSlice[0:len(arrL)])
	copy(arrR, tempSlice[len(arrL):len(arrL) + len(arrR)])
	// log.Print("tempSlice: ", tempSlice)

}


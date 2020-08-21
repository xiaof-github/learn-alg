package maina

import "fmt"

// 求两个数组的并集，去重
func main() {
	nums1 := []int{23,2,2,1,1,1,1,1,21,230,3,3,91}
	nums2 := []int{1,2,3,4,5,91,21,230,3,23,3,3,5}

	a := intersect(nums1, nums2)
	fmt.Printf("a:%+v", a)

}

func cross(n1 []int, n2 []int) []int{
	var i,j int
	var a []int
	
	for i = 0;i < len(n1);i++ {
		for j=0;j < len(n2) ;j++ {
			if (n1[i] == n2[j]){
				a = append(a, n1[i])
				break
			}
		}
		
	}
	return a
}

func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	var val []int
    for _,k := range nums1 {
		m0[k]++
	}
	for _, k:=range nums2 {
		if m0[k] > 0 {
			val = append(val, k)
			m0[k] = 0
		}
	}
	return val
}


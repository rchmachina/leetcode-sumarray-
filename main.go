package main

import "fmt"




func main(){
	angka := []int{1,3,1,4,5}
	
	fmt.Print(twoSum(angka,9))
}


func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++ {
			if (nums[i]+nums[j]== target){
				return []int{i,j}
			}
		}
	}
	return []int{}

}




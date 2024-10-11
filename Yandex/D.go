package main

// Работа со слайсами

func Mix(nums []int) []int {
	n := make([]int, len(nums))
	l := len(nums)
	for i := 0; i < l/2; i++ {
		n[2*i] = nums[i]
		n[2*i+1] = nums[l/2+i]
	}

	return n
}

func Join(nums1, nums2 []int) []int {
	n := make([]int, len(nums1)+len(nums2))
	copy(n[:len(nums1)], nums1)
	copy(n[len(nums1):], nums1)
	return n
}

func SliceCopy(nums []int) []int {
	n := make([]int, len(nums))
	copy(n, nums)

	return n
}

func Clean(nums []int, x int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != x {
			nums[j] = nums[i]
			j++
		}
	}

	return nums[:j]
}

func main() {

}

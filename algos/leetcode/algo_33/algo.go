package algo_33

func SearchAlgo33(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {
	defResult := -1
	if left > right {
		return defResult
	}

	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if target >= nums[left] && target < nums[mid] {
		return binarySearch(nums, target, left, mid-1)
	}
	return binarySearch(nums, target, mid+1, right)
}

package main

func main() {}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var mergedArray []int
	var i, j int

	for {
		if nums1[i] < nums2[j] {
			mergedArray = append(mergedArray, nums1[i])
			i++
		} else {
			mergedArray = append(mergedArray, nums2[j])
			j++
		}
		if i == m {
			return append(mergedArray, nums2[j:]...)
		}
		if j == n {
			return append(mergedArray, nums1[i:]...)
		}
	}
}

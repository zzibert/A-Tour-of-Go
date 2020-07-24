package main

type MovingAverage struct {
	queue []int
	size  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	queue := make([]int, 0)
	return MovingAverage{queue: queue, size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) < this.size {
		this.queue = append(this.queue, val)
	} else {
		this.queue = append(this.queue[1:], val)
	}

	return calculateAverage(this.queue)
}

func calculateAverage(array []int) float64 {
	sum := 0
	length := len(array)

	for _, val := range array {
		sum += val
	}

	return float64(sum) / float64(length)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

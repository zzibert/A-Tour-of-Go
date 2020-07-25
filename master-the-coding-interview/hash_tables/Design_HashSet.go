package main

type MyHashSet struct {
	buckets [][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	hashSet := MyHashSet{buckets: make([][]int, 5)}
	for i := 0; i < 5; i++ {
		hashSet.buckets[i] = make([]int, 0)
	}
	return hashSet
}

func (this *MyHashSet) hashFunction(key int) int {
	return key % 5
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		index := this.hashFunction(key)
		this.buckets[index] = append(this.buckets[index], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		index := this.hashFunction(key)
		bucket := this.buckets[index]
		var remove int

		for i, val := range bucket {
			if val == key {
				remove = i
				break
			}
		}
		if remove == len(bucket)-1 {
			this.buckets[index] = bucket[:remove]
		} else {
			this.buckets[index] = append(bucket[:remove], bucket[(remove+1):]...)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := this.hashFunction(key)
	bucket := this.buckets[index]

	for _, val := range bucket {
		if val == key {
			return true
		}
	}
	return false
}

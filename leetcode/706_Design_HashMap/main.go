package main

type MyHashMap struct {
	buckets [][][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{buckets: make([][][]int, 2069)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	bucket := this.buckets[key%2069]

	for i, keyValuePair := range bucket {
		if key == keyValuePair[0] {
			this.buckets[key%2069][i][1] = value
			return
		}
	}

	bucket = append(bucket, []int{key, value})
	this.buckets[key%2069] = bucket
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	bucket := this.buckets[key%2069]

	for _, keyValuePair := range bucket {
		if key == keyValuePair[0] {
			return keyValuePair[1]
		}
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	bucket := this.buckets[key%2069]

	for i, keyValuePair := range bucket {
		if key == keyValuePair[0] {
			if i == len(bucket)-1 {
				bucket = bucket[:i]
			} else {
				bucket = append(bucket[:i], bucket[i+1:]...)
			}
		}
	}
	this.buckets[key%2069] = bucket

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

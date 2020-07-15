package main

type MyHashMap struct {
	keyValuePairs [][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{keyValuePairs: [][]int{}}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {

	for _, keyValue := range this.keyValuePairs {
		if key == keyValue[0] {
			keyValue[1] = value
			return
		}
	}

	this.keyValuePairs = append(this.keyValuePairs, []int{key, value})

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	for _, keyValue := range this.keyValuePairs {
		if key == keyValue[0] {
			return keyValue[1]
		}
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	for i, keyValue := range this.keyValuePairs {
		if key == keyValue[0] {
			if i == len(this.keyValuePairs)-1 {
				this.keyValuePairs = this.keyValuePairs[:i]
			} else {
				this.keyValuePairs = append(this.keyValuePairs[:i], this.keyValuePairs[i+1:]...)
			}
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

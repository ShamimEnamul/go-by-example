package main

type LRUCache struct {
	capacity int
	arr      []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		arr:      []int{},
	}
}

func (this *LRUCache) Get(key int) int {
	for _, val := range this.arr {
		if key == val {

			this.arr = append(this.arr, key)
			return key
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

}
func main() {
	example(make([]string, 2, 4), "hello", 10)
}

func example(slice []string, str string, i int) {
	panic("Want stack trace")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

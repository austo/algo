package main

import "fmt"

func main() {
	h := newHashMap()
	// h.Put("amanaplanacanalpanama", 42)
	// fmt.Println(h.Get("amanaplanacanalpanama"))
	// fmt.Println(h.size)
	// h.Put("amanaplanacanalpanama", "hello")
	// fmt.Println(h.Get("amanaplanacanalpanama"))
	// fmt.Println(h.size)
	// fmt.Printf("%#v\n", h)

	names := []string{"Austin", "Yan", "Zoe", "Guifang", "Guanxing", "Ruth", "Robert"}

	for i, v := range names {
		h.Put(v, i)
	}

	fmt.Printf("size = %d\n", h.size)

	for _, v := range names {
		fmt.Printf("%s => %d\n", v, h.Get(v))
	}

	fmt.Printf("size = %d\n", h.size)

	h.ForEach(func(k string, v interface{}) {
		fmt.Printf("(ForEach) %s => %d\n", k, v)
	})

	fmt.Println(h.Keys())

	fmt.Printf("%#v\n", h)

}

const bucketSize = 2
const initialBuckets = 2

type value struct {
	key   string
	hash  int
	value interface{}
}

type hashMap struct {
	size     int
	nBuckets int
	buckets  [][]*value
}

func (h *hashMap) Put(k string, v interface{}) {
	hval := hash(k)
	b := hval % h.nBuckets
	if h.buckets[b] == nil {
		h.buckets[b] = make([]*value, 0, bucketSize)
	}

	for i := 0; i < len(h.buckets[b]); i++ {
		if h.buckets[b][i].key == k {
			h.buckets[b][i].value = v
			return
		}
	}

	if len(h.buckets[b]) >= bucketSize {
		h.resize()
		h.Put(k, v)
		return
	}

	h.buckets[b] = append(h.buckets[b], &value{
		key:   k,
		hash:  hval,
		value: v,
	})
	h.size++
}

func (h *hashMap) Get(k string) interface{} {
	hval := hash(k)
	b := hval % h.nBuckets
	if h.buckets[b] == nil {
		return nil
	}
	for i := 0; i < len(h.buckets[b]); i++ {
		if h.buckets[b][i].key == k {
			return h.buckets[b][i].value
		}
	}
	return nil
}

func (h *hashMap) ForEach(fn func(k string, v interface{})) {
	for i := 0; i < len(h.buckets); i++ {
		for j := 0; j < len(h.buckets[i]); j++ {
			v := h.buckets[i][j]
			fn(v.key, v.value)
		}
	}
}

func (h *hashMap) Keys() []string {
	keys := make([]string, 0, h.size)
	h.ForEach(func(k string, v interface{}) {
		keys = append(keys, k)
	})
	return keys
}

func (h *hashMap) resize() {
	fmt.Println("resizing...")
	oldNBuckets := h.nBuckets
	newNBuckets := oldNBuckets * 2
	newBuckets := make([][]*value, newNBuckets)
	for i := 0; i < len(h.buckets); i++ {
		for j := 0; j < len(h.buckets[i]); j++ {
			v := h.buckets[i][j]
			b := v.hash % newNBuckets
			if newBuckets[b] == nil {
				newBuckets[b] = make([]*value, 0, bucketSize)
			}
			newBuckets[b] = append(newBuckets[b], v)
		}
	}
	h.nBuckets = newNBuckets
	h.buckets = newBuckets
}

func newHashMap() *hashMap {
	h := &hashMap{
		size:     0,
		nBuckets: initialBuckets,
		buckets:  make([][]*value, initialBuckets),
	}
	return h
}

func hash(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = 31*result + int(s[i])
	}
	if result < 0 { // int overflow
		return -result
	}
	return result
}

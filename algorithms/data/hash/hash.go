package hash

import (
	"fmt"
	"hash/maphash"
)

const numBuckets = 256

type data struct {
	key   string
	value interface{}
}

type Hash struct {
	buckets [][]data
	hashmap maphash.Hash
	num     int
}

func NewHash(num int) *Hash {
	if num <= 0 {
		num = numBuckets
	}

	return &Hash{
		buckets: make([][]data, num),
		num:     num,
	}
}

// Store 保存一个键值哈希对，如果key存在则覆盖原来的值。
func (h *Hash) Store(key string, val interface{}) {
	// 找到对应的存储桶
	idx := h.hashIndex(key)
	bucket := h.buckets[idx]

	// 遍历判断是否已经存在
	for i := range bucket {
		// 如找到匹配的值，执行覆盖操作并返回
		if bucket[i].key == key {
			bucket[i].value = val
			return
		}
	}

	// 添加新元素
	h.buckets[idx] = append(h.buckets[idx], data{key: key, value: val})
}

// Delete todo.
func (h *Hash) Delete(key string) error {
	//idx := h.hashIndex(key)
	//bucket := h.buckets[idx]
	//
	//for i, item := range bucket {
	//	if item.key == key {
	//		return nil
	//	}
	//}

	return fmt.Errorf("%s is not found", key)
}

func (h *Hash) Get(key string) (interface{}, bool) {
	idx := h.hashIndex(key)
	bucket := h.buckets[idx]
	for _, item := range bucket {
		if item.key == key {
			return item.value, true
		}
	}

	return nil, false
}

// Len 统计哈希桶中的所有元素个数
func (h *Hash) Len() int {
	sum := 0
	for _, bucket := range h.buckets {
		sum += len(bucket)
	}

	return sum
}

// hashIndex 计算特定键的存储桶索引,用hashmap取余方法.
func (h *Hash) hashIndex(key string) int {
	h.hashmap.Reset()
	_, _ = h.hashmap.WriteString(key)
	sum := h.hashmap.Sum64()

	return int(sum % uint64(h.num))
}

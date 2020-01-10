package hashtable

// 本笔记只实现简单的 字符串 kv 存储  因为go没有泛型，不太好弄，whatever ...fuck it

type node struct {
	key 	string
	value 	string
}

type HashMap struct {
	buckets []interface{}
	length	int					// 哈希表元素个数
	size 	int					// 哈希表桶 buckets 容量
}

func New(size int) *HashMap{
	return &HashMap{
		buckets: nil,
		length:  0,
		size:    size,
	}
}

// TODO
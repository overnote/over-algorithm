package hashtable

// 哈希函数：通过key获取index，size是数组长度
func hash(key string, size int) int{
	hashCode := HashCode(key)				//  ，进行多项式运算
	return hashCode & (size - 1)			// 取模运算效率低：return int(hashCode) % size，这里使用位运算
}

// 计算哈希值的函数
func HashCode(key string) int{
	var hashCode rune
	unicodeArr := []rune(key)				// 得到字符串每个字符的 unicode值
	for i := 0; i < len(unicodeArr); i++ {
		// 霍纳法则多项式运算：哈希表中多使用质数 31 来进行运算，最后得到一个较大的hashCode
		hashCode = 31 * hashCode + unicodeArr[i]
	}
	return int(hashCode)
}

/**
	equals 函数：如果哈希表的key是对象，那么也需要实现一个比较key是否相等的equals函数，对对象内部成员进行比较
 */
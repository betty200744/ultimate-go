package hashtable_hashmap

/*
https://en.wikipedia.org/wiki/Hash_table
https://en.wikipedia.org/wiki/Associative_array#Hash_table_implementations


ciphers: 加密
hash: 本质是digest即摘要
hmac: 本质是digest即摘要
sign： 签名


hashtable， 就是一个存key, value的Associative array
hashtable, hashFunction, 通过hashFunction算法计算出一个index
hashtable, key ， 即index， 之后发value存入index对应的linked list
hashtable, collision resolution, 为什么存linked list, 因为计算出的index可以相同， 即

add
remove
lookup/find
change

all entries, O(n)
all keys, O(n)
all values, O(n)

size, O(1)
isEmpty, O(1)
isFull, O(1)
loadFactor, O(1)
*/

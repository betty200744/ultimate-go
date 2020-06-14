package hashtable_hashmap

/*
https://en.wikipedia.org/wiki/Hash_table
https://en.wikipedia.org/wiki/Associative_array#Hash_table_implementations

hashtable, key ， 通过hash算法计算出一个index， 之后发value存入index对应的linked list, 为什么存linked list, 因为计算出的index可以相同， 即collision resolution

hashtable， 就是一个存key, value的Associative array
hashtable, 存(key, value) pairs，or array
hashtable, The idea of hashing is to distribute the entries (key/value pairs) across an array of singly linked list
 associative array, map, symbol table, or dictionary


最好的情况, o(1)
最坏的情况，index 算出来都一样， 那么o(n)
*/

/*
	实际场景： redis ， 里面用到，因为快， HGET， HSET，HKEYS，HVALS， HEXISTS
*/
/*
Get(key interface{}) interface{}
Set(key, value interface)
Contains(key interface{})
*/

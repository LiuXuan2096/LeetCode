package Go_LeetCode

/*
设计有setAll功能的哈希表，put、get、setAll方法，时间复杂度O(1)
*/

type MyValue struct {
	value interface{}
	time  int64
}

type MyHashMap struct {
	m      map[interface{}]*MyValue
	Time   int64
	setAll *MyValue
}

func NewMyHashMap() *MyHashMap {
	return &MyHashMap{
		m:      make(map[interface{}]*MyValue),
		Time:   0,
		setAll: &MyValue{value: nil, time: -1},
	}
}

func (hm *MyHashMap) Put(key interface{}, value interface{}) {
	hm.m[key] = &MyValue{value: value, time: hm.Time}
	hm.Time++
}

func (hm *MyHashMap) SetAll(value interface{}) {
	hm.setAll = &MyValue{value: value, time: hm.Time}
	hm.Time++
}

func (hm *MyHashMap) Get(key interface{}) interface{} {
	if val, ok := hm.m[key]; ok {
		if val.time > hm.setAll.time {
			return val.value
		} else {
			return hm.setAll.value
		}
	}
	return nil
}

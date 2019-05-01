# Go 慢慢来

### slice
slice 可向后扩展 是数组的视图
len() slice总长度
cap() 不可超越数组长度
添加元素如果超越cap,系统会重新分配更大的底层数组
`s2 := make([]int,16,25)`
len=16,切片的长度
cap=25切片的容量
```
	s2 := make([]int,16,25)
	printSlice(s2)

	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3],s2[4:]...)

	printSlice(s2)

	fmt.Println("Popping from front")
	s2 = s2[1:]
	printSlice(s2)
	fmt.Println("Popping from back")
	s2 = s2[:(len(s2)-1)]
	printSlice(s2)
```

### Map
```
m := map[keyType]valueType {
    "key":"value",
    "key":"value",
    "key":"value",
    "key":"value",
}


m := map[string]string {
		"name":"dollarkiller",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	fmt.Println(m["name"])
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int // m2 = nil
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for i:=range m{
		fmt.Println(i," : ",m[i])
	}
	fmt.Println("Getting values")
	fmt.Println()
	if 	couresName,ok := m["cours e"];ok{
		fmt.Println(couresName)
	}else{
		fmt.Println("k is empty")
	}

	a,_ := returnTest(1,5)
	fmt.Println(a)
	fmt.Println("Delete values")
	delete(m,"name")
	fmt.Println(m)
```
- 使用range遍历key,或者遍历key,value
- 不保证遍历顺序,如需顺序,需手动对key排序
- 使用len获得元素个数
- 实例:寻找最长不含有重复字符的子串
```
package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i,ch:=range []byte(s){
		fmt.Println(ch)
		fmt.Println(lastOccurred)
		if lastI,ok := lastOccurred[ch];ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 >maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main()  {
	fmt.Println(lengthOfNonRepeatingSubStr("asdrgerg"))
}
```
#### rune
- 使用range遍历pos,rune对
- 使用utf8.RuneCountInString获得字符数量
- len获得字节的长度
- []byte获得字节
- []rune 获得rune 会重新开辟一个数组 放入rune类型
- 其他操作strings.func()
  - 返回在str出现的次数
  ```
    s = "12314564"
	fmt.Println(strings.Count(s,"1"))
  ```

package main

import (
	"github.com/willf/bitset"
	"fmt"
)

func main()  {
	var b bitset.BitSet // 定义一个BitSet对象
	b.Set(10).Set(11) // 给这个set新增两个值10和11
	b.Set(100)
	if b.Test(1000) { // 查看set中是否有1000这个值（我觉得Test这个名字起得是真差劲，为啥不叫Exist）
		b.Clear(1000) // 情况set
	}
	for i,e := b.NextSet(0); e; i,e = b.NextSet(i + 1) { // 遍历整个Set
		fmt.Println("b The following bit is set:", i)
	}
	
	fmt.Println("---")
	
	var c = bitset.New(100).Set(10).Set(11)
	for i, e := c.NextSet(0); e; i, e = c.NextSet(i + 1) {
		fmt.Println("c The following bit is set:", i)
	}
	
	fmt.Println("---")

	// set求交集
	if b.Intersection(c).Count() > 1 {
		fmt.Println("Intersection works.")
	}
	
	var tmpSameKeyBS *bitset.BitSet
	fmt.Println(tmpSameKeyBS)
	
}
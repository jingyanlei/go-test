package main

//17. string 与 byte slice 之间的转换

//当进行 string 和 byte slice 相互转换时，参与转换的是拷贝的原始值。这种转换的过程，与其他编程语的强制类型转换操作不同，也和新 slice 与旧 slice 共享底层数组不同。
//
//Go 在 string 与 byte slice 相互转换上优化了两点，避免了额外的内存分配：
//
//在 map[string] 中查找 key 时，使用了对应的 []byte，避免做 m[string(key)] 的内存分配
//使用 for range 迭代 string 转换为 []byte 的迭代：for i,v := range []byte(str) {...}
//雾：参考原文
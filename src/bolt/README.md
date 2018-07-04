## BoltDb

```
Bolt就是这么一个纯粹的Go语言版的嵌入式key/value的数据库，而且在Go的应用中很方便地去用作持久化。Bolt类似于LMDB，这个被认为是在现代kye/value存储中最好的。但是又不同于LevelDB，BoltDB支持完全可序列化的ACID事务，也不同于SQLlite，BoltDB没有查询语句，对于用户而言，更加易用。
BoltDB将数据保存在一个单独的内存映射的文件里。它没有wal、线程压缩和垃圾回收；它仅仅安全地处理一个文件。
```

* 安装

```
go get github.com/boltdb/bolt
```

```
Golang 优化之路——bitset
http://blog.cyeam.com/golang/2017/01/18/go-optimize-bitset
```
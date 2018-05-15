package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"time"
	"encoding/json"
)

// 打开数据库并初始化事务
func main()  {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	//打开之后，你有两种处理它的方式：读-写和只读操作，读-写方式开始于db.Update方法
	//可以看到，你传入了db.update函数一个参数，在函数内部你可以get/set数据和处理error。如果返回为nil，事务就会从数据库得到一个commit，但是如果你返回一个实际的错误，则会做回滚，你在函数中做的任何事情都不会commit到磁盘上。这很方便和自然，因为你不需要人为地去关心事务的回滚，只需要返回一个错误，其他的由Bolt去帮你完成。
	err = db.Update(func(tx *bolt.Tx) error {
		// ...read or write...
		return nil
	})
	
	//只读事务在db.View函数之中：
	//在函数中你可以读取，但是不能做修改。
	err = db.View(func(tx *bolt.Tx) error {
		//...
		return nil
	})
	
	//存储数据
	//Bolt是一个k/v的存储并提供了一个映射表，这意味着你可以通过name拿到值，就像Go原生的map，但是另外因为key是有序的，所以你可以通过key来遍历。
	//这里还有另外一层：k-v存储在bucket中，你可以将bucket当做一个key的集合或者是数据库中的表。（顺便提一句，buckets中可以包含其他的buckets，这将会相当有用）
	//我们新建了一个名为“posts”的bucket，然后将key为“2014-01-01”的vaue置为“My New Year Post”。注意bucket的名字、key和value都是bytes的slices。
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		return b.Put([]byte("2017-12-18"), []byte("My Test"))
	})
	
	//读取数据
	//你可以在Bolt的Readme里读到更多处理事务和遍历key的例子。https://github.com/boltdb/bolt#iterating-over-keys
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		v := b.Get([]byte("2017-12-18"))
		fmt.Printf("%s\n", v)
		return nil
	})
	
	//Go中的简单持久化
	//Bolt存储bytes，但是如果我们想存储一些结构体呢？这很容易通过Go的标准库实现，我们可以存储Json或是Gob编码后的结构化数据。或者，可以不限你自己使用标准库，你也可以使用Protocal Buffer或是其他的序列化方法。
	
	//我们可以先编码，例如使用Json，然后将编码后的bytes存储到BoltDB中：
	post := &Post{
		Created: time.Now(),
		Title:   "My first post",
		Content: "Hello, this is my first post.",
	}
	
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(post)
		if err != nil {
			return err
		}
		return b.Put([]byte(post.Created.Format(time.RFC3339)), encoded)
	})
	
	//当读取的时候，只需要将bytes unmarshal为结构体。
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		v := b.Get([]byte(post.Created.Format(time.RFC3339)))
		fmt.Printf("%s\n", v)
		data := map[string]interface{}{}
		err2 := json.Unmarshal(v, &data)
		if err2 != nil {
			fmt.Printf("%s\n", err2)
		}
		fmt.Printf("%s\n", data["Created"])
		fmt.Printf("%s\n", data["Title"])
		fmt.Printf("%s\n", data["Content"])
		return nil
	})
}

type Post struct {
	Created time.Time
	Title   string
	Content string
}
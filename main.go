package main

import (
	"github.com/xiazemin/BloomFilter/bloom"
	"go/src/fmt"
)

func hash1(key int)int  {
	key1:= key*33*31+key%33+key*3
	for key1<0{
		key1+=key
	}
	//fmt.Println(key1)
	return key1
}
func hash2(key int)int  {
	key1:= (key*23+79)*13+key%23
	for key1<0{
		key1+=key
	}
	//fmt.Println(key1)
	return key1
}

func main() {
	filter := bloom.NewBloom(80,2,[2]func(key int) int{hash1,hash2},[]int{34,5,99,17})
	fmt.Println(filter)
	filter.Set(12)
	filter.Set(19)
	filter.Set(12345678767532345)
	filter.Set(345689)
	filter.Set(80235674)
	fmt.Println(filter)
	fmt.Println(filter.Get(12))
	fmt.Println(filter.Get(13))
	fmt.Println(filter.Get(17))
	fmt.Println(filter.Get(19))
	fmt.Println(filter.Get(12345678767532345))
	fmt.Println(filter.Get(345689))
	fmt.Println(filter.Get(80235674))
}

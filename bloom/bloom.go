package bloom

import "go/src/fmt"

type Bloom struct {
	len int
	data []byte
	factor int //调节碰撞概率
	hash [2]func(key int) int
	seed []int
}

func NewBloom(dataSize,factor int,hash [2]func(key int) int,seed []int)*Bloom  {
    return  &Bloom{
	    len:dataSize/8,
	    factor:factor,
	    data:make([]byte,dataSize/8*factor),
	    hash:hash,
	    seed:seed,
    }
}

func (b*Bloom)fix(key int)int  {
	//2^31-1 = 2147483647
	max:=1<<31-1
	for key<0{
		//fmt.Println(key+max)
		key=(key>>2+max)
	}
	return  key%(b.len*b.factor)
}

func (b*Bloom)getValues(key int) []int {
	var values []int
	if len(b.seed)==0{
		values=[]int{
			b.fix(b.hash[0](key)),
			b.fix(b.hash[1](key)),
			b.fix(b.hash[0](key)+b.hash[1](key)),
		}
	}else{
		for _,v:=range b.seed{
			val:=b.fix(b.hash[0](key)+b.hash[1](key)*v)
			values=append(values,val)
		}
	}
	return  values
}

func (b*Bloom)Set(key int)  {
	valus:=b.getValues(key)
	for _,v:=range valus{
		index:=v/8
		fmt.Println(index)
		bit:=uint(7-v%8)
		b.data[index]|=1<<bit
	}
}

func (b*Bloom)Get(key int)bool  {
	valus:=b.getValues(key)
	r:=true
	fmt.Println(valus)
	for _,v:=range valus{
		index:=v/8
		bit:=uint(7-v%8)
                r=r&&(b.data[index]&(1<<bit)==1<<bit)
	}
	return r
}


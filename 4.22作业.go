package main

import (
	"fmt"
	"sync"
	"time"
)
//作业一：模拟工厂生产和消费者消费的场景。有一个工厂，生产日用品，
//每8ms生产一件商品，然后进入库存；另外有三个销售渠道来销售这批产品，
//三个销售渠道均是每20ms销售一件商品。当库存达到300件时，仓库满了，
//停止生产程序结束；当库存为0时，没有库存，程序也停止执行。编程模拟实现上述场景。
var a chan int=make(chan int)
var n int
var w sync.WaitGroup
var m sync.Mutex
func main() {
	n = 0
	go f1()
	w.Add(3)
	go f2()
	go f3()
	go f4()
	w.Wait()
	fmt.Println("over 剩余件数:",n)
	rs := <- a
	if rs >=300||rs<0{
		fmt.Println("子routine执行结束....")
	}
	fmt.Println("over ....")
}
func f1(){
	for {
		time.Sleep(8 * time.Millisecond)
		n++
		fmt.Printf("生产第%d件商品\n",n)
		if n>300 || n<0{
			a<-n
		}
	}
}
func f2(){
	for {
		m.Lock()
		if n>0{
			time.Sleep(20 * time.Millisecond)
			n--
			fmt.Printf("第一销售点售出一件，剩余库存%d件\n",n)//出票
		}else if n<0 || n>300{
			w.Done()
			m.Unlock()
			break
		}
		m.Unlock()
	}
}

func f3(){
	for {
		m.Lock()
		if n>0{
			time.Sleep(20 * time.Millisecond)
			n--
			fmt.Printf("第二销售点售出一件，剩余库存%d件\n",n)
		}else if n<0 || n>300{
			w.Done()
			m.Unlock()
			break
		}
		m.Unlock()
	}
}

func f4(){
for{
m.Lock()
if n > 0{
time.Sleep(20 * time.Millisecond)
n--
fmt.Printf("第三销售点售出一件，剩余库存%d件\n",n)
}else if n<0 || n>300{
w.Done()
m.Unlock()
break
}
m.Unlock()
}
}
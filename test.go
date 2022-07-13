// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	GuessNumber()
	//testTime()
	//testJson()
	//testStruct()
	//testPrint()
	//testFor()
	//testFunc()
}
func GuessNumber() {
	maxNumber := 100
	rand.Seed(time.Now().Unix())
	result := rand.Intn(maxNumber)
	fmt.Println("result=", result)

}

func testTime() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05")) //格式化输出只能用这个时间，没有为什么
	fmt.Println(now.Format("2006-01-02"))
}

type people struct {
	Name string `json:"Name"` //Struct tag 可以决定 Marshal 和 Unmarshal 函数如何序列化和反序列化数据。
	Age  int    `json:"-"`    //序列化时忽略此字段
}
type peopleTwo struct {
	Name   string `json:"Name"` //一般来说，json不区分大小写，推荐使用小写的key
	Weight int    `json:"Age"`  //通过额外的字段指定序列化时关键字
}

func (p people) say_hello() {
	fmt.Println(p.Name, "  hello")
}

func changeName(p people) {
	p.Name = "change success"
}

func changeNameTwo(p *people) {
	(*p).Name = "changeTwo success"
}

func testJson() {
	var a = people{
		Name: "Hui",
		Age:  24,
	}
	//只有 struct 中支持导出的 field 才能被 JSON package 序列化，即首字母大写的 field。
	buf, err := json.Marshal(a) //调用函数进行序列化操作，返回byte字节数组与err信息
	if err != nil {
		fmt.Printf("err=%v\n", err)
	} else {
		fmt.Println(buf)         //go语言序列化后返回的是byte字节数组
		fmt.Println(string(buf)) //转string输出正常的值
	}

	var b peopleTwo
	err = json.Unmarshal(buf, &b)
	fmt.Printf("unmarshal result,b=%+v\n", b)

}

func testPrint() {
	var a people
	a.Name = "Hui"
	a.Age = 24
	fmt.Println(a)
	fmt.Printf("%v\n", a)  //输出值
	fmt.Printf("%+v\n", a) //输出结构体字段，输出值
	fmt.Printf("%#v\n", a) //输出结构体名称，输出结构体字段，输出值
}

func testStruct() {
	a := people{
		Name: "Hui",
		Age:  20,
	}
	fmt.Println(a.Name, " Age = ", a.Age)
	var b people
	b.Name = "zhang"
	b.Age = 20
	b.say_hello()
	changeName(a)
	a.say_hello()
	changeNameTwo(&a)
	a.say_hello()
}

func testFor() {
	var a = 3
	//fmt.Println(reflect.TypeOf(a))
	sum := make([]int, 10)
	for j := 2; j < 8; j++ {
		sum[j] = j * a
		fmt.Print(sum[j], " ")
	}
	fmt.Println()

}

func testFunc() {
	sum := make([]int, 10)
	for i, num := range sum {
		val, ok := add(i, num)
		if ok == "ok" {
			fmt.Print(ok, "  i+num=", val)
		}
	}
}

func add(a int, b int) (int, string) {
	return a + b, "ok"
}

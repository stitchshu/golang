package main

type Person struct {
	Name   string
	EnName string
}

func getPersonNameFunc(p *Person) func() string {
	return func() string {
		return p.Name
	}
}

func main() {
	badBoy := &Person{Name: "法外狂徒张三", EnName: "ZhangSan"}
	goodBoy := &Person{Name: "李四", EnName: "LiSi"}

	defer func() {
		_ = badBoy
	}()
	getPersonNameFunc(goodBoy)
}

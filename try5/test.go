package main

type data struct{}

func (temp data) testValue()   {}
func (temp data) testPointer() {}

func main() {
	var i data
	(*data).testPointer(&i)
	data.testValue(i)
	data.testPointer(i)
}

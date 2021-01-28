package myLab

var Done = make(chan bool)
var Msg string
func AGoroutine()  {
	Msg = "helloworld"
	<- Done
}

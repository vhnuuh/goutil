package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"

	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}

func (log *Log) Add(s string) {
	log.msg += "\n" + s
}

func (log *Log) String() string {
	return log.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

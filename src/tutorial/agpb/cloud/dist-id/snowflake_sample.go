package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

func main() {
	//snowflake.NodeBits = 8
	//snowflake.StepBits = 14
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node: ", id.Node(),
			"step: ", id.Step(),
			"time: ", id.Time(),
			"\n")
	}

}

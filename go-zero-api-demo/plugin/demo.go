package main

import (
	"fmt"
	"plugin/tool"
)

func main() {

	plugin, err := tool.NewPlugin()
	if err != nil {
		panic(err)
	}

	if plugin.Api != nil {
		fmt.Printf("api: %+v \n", plugin.Api)
	}
	fmt.Printf("dir: %s \n", plugin.Dir)
	fmt.Println("Enjoy anything you want.")

}

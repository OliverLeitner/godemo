package main

import (
	"fmt" // import will be remembered if called!!!
	"godemo/src/libs/grabmyown"
) // we got some export and he knows that b...

func main() {
	// clear screen
	fmt.Print("\033[H\033[2J")
	// new class()...
	// var cn chucknorris.Export
	// ooptest
	// ooptest.CatchEmAll() // void funct call, no println...
	// add a newline...
	// fmt.Println()
	// return class().func...
	// fmt.Println("chuck quote: " + cn.GetChuck())
	// grab my own
	grabmyown.LoopThroughCustomers()
}

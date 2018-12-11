package main

import (
	"flag"
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tappend -data\tappend new block containing data")
	fmt.Print("\tprint\t\tprint all blockchain blocks")
}

func main() {
	printChainSet := flag.NewFlagSet("print", flag.ExitOnError)
	appendSet := flag.NewFlagSet("append", flag.ExitOnError)
	data := appendSet.String("data", "", "block data")

	if len(os.Args) == 1 {
		printUsage()
		return
	}
	switch os.Args[1] {
	case "append":
		if err := appendSet.Parse(os.Args[2:]); err == nil {
			if *data == "" {
				printUsage()
				return
			} else {
				fmt.Print(*data)
			}
		}

	case "print":
		if err := printChainSet.Parse(os.Args[2:]); err == nil {
			fmt.Print("打印所有区块链数据")
		}
	default:
		printUsage()
	}

}

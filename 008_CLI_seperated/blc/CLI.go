package blc

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tappend -data\tappend new block containing data")
	fmt.Println("\tprint\t\tprint all blockchain blocks")
	fmt.Println("\tinit -address\tinitialize blockchain with a genesis block")
	fmt.Print("\tbalanceOf -address\tquery balanceOf certain address")

}

func (cli *CLI) Run() {
	printChainSet := flag.NewFlagSet("print", flag.ExitOnError)
	appendSet := flag.NewFlagSet("append", flag.ExitOnError)
	initSet := flag.NewFlagSet("init", flag.ExitOnError)
	balanceOfSet := flag.NewFlagSet("balance", flag.ExitOnError)
	sendSet := flag.NewFlagSet("send", flag.ExitOnError)

	address := initSet.String("address", "", "miner's address")
	data := appendSet.String("data", "", "block data")
	balanceAddress := balanceOfSet.String("address", "", "who's address to check?")
	from := sendSet.String("from", "", "send from")
	to := sendSet.String("to", "", "send to")
	amount := sendSet.Int64("amount", 0, "value of transaction")
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "append":
		if !DBExists() {
			fmt.Print("数据库不存在，无法添加区块")
			os.Exit(1)
		}
		if err := appendSet.Parse(os.Args[2:]); err == nil {
			if *data == "" {
				printUsage()
				return
			} else {
				chain := BlockchainObject()
				defer chain.DB.Close()
				chain.AppendBlock([]*Transaction{})
				fmt.Println("添加区块成功！")
			}
		}
	case "print":
		if !DBExists() {
			fmt.Print("区块链不存在，无法打印。请使用init命令进行初始化")
			os.Exit(1)
		}
		if err := printChainSet.Parse(os.Args[2:]); err == nil {
			chain := BlockchainObject()
			defer chain.DB.Close()
			chain.PrintChain()
		}
	case "init":
		if DBExists() {
			fmt.Print("区块链已存在，不必初始化")
			os.Exit(1)
		} else {
			if err := initSet.Parse(os.Args[2:]); err == nil {
				if *address == "" {
					fmt.Println("未指明矿工，无法初始化")
					os.Exit(1)
				}
				chain := CreateBlockChain(*address)
				defer chain.DB.Close()
				fmt.Print("区块链初始化成功")
			}
		}

	case "balanceOf":
		if !DBExists() {
			fmt.Print("区块链不存在")
			os.Exit(1)
		} else {
			if err := balanceOfSet.Parse(os.Args[2:]); err == nil {
				if *balanceAddress == "" {
					fmt.Println("未指明查询账户")
					os.Exit(1)
				}
				chain := BlockchainObject()
				defer chain.DB.Close()

				var sum int64 = 0
				UTXOs := chain.FindUTXO(*balanceAddress)
				for _, out := range UTXOs {
					sum += out.Value
				}
				fmt.Printf("Balance Of %s is %d", *balanceAddress, sum)
			}
		}
	case "send":
		if !DBExists() {
			fmt.Print("区块链不存在")
			os.Exit(1)
		} else {
			if err := sendSet.Parse(os.Args[2:]); err == nil {
				if *from == "" || *to == "" {
					fmt.Println("未指明账户")
					os.Exit(1)
				}

				if *amount <= 0 {
					fmt.Println("转账金额不合法")
					os.Exit(1)
				}
				chain := BlockchainObject()
				defer chain.DB.Close()

				newTx := NewUTXOTransaction(*from, *to, *amount, chain)
				chain.AppendBlock([]*Transaction{newTx})
				fmt.Println("转账成功")

			}
		}

	default:
		printUsage()
	}
}

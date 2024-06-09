package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

type stack[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
}

func Stack[T any]() stack[T] {
	slice := make([]T, 0)
	return stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}

func bfc(cCtx *cli.Context) error {
	fmt.Println(cCtx.Args().Get(0))
	file, err := os.Open(cCtx.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	ram_size, err := strconv.Atoi(cCtx.Args().Get(1))
	if err != nil {
		panic(err)
	}
	stack := Stack[int]()

	var ram = make([]int, ram_size)
	var ip, dp = 0, 0
	for i := 0; i < len(data); i++ {
		switch string(data[i]) {
		case "+":
			fmt.Println("+")
			ram[dp] += 1
			ip += 1
		case "-":
			fmt.Println("-")
			ram[dp] -= 1
			ip += 1
		case ".":
			fmt.Println(string(ram[dp]))
			ip += 1
		case ",":
			fmt.Println(",")
		case "[":
			stack.Push(ip)
			ip += 1
			fmt.Println("[")
		case "]":
			ip = stack.Pop()
			ip += 1
			fmt.Println("]")
		case ">":
			ip += 1
			fmt.Println(">")
		case "<":
			ip -= 1
			fmt.Println("<")
		}
	}
	fmt.Println(ram)
	return nil
}

func main() {
	app := &cli.App{
		Name:   "bfc",
		Usage:  "bfc main.bf",
		Action: bfc,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

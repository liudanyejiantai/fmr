package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/earley/bnf"
	"io/ioutil"
	"os"
)

var inputs = []string{
	//"one", "two", "ten",
	/*
		"one plus one",
		"one plus two",
		"one plus three",
		"two plus two",
		"two plus three",
		"three plus one",
		"three plus minus two",
		"two plus two",
		"three minus two",
		"minus three minus two",
		"two times two",
		"two times three",
	*/
	"three plus three minus two",
	//"two times two plus three",
	//"minus four",
}

func main() {
	b, err := ioutil.ReadFile("arithmetic.grammar")
	if err != nil {
		glog.Fatal(err)
	}
	bnf.Debug = true
	g, err := bnf.Parse(string(b))
	if err != nil {
		glog.Fatal(err)
	}
	for _, input := range inputs {
		fmt.Println(input)
		p := bnf.NewParser(g, "number", input)
		trees := p.GetTrees()
		fmt.Printf("%+v\n", p)
		fmt.Println("tree number:", len(trees))
		for _, tree := range trees {
			tree.Print(os.Stdout)
			_, err := bnf.JsonMarshalIndent(tree, "", " ")
			if err != nil {
				glog.Fatal(err)
			}
			//fmt.Println(string(b))
		}
	}
}
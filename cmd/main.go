package main

import (
	"flag"
	"fmt"

	adp "github.com/ShinyaIshikawa/golang_design_pattern/adapter"
	bld "github.com/ShinyaIshikawa/golang_design_pattern/builder"
	fm "github.com/ShinyaIshikawa/golang_design_pattern/factory"
	itr "github.com/ShinyaIshikawa/golang_design_pattern/iterator"
	pr "github.com/ShinyaIshikawa/golang_design_pattern/prototype"
	sin "github.com/ShinyaIshikawa/golang_design_pattern/singleton"
	tm "github.com/ShinyaIshikawa/golang_design_pattern/template"
)

func main() {
	executeIterator()
	executeAdapter()
	executeTemplateMethod()
	executeFactoryMethod()
	executeSingleton()
	executePrototype()
	executeBuilder()
}

func executeIterator() {
	bookShelf := itr.BookShelf{Last: 0}
	bookShelf.AppendBook(itr.Book{Name: "Around the World in 80 days"})
	bookShelf.AppendBook(itr.Book{Name: "Bidle"})
	bookShelf.AppendBook(itr.Book{Name: "Cinderella"})
	it := bookShelf.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}

func executeAdapter() {
	pb := adp.NewPrintBanner("Black Friday SALE!")
	pb.PrintWeak()
	pb.PrintStrong()
}

func executeTemplateMethod() {
	cd := tm.NewCharDisplay("go")
	cdc := tm.NewDisplay(cd)
	cdc.Display()
	sd := tm.NewStringDisplay("gopher is nice.", 10)
	sdc := tm.NewDisplay(sd)
	sdc.Display()
}

func executeFactoryMethod() {
	//create IDCardFactory instance.
	cf := fm.NewIDCardFactory()
	//create IDCard instance.
	idc := fm.CreateInstance("ゆうきひろし", cf)
	//IDcard use.
	idc.Use()
}

func executeSingleton() {
	sin.GetInstance()
}

func executePrototype() {
	m := pr.NewManager()
	// create instance
	upen := pr.NewUnderlinePen("~")
	mbox := pr.NewMessageBox("*")
	sbox := pr.NewMessageBox("-")

	// register instance to manager
	m.Register("strong message", upen)
	m.Register("warning box", mbox)
	m.Register("slash box", sbox)

	// manager clone instance
	p1 := m.Create("strong message")
	p1.Use("Hello, world")
	p2 := m.Create("warning box")
	p2.Use("Hello, world")
	p3 := m.Create("slash box")
	p3.Use("Hello, world")
}

func executeBuilder() {
	flag.Parse()
	args := flag.Args()
	if args[0] == "plain" {
		b := bld.NewTextBuilder()
		d := bld.NewDirector(b)
		d.Construct()
		result := b.GetResult()
		fmt.Println(result)
	} else if args[0] == "html" {
		b := bld.NewHTMLBuilder()
		d := bld.NewDirector(b)
		d.Construct()
		fileNm := b.GetResult()
		fmt.Println(fileNm + " が作成されました")
	}
}

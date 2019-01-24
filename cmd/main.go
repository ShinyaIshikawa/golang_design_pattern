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
	executeTemplateMethod2()
	executeFactoryMethod()
	executeSingleton()
	executePrototype()
	executeBuilder()
}

func executeIterator() {
	start("Iteraotr")
	bookShelf := itr.BookShelf{Last: 0}
	bookShelf.AppendBook(itr.Book{Name: "Around the World in 80 days"})
	bookShelf.AppendBook(itr.Book{Name: "Bidle"})
	bookShelf.AppendBook(itr.Book{Name: "Cinderella"})
	it := bookShelf.Iterator()
	for it.HasNext() {
		// print book name
		fmt.Println(it.Next())
	}
	end("Iteraotr")
}

func executeAdapter() {
	start("Adapter")
	pb := adp.NewPrintBanner("Black Friday SALE!")
	pb.PrintWeak()
	pb.PrintStrong()
	end("Adapter")
}

func executeTemplateMethod() {
	start("Template method")
	cd := tm.NewCharDisplay("go")
	cdc := tm.NewDisplay(cd)
	cdc.Display()
	sd := tm.NewStringDisplay("gopher is nice.", 10)
	sdc := tm.NewDisplay(sd)
	sdc.Display()
	end("Template method")
}

func executeTemplateMethod2() {
	cd := tm.NewCharDisplay("James Gosling")
	cdc := tm.NewDisplay(cd)
	cdc.Display()
	sd := tm.NewStringDisplay("programming interface.", 10)
	sdc := tm.NewDisplay(sd)
	sdc.Display()
}

func executeFactoryMethod() {
	start("Factory method")
	cf := fm.NewIDCardFactory()
	idc := fm.CreateInstance("ゆうきひろし", cf)
	idc.Use()
	end("Factory method")
}

func executeSingleton() {
	start("Singleton")
	sin.GetInstance()
	end("Singleton")
}

func executePrototype() {
	start("Prototype")
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
	end("Prototype")
}

func executeBuilder() {
	start("Builder")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Builder is required parameter.")
		fmt.Println("Use plain or html.")
		return
	}
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
	end("Builder")
}

func start(name string) {
	fmt.Println("Start " + name + " pattern. ")
}

func end(name string) {
	fmt.Println("End " + name + " pattern. \n")
}

package prototype

import (
	"fmt"
)

//
type interface Product{
	createClone() Product
}

//
type struct MessageBox{
	decochar char
}

func (m MessageBoc) use(s string){
	len = len(s)
	for i=0;i < len;i++{
		fmt.Println(m.decochar)
	}
	fmt.Println("")
	fmt.Println(m.decochar)
}
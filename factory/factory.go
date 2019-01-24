package factory

import "fmt"

// InstanceFunc interface
type InstanceFunc interface {
	createProduct(owner string) Product
	registerProduct(p Product)
}

// CreateInstance use InstanceFunc
func CreateInstance(owner string, ins InstanceFunc) Product {
	p := ins.createProduct(owner)
	ins.registerProduct(p)
	return p
}

// IDCardFactory has Products.
type IDCardFactory struct {
	owners []Product
}

// NewIDCardFactory is constractor of IDCardFactory
func NewIDCardFactory() *IDCardFactory {
	return &IDCardFactory{}
}

// CreateProduct
func (icf IDCardFactory) createProduct(owner string) Product {
	return NewIDCard(owner)
}

// RegisterProduct
func (icf *IDCardFactory) registerProduct(p Product) {
	icf.owners = append(icf.owners, p)
}

// GetOwners return owners
func (icf *IDCardFactory) GetOwners() []Product {
	return icf.owners
}

// IDCard implement Product. IDCard can use.
type IDCard struct {
	owner string
}

// NewIDCard constructor
func NewIDCard(owner string) *IDCard {
	return &IDCard{owner}
}

// Use is Product's function
func (i IDCard) Use() {
	fmt.Println(i.owner + "のカードを作ります")
}

// GetOwner return owner string
func (i IDCard) GetOwner() string {
	return i.owner
}

// Product interface. Product can use.
type Product interface {
	Use()
}

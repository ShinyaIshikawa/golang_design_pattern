package factory

// Create is template
func Create(owner string, fh FactoryHelper) Product {
	p := fh.createProduct(owner)
	fh.registerProduct(p)
	return p
}

// FactoryHelper interface
type FactoryHelper interface {
	createProduct(owner string) Product
	registerProduct(p Product)
}

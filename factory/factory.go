package factory

// Create is template
func Create(owner string, fh factoryHelper) Product {
	p := fh.createProduct(owner)
	fh.registerProduct(p)
	return p
}

// FactoryHelper interface
type factoryHelper interface {
	createProduct(owner string) Product
	registerProduct(p Product)
}

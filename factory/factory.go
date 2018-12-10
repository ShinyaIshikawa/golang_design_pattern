package factory

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

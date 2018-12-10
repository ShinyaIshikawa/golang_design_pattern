package factory

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

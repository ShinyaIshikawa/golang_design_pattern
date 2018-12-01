package adapter

// PrintBanner is adpter of Banner.
type PrintBanner struct {
	banner Banner
}

// NewPrintBanner is constructer.
func NewPrintBanner(title string) *PrintBanner {
	//new
	return &PrintBanner{banner: Banner{title: title}}
}

// PrintWeak is adapter of Banner function.
func (pb PrintBanner) PrintWeak() {
	pb.banner.showWithParen()
}

// PrintStrong is adapter of Banner function.
func (pb PrintBanner) PrintStrong() {
	pb.banner.showWithAster()
}

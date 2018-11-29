package template

// DisplayController has display.
type DisplayController struct {
	display Display
}

// NewDisplayController is constructor.
func NewDisplayController(display Display) *DisplayController {
	return &DisplayController{display: display}
}

// Display function is template function
func (dc DisplayController) Display() {
	dc.display.Open()
	for i := 0; i < 5; i++ {
		dc.display.Print()
	}
	dc.display.Close()
}

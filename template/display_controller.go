package template

/*
Java言語など抽出クラスがある言語では、抽象クラスにテンプレートの関数を定義しますが
Go言語には継承がないため、コンポジション（委譲）で、テンプレート関数を共通化させました。
*/

// DisplayController has display.
type DisplayController struct {
	display Display
}

// NewDisplayController is constructor.
func NewDisplayController(display Display) *DisplayController {
	dc := new(DisplayController)
	dc.display = display
	return dc
}

// Display function is template function
func (dc DisplayController) Display() {
	dc.display.Open()
	for i := 0; i < 5; i++ {
		dc.display.Print()
	}
	dc.display.Close()
}

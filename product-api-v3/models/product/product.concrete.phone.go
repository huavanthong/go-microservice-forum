package models

type Product_phone struct {
	Product
	Model      string
	RAM        string
	ROM        string
	ScreenSize string
	Wifi       string
}

func NewProductPhone() iProduct {
	return &Product_phone{
		Product:    Product{},
		Model:      "iphone 13 128g",
		RAM:        "4GB",
		ROM:        "128GB",
		ScreenSize: "6.1inch",
		Wifi:       "Wifi 6",
	}
}

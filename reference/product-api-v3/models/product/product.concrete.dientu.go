package models

type Product_dientu struct {
	Product
	Model      string
	PortHDMI   string
	Wifi       string
	Resolution string
	ScreenSize string
}

func NewProductDienTu() iProduct {
	return &Product_dientu{
		Product:    Product{},
		Model:      "XR-55X90J",
		PortHDMI:   "Co",
		Wifi:       "Wi-Fi 802.11a/b/g/n/ac được chứng nhận",
		Resolution: "3840 x 2160 pixels",
		ScreenSize: "55 inch",
	}
}

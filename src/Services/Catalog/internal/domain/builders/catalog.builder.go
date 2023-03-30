type IShapeBuilder interface {
	SetObjectName(objectName string) IShapeBuilder
	SetOwnerName(ownerName string) IShapeBuilder
	SetDimensions(dimensions common.Dimensions) IShapeBuilder
	SetColor(color common.Color) IShapeBuilder
	SetPosition(position common.Position) IShapeBuilder
	SetBorderSize(borderSize int) IShapeBuilder
	Build() entity.Shape
}

// Mẫu thiết kế mới sử dụng kiểu trả về interface{} cho phương thức Build()
// để cho phép trả về các đối tượng hình dạng khác nhau như Circle và Rectangle.
func GetShapeBuilder(builderType string) interface{} {
	switch builderType {
	case "circle":
		return &CircleBuilder{}
	case "rectangle":
		return &RectangleBuilder{}
	default:
		return nil
	}
	return nil
}

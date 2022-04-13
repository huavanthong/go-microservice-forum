package handlers

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func (p *Products) NewProducts(l *log.Logger) *Products {
	return &Products{l}
}


package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/utils"
)

// import (
// 	"context"
// 	"net/http"

// 	"github.com/huavanthong/microservice-golang/product-api-v3/models"
// )

// // MiddlewareValidateProduct validates the product in the request and calls next if ok
// func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		rw.Header().Add("Content-Type", "application/json")

// 		prod := &data.Product{}

// 		err := data.FromJSON(prod, r.Body)
// 		if err != nil {
// 			p.l.Error("Deserializing product", "error", err)

// 			rw.WriteHeader(http.StatusBadRequest)
// 			data.ToJSON(&GenericError{Message: err.Error()}, rw)
// 			return
// 		}

// 		// validate the product
// 		errs := p.v.Validate(prod)
// 		if len(errs) != 0 {
// 			p.l.Error("Validating product", "error", errs)

// 			// return the validation messages as an array
// 			rw.WriteHeader(http.StatusUnprocessableEntity)
// 			models.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
// 			return
// 		}

// 		// add the product to the context
// 		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
// 		r = r.WithContext(ctx)

// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		next.ServeHTTP(rw, r)
// 	})
// }

func SelectProductType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Use Factory Design Pattern to get product following product type
		productType, perr := entities.GetProductType(entities.ProductType(pr.ProductType))
		if perr != nil {
			return perr
		}

		switch utils.TypeOfModel(productType) {
		case "phone":
			productPhone, _ := productType.(*entities.Product_phone)
			break
		case "dien-tu":
			productDienTu, _ := productType.(*entities.Product_dientu)
			break
		case "thoi-trang":
			productThoiTrang, _ := productType.(*entities.Product_thoitrang)
		default:
			return fmt.Errorf("Wrong product type passed")
		}

	}
}

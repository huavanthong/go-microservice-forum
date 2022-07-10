/*
 * @File: payload.ApiResponse.go
 * @Description: Defines Error information will be returned to the clients
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 * @Reference: [here](https://github.com/googlemaps/google-maps-services-java/blob/main/src/main/java/com/google/maps/GeocodingApi.jav)
 */
package payload

type ApiResponse interface {
	getResult() interface{}
	Perimeter() float64
}

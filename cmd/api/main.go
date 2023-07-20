package main

import (
	"net/http"

	"github.com/PhilipFelipe/go-intensivo-jul/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	// chi
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", Order)
	// http.ListenAndServe(":8888", r)

	// echo server http
	e := echo.New()
	e.GET("/order", Order)
	e.Logger.Fatal(e.Start(":8888"))
}

// Order endpoint for echo framework
func Order(c echo.Context) error {
	order, err := entity.NewOrder("123", 10.0, 0.5)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	order.CalculateFinalPrice()
	return c.JSON(http.StatusOK, order)
}

// func Order(w http.ResponseWriter, r *http.Request) {
// 	order, err := entity.NewOrder("123", 10.0, 0.5)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}
// 	order.CalculateFinalPrice()
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(order)
// }

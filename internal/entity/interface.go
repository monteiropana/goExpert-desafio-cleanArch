// Entidades regras de negocios
package entity

//Regras de negocios
//coracao da application

type OrderRepositoryInterface interface {
	Save(order *Order) error
	// GetTotal() (int, error)
	GetListOfOrders() ([]*Order, error)
}

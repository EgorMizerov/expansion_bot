package response

//
//import (
//	"strings"
//	"time"
//
//	"github.com/EgorMizerov/expansion_bot/domain/shift/entity"
//)
//
//type GetOrdersResponse struct {
//	Orders []Order
//}
//
//type Order struct {
//	ID            string    `json:"id"`
//	ShortID       int       `json:"short_id"`
//	AddressFrom   Address   `json:"address_from"`
//	Status        string    `json:"status"`
//	Category      string    `json:"category"`
//	PaymentMethod string    `json:"payment_method"`
//	Price         string    `json:"price"`
//	CreatedAt     time.Time `json:"created_at"`
//	BookedAt      time.Time `json:"booked_at"`
//	EndedAt       time.Time `json:"ended_at"`
//	Amenities     []string  `json:"amenities"`
//	Car           CarID     `json:"car"`
//}
//
//type Address struct {
//	Address   string  `json:"address"`
//	Latitude  float64 `json:"lat"`
//	Longitude float64 `json:"lon"`
//}
//
//type CarID struct {
//	ID string `json:"id"`
//}
//
//func (self *GetOrdersResponse) ToEntities() []*entity.Order {
//	var orders []*entity.Order
//	for _, order := range self.Orders {
//		orders = append(orders, order.ToEntity())
//	}
//	return orders
//}
//
//func (self Order) ToEntity() *entity.Order {
//	sPrice := strings.Replace(self.Price, ".", "", 1)
//	price, _ := entity.NewPrice(string([]rune(sPrice)[0 : len([]rune(sPrice))-2]))
//	return &entity.Order{
//		ID:      self.ID,
//		ShortID: self.ShortID,
//		AddressFrom: entity.Address{
//			Name:      self.AddressFrom.Address,
//			Latitude:  self.AddressFrom.Latitude,
//			Longitude: self.AddressFrom.Longitude,
//		},
//		AddressTo:     entity.Address{},
//		Status:        entity.OrderStatus(self.Status),
//		Category:      entity.OrderCategory(self.Category),
//		PaymentMethod: entity.OrderPaymentMethod(self.PaymentMethod),
//		Price:         price,
//		CarID:         self.Car.ID,
//		CreatedAt:     self.CreatedAt,
//		BookedAt:      self.BookedAt,
//		EndedAt:       self.EndedAt,
//	}
//}

package cartdto

type CreateCart struct {
	ID        int `json:"id"`
	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
	SubTotal  int `json:"Subtotal"`
	Shipping  int `json:"shipping"`
	Total     int `json:"total"`
	Status    int `json:"Status"`
}

type UpdateCart struct {
	ID        int `json:"id"`
	SubTotal  int `json:"Subtotal"`
	Shipping  int `json:"shipping"`
	Total     int `json:"total"`
	Status    int `json:"Status"`
}
type cartResponse struct{
	ID        int `json:"id"`
	SubTotal  int `json:"Subtotal"`
	Shipping  int `json:"shipping"`
	Total     int `json:"total"`
}
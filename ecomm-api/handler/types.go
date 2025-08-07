package handler

import "time"

type ProductReq struct {
	ID           int64      `json:"id"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
	Rating       int64     `json:"rating"`
	NumReviews   int64     `json:"num_reviews"`
	Price        float32 `json:"price"`
	CountInStock int64     `json:"count_in_stock"`
}

type ProductRes struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Image        string     `json:"image"`
	Category     string     `json:"category"`
	Description  string     `json:"description"`
	Rating       int64        `json:"rating"`
	NumReviews   int64        `json:"num_reviews"`
	Price        float32    `json:"price"`
	CountInStock int64        `json:"count_in_stock"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

//* ORDERS
type OrderReq struct {
	Items         []*OrderItem `json:"items"`
	PaymentMethod string      `json:"payment_method"`
	TaxPrice      float32     `json:"tax_price"`
	ShippingPrice float32     `json:"shipping_price"`
	TotalPrice    float32     `json:"total_price"`
}

type OrderItem struct {
	Name      string  `json:"name"`
	Quantity  int64     `json:"quantity"`
	Image     string  `json:"image"`
	Price     float32 `json:"price"`
	ProductID int64   `json:"product_id"`
}

type OrderRes struct {
	ID            int64       `json:"id"`
	Items         []*OrderItem `json:"items"`
	PaymentMethod string      `json:"payment_method"`
	TaxPrice      float32     `json:"tax_price"`
	ShippingPrice float32     `json:"shipping_price"`
	TotalPrice    float32     `json:"total_price"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at"`
}

//* USERS

type UserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserRes struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

type ListUserRes struct {
	Users []UserRes `json:"users"`
}

//! AUTh USERS TYPE

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	SessionID   string  `json:"session_id"`
	AccessToken string  `json:"access_token"` 
	RefreshToken string  `json:"refresh_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User        UserRes `json:"user"`
}
/*
* AccessToken - токен для доступа к ресурсам
* RefreshToken - токен для обновления сессии
* SessionID - идентификатор сессии
* AccessTokenExpiresAt - токен доступа чей срок истекает
* RefreshTokenExpiresAt - токен обновления чей срок истекает
*/

//* SESSIONS
// * структура для запроса токена обновления доступа 
//* принимает только токен обновления в качестве запроса  и возвращает токен доступа
type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

// * структуру ответа токена  обновления доступа 
// * где вернем токен доступа
// * также срок действия токена доступа истекает 
type RenewAccessTokenRes struct {
	AccessToken string `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

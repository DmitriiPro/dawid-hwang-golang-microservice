package handler

import (
	"context"
	// "davidHwang/ecomm/ecomm-api/server"
	"davidHwang/ecomm/ecomm-grpc/pb"

	// "davidHwang/ecomm/ecomm-api/storer"
	"davidHwang/ecomm/token"
	"davidHwang/ecomm/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// * rest api
// type handler struct {
// 	ctx        context.Context
// 	server     *server.Server
// 	TokenMaker *token.JWTMaker
// }

// func NewHandler(server *server.Server, secretKey string) *handler {
// 	return &handler{
// 		ctx:        context.Background(),
// 		server:     server,
// 		TokenMaker: token.NewJWTMaker(secretKey),
// 	}
// }

// ! GRPC CLIENT
type handler struct {
	ctx        context.Context
	client     pb.EcommClient
	TokenMaker *token.JWTMaker
}

func NewHandler(client pb.EcommClient, secretKey string) *handler {
	return &handler{
		ctx:        context.Background(),
		client:     client,
		TokenMaker: token.NewJWTMaker(secretKey),
	}
}

// ! GRPC CLIENT

//* PRODUCTS

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p ProductReq

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	product, err := h.client.CreateProduct(h.ctx, toPBProductReq(p))

	if err != nil {
		http.Error(w, "error creating product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// product/{id}
func (h *handler) getProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	product, err := h.client.GetProduct(h.ctx, &pb.ProductReq{Id: i})

	if err != nil {

		http.Error(w, "error getting product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	lpr, err := h.client.ListProducts(h.ctx, &pb.ProductReq{})

	if err != nil {
		http.Error(w, "error listing products", http.StatusInternalServerError)
		return
	}

	var res []ProductRes

	for _, p := range lpr.GetProducts() {
		res = append(res, toProductRes(p))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	var p ProductReq
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	p.ID = i

	// product, err := h.client.GetProduct(h.ctx, &pb.ProductReq{Id: i})

	// if err != nil {
	// 	http.Error(w, "error getting product", http.StatusInternalServerError)
	// 	return
	// }

	// //* patch our product request
	// patchProductReq(product, p)
	updatedProduct, err := h.client.UpdateProduct(h.ctx, toPBProductReq(p))

	if err != nil {
		http.Error(w, "error updating product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(updatedProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}
	_, err = h.client.DeleteProduct(h.ctx, &pb.ProductReq{Id: i})
	if err != nil {
		http.Error(w, "error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// //////////////
// func toStoreProduct(p ProductReq) *storer.Product {
// 	return &storer.Product{
// 		Name:         p.Name,
// 		Image:        p.Image,
// 		Category:     p.Category,
// 		Description:  p.Description,
// 		Rating:       p.Rating,
// 		NumReviews:   p.NumReviews,
// 		Price:        p.Price,
// 		CountInStock: p.CountInStock,
// 	}
// }

// func toProductRes(p *storer.Product) ProductRes {
// 	return ProductRes{
// 		ID:           p.ID,
// 		Name:         p.Name,
// 		Image:        p.Image,
// 		Category:     p.Category,
// 		Description:  p.Description,
// 		Rating:       p.Rating,
// 		NumReviews:   p.NumReviews,
// 		Price:        p.Price,
// 		CountInStock: p.CountInStock,
// 		CreatedAt:    p.CreatedAt,
// 		UpdatedAt:    p.UpdatedAt,
// 	}
// }

// func patchProductReq(product *storer.Product, p ProductReq) {
// 	if p.Name != "" {
// 		product.Name = p.Name
// 	}

// 	if p.Image != "" {
// 		product.Image = p.Image
// 	}

// 	if p.Category != "" {
// 		product.Category = p.Category
// 	}

// 	if p.Description != "" {
// 		product.Description = p.Description
// 	}

// 	if p.Rating != 0 {
// 		product.Rating = p.Rating
// 	}

// 	if p.NumReviews != 0 {
// 		product.NumReviews = p.NumReviews
// 	}

// 	if p.Price != 0.0 {
// 		product.Price = p.Price
// 	}

// 	if p.CountInStock != 0 {
// 		product.CountInStock = p.CountInStock
// 	}

// 	product.UpdatedAt = toTimePtr(time.Now())

// }

// func toTimePtr(t time.Time) *time.Time {
// 	return &t
// }

/////////////////////////////

// * ORDERS
func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o OrderReq

	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	// so := toStoreOrder(o)
	// so.UserID = claims.ID
	po := toPBOrderReq(o)
	po.UserId = claims.ID

	created, err := h.client.CreateOrder(h.ctx, po)

	if err != nil {
		http.Error(w, "HANDLER - CreateOrder: error creating order", http.StatusInternalServerError)
		return
	}

	res := toOrderRes(created)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// get order
func (h *handler) getOrder(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	// id := chi.URLParam(r, "id")
	// i, err := strconv.ParseInt(id, 10, 64)

	// if err != nil {
	// 	panic(err)
	// }

	order, err := h.client.GetOrder(h.ctx, &pb.OrderReq{UserId: claims.ID})

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := toOrderRes(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// list orders
func (h *handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.client.ListOrders(h.ctx, &pb.OrderReq{})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	var res []OrderRes
	for _, o := range orders.Orders {
		res = append(res, toOrderRes(o))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// delete

func (h *handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		panic(err)
	}

	_, err = h.client.DeleteOrder(h.ctx, &pb.OrderReq{Id: i})

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// func toStoreOrder(o OrderReq) *storer.Order {
// 	return &storer.Order{
// 		PaymentMethod: o.PaymentMethod,
// 		TaxPrice:      o.TaxPrice,
// 		ShippingPrice: o.ShippingPrice,
// 		TotalPrice:    o.TotalPrice,
// 		Items:         toStoreOrderItems(o.Items),
// 	}
// }

// func toStoreOrderItems(items []OrderItem) []storer.OrderItem {
// 	var res []storer.OrderItem

// 	for _, i := range items {
// 		res = append(res, storer.OrderItem{
// 			Name:      i.Name,
// 			Quantity:  i.Quantity,
// 			Image:     i.Image,
// 			Price:     i.Price,
// 			ProductID: i.ProductID,
// 		})
// 	}
// 	return res
// }

// func toOrderRes(o *storer.Order) OrderRes {
// 	return OrderRes{
// 		ID:            o.ID,
// 		Items:         toOrderItems(o.Items),
// 		PaymentMethod: o.PaymentMethod,
// 		TaxPrice:      o.TaxPrice,
// 		ShippingPrice: o.ShippingPrice,
// 		TotalPrice:    o.TotalPrice,
// 		CreatedAt:     o.CreatedAt,
// 		UpdatedAt:     o.UpdatedAt,
// 	}
// }

// items
// func toOrderItems(items []storer.OrderItem) []OrderItem {
// 	var res []OrderItem

// 	for _, i := range items {
// 		res = append(res, OrderItem{
// 			Name:      i.Name,
// 			Quantity:  i.Quantity,
// 			Image:     i.Image,
// 			Price:     i.Price,
// 			ProductID: i.ProductID,
// 		})
// 	}

// 	return res
// }

//* USERS

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u UserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	//* hash password
	hashedPass, err := util.HashPassword(u.Password)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}
	u.Password = hashedPass
	//* hash password end

	created, err := h.client.CreateUser(h.ctx, toPBUserReq(u))

	if err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	res := toUserRes(created)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.client.ListUsers(h.ctx, &pb.UserReq{})

	if err != nil {
		http.Error(w, "error listing users", http.StatusInternalServerError)
		return
	}

	var res ListUserRes
	for _, u := range users.GetUsers() {
		res.Users = append(res.Users, toUserRes(u))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u UserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	claims := r.Context().Value(authKey{}).(*token.UserClaims)
	u.Email =  claims.Email

	// user, err := h.client.GetUser(h.ctx, &pb.UserReq{Id: claims.ID})

	// if err != nil {
	// 	http.Error(w, "error getting user", http.StatusInternalServerError)
	// 	return
	// }

	// //* patch our user request
	// patchUserReq(user, u)

	// if user.Email == "" {
	// 	user.Email = claims.Email
	// }

	updatedUser, err := h.client.UpdateUser(h.ctx, toPBUserReq(u))

	if err != nil {
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}

	res := toUserRes(updatedUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}
	_, err = h.client.DeleteUser(h.ctx, &pb.UserReq{Id: i})
	if err != nil {
		http.Error(w, "error deleting user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// func patchUserReq(user *storer.User, u UserReq) {
// 	if u.Name != "" {
// 		user.Name = u.Name
// 	}

// 	if u.Email != "" {
// 		user.Email = u.Email
// 	}

// 	if u.Password != "" {
// 		hashed, err := util.HashPassword(u.Password)
// 		if err != nil {
// 			panic(err)
// 		}

// 		user.Password = hashed
// 	}

// 	if u.IsAdmin {
// 		user.IsAdmin = u.IsAdmin
// 	}

// 	user.UpdatedAt = handlerToTimePtr(time.Now())

// }

// func handlerToTimePtr(t time.Time) *time.Time {
// 	return &t
// }

// func toUserRes(user *storer.User) UserRes {
// 	return UserRes{
// 		Name:    user.Name,
// 		Email:   user.Email,
// 		IsAdmin: user.IsAdmin,
// 	}
// }

// func toStoreUser(userReq UserReq) *storer.User {
// 	return &storer.User{
// 		Name:     userReq.Name,
// 		Email:    userReq.Email,
// 		Password: userReq.Password,
// 		IsAdmin:  userReq.IsAdmin,
// 	}
// }

//! AUTH USERS

func (h *handler) loginUser(w http.ResponseWriter, r *http.Request) {

	var u LoginUserReq

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	gu, err := h.client.GetUser(h.ctx, &pb.UserReq{Email: u.Email})

	if err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}

	err = util.CheckPassword(u.Password, gu.GetPassword())

	if err != nil {
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}

	// * если пароль верный мы можем создать токен и вернуть в качестве ответа
	//* json web token (jwt)

	accessToken, accessClaims, err := h.TokenMaker.CreateToken(gu.GetId(), gu.GetEmail(), gu.GetIsAdmin(), time.Minute*15)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	//* метод для создания токена обновления доступа
	refreshToken, refreshClaims, err := h.TokenMaker.CreateToken(gu.GetId(), gu.GetEmail(), gu.GetIsAdmin(), time.Hour*24)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	//* создать сессию для хранения токена обновления в базе данных
	session, err := h.client.CreateSession(h.ctx, &pb.SessionReq{
		Id:           refreshClaims.RegisteredClaims.ID,
		UserEmail:    gu.GetEmail(),
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    timestamppb.New(refreshClaims.RegisteredClaims.ExpiresAt.Time),
	})

	if err != nil {
		log.Printf("Error creating session: %v", err)
		http.Error(w, "error creating session", http.StatusInternalServerError)
		return
	}

	res := LoginUserRes{
		SessionID:             session.GetId(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
		User:                  toUserRes(gu),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

// * бработчик выхода из системы
func (h *handler) logoutUser(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	// id := chi.URLParam(r, "id")
	// if id == "" {
	// 	http.Error(w, "missing session ID", http.StatusBadRequest)
	// 	return
	// }

	_, err := h.client.DeleteSession(h.ctx, &pb.SessionReq{Id: claims.RegisteredClaims.ID})
	if err != nil {
		http.Error(w, "error deleting session", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//* обновления токена доступа

func (h *handler) renewAccessToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	//* для проверки токена обновления
	refreshClaims, err := h.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "error verifying token", http.StatusUnauthorized)
		return
	}

	//* получим сессии из базы данных
	session, err := h.client.GetSession(h.ctx, &pb.SessionReq{Id: refreshClaims.RegisteredClaims.ID})

	if err != nil {
		http.Error(w, "error getting session", http.StatusInternalServerError)
		return
	}

	//* проверим не отозвана ли эта сессия
	if session.IsRevoked {
		http.Error(w, "session revoked", http.StatusUnauthorized)
		return
	}

	//* проверка совпадает ли адрес электронной почты с адресом в refreshClaims.Email
	if session.GetUserEmail() != refreshClaims.Email {
		http.Error(w, "session revoked", http.StatusUnauthorized)
		return
	}

	//* получим данные для токена доступа и ошибку
	accessToken, accessClaims, err := h.TokenMaker.CreateToken(refreshClaims.ID, refreshClaims.Email, refreshClaims.IsAdmin, time.Minute*15)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}
	//* создадим ответ с токеном доступа
	res := RenewAccessTokenRes{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

// * отзыв токена доступа - отмена сессии
func (h *handler) revokeSession(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	// id := chi.URLParam(r, "id")
	// if id == "" {
	// 	http.Error(w, "missing session ID", http.StatusBadRequest)
	// 	return
	// }

	_, err := h.client.RevokeSession(h.ctx, &pb.SessionReq{Id: claims.RegisteredClaims.ID})
	if err != nil {
		http.Error(w, "error revoking session", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

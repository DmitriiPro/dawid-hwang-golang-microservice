package handler

// import (
// 	"context"
// 	"davidHwang/ecomm/ecomm-grpc/pb"
// 	"davidHwang/ecomm/token"
// )

// ! GRPC CLIENT
// type handlerGRPC struct {
// 	ctx        context.Context
// 	client     pb.EcommClient
// 	TokenMaker *token.JWTMaker
// }

// func NewHandlerGRPC(client pb.EcommClient, secretKey string) *handlerGRPC {
// 	return &handlerGRPC{
// 		ctx:        context.Background(),
// 		client:     client,
// 		TokenMaker: token.NewJWTMaker(secretKey),
// 	}
// }

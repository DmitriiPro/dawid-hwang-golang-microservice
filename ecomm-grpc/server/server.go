package server

import (
	"context"
	"davidHwang/ecomm/ecomm-grpc/pb"
	"davidHwang/ecomm/ecomm-grpc/storer"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	storer *storer.MySQLStorer
	pb.UnimplementedEcommServer
}

func NewServer(storer *storer.MySQLStorer) *Server {
	return &Server{storer: storer}
}

// * PRODUCTS
func (s *Server) CreateProduct(ctx context.Context, req *pb.ProductReq) (*pb.ProductRes, error) {

	pr, err := s.storer.CreateProduct(ctx, toStorerProduct(req))
	if err != nil {
		return nil, err
	}

	return toPBProductRes(pr), nil
}

func (s *Server) GetProduct(ctx context.Context, p *pb.ProductReq) (*pb.ProductRes, error) {
	pr, err := s.storer.GetProduct(ctx, p.GetId())

	if err != nil {
		return nil, err
	}

	return toPBProductRes(pr), nil
}

func (s *Server) ListProducts(ctx context.Context, p *pb.ProductReq) (*pb.ListProductRes, error) {
	prs, err := s.storer.ListProducts(ctx)

	if err != nil {
		return nil, err
	}
	var lpr []*pb.ProductRes

	for _, lp := range prs {
		lpr = append(lpr, toPBProductRes(lp))
	}

	return &pb.ListProductRes{Products: lpr}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, p *pb.ProductReq) (*pb.ProductRes, error) {

	product, err := s.storer.GetProduct(ctx, p.GetId())

	if err != nil {
		return nil, err
	}

	patchProductReq(product, p)

	pr, err := s.storer.UpdateProduct(ctx, product)

	if err != nil {
		return nil, err
	}

	return toPBProductRes(pr), nil
}

func (s *Server) DeleteProduct(ctx context.Context, p *pb.ProductReq) (*pb.ProductRes, error) {
	err := s.storer.DeleteProduct(ctx, p.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.ProductRes{}, nil
}

// * ORDERS
func (s *Server) CreateOrder(ctx context.Context, o *pb.OrderReq) (*pb.OrderRes, error) {
	or, err := s.storer.CreateOrder(ctx, toStorerOrder(o))

	if err != nil {
		return nil, err
	}

	return toPBOrderRes(or), nil
}

func (s *Server) GetOrder(ctx context.Context, o *pb.OrderReq) (*pb.OrderRes, error) {
	or, err := s.storer.GetOrder(ctx, o.GetUserId())

	if err != nil {
		return nil, err
	}

	return toPBOrderRes(or), nil
}

func (s *Server) ListOrders(ctx context.Context, o *pb.OrderReq) (*pb.ListOrderRes, error) {
	orders, err := s.storer.ListOrders(ctx)

	if err != nil {
		return nil, err
	}

	var lor []*pb.OrderRes

	for _, order := range orders {
		lor = append(lor, toPBOrderRes(order))
	}

	return &pb.ListOrderRes{Orders: lor}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, o *pb.OrderReq) (*pb.OrderRes, error) {
	err := s.storer.DeleteOrder(ctx, o.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.OrderRes{}, nil
}

// * USERS
func (s *Server) CreateUser(ctx context.Context, u *pb.UserReq) (*pb.UserRes, error) {
	usr, err := s.storer.CreateUser(ctx, toStorerUser(u))

	if err != nil {
		return nil, err
	}

	return toPBUserRes(usr), nil
}

func (s *Server) GetUser(ctx context.Context, u *pb.UserReq) (*pb.UserRes, error) {
	usr, err := s.storer.GetUser(ctx, u.GetEmail())

	if err != nil {
		return nil, err
	}

	return toPBUserRes(usr), nil
}

func (s *Server) ListUsers(ctx context.Context, u *pb.UserReq) (*pb.ListUserRes, error) {
	users, err := s.storer.ListUsers(ctx)

	if err != nil {
		return nil, err
	}
	var lu []*pb.UserRes

	for _, user := range users {
		lu = append(lu, toPBUserRes(user))
	}

	return &pb.ListUserRes{Users: lu}, nil
}

func (s *Server) UpdateUser(ctx context.Context, u *pb.UserReq) (*pb.UserRes, error) {
	user, err := s.storer.GetUser(ctx, u.GetEmail())

	if err != nil {
		return nil, err
	}

	patchUserReq(user, u)

	usr, err := s.storer.UpdateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	return toPBUserRes(usr), nil
}

func (s *Server) DeleteUser(ctx context.Context, u *pb.UserReq) (*pb.UserRes, error) {
	err := s.storer.DeleteUser(ctx, u.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.UserRes{}, nil
}

//* SESSIONS

func (s *Server) CreateSession(ctx context.Context, sr *pb.SessionReq) (*pb.SessionRes, error) {
	session, err := s.storer.CreateSession(ctx, &storer.Session{
		ID:           sr.GetId(),
		UserEmail:    sr.GetUserEmail(),
		RefreshToken: sr.GetRefreshToken(),
		IsRevoked:    sr.GetIsRevoked(),
		ExpiresAt:    toTimePtr(sr.GetExpiresAt().AsTime()),
	})

	if err != nil {
		return nil, err
	}

	return &pb.SessionRes{
		Id:           session.ID,
		UserEmail:    session.UserEmail,
		RefreshToken: session.RefreshToken,
		IsRevoked:    session.IsRevoked,
		ExpiresAt:    timestamppb.New(*session.ExpiresAt),
	}, nil
}

func (s *Server) GetSession(ctx context.Context, sr *pb.SessionReq) (*pb.SessionRes, error) {
	session, err := s.storer.GetSession(ctx, sr.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.SessionRes{
		Id:           session.ID,
		UserEmail:    session.UserEmail,
		RefreshToken: session.RefreshToken,
		IsRevoked:    session.IsRevoked,
		ExpiresAt:    timestamppb.New(*session.ExpiresAt),
	}, nil
}

func (s *Server) RevokeSession(ctx context.Context, sr *pb.SessionReq) (*pb.SessionRes, error) {
	err := s.storer.RevokeSession(ctx, sr.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.SessionRes{}, nil
}


func (s *Server) DeleteSession(ctx context.Context, sr *pb.SessionReq) (*pb.SessionRes, error) {
	err := s.storer.DeleteSession(ctx, sr.GetId())

	if err != nil {
		return nil, err
	}

	return &pb.SessionRes{}, nil
}

//* 21 : 42
//* https://www.youtube.com/watch?v=D1a7ny_imUw

package grpcapi

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	_ "log/slog"
	pb "preproj/internal/handler/grpcapi/gen/product"
	"preproj/internal/models"
	"preproj/internal/service"
	"time"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	Services *service.Service
}

func (p *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	product := models.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		UserID:      req.GetUserid(),
	}

	if product.Name == "" || product.Description == "" || product.Price == 0.0 || product.UserID == 0 {
		slog.Error("Bad request data")
		return nil, status.Errorf(codes.InvalidArgument, "name or description or price or userid are empty")
	}

	productID, err := p.Services.Product.Create(ctx, &product)
	if err != nil {
		slog.Error("failed to create product", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed create product: %v", err)
	}
	return &pb.CreateProductResponse{
		Id: productID,
	}, nil
}

func (p *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	productID := req.GetId()

	product, err := p.Services.Product.GetByID(ctx, productID)
	if err != nil {
		slog.Error("failed to get product", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			UserId:      product.UserID,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		},
	}, nil
}
func (p *ProductService) GetAllProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	products, err := p.Services.Product.GetAll(ctx)
	if err != nil {
		slog.Error("failed to get all products", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get products: %v", err)
	}

	var pbProducts []*pb.Product

	for _, product := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			UserId:      product.UserID,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		})
	}
	if len(pbProducts) == 0 {
		slog.Info("all products array is empty")
		return &pb.GetProductsResponse{Products: pbProducts}, nil
	}
	return &pb.GetProductsResponse{Products: pbProducts}, nil
}
func (p *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	productID := req.GetId()

	err := p.Services.Product.Delete(ctx, productID)
	if err != nil {
		slog.Error("failed to delete product", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to delete product: %v", err)
	}
	return &pb.DeleteProductResponse{}, nil
}
func (p *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	pbProduct := req.GetProduct()

	product := models.Product{
		Name:        pbProduct.Name,
		Description: pbProduct.Description,
		Price:       pbProduct.Price,
		UserID:      pbProduct.UserId,
	}

	productID, err := p.Services.Product.Update(ctx, product)
	if err != nil {
		slog.Error("failed to update product", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}

	return &pb.UpdateProductResponse{
		Id: productID,
	}, nil
}

func (p *ProductService) GetAllByUserID(ctx context.Context, req *pb.GetAllProductsByUserIDRequest) (*pb.GetAllProductsByUserIDResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	userID := req.GetUserId()
	products, err := p.Services.Product.GetAllByUserID(ctx, userID)
	if err != nil {
		slog.Error("failed to get all products by userID", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get products: %v", err)
	}

	var pbProducts []*pb.Product

	for _, product := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			UserId:      product.UserID,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		})
	}
	if len(pbProducts) == 0 {
		slog.Info("all products by userID array is empty")
		return &pb.GetAllProductsByUserIDResponse{Products: pbProducts}, nil
	}
	return &pb.GetAllProductsByUserIDResponse{Products: pbProducts}, nil

}

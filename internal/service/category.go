package service

import (
	"context"
	"io"

	"github.com/giovane-aG/go_grpc/internal/database"
	"github.com/giovane-aG/go_grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, input *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var pbCategories []*pb.Category

	for _, value := range categories {
		pbCategories = append(pbCategories, &pb.Category{Id: value.ID, Name: value.Name, Description: value.Description})
	}

	return &pb.CategoryList{Categories: pbCategories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, input *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindById(input.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := pb.CategoryList{}

	for {
		categoryRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&categories)
		}
		if err != nil {
			return err
		}
		categoryResult, err := c.CategoryDB.Create(categoryRequest.Name, categoryRequest.Description)
		if err != nil {
			return err
		}
		newCategory := &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		}
		categories.Categories = append(categories.Categories, newCategory)
	}
}

package service

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}
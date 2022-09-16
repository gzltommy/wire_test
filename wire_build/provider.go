package wire_build

// repo

// IPostRepo IPostRepo
type IPostRepo interface{}

// NewPostRepo NewPostRepo
func NewPostRepo() IPostRepo {
	return new(IPostRepo)
}

// use case

// IPostUseCase IPostUseCase
type IPostUseCase interface{}
type postUseCase struct {
	repo IPostRepo
}

// NewPostUseCase
func NewPostUseCase(repo IPostRepo) IPostUseCase {
	return postUseCase{repo: repo}
}

// service

// PostService
type PostService struct {
	useCase IPostUseCase
}

// NewPostService
func NewPostService(u IPostUseCase) *PostService {
	return &PostService{useCase: u}
}

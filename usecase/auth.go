package usecase


type AuthUsecase struct {
}

func NewAuthUsecase(starStorage StarStorage) *StarUsecase {
	return &StarUsecase{
		starStorage: starStorage,
	}
}

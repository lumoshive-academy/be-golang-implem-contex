package adaptor

import (
	"net/http"
	"session-26/internal/usecase"
)

type UserAdaptor struct {
	UserUseCase usecase.UserUseCaseInterface
}

func NewUserAdaptor(userUseCase usecase.UserUseCaseInterface) *UserAdaptor {
	return &UserAdaptor{UserUseCase: userUseCase}
}

func (userAdaptor *UserAdaptor) Get(w http.ResponseWriter, r *http.Request) {

}

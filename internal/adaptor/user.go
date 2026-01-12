package adaptor

import (
	"net/http"
	"session-26/internal/usecase"
	"session-26/pkg/utils"
)

type UserAdaptor struct {
	UserUseCase usecase.UserUseCaseInterface
}

func NewUserAdaptor(userUseCase usecase.UserUseCaseInterface) *UserAdaptor {
	return &UserAdaptor{UserUseCase: userUseCase}
}

func (userAdaptor *UserAdaptor) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := userAdaptor.UserUseCase.GetUser(ctx)
	if err != nil {
		// return error
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// return success
	utils.ResponseSuccess(w, http.StatusOK, "success get data", users)
}

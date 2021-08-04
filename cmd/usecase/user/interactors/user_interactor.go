package interactors

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/values"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/inputs"
)

type UserInteractor struct {
	userRepository repositories.UserRepository
}

// NewUserInteractor コンストラクタ
func NewUserInteractor(userRepository repositories.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

// GetUser ユーザを取得します．
func (ui *UserInteractor) GetUser(input *inputs.GetUserInput) (*entities.User, error) {
	return ui.userRepository.FindById(ids.UserId(input.UserId))
}

// CreateUser ユーザを作成します．
func (ui *UserInteractor) CreateUser(cui *inputs.CreateUserInput) (*entities.User, error) {
	user := entities.NewUser(
		ids.UserId(cui.UserId),
		values.NewUserName(cui.UserLastName, cui.UserFirstName, cui.UserLastKanaName, cui.UserFirstKanaName),
		values.UserGenderType(cui.UserGenderType),
	)

	return ui.userRepository.Create(user)
}

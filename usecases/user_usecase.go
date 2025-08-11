package usecases

import (
	"Blog/domain"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(user *domain.User) (primitive.ObjectID, error) {
	return u.repo.Create(user)
}

func (u *UserUsecase) Authenticate(usernameOrEmail, password string) (*domain.User, error) {
	return u.repo.Authenticate(usernameOrEmail, password)
}

func (u *UserUsecase) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	return u.repo.GetAll()
}

func (u *UserUsecase) PromoteUser(userID primitive.ObjectID, newRole string) error {
	if newRole == "" {
		return errors.New("new role must be provided")
	}
	return u.repo.PromoteUser(userID, newRole)
}

func (u *UserUsecase) DeleteUserByID(id primitive.ObjectID) error {
	return u.repo.DeleteByID(id)
}

func (u *UserUsecase) ChangePassword(id primitive.ObjectID, oldPassword string, newPassword string) error {
	return u.repo.ChangePassword(id, oldPassword, newPassword)
}

func (u *UserUsecase) ForgotPassword(email string) error {
	token := primitive.NewObjectID().Hex() // Could use UUID for better randomness
	expiry := time.Now().Add(1 * time.Hour)
	return u.repo.SetResetToken(email, token, expiry)
}

func (u *UserUsecase) ResetPassword(token, newPassword string) error {
	return u.repo.ResetPasswordUsingToken(token, newPassword)
}

func (u *UserUsecase) Logout(userID primitive.ObjectID) error {
	// If you store tokens in DB, delete them here
	return u.repo.ClearTokens(userID)
}

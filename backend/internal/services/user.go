package services

import (
	"context"
	"errors"

	"find_a_walk/internal/domain"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.UserIn) (*domain.User, error)
	IsUserExists(ctx context.Context, email string) (bool, error)
}

type UserService struct {
	repo        UserRepository
	jwtConfig   *jwtauth.JWTAuth
}

func NewDefaultUserService(repo UserRepository, tokenConfig *jwtauth.JWTAuth) *UserService {
	return &UserService{repo: repo, jwtConfig: tokenConfig}
}

func (s *UserService) GetJWTConfig() (*jwtauth.JWTAuth) {
	return s.jwtConfig
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.UserIn) (*domain.User, error) {
	userExistsStatus, err := s.repo.IsUserExists(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if userExistsStatus {
		return nil, errors.New("user with this email already exists")
	}
	hashedPassword, err := s.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, user *domain.UserAuth) (*domain.Token, error) {
	userExistsStatus, err := s.repo.IsUserExists(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if !userExistsStatus {
		return nil, errors.New("user with this email doesnt exists")
	}

	var userSchema *domain.User
	userSchema, err = s.repo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		return nil, err
	}

	passwordStatus := s.CheckPasswordHash(user.Password, userSchema.Password)
	if passwordStatus {
		_, tokenString, err := s.jwtConfig.Encode(map[string]interface{}{"user_id": userSchema.ID})
		if err != nil {
			return nil, err
		}

		token := domain.Token{Token: tokenString}
		return &token, nil
	}
	return nil, errors.New("password incorrect")
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (s *UserService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

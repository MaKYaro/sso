package auth

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/MaKYaro/sso/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

// UserSaver is an interface that defines the method for saving user information.
type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

// UserProvider interface defines methods for user management and authorization in a system.
type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

// AppProvider defines the interface for retrieving application details.
type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// Auth represents an authentication structure that encapsulates the
// necessary components for user authentication and management.
type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvider
	tokenTTL     time.Duration
}

// New returns a new instance of the Auth service.
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

// Login checks if the user with given credentials exists in the system and returns access token.
//
// If user exists but password is incorrect returns error.
// If user doesn't exists returns error.
func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appID int,
) (string, error) {
	panic("not implemented")
}

// RegisterNewUser registers new user with given credentials.
//
// If the user exists returns error.
func (a *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "internal.service.auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(
			"failed to generate password hash",
			slog.String("error", err.Error()),
		)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	userID, err := a.userSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error(
			"failed to save user",
			slog.String("error", err.Error()),
		)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userID, nil
}

// Auth checks if the user with given userID is an admin
//
// If user isn't an admin returns error.
func (a *Auth) IsAdmin(
	ctx context.Context,
	userID int64,
) (bool, error) {
	panic("not implemented")
}

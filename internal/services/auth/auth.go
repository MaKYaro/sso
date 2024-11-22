package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/MaKYaro/sso/internal/domain/models"
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

// Register registers new user with given credentials.
//
// If the user exists returns error.
func (a *Auth) Register(
	ctx context.Context,
	email string,
	password string,
) (int, error) {
	panic("not impelemented")
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

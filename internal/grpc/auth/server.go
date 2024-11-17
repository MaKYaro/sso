package auth

import (
	"context"

	ssov1 "github.com/MaKYaro/protos/gen/go/sso"
	"github.com/badoux/checkmail"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
	IsAdmin(
		ctx context.Context,
		userID int64,
	) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

// Register registers auth server
func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

const (
	emptyEmail     = ""
	emptyAppID     = 0
	minEntropyBits = 50
)

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {

	if err := validateLogin(req); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}

func validateLogin(req *ssov1.LoginRequest) error {
	email := req.GetEmail()

	if email == emptyEmail {
		return status.Error(codes.InvalidArgument, "email required")
	}
	if err := checkmail.ValidateFormat(email); err != nil {
		return status.Error(codes.InvalidArgument, checkmail.ErrBadFormat.Error())
	}
	if err := checkmail.ValidateHost(email); err != nil {
		return status.Error(codes.InvalidArgument, checkmail.ErrUnresolvableHost.Error())
	}

	password := req.GetPassword()
	if err := passwordvalidator.Validate(password, minEntropyBits); err != nil {
		return status.Error(codes.InvalidArgument, "weak password")
	}

	appID := req.GetAppId()

	if appID == emptyAppID {
		return status.Error(codes.InvalidArgument, "app_id required")
	}

	return nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement Register method")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement IsAdmin method")
}

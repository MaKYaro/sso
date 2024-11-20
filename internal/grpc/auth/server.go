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
	emptyPassword  = ""
	emptyUserID    = 0
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

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}
	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())

	// TODO: handle service layer errors
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	if err := validateIsAdmin(req); err != nil {
		return nil, err
	}
	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())

	// TODO: implement handling of service layer error
	if err != nil {
		return nil, status.Error(codes.Internal, "internal erros")
	}

	return &ssov1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

func validateLogin(req *ssov1.LoginRequest) error {
	if req.GetEmail() == emptyEmail {
		return status.Error(codes.InvalidArgument, "email required")
	}

	if req.GetPassword() == emptyPassword {
		return status.Error(codes.InvalidArgument, "password required")
	}

	if req.GetAppId() == emptyAppID {
		return status.Error(codes.InvalidArgument, "app_id required")
	}

	return nil
}

func validateRegister(req *ssov1.RegisterRequest) error {
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

	return nil
}

func validateIsAdmin(req *ssov1.IsAdminRequest) error {
	if req.GetUserId() == emptyAppID {
		return status.Error(codes.InvalidArgument, "user id required")
	}
	return nil
}

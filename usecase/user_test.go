package usecase

import (
	"context"
	"errors"
	"self-payrol/model"
	"self-payrol/model/mocks"
	"self-payrol/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_userUsecase_WithdrawSalary(t *testing.T) {
	type args struct {
		ctx context.Context
		req *request.WithdrawRequest
	}
	type repoUserResponse struct {
		user *model.User
		err  error
	}
	type repoCompanyResponse struct {
		err error
	}
	tests := []struct {
		name                string
		args                args
		repoUserID          int
		repoUserResponse    repoUserResponse
		repoCompanyResponse repoCompanyResponse
		expectedErr         error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully withdraw salary",
			args: args{
				ctx: context.TODO(),
				req: &request.WithdrawRequest{
					ID:       1,
					SecretID: "secret",
				},
			},
			repoUserID: 1,
			repoUserResponse: repoUserResponse{
				user: &model.User{
					ID:         1,
					Name:       "test",
					SecretID:   "secret",
					PositionID: 1,
					Position: &model.Position{
						ID:     1,
						Name:   "CEO",
						Salary: 5000,
					},
				},
				err: nil,
			},
			repoCompanyResponse: repoCompanyResponse{
				err: nil,
			},
			expectedErr: nil,
		},
		{
			name: "Invalid user id",
			args: args{
				ctx: context.TODO(),
				req: &request.WithdrawRequest{
					ID:       0,
					SecretID: "secret",
				},
			},
			repoUserID: 0,
			repoUserResponse: repoUserResponse{
				user: nil,
				err:  assert.AnError,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Invalid secret id",
			args: args{
				ctx: context.TODO(),
				req: &request.WithdrawRequest{
					ID:       1,
					SecretID: "not-a-secret",
				},
			},
			repoUserID: 1,
			repoUserResponse: repoUserResponse{
				user: &model.User{
					ID:         1,
					Name:       "test",
					SecretID:   "secret",
					PositionID: 1,
					Position: &model.Position{
						ID:     1,
						Name:   "CEO",
						Salary: 5000,
					},
				},
				err: nil,
			},
			repoCompanyResponse: repoCompanyResponse{
				err: nil,
			},
			expectedErr: errors.New("secret id not valid"),
		},
		{
			name: "Failed to debit balance",
			args: args{
				ctx: context.TODO(),
				req: &request.WithdrawRequest{
					ID:       1,
					SecretID: "secret",
				},
			},
			repoUserID: 1,
			repoUserResponse: repoUserResponse{
				user: &model.User{
					ID:         1,
					Name:       "test",
					SecretID:   "secret",
					PositionID: 1,
					Position: &model.Position{
						ID:     1,
						Name:   "CEO",
						Salary: 5000,
					},
				},
				err: nil,
			},
			repoCompanyResponse: repoCompanyResponse{
				err: assert.AnError,
			},
			expectedErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := new(mocks.UserRepository)
			mockPositionRepository := new(mocks.PositionRepository)
			mockCompanyRepository := new(mocks.CompanyRepository)

			//? should we assert mockUserRepository.FindByID and mockCompanyRepository.DebitBalance error?
			mockUserRepository.On("FindByID", mock.Anything, tt.repoUserID).
				Return(tt.repoUserResponse.user, tt.repoUserResponse.err)

			if tt.repoUserResponse.err == nil {
				mockCompanyRepository.On("DebitBalance", mock.Anything, tt.repoUserResponse.user.Position.Salary, tt.repoUserResponse.user.Name+" withdraw salary ").
					Return(tt.repoCompanyResponse.err)
			}

			p := NewUserUsecase(mockUserRepository, mockPositionRepository, mockCompanyRepository)

			err := p.WithdrawSalary(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedErr, err)

			mockUserRepository.AssertExpectations(t)
			mockPositionRepository.AssertExpectations(t)
			if err == nil {
				mockCompanyRepository.AssertExpectations(t)
			}
		})
	}
}

func Test_userUsecase_GetByID(t *testing.T) {
	type repoUserResponse struct {
		user *model.User
		err  error
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name             string
		args             args
		repoUserResponse repoUserResponse
		expectedUser     *model.User
		expectedErr      error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfull get user by id",
			args: args{ctx: context.Background(), id: 1},
			repoUserResponse: repoUserResponse{
				user: &model.User{
					ID:       1,
					SecretID: "secret",
					Name:     "test",
					Email:    "test@test.com",
					Phone:    "123456789",
					Address:  "test address",
				},
				err: nil,
			},
			expectedUser: &model.User{
				ID:       1,
				SecretID: "secret",
				Name:     "test",
				Email:    "test@test.com",
				Phone:    "123456789",
				Address:  "test address",
			},
			expectedErr: nil,
		},
		{
			name: "User by id not found",
			args: args{ctx: context.Background(), id: 2},
			repoUserResponse: repoUserResponse{
				user: nil,
				err:  gorm.ErrRecordNotFound,
			},
			expectedUser: nil,
			expectedErr:  gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := new(mocks.UserRepository)
			mockPositionRepository := new(mocks.PositionRepository)
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockUserRepository.On("FindByID", mock.Anything, tt.args.id).
				Return(tt.repoUserResponse.user, tt.repoUserResponse.err)

			p := NewUserUsecase(mockUserRepository, mockPositionRepository, mockCompanyRepository)

			user, err := p.GetByID(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.expectedUser, user)
			assert.Equal(t, tt.expectedErr, err)

			mockUserRepository.AssertExpectations(t)
			mockPositionRepository.AssertExpectations(t)
			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

func Test_userUsecase_FetchUser(t *testing.T) {
	type repoUserResponse struct {
		users []*model.User
		err   error
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name             string
		args             args
		repoUserResponse repoUserResponse
		expectedUsers    []*model.User
		expectedErr      error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Succesfully fetch user",
			args: args{
				ctx:    context.TODO(),
				limit:  10,
				offset: 0,
			},
			repoUserResponse: repoUserResponse{
				users: []*model.User{
					{
						ID:       1,
						SecretID: "secret",
						Name:     "test",
						Email:    "test@test.com",
						Phone:    "123456789",
						Address:  "test address",
					},
					{
						ID:       2,
						SecretID: "secret2",
						Name:     "test2",
						Email:    "test2@test.com",
						Phone:    "123456789",
						Address:  "test address2",
					},
				},
				err: nil,
			},
			expectedUsers: []*model.User{
				{
					ID:       1,
					SecretID: "secret",
					Name:     "test",
					Email:    "test@test.com",
					Phone:    "123456789",
					Address:  "test address",
				},
				{
					ID:       2,
					SecretID: "secret2",
					Name:     "test2",
					Email:    "test2@test.com",
					Phone:    "123456789",
					Address:  "test address2",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Failed to fetch users",
			args: args{
				ctx:    context.TODO(),
				limit:  10,
				offset: 0,
			},
			repoUserResponse: repoUserResponse{
				users: nil,
				err:   assert.AnError,
			},
			expectedUsers: nil,
			expectedErr:   assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := new(mocks.UserRepository)
			mockPositionRepository := new(mocks.PositionRepository)
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockUserRepository.On("Fetch", mock.Anything, tt.args.limit, tt.args.offset).
				Return(tt.repoUserResponse.users, tt.repoUserResponse.err)

			p := NewUserUsecase(mockUserRepository, mockPositionRepository, mockCompanyRepository)

			users, err := p.FetchUser(tt.args.ctx, tt.args.limit, tt.args.offset)

			assert.Equal(t, tt.expectedUsers, users)
			assert.Equal(t, tt.expectedErr, err)

			mockUserRepository.AssertExpectations(t)
			mockPositionRepository.AssertExpectations(t)
			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

func Test_userUsecase_DestroyUser(t *testing.T) {
	type repoUserResponse struct {
		err error
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name             string
		args             args
		repoUserResponse repoUserResponse
		expectedErr      error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully destroy user",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			repoUserResponse: repoUserResponse{
				err: nil,
			},
			expectedErr: nil,
		},
		{
			name: "Failed to destroy user",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			repoUserResponse: repoUserResponse{
				err: assert.AnError,
			},
			expectedErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := new(mocks.UserRepository)
			mockPositionRepository := new(mocks.PositionRepository)
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockUserRepository.On("Delete", mock.Anything, tt.args.id).
				Return(tt.repoUserResponse.err)

			p := NewUserUsecase(mockUserRepository, mockPositionRepository, mockCompanyRepository)

			err := p.DestroyUser(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.expectedErr, err)

			mockUserRepository.AssertExpectations(t)
			mockPositionRepository.AssertExpectations(t)
			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

func Test_userUsecase_EditUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
		req *request.UserRequest
	}
	type repoUserResponse struct {
		user *model.User
		err1 error
		err2 error
	}
	tests := []struct {
		name             string
		args             args
		repoUser         *model.User
		repoUserResponse repoUserResponse
		expectedUser     *model.User
		expectedErr      error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully edit user",
			args: args{
				ctx: context.Background(),
				id:  1,
				req: &request.UserRequest{
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 1,
				},
			},
			repoUser: &model.User{
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 1,
			},
			repoUserResponse: repoUserResponse{
				user: &model.User{
					ID:         1,
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 1,
				},
				err1: nil,
				err2: nil,
			},
			expectedUser: &model.User{
				ID:         1,
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 1,
			},
			expectedErr: nil,
		},
		{
			name: "Failed to find user by id",
			args: args{
				ctx: context.Background(),
				id:  1,
				req: &request.UserRequest{
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 1,
				},
			},
			repoUser: &model.User{
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 1,
			},
			repoUserResponse: repoUserResponse{
				user: nil,
				err1: assert.AnError,
				err2: nil,
			},
			expectedUser: nil,
			expectedErr:  assert.AnError,
		},
		{
			name: "Failed to update user by id",
			args: args{
				ctx: context.Background(),
				id:  1,
				req: &request.UserRequest{
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 1,
				},
			},
			repoUser: &model.User{
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 1,
			},
			repoUserResponse: repoUserResponse{
				user: nil,
				err1: nil,
				err2: assert.AnError,
			},
			expectedUser: nil,
			expectedErr:  assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := new(mocks.UserRepository)
			mockPositionRepository := new(mocks.PositionRepository)
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockUserRepository.On("FindByID", mock.Anything, tt.args.id).
				Return(tt.repoUserResponse.user, tt.repoUserResponse.err1)

			if tt.repoUserResponse.err1 == nil {
				mockUserRepository.On("UpdateByID", mock.Anything, tt.args.id, tt.repoUser).
					Return(tt.repoUserResponse.user, tt.repoUserResponse.err2)
			}

			p := NewUserUsecase(mockUserRepository, mockPositionRepository, mockCompanyRepository)

			user, err := p.EditUser(tt.args.ctx, tt.args.id, tt.args.req)

			assert.Equal(t, tt.expectedUser, user)
			assert.Equal(t, tt.expectedErr, err)

			mockUserRepository.AssertExpectations(t)
			mockPositionRepository.AssertExpectations(t)
			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

func Test_userUsecase_StoreUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *request.UserRequest
	}
	type repoPositionResponse struct {
		position *model.Position
		err      error
	}
	type repoUserResponse struct {
		user *model.User
		err  error
	}
	tests := []struct {
		name                 string
		args                 args
		repoUser             *model.User
		repoPositionResponse repoPositionResponse
		repoUserResponse     repoUserResponse
		expectedUser         *model.User
		expectedErr          error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully store a new user",
			args: args{
				ctx: context.TODO(),
				req: &request.UserRequest{
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 1,
				},
			},
			repoUser: &model.User{
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 1,
			},
			repoPositionResponse: repoPositionResponse{
				err: nil,
			},
			repoUserResponse: repoUserResponse{
				user: &model.User{
					ID:         1,
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 1,
				},
				err: nil,
			},
			expectedUser: &model.User{
				ID:         1,
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 1,
			},
			expectedErr: nil,
		},
		{
			name: "invalid position id",
			args: args{
				ctx: context.TODO(),
				req: &request.UserRequest{
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 0,
				},
			},
			repoUser: &model.User{
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 0,
			},
			repoPositionResponse: repoPositionResponse{
				err: gorm.ErrRecordNotFound,
			},
			expectedUser: nil,
			expectedErr:  errors.New("position id not valid "),
		},
		{
			name: "failed to find position by id",
			args: args{
				ctx: context.TODO(),
				req: &request.UserRequest{
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 2,
				},
			},
			repoUser: &model.User{
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 2,
			},
			repoPositionResponse: repoPositionResponse{
				err: assert.AnError,
			},
			expectedUser: nil,
			expectedErr:  assert.AnError,
		},
		{
			name: "Failed to store a new user",
			args: args{
				ctx: context.TODO(),
				req: &request.UserRequest{
					Name:       "test",
					SecretID:   "secret",
					Email:      "test@test.com",
					Phone:      "123456789",
					Address:    "test address",
					PositionID: 1,
				},
			},
			repoUser: &model.User{
				Name:       "test",
				SecretID:   "secret",
				Email:      "test@test.com",
				Phone:      "123456789",
				Address:    "test address",
				PositionID: 1,
			},
			repoPositionResponse: repoPositionResponse{
				err: nil,
			},
			repoUserResponse: repoUserResponse{
				user: nil,
				err:  assert.AnError,
			},
			expectedUser: nil,
			expectedErr:  assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := new(mocks.UserRepository)
			mockPositionRepository := new(mocks.PositionRepository)
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockPositionRepository.On("FindByID", mock.Anything, tt.args.req.PositionID).
				Return(tt.repoPositionResponse.position, tt.repoPositionResponse.err)

			if tt.repoPositionResponse.err == nil {
				mockUserRepository.On("Create", mock.Anything, tt.repoUser).
					Return(tt.repoUserResponse.user, tt.repoUserResponse.err)
			}

			p := NewUserUsecase(mockUserRepository, mockPositionRepository, mockCompanyRepository)

			user, err := p.StoreUser(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedUser, user)
			assert.Equal(t, tt.expectedErr, err)

			if err == nil {
				mockUserRepository.AssertExpectations(t)
			}
			mockPositionRepository.AssertExpectations(t)
			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

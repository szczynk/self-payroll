package usecase_test

import (
	"context"
	"errors"
	"self-payrol/model"
	"self-payrol/model/mocks"
	"self-payrol/request"
	"self-payrol/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_positionUsecase_GetByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name                 string
		args                 args
		repoResponsePosition *model.Position
		repoResponseErr      error
		expectedPosition     *model.Position
		expectedErr          error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully get position by id",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			repoResponsePosition: &model.Position{
				ID:     1,
				Name:   "CEO",
				Salary: 1000,
			},
			repoResponseErr: nil,
			expectedPosition: &model.Position{
				ID:     1,
				Name:   "CEO",
				Salary: 1000,
			},
			expectedErr: nil,
		},
		{
			name: "failed to get position by id",
			args: args{
				ctx: context.TODO(),
				id:  2,
			},
			repoResponsePosition: nil,
			repoResponseErr:      gorm.ErrRecordNotFound,
			expectedPosition:     nil,
			expectedErr:          gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPositionRepository := new(mocks.PositionRepository)

			mockPositionRepository.On("FindByID", mock.Anything, tt.args.id).
				Return(tt.repoResponsePosition, tt.repoResponseErr)

			p := usecase.NewPositionUsecase(mockPositionRepository)

			position, err := p.GetByID(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.expectedPosition, position)
			assert.Equal(t, tt.expectedErr, err)

			mockPositionRepository.AssertExpectations(t)
		})
	}
}

func Test_positionUsecase_FetchPosition(t *testing.T) {
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name                  string
		args                  args
		repoResponsePositions []*model.Position
		repoResponseErr       error
		expectedPositions     []*model.Position
		expectedErr           error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully fetch positions",
			args: args{
				ctx:    context.TODO(),
				limit:  10,
				offset: 0,
			},
			repoResponsePositions: []*model.Position{
				{
					ID:     1,
					Name:   "CEO",
					Salary: 1000,
				},
				{
					ID:     2,
					Name:   "CPO",
					Salary: 900,
				},
				{
					ID:     3,
					Name:   "CTO",
					Salary: 900,
				},
			},
			repoResponseErr: nil,
			expectedPositions: []*model.Position{
				{
					ID:     1,
					Name:   "CEO",
					Salary: 1000,
				},
				{
					ID:     2,
					Name:   "CPO",
					Salary: 900,
				},
				{
					ID:     3,
					Name:   "CTO",
					Salary: 900,
				},
			},
			expectedErr: nil,
		},
		{
			name: "Failed to fetch positions",
			args: args{
				ctx:    context.TODO(),
				limit:  10,
				offset: 0,
			},
			repoResponsePositions: nil,
			repoResponseErr:       assert.AnError,
			expectedPositions:     nil,
			expectedErr:           assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPositionRepository := new(mocks.PositionRepository)

			mockPositionRepository.On("Fetch", mock.Anything, tt.args.limit, tt.args.offset).
				Return(tt.repoResponsePositions, tt.repoResponseErr)

			p := usecase.NewPositionUsecase(mockPositionRepository)

			positions, err := p.FetchPosition(tt.args.ctx, tt.args.limit, tt.args.offset)

			assert.Equal(t, tt.expectedPositions, positions)
			assert.Equal(t, tt.expectedErr, err)

			mockPositionRepository.AssertExpectations(t)
		})
	}
}

func Test_positionUsecase_DestroyPosition(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name            string
		args            args
		repoResponseErr error
		expectedErr     error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully delete position",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			repoResponseErr: nil,
			expectedErr:     nil,
		},
		{
			name: "Failed to delete position",
			args: args{
				ctx: context.TODO(),
				id:  2,
			},
			repoResponseErr: assert.AnError,
			expectedErr:     assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPositionRepository := new(mocks.PositionRepository)

			mockPositionRepository.On("Delete", mock.Anything, tt.args.id).
				Return(tt.repoResponseErr)

			p := usecase.NewPositionUsecase(mockPositionRepository)

			err := p.DestroyPosition(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.expectedErr, err)

			mockPositionRepository.AssertExpectations(t)
		})
	}
}

func Test_positionUsecase_EditPosition(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
		req *request.PositionRequest
	}
	tests := []struct {
		name                 string
		args                 args
		repoPosition         *model.Position
		repoResponsePosition *model.Position
		repoResponseErr1     error
		repoResponseErr2     error
		expectedPosition     *model.Position
		expectedErr          error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully edit position",
			args: args{
				ctx: context.TODO(),
				id:  1,
				req: &request.PositionRequest{Name: "CEO", Salary: 1000},
			},
			repoPosition:         &model.Position{Name: "CEO", Salary: 1000},
			repoResponsePosition: &model.Position{ID: 1, Name: "CEO", Salary: 1000},
			repoResponseErr1:     nil,
			repoResponseErr2:     nil,
			expectedPosition:     &model.Position{ID: 1, Name: "CEO", Salary: 1000},
			expectedErr:          nil,
		},
		{
			name: "Failed to find position by id",
			args: args{
				ctx: context.TODO(),
				id:  1,
				req: &request.PositionRequest{Name: "CEO", Salary: 1000},
			},
			repoPosition:         nil,
			repoResponsePosition: nil,
			repoResponseErr1:     gorm.ErrRecordNotFound,
			repoResponseErr2:     nil,
			expectedPosition:     nil,
			expectedErr:          gorm.ErrRecordNotFound,
		},
		{
			name: "Failed to edit position by id",
			args: args{
				ctx: context.TODO(),
				id:  1,
				req: &request.PositionRequest{Name: "CEO", Salary: 1000},
			},
			repoPosition:         &model.Position{Name: "CEO", Salary: 1000},
			repoResponsePosition: nil,
			repoResponseErr1:     nil,
			repoResponseErr2:     assert.AnError,
			expectedPosition:     nil,
			expectedErr:          assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPositionRepository := new(mocks.PositionRepository)

			//? should we assert FindByID and UpdateByID error?
			mockPositionRepository.On("FindByID", mock.Anything, tt.args.id).
				Return(tt.repoResponsePosition, tt.repoResponseErr1)

			if tt.repoResponseErr1 == nil {
				mockPositionRepository.On("UpdateByID", mock.Anything, tt.args.id, tt.repoPosition).
					Return(tt.repoResponsePosition, tt.repoResponseErr2)
			}

			p := usecase.NewPositionUsecase(mockPositionRepository)

			position, err := p.EditPosition(tt.args.ctx, tt.args.id, tt.args.req)

			assert.Equal(t, tt.expectedPosition, position)
			assert.Equal(t, tt.expectedErr, err)

			mockPositionRepository.AssertExpectations(t)
		})
	}
}

func Test_positionUsecase_StorePosition(t *testing.T) {
	type args struct {
		ctx context.Context
		req *request.PositionRequest
	}
	tests := []struct {
		name                 string
		args                 args
		repoPosition         *model.Position
		repoResponsePosition *model.Position
		repoResponseErr      error
		expectedPosition     *model.Position
		expectedErr          error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: &request.PositionRequest{
					Name:   "Software Engineer",
					Salary: 10000000,
				},
			},
			repoPosition: &model.Position{
				Name:   "Software Engineer",
				Salary: 10000000,
			},
			repoResponsePosition: &model.Position{
				ID:     1,
				Name:   "Software Engineer",
				Salary: 10000000,
			},
			repoResponseErr: nil,
			expectedPosition: &model.Position{
				ID:     1,
				Name:   "Software Engineer",
				Salary: 10000000,
			},
			expectedErr: nil,
		},
		{
			name: "error repository",
			args: args{
				ctx: context.TODO(),
				req: &request.PositionRequest{
					Name:   "Data Analyst",
					Salary: 8000000,
				},
			},
			repoPosition: &model.Position{
				Name:   "Data Analyst",
				Salary: 8000000,
			},
			repoResponsePosition: nil,
			repoResponseErr:      errors.New("error repository"),
			expectedPosition:     nil,
			expectedErr:          errors.New("error repository"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPositionRepository := new(mocks.PositionRepository)

			mockPositionRepository.On("Create", mock.Anything, tt.repoPosition).
				Return(tt.repoResponsePosition, tt.repoResponseErr)

			p := usecase.NewPositionUsecase(mockPositionRepository)

			position, err := p.StorePosition(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedPosition, position)
			assert.Equal(t, tt.expectedErr, err)

			mockPositionRepository.AssertExpectations(t)
		})
	}
}

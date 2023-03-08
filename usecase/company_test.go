package usecase_test

import (
	"context"
	"net/http"
	"self-payrol/model"
	"self-payrol/model/mocks"
	"self-payrol/request"
	"self-payrol/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_companyUsecase_GetCompanyInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name                string
		args                args
		repoResponseCompany *model.Company
		repoResponseErr     error
		expectedCompany     *model.Company
		expectedStatusCode  int
		expectedErr         error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully retrieve company info",
			args: args{
				ctx: context.TODO(),
			},
			repoResponseCompany: &model.Company{
				ID:      1,
				Name:    "PT SEJAHTERA SELAMANYA",
				Address: "Jln. Malioboro",
				Balance: 20000000,
			},
			repoResponseErr: nil,
			expectedCompany: &model.Company{
				ID:      1,
				Name:    "PT SEJAHTERA SELAMANYA",
				Address: "Jln. Malioboro",
				Balance: 20000000,
			},
			expectedStatusCode: http.StatusOK,
			expectedErr:        nil,
		},
		{
			name: "Failed to retrieve company info",
			args: args{
				ctx: context.TODO(),
			},
			repoResponseCompany: nil,
			repoResponseErr:     gorm.ErrRecordNotFound,
			expectedCompany:     nil,
			expectedStatusCode:  http.StatusNotFound,
			expectedErr:         gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockCompanyRepository.On("Get", mock.Anything).
				Return(tt.repoResponseCompany, tt.repoResponseErr)

			c := usecase.NewCompanyUsecase(mockCompanyRepository)

			company, statusCode, err := c.GetCompanyInfo(tt.args.ctx)

			assert.Equal(t, tt.expectedCompany, company)
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			assert.Equal(t, tt.expectedErr, err)

			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

func Test_companyUsecase_CreateOrUpdateCompany(t *testing.T) {
	type args struct {
		ctx context.Context
		req request.CompanyRequest
	}
	tests := []struct {
		name                string
		args                args
		repoCompany         *model.Company
		repoResponseCompany *model.Company
		repoResponseErr     error
		expectedCompany     *model.Company
		expectedStatusCode  int
		expectedErr         error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully create company",
			args: args{
				ctx: context.TODO(),
				req: request.CompanyRequest{
					Name:    "PT SEJAHTERA SELAMANYA",
					Balance: 20000000,
					Address: "Jln. Malioboro",
				},
			},
			repoCompany: &model.Company{
				Name:    "PT SEJAHTERA SELAMANYA",
				Balance: 20000000,
				Address: "Jln. Malioboro",
			},
			repoResponseCompany: &model.Company{
				ID:      1,
				Name:    "PT SEJAHTERA SELAMANYA",
				Balance: 20000000,
				Address: "Jln. Malioboro",
			},
			repoResponseErr: nil,
			expectedCompany: &model.Company{
				ID:      1,
				Name:    "PT SEJAHTERA SELAMANYA",
				Balance: 20000000,
				Address: "Jln. Malioboro",
			},
			expectedStatusCode: http.StatusOK,
			expectedErr:        nil,
		},
		{
			name: "Successfully update company",
			args: args{
				ctx: context.TODO(),
				req: request.CompanyRequest{
					Name:    "PT MANTAP MANTAP",
					Balance: 20000000,
					Address: "Jln. Malioboro No.42",
				},
			},
			repoCompany: &model.Company{
				Name:    "PT MANTAP MANTAP",
				Balance: 20000000,
				Address: "Jln. Malioboro No.42",
			},
			repoResponseCompany: &model.Company{
				ID:      1,
				Name:    "PT MANTAP MANTAP",
				Balance: 20000000,
				Address: "Jln. Malioboro No.42",
			},
			repoResponseErr: nil,
			expectedCompany: &model.Company{
				ID:      1,
				Name:    "PT MANTAP MANTAP",
				Balance: 20000000,
				Address: "Jln. Malioboro No.42",
			},
			expectedStatusCode: http.StatusOK,
			expectedErr:        nil,
		},
		{
			name: "Failed to create or update Company",
			args: args{
				ctx: context.TODO(),
				req: request.CompanyRequest{
					Name:    "PT MANTAP MANTAP",
					Balance: 20000000,
					Address: "Jln. Malioboro No.42",
				},
			},
			repoCompany: &model.Company{
				Name:    "PT MANTAP MANTAP",
				Balance: 20000000,
				Address: "Jln. Malioboro No.42",
			},
			repoResponseCompany: nil,
			repoResponseErr:     assert.AnError,
			expectedCompany:     nil,
			expectedStatusCode:  http.StatusUnprocessableEntity,
			expectedErr:         assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockCompanyRepository.On("CreateOrUpdate", mock.Anything, tt.repoCompany).
				Return(tt.repoResponseCompany, tt.repoResponseErr)

			c := usecase.NewCompanyUsecase(mockCompanyRepository)

			company, statusCode, err := c.CreateOrUpdateCompany(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedCompany, company)
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			assert.Equal(t, tt.expectedErr, err)

			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

func Test_companyUsecase_TopupBalance(t *testing.T) {
	type args struct {
		ctx context.Context
		req request.TopupCompanyBalance
	}
	tests := []struct {
		name                string
		args                args
		repoCompanyBalance  int
		repoResponseCompany *model.Company
		repoResponseErr     error
		expectedCompany     *model.Company
		expectedStatusCode  int
		expectedErr         error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "Successfully topup balance company",
			args: args{
				ctx: context.TODO(),
				req: request.TopupCompanyBalance{
					Balance: 5000000,
				},
			},
			repoResponseCompany: &model.Company{
				ID:      1,
				Name:    "PT SEJAHTERA SELAMANYA",
				Balance: 25000000,
				Address: "Jln. Malioboro",
			},
			repoResponseErr: nil,
			expectedCompany: &model.Company{
				ID:      1,
				Name:    "PT SEJAHTERA SELAMANYA",
				Balance: 25000000,
				Address: "Jln. Malioboro",
			},
			expectedStatusCode: http.StatusOK,
			expectedErr:        nil,
		},
		{
			name: "Failed topup balance company",
			args: args{
				ctx: context.TODO(),
				req: request.TopupCompanyBalance{
					Balance: 5000000,
				},
			},
			repoResponseCompany: nil,
			repoResponseErr:     assert.AnError,
			expectedCompany:     nil,
			expectedStatusCode:  http.StatusUnprocessableEntity,
			expectedErr:         assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCompanyRepository := new(mocks.CompanyRepository)

			mockCompanyRepository.On("AddBalance", mock.Anything, tt.args.req.Balance).
				Return(tt.repoResponseCompany, tt.repoResponseErr)

			c := usecase.NewCompanyUsecase(mockCompanyRepository)

			company, statusCode, err := c.TopupBalance(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedCompany, company)
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			assert.Equal(t, tt.expectedErr, err)

			mockCompanyRepository.AssertExpectations(t)
		})
	}
}

package usecase_test

import (
	"context"
	"net/http"
	"self-payrol/model"
	"self-payrol/model/mocks"
	"self-payrol/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_transactionUsecase_Fetch(t *testing.T) {
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name                     string
		args                     args
		repoResponseTransactions []*model.Transaction
		repoResponseErr          error
		expectedTransactions     []*model.Transaction
		expectedStatusCode       int
		expectedErr              error
	}{
		// TODO(Bagus): Add test cases.
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				limit:  10,
				offset: 0,
			},
			repoResponseTransactions: []*model.Transaction{
				{
					ID:     1,
					Amount: 20000000,
					Note:   "Topup balance company",
					Type:   "credit",
				},
				{
					ID:     2,
					Amount: 100,
					Note:   "test withdraw salary ",
					Type:   "debit",
				},
				{
					ID:     3,
					Amount: 200,
					Note:   "test2 withdraw salary ",
					Type:   "debit",
				},
			},
			repoResponseErr:    nil,
			expectedStatusCode: http.StatusOK,
			expectedTransactions: []*model.Transaction{
				{
					ID:     1,
					Amount: 20000000,
					Note:   "Topup balance company",
					Type:   "credit",
				},
				{
					ID:     2,
					Amount: 100,
					Note:   "test withdraw salary ",
					Type:   "debit",
				},
				{
					ID:     3,
					Amount: 200,
					Note:   "test2 withdraw salary ",
					Type:   "debit",
				},
			},
			expectedErr: nil,
		},
		{
			name: "error",
			args: args{
				ctx:    context.Background(),
				limit:  10,
				offset: 0,
			},
			repoResponseTransactions: nil,
			repoResponseErr:          assert.AnError,
			expectedStatusCode:       http.StatusInternalServerError,
			expectedTransactions:     nil,
			expectedErr:              assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTransactionRepository := new(mocks.TransactionRepository)

			mockTransactionRepository.On("Fetch", mock.Anything, tt.args.limit, tt.args.offset).
				Return(tt.repoResponseTransactions, tt.repoResponseErr)

			tr := usecase.NewTransactionUsecase(mockTransactionRepository)

			transactions, statusCode, err := tr.Fetch(tt.args.ctx, tt.args.limit, tt.args.offset)

			assert.Equal(t, tt.expectedTransactions, transactions)
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			assert.Equal(t, tt.expectedErr, err)

			mockTransactionRepository.AssertExpectations(t)
		})
	}
}

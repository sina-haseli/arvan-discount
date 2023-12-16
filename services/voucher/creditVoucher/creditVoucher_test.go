package creditVoucher

import (
	"context"
	"discount/models"
	"discount/repositories"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
	*repositories.Repository
	// Add other necessary mock methods here
}

func (m *MockRepository) FindVoucherByCode(ctx context.Context, code string) (models.VoucherModel, error) {
	args := m.Called(ctx, code)
	return args.Get(0).(models.VoucherModel), args.Error(1)
}

// Similarly mock other methods like RedeemVoucher, etc.

func TestRedeem(t *testing.T) {
	mockRepo := new(MockRepository)
	cv := NewCreditVoucher(mockRepo.Repository, "testQueue")

	// Test case: When FindVoucherByCode returns an error
	mockRepo.On("FindVoucherByCode", mock.Anything, mock.Anything).Return(models.VoucherModel{}, errors.New("not found"))
	err := cv.Redeem(context.Background(), "testUser", "testCode")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Add more test cases here
}

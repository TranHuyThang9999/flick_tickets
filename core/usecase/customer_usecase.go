package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/utils"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
)

type UseCaseCustomer struct {
	cus   domain.RepositoryCustomers
	trans domain.RepositoryTransaction
}

func NewUseCaseCustomer(
	cus domain.RepositoryCustomers,
	trans domain.RepositoryTransaction,
) *UseCaseCustomer {
	return &UseCaseCustomer{
		cus:   cus,
		trans: trans,
	}
}
func (e *UseCaseCustomer) SendOtpToEmail(ctx context.Context, email string) (*entities.SendOtpResponse, error) {

	if email == "" {
		return &entities.SendOtpResponse{
			Result: entities.Result{
				Code:    enums.INVALID_REQUEST_CODE,
				Message: enums.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	codeOtp := utils.GenerateOtp()

	tx, err := e.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.SendOtpResponse{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
			CreatedAt: utils.GenerateTimestamp(),
		}, nil
	}

	err = e.cus.RegisterCustomers(ctx, tx, &domain.Customers{
		ID:        utils.GenerateUniqueKey(),
		OTP:       codeOtp,
		Email:     email,
		CreatedAt: utils.GenerateTimestamp(),
		UpdatedAt: utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.SendOtpResponse{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
			CreatedAt: utils.GenerateTimestamp(),
		}, nil
	}
	err = utils.SendOtpToEmail(email, "Mã Xác Minh Tài Khoản - Vui lòng không chia sẻ với người khác", codeOtp)
	if err != nil {
		return &entities.SendOtpResponse{
			Result: entities.Result{
				Code:    enums.SEND_EMAIL_ERR_CODE,
				Message: enums.SEND_EMAIL_ERR_MESS,
			},
			CreatedAt: utils.GenerateTimestamp(),
		}, nil
	}
	tx.Commit()
	return &entities.SendOtpResponse{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		CreatedAt: utils.GenerateTimestamp(),
	}, nil
}
func (e *UseCaseCustomer) CheckOtp(ctx context.Context, req *entities.CustomersReqVerifyOtp) (*entities.CustomersRespVerifyOtp, error) {

	customer, err := e.cus.GetCustomersByEmail(ctx, req.Email, req.Otp)
	if err != nil {
		return &entities.CustomersRespVerifyOtp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if customer == nil {
		return &entities.CustomersRespVerifyOtp{
			Result: entities.Result{
				Code:    enums.OTP_ERR_VERIFY_CODE,
				Message: enums.OTP_ERR_VERIFY_MESS,
			},
		}, nil
	}
	err = e.cus.UpdateWhenCheckOtp(ctx, req.Otp, req.Email)
	if err != nil {
		return &entities.CustomersRespVerifyOtp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.CustomersRespVerifyOtp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

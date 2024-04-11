package usecase

import (
	"context"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"flick_tickets/core/mapper"
	"time"
)

type UseCaseCustomer struct {
	cus    domain.RepositoryCustomers
	trans  domain.RepositoryTransaction
	aes    *UseCaseAes
	config *configs.Configs
	jwt    *UseCaseJwt
}

func NewUseCaseCustomer(
	cus domain.RepositoryCustomers,
	trans domain.RepositoryTransaction,
	aes *UseCaseAes,
	config *configs.Configs,
	jwt *UseCaseJwt,

) *UseCaseCustomer {
	return &UseCaseCustomer{
		cus:    cus,
		trans:  trans,
		aes:    aes,
		config: config,
		jwt:    jwt,
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

	customer, err := e.cus.GetCustomerByEmail(ctx, email)
	if err != nil {
		return &entities.SendOtpResponse{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
			CreatedAt: utils.GenerateTimestamp(),
		}, nil
	}
	if customer == nil {
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
	} else {
		err = e.cus.UpdateProfile(ctx, tx, &domain.Customers{
			ID:        customer.ID,
			OTP:       codeOtp,
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

	customer, err := e.cus.GetCustomersByEmailUseCheckOtp(ctx, req.Email, req.Otp)
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
func (e *UseCaseCustomer) RegisterManager(ctx context.Context, req *entities.CustomersReqRegister) (*entities.CustomersReqRegisterResp, error) {

	id := utils.GenerateUniqueKey()

	tx, err := e.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.CustomersReqRegisterResp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}

	listCustomers, err := e.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		UserName: req.UserName,
		Role:     enums.ROLE_ADMIN,
	})
	if err != nil {
		tx.Rollback()
		return &entities.CustomersReqRegisterResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listCustomers) > 0 {
		return &entities.CustomersReqRegisterResp{
			Result: entities.Result{
				Code:    enums.USER_EXITS_CODE,
				Message: enums.USER_EXITS_CODE_MESS,
			},
		}, nil
	}
	respFile, err := utils.SetByCurlImage(ctx, req.File)
	if respFile.Result.Code != 0 || err != nil {
		return &entities.CustomersReqRegisterResp{
			Result: respFile.Result,
		}, nil
	}

	err = e.cus.RegisterCustomers(ctx, tx, &domain.Customers{
		ID:          id,
		UserName:    req.UserName,
		Password:    utils.GeneratePassword(),
		AvatarUrl:   respFile.URL,
		Address:     req.Address, //[]string
		Age:         req.Age,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		IsActive:    true,
		Role:        enums.ROLE_ADMIN,
		CreatedAt:   utils.GenerateTimestamp(),
		UpdatedAt:   utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.CustomersReqRegisterResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &entities.CustomersReqRegisterResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Id: id,
	}, nil
}
func (e *UseCaseCustomer) LoginCustomerManager(ctx context.Context, req *entities.CustomerReqLogin) (*entities.CustomerRespLogin, error) {

	// keyAes := e.config.KeyAES128

	listCustomers, err := e.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		UserName: req.UserName,
	})
	if err != nil {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listCustomers) == 0 {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	if req.Password != listCustomers[0].Password {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.LOGIN_ERR_CODE,
				Message: enums.LOGIN_ERR_MESS,
			},
		}, nil
	}
	if !listCustomers[0].IsActive {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.ACCOUNT_STAFF_LOCK_CODE,
				Message: enums.ACCOUNT_STAFF_LOCK_MESS,
			},
		}, nil
	}
	respToken, err := e.jwt.generateToken(enums.ROLE_STAFF, req.UserName)
	if err != nil {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.CREATE_TOKEN,
				Message: enums.CREATE_TOKEN_MESS,
			},
		}, nil
	}

	return &entities.CustomerRespLogin{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		JwtToken: respToken,
	}, nil
}
func (e *UseCaseCustomer) CreateAccountAdminManagerForStaff(ctx context.Context, req *entities.CustomersReqRegisterAdminForStaff) (*entities.CustomersRespRegisterAdmin, error) {

	id := utils.GenerateUniqueKey()
	keyPassword := utils.GeneratePassword()
	// Tạo một ngữ cảnh với thời gian chờ 7 giây
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()

	// Tạo một kênh để nhận thông báo khi quá thời gian chờ

	tx, err := e.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.CustomersRespRegisterAdmin{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}

	listCustomers, err := e.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		UserName: req.UserName,
		Role:     enums.ROLE_STAFF,
	})

	if err != nil {
		return &entities.CustomersRespRegisterAdmin{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	log.Infof("list", listCustomers)
	if len(listCustomers) > 0 {
		return &entities.CustomersRespRegisterAdmin{
			Result: entities.Result{
				Code:    enums.USER_EXITS_CODE,
				Message: enums.USER_EXITS_CODE_MESS,
			},
		}, nil
	}

	respFile, err := utils.SetByCurlImage(ctx, req.File)
	if respFile.Result.Code != 0 || err != nil {
		return &entities.CustomersRespRegisterAdmin{
			Result: respFile.Result,
		}, nil
	}
	err = e.cus.RegisterCustomers(ctx, tx, &domain.Customers{
		ID:          id,
		UserName:    req.UserName,
		Password:    keyPassword,
		AvatarUrl:   respFile.URL,
		Address:     req.Address, //[]string
		Age:         req.Age,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		IsActive:    true,
		Role:        enums.ROLE_STAFF,
		CreatedAt:   utils.GenerateTimestamp(),
		UpdatedAt:   utils.GenerateTimestamp(),
	})

	tx.Commit()
	if err != nil {
		return &entities.CustomersRespRegisterAdmin{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = utils.SendPasswordToEmail(req.Email, "Dạp phim gửi bạn tài khoản đăng nhập", keyPassword)
	if err != nil {
		return &entities.CustomersRespRegisterAdmin{
			Result: entities.Result{
				Code:    enums.SEND_EMAIL_ERR_CODE,
				Message: enums.SEND_EMAIL_ERR_MESS,
			},
		}, err
	}
	return &entities.CustomersRespRegisterAdmin{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Id: id,
	}, nil
}
func (e *UseCaseCustomer) LoginCustomerForStaff(ctx context.Context, req *entities.CustomerReqLogin) (*entities.CustomerRespLogin, error) {

	listCustomers, err := e.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		UserName: req.UserName,
	})
	if err != nil {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listCustomers) == 0 {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}

	if req.Password != listCustomers[0].Password {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.LOGIN_ERR_CODE,
				Message: enums.LOGIN_ERR_MESS,
			},
		}, nil
	}
	if !listCustomers[0].IsActive {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.ACCOUNT_STAFF_LOCK_CODE,
				Message: enums.ACCOUNT_STAFF_LOCK_MESS,
			},
		}, nil
	}
	respToken, err := e.jwt.generateToken(enums.ROLE_STAFF, req.UserName)
	if err != nil {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.CREATE_TOKEN,
				Message: enums.CREATE_TOKEN_MESS,
			},
		}, nil
	}

	return &entities.CustomerRespLogin{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		JwtToken: respToken,
	}, nil

}

// func (e *UseCaseCustomer) UpdateStaff(ctx context.Context)
func (e *UseCaseCustomer) GetAllStaff(ctx context.Context) (*entities.CustomersFindByFormResp, error) {

	listCustomer, err := e.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		Role: enums.ROLE_STAFF,
	})

	if err != nil {
		return &entities.CustomersFindByFormResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listCustomer) == 0 {
		return &entities.CustomersFindByFormResp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	return &entities.CustomersFindByFormResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Customers: mapper.ConvertListCustomerDomainToListCustomerEntity(listCustomer),
	}, nil
}
func (e *UseCaseCustomer) DeleteStaffByName(ctx context.Context, name string) (*entities.CustomerDeleteResp, error) {
	err := e.cus.DeleteStaffByName(ctx, name)

	if err != nil {
		return &entities.CustomerDeleteResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.CustomerDeleteResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

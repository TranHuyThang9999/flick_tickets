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
	"fmt"
	"time"

	"gorm.io/gorm"
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
	log.Infof("req : ", email)
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
			UserName:  utils.GetUsernameFromEmail(email),
			Role:      enums.ROLE_CUSTOMER,
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

	err = e.cus.RegisterCustomers(ctx, tx, &domain.Customers{
		ID:          id,
		UserName:    req.UserName,
		Password:    utils.GeneratePassword(),
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

	keyPassword := utils.GeneratePassword()
	account := fmt.Sprintf("user name : %s \n password : %s ", req.UserName, keyPassword)
	err = utils.SendPasswordToEmail(req.Email, "Dạp phim gửi bạn tài khoản đăng nhập", account)
	if err != nil {
		tx.Rollback()
		return &entities.CustomersReqRegisterResp{
			Result: entities.Result{
				Code:    enums.SEND_EMAIL_ERR_CODE,
				Message: enums.SEND_EMAIL_ERR_MESS,
			},
		}, err
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
func (e *UseCaseCustomer) Login(ctx context.Context, req *entities.CustomerReqLogin) (*entities.CustomerRespLogin, error) {

	customer, err := e.cus.FindCustomersByRole(ctx, req.UserName, req.Password, req.Role)
	if err != nil {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if customer == nil {
		return &entities.CustomerRespLogin{
			Result: entities.Result{
				Code:    enums.USER_NOT_EXIST_CODE,
				Message: enums.USER_NOT_EXIST_MESS,
			},
		}, nil
	}
	respToken, err := e.jwt.generateToken(customer.ID, req.Role, req.UserName)
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
		JwtToken:  respToken,
		Email:     customer.Email,
		UserName:  customer.UserName,
		CreatedAt: utils.GenerateTimestamp(),
	}, nil
}

// tao tk cho nhan vien
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

	if err != nil {
		return &entities.CustomersRespRegisterAdmin{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	// account := fmt.Sprintf("user name : %s \n password : %s ", req.UserName, keyPassword)
	// err = utils.SendPasswordToEmail(req.Email, "Dạp phim gửi bạn tài khoản đăng nhập", account)
	// if err != nil {
	// 	return &entities.CustomersRespRegisterAdmin{
	// 		Result: entities.Result{
	// 			Code:    enums.SEND_EMAIL_ERR_CODE,
	// 			Message: enums.SEND_EMAIL_ERR_MESS,
	// 		},
	// 	}, err
	// }
	tx.Commit()

	return &entities.CustomersRespRegisterAdmin{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Id: id,
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
func (e *UseCaseCustomer) CheckAccountAndSendOtp(ctx context.Context, req *entities.CheckAccountAndSendOtpReq) (
	*entities.CheckAccountAndSendOtpResp, error) {
	otp := utils.GenerateOtp()

	if req == nil {
		return &entities.CheckAccountAndSendOtpResp{
			Result: entities.Result{
				Code:    enums.INVALID_REQUEST_CODE,
				Message: enums.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	tx, err := e.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.CheckAccountAndSendOtpResp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}
	sumAcount, err := e.cus.FindAccountResetPassWord(ctx, req.UserName, req.Email, enums.ROLE_CUSTOMER)
	log.Infof("data ", req, enums.ROLE_CUSTOMER)
	if err != nil {
		return &entities.CheckAccountAndSendOtpResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if sumAcount == 0 {
		return &entities.CheckAccountAndSendOtpResp{
			Result: entities.Result{
				Code:    enums.DATA_EMPTY_ERR_CODE,
				Message: enums.DATA_EMPTY_ERR_MESS,
			},
		}, nil
	}
	account, err := e.cus.FindByUserName(ctx, req.Email)

	if err != nil {
		return &entities.CheckAccountAndSendOtpResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = e.cus.UpdateProfile(ctx, tx, &domain.Customers{
		ID:        account.ID,
		UserName:  req.UserName,
		OTP:       otp,
		UpdatedAt: utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.CheckAccountAndSendOtpResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = utils.SendOtpToEmail(req.Email, "Gửi bạn mã OTP để lấy lại password", otp)
	if err != nil {
		tx.Rollback()
		return &entities.CheckAccountAndSendOtpResp{
			Result: entities.Result{
				Code:    enums.SEND_EMAIL_ERR_CODE,
				Message: enums.SEND_EMAIL_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &entities.CheckAccountAndSendOtpResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (e *UseCaseCustomer) VerifyOtpByEmailAndResetPassword(ctx context.Context,
	req *entities.VerifyOtpByEmailReq) (*entities.VerifyOtpByEmailResp, error) {

	listAccount, err := e.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		Email: req.Email,
		OTP:   req.OTP,
	})
	if err != nil {
		return &entities.VerifyOtpByEmailResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(listAccount) == 0 {
		return &entities.VerifyOtpByEmailResp{
			Result: entities.Result{
				Code:    enums.OTP_ERR_VERIFY_CODE,
				Message: enums.OTP_ERR_VERIFY_MESS,
			},
		}, nil
	}
	err = e.cus.UpdateProfile(ctx, &gorm.DB{}, &domain.Customers{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.PasswordNew,
	})
	if err != nil {
		return &entities.VerifyOtpByEmailResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	return &entities.VerifyOtpByEmailResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

func (c *UseCaseCustomer) RegisterAccountCustomer(ctx context.Context,
	req *entities.RegisterAccountCustomerReq) (*entities.RegisterAccountCustomerResp, error) {

	userId := utils.GenerateUniqueKey()

	tx, err := c.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.RegisterAccountCustomerResp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}
	informationUser, err := c.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		UserName: req.UserName,
		Role:     enums.ROLE_CUSTOMER,
	})
	if err != nil {
		return &entities.RegisterAccountCustomerResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(informationUser) > 0 {
		return &entities.RegisterAccountCustomerResp{
			Result: entities.Result{
				Code:    enums.USER_EXITS_CODE,
				Message: enums.USER_EXITS_CODE_MESS,
			},
		}, nil
	}
	resp, err := utils.SetByCurlImage(ctx, req.File)
	if err != nil {
		return &entities.RegisterAccountCustomerResp{
			Result: resp.Result,
		}, nil
	}
	err = c.cus.RegisterCustomers(ctx, tx, &domain.Customers{
		ID:          userId,
		UserName:    req.UserName,
		Password:    req.Password,
		AvatarUrl:   resp.URL,
		Address:     req.Address,
		Age:         req.Age,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		OTP:         0,
		IsActive:    true,
		ExpiredTime: 0,
		Role:        enums.ROLE_CUSTOMER,
		CreatedAt:   utils.GenerateTimestamp(),
		UpdatedAt:   utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.RegisterAccountCustomerResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &entities.RegisterAccountCustomerResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (c *UseCaseCustomer) UpdateProfileCustomerByUserName(ctx context.Context,
	req *entities.UpdateProfileCustomerByUserNameReq,
) (*entities.UpdateProfileCustomerByUserNameResp, error) {

	tx, err := c.trans.BeginTransaction(ctx)
	if err != nil {
		return &entities.UpdateProfileCustomerByUserNameResp{
			Result: entities.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}
	informationUser, err := c.cus.FindCustomers(ctx, &domain.CustomersFindByForm{
		UserName: req.UserName,
		Role:     enums.ROLE_CUSTOMER,
	})
	log.Infof("data : ", informationUser)
	if err != nil {
		return &entities.UpdateProfileCustomerByUserNameResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(informationUser) == 0 {
		return &entities.UpdateProfileCustomerByUserNameResp{
			Result: entities.Result{
				Code:    enums.USER_NOT_EXIST_CODE,
				Message: enums.USER_NOT_EXIST_MESS,
			},
		}, nil
	}
	// Kiểm tra xem có hình ảnh được gửi đi không và cập nhật URL của ảnh đại diện
	var url string
	if req.File == nil {
		url = informationUser[0].AvatarUrl
	} else {
		resp, err := utils.SetByCurlImage(ctx, req.File)
		if err != nil {
			log.Error(err, "error image")
			return &entities.UpdateProfileCustomerByUserNameResp{
				Result: resp.Result,
			}, err // Trả về lỗi nếu có lỗi xử lý hình ảnh
		}
		url = resp.URL
	}

	err = c.cus.UpdateProfile(ctx, tx, &domain.Customers{
		ID:          informationUser[0].ID,
		UserName:    req.UserName,
		AvatarUrl:   url,
		Address:     req.Address,
		Age:         req.Age,
		Email:       req.Address,
		PhoneNumber: req.PhoneNumber,
		UpdatedAt:   utils.GenerateTimestamp(),
	})
	if err != nil {
		tx.Rollback()
		return &entities.UpdateProfileCustomerByUserNameResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &entities.UpdateProfileCustomerByUserNameResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (c *UseCaseCustomer) GetCustomerByUseName(ctx context.Context, req *entities.GetCustomerByUseNameReq) (*entities.GetCustomerByUseNameResp, error) {
	resp, err := c.cus.FindCustomersByUsename(ctx, req.UserName)
	if err != nil {
		return &entities.GetCustomerByUseNameResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.GetCustomerByUseNameResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Customer: resp,
	}, nil
}
func (c *UseCaseCustomer) GenTokenByEmail(ctx context.Context, email string) (*entities.CreateTokenRespWhenLoginWithEmail, error) {

	if email == "" {
		return &entities.CreateTokenRespWhenLoginWithEmail{
			Result: entities.Result{
				Code:    enums.CLIENT_ERROR_CODE,
				Message: enums.CLIENT_ERROR_MESS,
			},
		}, nil
	}
	customer, err := c.cus.GetCustomerByEmail(ctx, email)
	if err != nil {
		return &entities.CreateTokenRespWhenLoginWithEmail{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tokenEd, err := c.jwt.generateToken(customer.ID, enums.ROLE_CUSTOMER, customer.UserName)
	if err != nil {
		return &entities.CreateTokenRespWhenLoginWithEmail{
			Result: entities.Result{
				Code:    enums.CREATE_TOKEN,
				Message: enums.CREATE_TOKEN_MESS,
			},
		}, nil
	}
	return &entities.CreateTokenRespWhenLoginWithEmail{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		JwtToken: tokenEd,
	}, nil
}
func (c *UseCaseCustomer) UpdatePassWordByUsername(ctx context.Context, req *entities.UpdatePassWordReq) (*entities.UpdatePassWordResp, error) {
	err := c.cus.UpdatePassWord(ctx, req.UserName, req.NewPassword)
	if err != nil {
		return &entities.UpdatePassWordResp{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &entities.UpdatePassWordResp{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}

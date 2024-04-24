package domain

import "context"

type PayOSResponseType struct {
	Code      string      `json:"code"`
	Desc      string      `json:"desc"`
	Data      interface{} `json:"data"`
	Signature *string     `json:"signature"`
}

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type CheckoutRequestType struct {
	OrderCode    int64   `json:"orderCode"`
	Amount       int     `json:"amount"`
	Description  string  `json:"description"`
	CancelUrl    string  `json:"cancelUrl"`
	ReturnUrl    string  `json:"returnUrl"`
	Signature    *string `json:"signature"`
	Items        []Item  `json:"items"`
	BuyerName    *string `json:"buyerName"`
	BuyerEmail   *string `json:"buyerEmail"`
	BuyerPhone   *string `json:"buyerPhone"`
	BuyerAddress *string `json:"buyerAddress"`
	ExpiredAt    *int    `json:"expiredAt"`
}

type CheckoutResponseDataType struct {
	Bin           string `json:"bin"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	Amount        int    `json:"amount"`
	Description   string `json:"description"`
	OrderCode     int64  `json:"orderCode"`
	Currency      string `json:"currency"`
	PaymentLinkId string `json:"paymentLinkId"`
	Status        string `json:"status"`
	CheckoutUrl   string `json:"checkoutUrl"`
	QRCode        string `json:"qrCode"`
}

type CancelPaymentLinkRequestType struct {
	CancellationReason *string `json:"cancellationReason"`
}

type ConfirmWebhookRequestType struct {
	WebhookUrl string `json:"webhookUrl"`
}

type PaymentLinkDataType struct {
	Id                 string            `json:"id"`
	OrderCode          int64             `json:"orderCode"`
	Amount             int               `json:"amount"`
	AmountPaid         int               `json:"amontPaid"`
	AmountRemaining    int               `json:"amountRemaining"`
	Status             string            `json:"status"`
	CreateAt           string            `json:"createAt"`
	Transactions       []TransactionType `json:"transactions"`
	CancellationReason *string           `json:"cancellationReason"`
	CancelAt           *string           `json:"cancelAt"`
}

type TransactionType struct {
	Reference              string  `json:"reference"`
	Amount                 int     `json:"amount"`
	AccountNumber          string  `json:"accountNumber"`
	Description            string  `json:"description"`
	TransactionDateTime    string  `json:"transactionDateTime"`
	VirtualAccountName     *string `json:"virtualAccountName"`
	VirtualAccountNumber   *string `json:"virtualAccountNumber"`
	CounterAccountBankId   *string `json:"counterAccountBankId"`
	CounterAccountBankName *string `json:"counterAccountBankName"`
	CounterAccountName     *string `json:"counterAccountName"`
	CounterAccountNumber   *string `json:"counterAccountNumber"`
}

type WebhookType struct {
	Code      string           `json:"code"`
	Desc      string           `json:"desc"`
	Data      *WebhookDataType `json:"data"`
	Signature string           `json:"signature"`
}

type WebhookDataType struct {
	OrderCode              int64   `json:"orderCode"`
	Amount                 int     `json:"amount"`
	Description            string  `json:"description"`
	AccountNumber          string  `json:"accountNumber"`
	Reference              string  `json:"reference"`
	TransactionDateTime    string  `json:"transactionDateTime"`
	Currency               string  `json:"currency"`
	PaymentLinkId          string  `json:"paymentLinkId"`
	Code                   string  `json:"code"`
	Desc                   string  `json:"desc"`
	CounterAccountBankId   *string `json:"counterAccountBankId"`
	CounterAccountBankName *string `json:"counterAccountBankName"`
	CounterAccountName     *string `json:"counterAccountName"`
	CounterAccountNumber   *string `json:"counterAccountNumber"`
	VirtualAccountName     *string `json:"virtualAccountName"`
	VirtualAccountNumber   *string `json:"virtualAccountNumber"`
}
type Payment struct {
	ID            int64   `json:"id"`
	OrderID       int64   `json:"order_id"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
	PaymentDate   int     `json:"payment_date"`
}
type RepositoryPayment interface {
	AddPayment(ctx context.Context, req *Payment) error
	GetOrderByOrderid(ctx context.Context, OrderId int64) (*Payment, error)
}

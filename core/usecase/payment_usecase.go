package usecase

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flick_tickets/api/resources"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/core/entities"
	"flick_tickets/core/mapper"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type UseCasePayment struct {
	order *UseCaseOrder
}

func NewUseCasePayment(
	order *UseCaseOrder,
) *UseCasePayment {
	return &UseCasePayment{
		order: order,
	}

}

var PayOSClientId string
var PayOSApiKey string
var PayOSChecksumKey string

// Set ClientId, APIKey, ChecksumKey
func InitKeyPayPos(clientId string, apiKey string, checksumKey string) error {
	if clientId == "" || apiKey == "" || checksumKey == "" {
		return errors.New("invalid key")
	}
	PayOSClientId = clientId
	PayOSApiKey = apiKey
	PayOSChecksumKey = checksumKey
	return nil
}

const PayOSBaseUrl = "https://api-merchant.payos.vn/"

func (u *UseCasePayment) CreatePayment(ctx context.Context, paymentData entities.CheckoutRequestType) (*entities.CheckoutResponseDataType, error) {

	resp, err := u.order.RegisterTicket(ctx, &entities.OrdersReq{
		Id:         paymentData.OrderCode,
		ShowTimeId: paymentData.ShowTimeId,
		Email:      *paymentData.BuyerEmail,
		Seats:      paymentData.Seats,
	})
	if err != nil {
		log.Error(err, "error payment")
		return &entities.CheckoutResponseDataType{RespOrder: resp}, nil
	}
	if resp.Result.Code == enums.SHOW_TIME_ORDER_CODE {
		log.Infof("order regietrd")
		return &entities.CheckoutResponseDataType{
			RespOrder: enums.SHOW_TIME_ORDER_CODE,
		}, nil
	}
	if resp.Result.Code != 0 {
		log.Error(err, "error payment")
		return &entities.CheckoutResponseDataType{RespOrder: resp}, nil
	}

	if paymentData.OrderCode == 0 || paymentData.Amount == 0 || paymentData.Description == "" || paymentData.CancelUrl == "" || paymentData.ReturnUrl == "" {
		requiredPaymentData := entities.CheckoutRequestType{
			OrderCode:   paymentData.OrderCode,
			Amount:      paymentData.Amount,
			ReturnUrl:   paymentData.ReturnUrl,
			CancelUrl:   paymentData.CancelUrl,
			Description: paymentData.Description,
		}
		requiredKeys := []string{"OrderCode", "Amount", "ReturnUrl", "CancelUrl", "Description"}
		keysError := []string{}
		for _, key := range requiredKeys {
			switch key {
			case "OrderCode":
				if requiredPaymentData.OrderCode == 0 {
					keysError = append(keysError, key)
				}
			case "Amount":
				if requiredPaymentData.Amount == 0 {
					keysError = append(keysError, key)
				}
			case "ReturnUrl":
				if requiredPaymentData.ReturnUrl == "" {
					keysError = append(keysError, key)
				}
			case "CancelUrl":
				if requiredPaymentData.CancelUrl == "" {
					keysError = append(keysError, key)
				}
			case "Description":
				if requiredPaymentData.Description == "" {
					keysError = append(keysError, key)
				}
			}
		}

		if len(keysError) > 0 {
			msgError := fmt.Sprintf("%s %s must not be undefined or null.", enums.InvalidParameterErrorMessage, strings.Join(keysError, ", "))
			return nil, resources.NewPayOSError(enums.InvalidParameterErrorCode, msgError)
		}
	}
	if paymentData.OrderCode < -9007199254740991 || paymentData.OrderCode > 9007199254740991 {
		return nil, resources.NewPayOSError(enums.InvalidParameterErrorCode, enums.OrderCodeOuOfRange)
	}
	url := fmt.Sprintf("%s/v2/payment-requests", PayOSBaseUrl)
	signaturePaymentRequest, _ := u.CreateSignatureOfPaymentRequest(paymentData, PayOSChecksumKey)
	paymentData.Signature = &signaturePaymentRequest
	checkoutRequest, err := json.Marshal(paymentData)
	if err != nil {
		log.Error(err, "error 1")
		return nil, resources.NewPayOSError(enums.InternalServerErrorErrorCode, err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(checkoutRequest))
	if err != nil {
		log.Error(err, "error 2")
		return nil, resources.NewPayOSError(enums.InternalServerErrorErrorCode, err.Error())
	}

	req.Header.Set("x-client-id", PayOSClientId)
	req.Header.Set("x-api-key", PayOSApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Error(err, "error 3")
		return nil, resources.NewPayOSError(enums.InternalServerErrorErrorCode, err.Error())
	}
	defer res.Body.Close()

	var paymentLinkRes entities.PayOSResponseType
	resBody, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(resBody, &paymentLinkRes)
	if err != nil {
		log.Error(err, "error 4")
		return nil, resources.NewPayOSError(enums.InternalServerErrorErrorCode, err.Error())
	}

	if paymentLinkRes.Code == "00" {
		paymentLinkResSignatute, _ := u.CreateSignatureFromObj(paymentLinkRes.Data, PayOSChecksumKey)
		if paymentLinkResSignatute != *paymentLinkRes.Signature {
			log.Error(err, "error 5")
			return nil, resources.NewPayOSError(enums.DataNotIntegrityErrorCode, enums.DataNotIntegrityErrorMessage)
		}
		if paymentLinkRes.Data != nil {
			jsonData, err := json.Marshal(paymentLinkRes.Data)
			if err != nil {
				log.Error(err, "error 6")
				return nil, resources.NewPayOSError(enums.InternalServerErrorErrorCode, enums.InternalServerErrorErrorMessage)
			}

			var paymentLinkData entities.CheckoutResponseDataType
			err = json.Unmarshal(jsonData, &paymentLinkData)
			if err != nil {
				log.Error(err, "error 7")
				return nil, resources.NewPayOSError(enums.InternalServerErrorErrorCode, enums.InternalServerErrorErrorMessage)
			}
			log.Infof("next url", paymentLinkData.CheckoutUrl)

			//them check
			return &paymentLinkData, nil
		}

	}

	return nil, resources.NewPayOSError(paymentLinkRes.Code, paymentLinkRes.Desc)
}

func (u *UseCasePayment) GetOrderByIdFromPayOs(ctx context.Context, orderID string) (*entities.PayMentResponseCheckOrder, error) {
	url := fmt.Sprintf("https://api-merchant.payos.vn/v2/payment-requests/%s", orderID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}

	req.Header.Set("x-client-id", "c84c857d-160c-456a-91f2-384526d7a360")
	req.Header.Set("x-api-key", "f74461b1-d7d3-4fca-b918-fcb39524ce8c")
	req.Header.Set("Cookie", "connect.sid=s%3A-Sat8d9c-WFoxLgE3cJZTb9bi3oSwFC2.uiWQpbtmdJc8ARx1PevsohQW62U4QiaOgBPfX85%2F91s")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	defer resp.Body.Close()

	// Cập nhật phần "Cookie" từ phản hồi API trước đó
	newCookie := resp.Header.Get("Set-Cookie")
	req.Header.Set("Cookie", newCookie)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}

	var response entities.PayMentResponseCheckOrder
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error(err, "error convert string to json")
		return nil, fmt.Errorf(enums.ERROR_CONVERT_JSON_MESS+"%s", err)
	}

	if response.Code == "101" {
		log.Info("Mã thanh toán không tồn tại")
		a := []entities.Transaction{}
		return &entities.PayMentResponseCheckOrder{
			Code: "101",
			Desc: "Mã thanh toán không tồn tại",
			Data: entities.Data{
				ID:                 orderID,
				OrderCode:          0,
				Amount:             0,
				AmountPaid:         0,
				AmountRemaining:    0,
				Status:             "",
				CreatedAt:          "",
				Transactions:       a,
				CanceledAt:         nil,
				CancellationReason: nil,
			},
			Signature: "",
		}, nil
	}

	if response.Data.Status == "EXPIRED" {
		a := []entities.Transaction{}
		log.Info("Đơn hàng đã quá hạn thời gian")
		return &entities.PayMentResponseCheckOrder{
			Code: "103",
			Desc: "Đơn hàng đã quá hạn thời gian",
			Data: entities.Data{
				ID:                 "",
				OrderCode:          mapper.ConvertStringToInt(orderID),
				Amount:             0,
				AmountPaid:         0,
				AmountRemaining:    0,
				Status:             "",
				CreatedAt:          "",
				Transactions:       a,
				CanceledAt:         nil,
				CancellationReason: nil,
			},
		}, nil
	} else if response.Data.Status == "CANCELLED" {
		a := []entities.Transaction{}
		return &entities.PayMentResponseCheckOrder{
			Code: "102",
			Desc: "Đơn hàng đã hủy",
			Data: entities.Data{
				ID:                 "0",
				OrderCode:          mapper.ConvertStringToInt(orderID),
				Amount:             0,
				AmountPaid:         0,
				AmountRemaining:    0,
				Status:             "",
				CreatedAt:          "",
				Transactions:       a,
				CanceledAt:         nil,
				CancellationReason: nil,
			},
			Signature: "",
		}, nil
	}

	return &response, nil
}

func (u *UseCasePayment) CreateSignatureFromObj(obj interface{}, key string) (string, error) {
	sortedObj, err := u.SortObjByKey(obj)
	if err != nil {
		return "", err
	}

	keyBytes := []byte(key)

	hasher := hmac.New(sha256.New, keyBytes)

	hasher.Write([]byte(sortedObj))

	signature := hex.EncodeToString(hasher.Sum(nil))

	return signature, nil
}

func (u *UseCasePayment) SortObjByKey(obj interface{}) (string, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	var sortedPairs []string

	var jsonObj map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		return "", err
	}

	keys := make([]string, 0, len(jsonObj))
	for key := range jsonObj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := jsonObj[key]
		if value != nil {
			sortedPairs = append(sortedPairs, fmt.Sprintf("%s=%v", key, u.convertToString(value)))
		} else {
			sortedPairs = append(sortedPairs, fmt.Sprintf("%s=", key))
		}
	}

	sortedObj := strings.Join(sortedPairs, "&")

	return sortedObj, nil
}

func (u *UseCasePayment) CreateSignatureOfPaymentRequest(data entities.CheckoutRequestType, key string) (string, error) {
	dataStr := fmt.Sprintf("amount=%s&cancelUrl=%s&description=%s&orderCode=%s&returnUrl=%s",
		strconv.Itoa(data.Amount), data.CancelUrl, data.Description, strconv.FormatInt(data.OrderCode, 10), data.ReturnUrl)

	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(dataStr))
	signature := hex.EncodeToString(hasher.Sum(nil))

	return signature, nil
}

func (u *UseCasePayment) convertToString(value interface{}) string {
	switch v := value.(type) {
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprint(value)
	}
}

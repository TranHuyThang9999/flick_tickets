package main

import (
	"io"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

// Hàm mô phỏng use case khác để lấy dữ liệu
func GetDataFromUseCase() string {
	// Lấy dữ liệu từ use case khác
	data := "Hello from use case!"
	return data
}

func main() {
	r := gin.Default()

	r.GET("/api/sse", func(c *gin.Context) {

		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("Access-Control-Allow-Origin", "*")

		// Tạo kênh SSE
		ch := make(chan sse.Event)
		defer close(ch)

		// Goroutine để lấy dữ liệu từ use case khác và gửi SSE
		go func() {
			for {
				// Lấy dữ liệu từ use case khác
				data := GetDataFromUseCase()

				event := sse.Event{
					Event: "message",
					Data:  []byte(data),
				}
				ch <- event
				time.Sleep(1 * time.Second)
			}
		}()

		// Gửi dữ liệu SSE đến client
		c.Stream(func(w io.Writer) bool {
			if _, ok := <-ch; ok {
				// sse.Encode(w, event)
				sse.Encode(w, sse.Event{
					Data: GetDataFromUseCase(),
				})
				w.Write([]byte("\n"))
				return true
			}
			return false
		})

		// Xử lý sự kiện đóng kết nối từ client
		<-c.Writer.CloseNotify()
		// Thực hiện các tác vụ khi kết nối bị đóng
		println("Connection closed by client")
	})

	r.Run(":8080")
}
package main

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/payOSHQ/payos-lib-golang"
)

func main() {
	clientID := "c84c857d-160c-456a-91f2-384526d7a360"
	apiKey := "f74461b1-d7d3-4fca-b918-fcb39524ce8c"
	checksumKey := "a861fb19b44c840efe2632b492140200e6a2e496640e1312fb5b63d5bf54a47c"

	payos.Key(clientID, apiKey, checksumKey)

	name := "product_name"
	quantityStr := "10"
	totalAmountStr := "2000"

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		log.Fatal(err)
	}

	totalAmount, err := strconv.Atoi(totalAmountStr)
	if err != nil {
		log.Fatal(err)
	}

	item := payos.Item{
		Name:     name,
		Quantity: quantity,
		Price:    totalAmount,
	}

	items := []payos.Item{item}

	description := "Thanh toán tiền nhậu"
	orderID := 2005
	totalPrice := 2000

	paymentData := payos.CheckoutRequestType{
		OrderCode:   int64(orderID),
		Amount:      totalPrice,
		Description: description,
		Items:       items,
		ReturnUrl:   "https://localhost:3002",
		CancelUrl:   "https://localhost:3002",
	}

	createPayment, err := payos.CreatePaymentLink(paymentData)
	if err != nil {
		log.Fatal(err)
	}

	linkCheckout := createPayment.CheckoutUrl
	cmd := exec.Command("google-chrome", linkCheckout)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

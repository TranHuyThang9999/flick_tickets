package main

import (
	"flick_tickets/common/log"
	"flick_tickets/configs"
	"flick_tickets/core/usecase"
	"fmt"
	"image/png"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	//r := gin.Default()
	// r.GET("/api/sse/:id", SSEHandler)

	//	fmt.Println(result)

	// r.Use(cors.AllowAll())
	// r.GET("/load", LoadFileHtml)
	// r.Run(":8080")
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(utils.GeneratePassword())

	// }
	configs.LoadConfig("./configs/configs.json")
	log.LoadLogger()
	resp, _ := usecase.NewUseCaseAes(configs.Get())
	data := []byte("7452143")
	key, _ := resp.EncryptAes(data, []byte(configs.Get().KeyAES128))
	dataEncrypt := "K3DSu2akRFxbZ/x1I+0WuA=="
	statusSData, _ := resp.DecryptAes(dataEncrypt, []byte(configs.Get().KeyAES128))
	fmt.Println(key)
	fmt.Println(string(statusSData))

	// Tạo mã QR từ dữ liệu đã mã hóa
	qrCode, err := qrcode.New(dataEncrypt, qrcode.Medium)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Lưu mã QR vào file PNG
	file, err := os.Create("qrcode.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Lưu hình ảnh mã QR vào file PNG
	err = png.Encode(file, qrCode.Image(256))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("QR code generated successfully!")
}

func LoadFileHtml(c *gin.Context) {
	path := "cmd/test/index.html"
	htmlBytes, err := os.ReadFile(path)
	if err != nil {
		// Xử lý lỗi nếu có
		c.String(http.StatusInternalServerError, "Lỗi khi đọc tệp HTML")
		return
	}

	// Trả về trang HTML
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlBytes)
}

func SSEHandler(c *gin.Context) {
	id := c.Param("id")

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
			data := GetDataFromUseCase(id)
			event := sse.Event{
				Event: "message",
				Data:  data,
			}
			ch <- event
			time.Sleep(1 * time.Second)
		}
	}()

	// Gửi dữ liệu SSE đến client
	c.Stream(func(w io.Writer) bool {
		if _, ok := <-ch; ok {
			sse.Encode(w, sse.Event{
				Data: GetDataFromUseCase(id),
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
}

type SinhVien struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var data = []SinhVien{
	SinhVien{
		Id:   "1",
		Name: "A 11",
	},
	SinhVien{
		Id:   "2",
		Name: "B 22",
	},
	SinhVien{
		Id:   "3",
		Name: "B 3",
	},
}

func GetDataFromUseCase(id string) []SinhVien {
	var list = make([]SinhVien, 0)

	for i := 0; i < len(data); i++ {
		if data[i].Id == id {
			list = append(list, data[i])
		}
	}
	return list
}

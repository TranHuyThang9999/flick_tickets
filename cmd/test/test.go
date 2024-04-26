package main

import (
	"flick_tickets/common/log"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
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
	log.LoadLogger()
	// Thời điểm hiện tại
	currentTime := time.Now().Unix()

	// Thời điểm tạo của đơn hàng
	orderCreatedAt := int64(1714144681) // Đây là thời gian tạo của đơn hàng trong dữ liệu của bạn

	// Tính thời gian cách đây bao nhiêu giây
	timeDifference := currentTime - orderCreatedAt

	// Kiểm tra nếu thời gian cách đây lớn hơn 15 phút (15 * 60 giây)
	if timeDifference > 15*60 {
		fmt.Println("Đơn hàng này đã được tạo cách đây hơn 15 phút.")
	} else {
		fmt.Println("Đơn hàng này được tạo trong vòng 15 phút gần đây.")
	}
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

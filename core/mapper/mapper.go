package mapper

import (
	"flick_tickets/common/log"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"strconv"
	"strings"
)

func ParseToIntSlice(input string) ([]int, error) {
	if input == "" {
		// Trả về slice rỗng nếu chuỗi đầu vào là rỗng
		return []int{}, nil
	}

	// Tách chuỗi thành các phần tử
	parts := strings.Split(input, ",")

	// Khởi tạo slice để lưu trữ các số nguyên
	var list []int

	// Chuyển đổi từng phần tử thành số nguyên và thêm vào slice
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			// Trả về lỗi nếu có lỗi xảy ra
			return nil, err
		}
		list = append(list, num)
	}

	// Trả về slice và không có lỗi
	return list, nil
}

func ConvertIntArrayToString(intArray []int) string {
	stringArray := make([]string, len(intArray))
	for i, num := range intArray {
		stringArray[i] = strconv.Itoa(num)
	}
	result := strings.Join(stringArray, ",")
	return result
}
func HasDuplicate(list []int) bool {
	n := len(list)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[i] == list[j] {
				// Nếu tìm thấy hai phần tử giống nhau, trả về true
				return true
			}
		}
	}
	// Nếu không tìm thấy phần tử nào giống nhau, trả về false
	return false
}

// ["a","Hi"]

func ParseToStringSlice(input string) ([]string, error) {
	if input == "" {
		// Trả về slice rỗng nếu chuỗi đầu vào là rỗng
		return []string{}, nil
	}

	// Loại bỏ khoảng trắng dư thừa từ chuỗi
	input = strings.TrimSpace(input)

	// Loại bỏ dấu ngoặc vuông [ và ]
	input = strings.TrimPrefix(input, "[")
	input = strings.TrimSuffix(input, "]")

	// Sử dụng hàm strings.Fields để tách chuỗi thành các từ dựa trên dấu cách
	parts := strings.Fields(input)

	// Thêm dấu ngoặc kép xung quanh từng từ trong slice
	for i, part := range parts {
		parts[i] = `"` + part + `"`
	}

	// Trả về slice kết quả và không có lỗi
	return parts, nil
}

// letters := []string{"ayy", "Hi  b", "co VIP", "dapter jh"}
func ConvertListToStringSlice(list string) []string {
	// Loại bỏ khoảng trắng ở đầu và cuối chuỗi
	list = strings.TrimSpace(list)

	// Kiểm tra nếu chuỗi không có ký tự nào
	if list == "" {
		return []string{}
	}

	// Tách chuỗi thành các từ dựa trên dấu phẩy
	words := strings.Split(list, ",")

	// Loại bỏ khoảng trắng ở đầu và cuối mỗi từ
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}

	return words
}
func ConvertCustomerDomainToCustomerEntity(req *domain.Customers) *entities.CustomersRespFindByForm {
	return &entities.CustomersRespFindByForm{
		ID:          req.ID,
		UserName:    req.UserName,
		AvatarUrl:   req.AvatarUrl,
		Address:     req.Address,
		Age:         req.Age,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		IsActive:    req.IsActive,
		CreatedAt:   req.CreatedAt,
	}
}

func ConvertListCustomerDomainToListCustomerEntity(reqList []*domain.Customers) []*entities.CustomersRespFindByForm {
	entityList := make([]*entities.CustomersRespFindByForm, len(reqList))
	for i, req := range reqList {
		entity := ConvertCustomerDomainToCustomerEntity(req)
		entityList[i] = entity
	}
	return entityList
}
func ConvertStringToInt(id string) int {
	resp, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err, "error convert string to daate")
		return 0
	}
	return resp
}

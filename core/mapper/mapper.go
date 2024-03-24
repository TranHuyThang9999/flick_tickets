package mapper

import (
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

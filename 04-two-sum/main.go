package basicconcurrencyworkerpool

import "fmt"

func TwoSum(nums []int, target int) []int {
	// สร้าง Map สำหรับเก็บค่าที่เคยผ่านมาแล้ว (Value -> Index)
	indexMap := make(map[int]int)

	for i, num := range nums {
		// 1. คำนวณหา "ตัวเลขที่ขาดไป" เพื่อให้รวมกันได้ target
		complement := target - num

		// 2. ตรวจสอบใน Map ว่าเคยเจอไอ้เจ้า complement นี้มาก่อนหรือยัง
		if idx, found := indexMap[complement]; found {
			// ถ้าเจอแล้ว แปลว่าเราพบคู่ที่บวกกันได้ target พอดี
			return []int{idx, i}
		}

		// 3. ถ้ายังไม่เจอ ให้บันทึกตัวเลขปัจจุบันและ index ลงใน Map เพื่อรอเป็นคู่ให้คนอื่น
		indexMap[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(nums, target)

	if result != nil {
		fmt.Printf("Indices: %v (Values: %d + %d = %d)\n", 
			result, nums[result[0]], nums[result[1]], target)
	} else {
		fmt.Println("No solution found.")
	}
}
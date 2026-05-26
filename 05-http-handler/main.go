package basicconcurrencyworkerpool
// ด่านตรวจของ API (Request Validation Pipeline)
// ตัวอย่างการนำเข้าที่ถูกต้อง
// curl -X POST http://localhost:8080/hello \
//      -H "Content-Type: application/json" \
//      -d '{"name":"Somchai"}'
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// โครงสร้างข้อมูลสำหรับรับเข้า
type HelloRequest struct {
	Name string `json:"name"`
}

// โครงสร้างข้อมูลสำหรับส่งออก
type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1. ตรวจสอบ HTTP Method (ต้องเป็น POST เท่านั้น)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. แกะค่า JSON จาก Body
	var req HelloRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 3. เตรียมข้อมูลสำหรับตอบกลับ
	res := HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}

	// 4. ตั้งค่า Header ให้เป็น JSON และส่ง Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func main() {
	// กำหนด Endpoint
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server starting at :8080...")
	// รัน Server บน Port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

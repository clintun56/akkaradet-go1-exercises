package basicconcurrencyworkerpool
// จัดการการใช้งานพร้อมกันหลายๆตัวเช่นการกดไลค์หลายๆเครื่องพร้อมกัน
// ทำงานเป็นลำดับ
import (
	"fmt"
	"sync"
)

// SafeCounter เก็บค่า count และ Mutex เพื่อความปลอดภัยของข้อมูล
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc เพิ่มค่า count ทีละ 1 ทำงาน
func (c *SafeCounter) Inc() {
	// ล็อค Mutex ก่อนเข้าถึงข้อมูล
	c.mu.Lock()
	// มั่นใจว่า Mutex จะถูกคลายล็อคเมื่อฟังก์ชันทำงานเสร็จ
	defer c.mu.Unlock()
	
	c.count++
}

// Value คืนค่า count ปัจจุบัน แสดง
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	return c.count
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	numGoroutines := 1000

	// รัน 1,000 Goroutines เพื่อเพิ่มค่าพร้อมกัน
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	// รอให้ทุก Goroutine ทำงานเสร็จ
	wg.Wait()

	fmt.Printf("Final Count: %d\n", counter.Value())
}

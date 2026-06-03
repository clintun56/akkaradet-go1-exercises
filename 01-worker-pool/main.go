package basicconcurrencyworkerpool

import (
	"fmt"
	"sync"
	"time"
)


func RunWorkers(numWorkers int, numJobs int) {
	jobs := make(chan int, numJobs) // สร้าง Channel สำหรับส่งงาน
	var wg sync.WaitGroup           // ใช้สำหรับรอให้ทุก Worker ทำงานเสร็จ

	// 1. เริ่มสร้าง Workers ตามจำนวนที่กำหนด
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // เพิ่ม Counter ใน WaitGroup
		go worker(i, jobs, &wg)
	}

	// 2. ส่งงานเข้าไปใน Channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// 3. ปิด Channel เมื่อส่งงานครบ เพื่อให้ Worker รู้ว่าไม่มีงานใหม่แล้ว
	close(jobs)

	// 4. รอให้ทุก Worker ทำงานในคิวของตัวเองจนเสร็จ (wg.Done จนครบ)
	wg.Wait()
}

	// ฟังก์ชันสำหรับ Worker แต่ละตัว
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	// เสร็จแล้ว
	defer wg.Done()

	// วนลูปรับงานจาก Channel จนกว่า Channel จะถูกปิด
	for job := range jobs {
		fmt.Printf("Worker [%d] processing job [%d]\n", id, job)
		// จำลองการทำงานด้วยการ Sleep 1 วินาที
		time.Sleep(time.Second)
	}
}



func main() {
	numWorkers := 3
	numJobs := 5

	fmt.Printf("Starting %d workers to handle %d jobs...\n", numWorkers, numJobs)
	start := time.Now()

	RunWorkers(numWorkers, numJobs)

	fmt.Printf("All jobs finished. Time taken: %v\n", time.Since(start))
}
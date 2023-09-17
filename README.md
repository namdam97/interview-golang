# Golang interview questions  

## I. Golang
### 1. Concurrency trong golang ? Làm sao để dừng một chương trình goroutine
#### - Concurrency (tính đồng thời) là khả năng xử lí nhiều tác vụ cùng 1 lúc, nhưng CPU không xử lí hết 1 tác vụ rồi mới đến tác vụ khác, mà sẽ dành 1 lúc cho tác vụ này, 1 lúc cho tác vụ kia. Do vậy, chúng ta có cảm giác máy tính thực hiện nhiều tác vụ cùng 1 lúc, nhưng thực chất chỉ có 1 tác vụ được xử lí tại 1 thời điểm. Những thực thi đồng thời được gọi là Goroutines trong Go (Golang)
### - Sử dụng channels để gửi tín hiệu dừng từ chương trình chính đến goroutine. Goroutine sẽ thực hiện việc xem xét channel để biết khi nào nên kết thúc.
### 2. Channel là gì ? sự khác nhau giữa channel buffer và unbuffer
#### - Channel là các kênh giao tiếp trung gian giữa các Goroutines trong Golang. Channel giúp các goroutines có thể gởi và nhận được dữ liệu cho nhau một cách an toàn thông qua cơ chế lock-free. Mặc định, Channel là kênh giao tiếp 2 chiều. Nghĩa là Channel có thể dùng cho cả gởi và nhận dữ liệu.
#### - Buffered Channel là một channel trong Golang có khả năng lưu trữ được dữ liệu bên trong nó. Khả năng này được mô tả như sức chứa (capacity) của channel.
#### - Các đặc tính của buffered chanel :
##### Buffered Channel có len và cap. 
##### Buffered Channel sẽ không block goroutine nếu sức chứa vẫn còn, mà không cần phải có một goroutine khác lấy dữ liệu từ channel. 
##### Buffered Channel sẽ block goroutine hiện tại nếu vượt sức chứa
##### Lấy dữ liệu từ empty buffered channel sẽ block goroutine
##### Lưu trữ dữ liệu theo thứ tự FIFO (First-In-First-Out)
#### - Khác biệt giữa buffered channel và unbuffered channel
##### Sự khác biệt lớn nhất giữa buffered channel và unbuffered channel đó là về capacity. Buffered channel sẽ yêu cầu khi báo capacity lúc khởi tạo channel, unbuffer channel thì không cần.



### 3. Tại sao dùng channel truyền data trong khi có thể implement truyền thẳng trên go func
### 4. Nil channel để làm gì?
### 5. Select trong golang
### 6. Race condition là gì ? Sự khác nhau giữa data race và race condition
### 7. Deadlock là gì
### 8. Pointer trong golang
### 9. Phân biệt Array, Slice, Map trong golang
### 10. GRPC là gì
### 11. Generics trong golang là gì
### 12. Defer trong go để làm gì 
### 13. Các trường hợp memory leak
### 14. Khi nào sử dụng panic
### 15. Viết unit testing như thế nào
### 16. Giải thích method và interface ? method trong golang khác gì với function
### 17. Thiết kế chức năng graceful shutdown golang
### 18. Garbage collection 
### 19. Giải thích go runtime, go schedule 
### 20. Làm sao để monitor được RAM
### 21. Manual management heap như nào cho đỡ lâu
### 22. Golang có OOP không
### 23. Sự khác nhau giữa nil và empty slice
### 24. Giải thích sự khác nhau giữa map và struct
### 25. Tại sao dùng golang
### 26. Phương pháp xử lý lỗi trong golang



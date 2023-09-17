# Golang interview questions  

## I. Golang
### 1. Concurrency trong golang ? Làm sao để dừng một chương trình goroutine
#### - Concurrency (tính đồng thời) là khả năng xử lí nhiều tác vụ cùng 1 lúc, nhưng CPU không xử lí hết 1 tác vụ rồi mới đến tác vụ khác, mà sẽ dành 1 lúc cho tác vụ này, 1 lúc cho tác vụ kia. Do vậy, chúng ta có cảm giác máy tính thực hiện nhiều tác vụ cùng 1 lúc, nhưng thực chất chỉ có 1 tác vụ được xử lí tại 1 thời điểm. Những thực thi đồng thời được gọi là Goroutines trong Go (Golang)
### - Sử dụng channels để gửi tín hiệu dừng từ chương trình chính đến goroutine. Goroutine sẽ thực hiện việc xem xét channel để biết khi nào nên kết thúc.
### 2. Channel là gì ? sự khác nhau giữa channel buffer và unbuffer

### 3. Giải thích method và interface ? method trong golang khác gì với function
### 4. Phân biệt Array, Slice, Map trong golang
### 5. Pointer trong golang
### 6. 
### 7. Giải thích sự khác nhau giữa map và struct
### 8. Race condition là gì ? Sự khác nhau giữa data race và race condition
### 9. Deadlock là gì
### 10. GRPC là gì
### 11. Generics trong golang là gì
### 12. Defer trong go để làm gì 
### 13. Select trong golang
### 14. Khi nào sử dụng panic
### 15. Viết unit testing như thế nào
### 16. Các trường hợp memory leak
### 17. Thiết kế chức năng graceful shutdown golang
### 18. Garbage collection 
### 19. Giải thích go runtime, go schedule 
### 20. Làm sao để monitor được RAM
### 21. Manual management heap như nào cho đỡ lâu
### 22. Golang có OOP không
### 23. Sự khác nhau giữa nil và empty slice
### 24. Tại sao dùng channel truyền data trong khi có thể implement truyền thẳng trên go func
### 25. Nil channel để làm gì?
### 26. Tại sao dùng golang
### 27. Phương pháp xử lý lỗi trong golang



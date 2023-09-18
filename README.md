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
#### - Select trong golang được sử dụng để chọn channel đã sẵn sàng, select sẽ block cho tới khi có case có thể được thực thi trên các channel
### 6. Race condition là gì ? Sự khác nhau giữa data race và race condition
#### - Data race xảy ra khi có từ 2 goroutine trở lên cùng truy cập vào một vùng nhớ chung (shared resource) với ít nhất 1 goroutine thực hiện việc thay đổi giá trị trên vùng nhớ đó
#### - Race condition là vấn đề sai sót về mặt thời gian hoặc thứ tự thực thi của các goroutine trong chương trình.
### 7. Deadlock là gì
#### - Deadlock xảy ra khi một nhóm goroutines đang đợi nhau và không ai trong số đó có thể tiến hành. Ví dụ như các groutine đợi nhau hoặc deadlock xảy ra trong channel
### 8. Pointer trong golang
#### - Con trỏ (Pointer) là một biến chứa địa chỉ bộ nhớ của biến khác.
### 9. Phân biệt Array, Slice, Map trong golang
#### - Array hay mảng là một tập hợp các phần tử có cùng kiểu dữ liệu nằm liên tiếp nhau, hay nói cách khác thì Array là một tập hợp có thứ tự, vì thế chúng ta có thể truy cập Array bằng chỉ số (index) của phần tử đó trong mảng.
#### - Slice về bản chất là các tham chiếu đến mảng hiện có, nó mô tả một phần hoặc toàn bộ Array, vì thế nó có kích thước động (thay đổi được). Thông thường Slice sẽ được tạo từ 1 Array bằng cách lấy từ vị trí phần tử bắt đầu và kết thúc trên Array đó. Một slice gồm chiều dài, sức chứa và một con trỏ đến phần tử con của mảng .Độ dài của một slice chính là số lượng của các phần tử có trong slice đó, Sức chứa của slice là chỉ số lượng phần tử cơ bản bắt đầu từ mục mà slice đó được tạo ra
#### - Map cũng là một kiểu dữ liệu tập hợp, tuy nhiên các phần tử trong nó không có thứ tự, tức là chúng ta không thể truy xuất đến phần tử trong Map bằng chỉ số như Slice hay Array. Map sẽ chứa các phần tử dạng key-value, từ đó việc truy xuất sẽ thực hiện qua các key.
### 10. GRPC, GRPC Gateway là gì
#### - GRPC là một khung dựa trên RPC framework (Remote Procedure Call). GRPC nên dùng để giao tiếp backend to backend. CPU không gánh nhiều cost cho encode/decoding mỗi đầu nữa. Tuy nhiên đặc tính mỗi đầu cần import file model chung (gen từ protobuf file) nên nếu update thì phải update đủ. Việc này vô tình tạo dependency cho các bên sử dụng, có thể nhiều bạn sẽ không thích điều này.
#### - GRPC Gateway : là một công cụ mã nguồn mở được sử dụng để tạo ra một RESTful API từ một service gRPC, nó hoạt động như một reverse proxy để chuyển tiếp request từ client đến service gRPC
### 11. Generics trong golang là gì
### 12. Defer trong go để làm gì 
#### - Defer nó cho phép một câu lệnh được gọi ra nhưng không thực thi ngay mà hoãn lại đến khi các lệnh xung quanh trả về kết quả. Các lệnh được gọi qua từ khóa defer sẽ được đưa vào một stack, tức là hoạt động theo cơ chế vào sau ra trước (last-in-first-out). Lệnh nào defer sau sẽ được thực thi trước, giống như xếp 1 chồng đĩa thì chiếc đĩa sau cùng (ở trên cùng) sẽ được lấy ra trước
### 13. Các trường hợp memory leak
### 14. Khi nào sử dụng panic
### 15. Viết unit testing như thế nào
### 16. Giải thích method và interface ? method trong golang khác gì với function
#### - Interface trong Golang có 2 đặc điểm : Nó là 1 tập hợp các phương thức (methods), nhưng cũng là 1 kiểu
#### - Methods (phương thức) – nó là một hàm(function) được khai báo cho riêng một kiểu dữ liệu đặc biệt được gọi là receiver.
#### - Điểm khác cơ bản của methods so với function là việc khai báo receiver, từ đó cho phép khai báo trùng tên và chỉ cần khác kiểu dữ liệu nhận (receiver).
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



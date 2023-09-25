# Golang interview questions  

## I. Golang
### 1. Concurrency trong golang ? Làm sao để dừng một chương trình goroutine
#### - Concurrency (tính đồng thời) là khả năng xử lí nhiều tác vụ cùng 1 lúc, nhưng CPU không xử lí hết 1 tác vụ rồi mới đến tác vụ khác, mà sẽ dành 1 lúc cho tác vụ này, 1 lúc cho tác vụ kia. Do vậy, chúng ta có cảm giác máy tính thực hiện nhiều tác vụ cùng 1 lúc, nhưng thực chất chỉ có 1 tác vụ được xử lí tại 1 thời điểm. Những thực thi đồng thời được gọi là Goroutines trong Go (Golang)
#### - Sử dụng channels để gửi tín hiệu dừng từ chương trình chính đến goroutine. Goroutine sẽ thực hiện việc xem xét channel để biết khi nào nên kết thúc.
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
#### - Golang có OOP k ? câu trả lời là vừa có vừa không. Mặc dù Go có các kiểu và phương thức cho phép viết code theo kiểu lập trình hướng đối tượng, nhưng lại không có type phân cấp (tức là không có kế thừa). Ngoài ra trong Go còn có các cách để nhúng 1 loại dữ liệu này vào trong 1 kiểu dữ liệu khác
#### - Cách go hướng đối tượng
##### + Tính đóng gói : Go đóng gói mọi thứ ở cấp độ package. Với cách đặt tên  bắt đầu bằng một chữ cái viết thường thì chỉ truy cập được trong package đó (tức là private). Tên bắt đầu bằng chữ hoa sẽ là public
##### + Tính kế thừa : Go cho phép nhúng 1 type vào 1 type khác.
##### + Tính đa hình : Bản chất của đa hình của lập trình hướng đối tượng là khả năng xử lý các đối tượng thuộc các loại khác nhau miễn là chúng tuân thủ cùng một interface.
### 23. Sự khác nhau giữa nil và empty slice
### 24. Giải thích sự khác nhau giữa map và struct
### 25. Tại sao dùng golang
### 26. Phương pháp xử lý lỗi trong golang
### 27. Nguyên lý SOLID
### 28. 7 cấu trúc dữ liệu cơ bản
#### - Mảng (Array)
##### Mảng là một tập hợp các phần tử cùng kiểu dữ liệu được lưu trữ liên tiếp trong bộ nhớ và được truy cập thông qua chỉ số. Mảng giúp lưu trữ và quản lý dữ liệu theo cách có thứ tự
#### - Danh sách liên kết (Linked list)
##### Danh sách liên kết là cấu trúc dữ liệu mà mỗi phần tử gọi là nút chứa dữ liệu và một tham chiếu đến phần tử tiếp theo trong danh sách. Danh sách liên kết có thể là liên kết đơn hoặc liên kết kép
#### - Ngăn xếp (Stack)
##### Ngăn xếp là một cấu trúc dữ liệu dạng ngăn chứa, các phần tử được thêm và loại bỏ theo cơ chế "Last-In-First-Out" (LIFO). Phần tử cuối cùng được thêm vào ngăn sẽ được loại bỏ trước
#### - Hàng đợi (Queue)
##### Hàng đợi là một cấu trúc dữ liệu dạng ngăn chứa, các phần tử được thêm vào cuối và loại bỏ ở đầu theo cơ chế "First-In-First-Out" (FIFO). Phần tử đầu tiên được thêm vào hàng đợi sẽ được loại bỏ đầu tiên
#### - Cây (Tree)
##### Cây là một cấu trúc dữ liệu phân cấp bao gồm các nút có mối quan hệ cha-con. Cây thường được sử dụng để biểu diễn dữ liệu có cấu trúc phân cấp như cây gia phả hoặc cây thư mục trên máy tính
#### - Đồ thị (Graph)
##### Đồ thị là một cấu trúc dữ liệu bao gồm các đỉnh (node) và các cạnh (edge) để biểu diễn mối quan hệ giữa chúng. Đồ thị có nhiều ứng dụng, bao gồm mạng xã hội, mạng máy tính và tối ưu hóa đường đi
#### - Hàm băm (Hash table)
##### Hàm băm là một cấu trúc dữ liệu sử dụng hàm băm để ánh xạ khóa vào một giá trị trong bảng. Hàm băm giúp tìm kiếm và truy cập dữ liệu với thời gian độ phức tạp gần như là O(1)



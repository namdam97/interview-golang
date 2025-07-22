# Phỏng vấn Golang Toàn tập

## I. Golang Cơ bản (Go Fundamentals)

### 1. Con trỏ (Pointer) trong Golang là gì?
Con trỏ (Pointer) là một biến chứa địa chỉ bộ nhớ của một biến khác. Sử dụng con trỏ cho phép các hàm có thể sửa đổi trực tiếp giá trị của biến được truyền vào, thay vì tạo ra một bản sao. Điều này rất hữu ích để tối ưu hiệu năng khi làm việc với các cấu trúc dữ liệu lớn.

### 2. Phân biệt Array, Slice, và Map trong Golang
- **Array (Mảng):** Là một tập hợp các phần tử có cùng kiểu dữ liệu và có kích thước cố định, được xác định lúc biên dịch. Kích thước là một phần của kiểu dữ liệu của mảng (`[4]int` và `[5]int` là hai kiểu khác nhau).
- **Slice:** Về bản chất, Slice là một cấu trúc dữ liệu linh hoạt hơn, nó "nhìn" vào một phần của một mảng (underlying array). Một slice được định nghĩa bởi ba thành phần: con trỏ tới mảng, độ dài (length - số phần tử slice đang chứa), và sức chứa (capacity - số phần tử của mảng tính từ vị trí con trỏ). Slice có kích thước động và được sử dụng phổ biến hơn nhiều so với Array.
- **Map:** Là một tập hợp các cặp key-value không có thứ tự. Map cung cấp khả năng truy xuất, cập nhật và xóa phần tử với độ phức tạp thời gian trung bình là O(1). Tất cả các key trong map phải có cùng kiểu dữ liệu, và tất cả các value cũng vậy.

### 3. Method và Interface là gì? Method khác Function như thế nào?
- **Function:** Là một khối mã thực hiện một tác vụ cụ thể, có thể được gọi từ bất kỳ đâu trong package (nếu public).
- **Method:** Là một function được gắn với một kiểu dữ liệu cụ thể, được gọi là `receiver`. Về cơ bản, nó là một "hành vi" của kiểu dữ liệu đó.
- **Sự khác biệt:** Điểm khác biệt chính là `receiver`. Methods cho phép chúng ta định nghĩa hành vi trên các `struct` và các kiểu dữ liệu khác, là nền tảng cho việc triển khai interface trong Go.
- **Interface:** Là một tập hợp các `method signature` (chữ ký phương thức). Một kiểu dữ liệu được coi là "triển khai" một interface nếu nó định nghĩa tất cả các phương thức mà interface đó yêu cầu. Interface trong Go là `implicit` (ngầm), không cần từ khóa `implements`. Điều này tạo ra sự linh hoạt và découpage (tách rời) mạnh mẽ trong thiết kế hệ thống.

### 4. Cách Go tiếp cận Lập trình Hướng đối tượng (OOP)
Go không phải là ngôn ngữ OOP thuần túy như Java hay C++, nhưng nó hỗ trợ các nguyên lý OOP theo cách riêng:
- **Tính đóng gói (Encapsulation):** Go sử dụng package để đóng gói. Các định danh (biến, hàm, kiểu...) bắt đầu bằng chữ cái viết hoa sẽ được `export` (public), có thể truy cập từ package khác. Các định danh bắt đầu bằng chữ thường là `private`, chỉ truy cập được bên trong package.
- **Tính kế thừa (Inheritance):** Go không có kế thừa class. Thay vào đó, nó sử dụng **composition (kết hợp)** thông qua `struct embedding`. Bằng cách nhúng một struct vào một struct khác, struct bên ngoài có thể truy cập trực tiếp các trường và phương thức của struct được nhúng, mô phỏng lại hành vi của kế thừa.
- **Tính đa hình (Polymorphism):** Go đạt được tính đa hình thông qua `interfaces`. Một hàm có thể nhận một interface làm tham số, và nó có thể hoạt động với bất kỳ kiểu dữ liệu nào triển khai interface đó, cho phép các đối tượng khác nhau được xử lý theo cùng một cách.

### 5. Sự khác nhau giữa `nil` slice và `empty` slice
- **Nil Slice:** Một `nil` slice là slice chưa được khởi tạo. Con trỏ tới mảng của nó là `nil`. Length và capacity đều bằng 0.
  ```go
  var s []int // s is nil
  ```
- **Empty Slice:** Một `empty` slice đã được khởi tạo nhưng không chứa phần tử nào. Nó có con trỏ tới một vùng nhớ, nhưng length bằng 0. Capacity có thể bằng 0 hoặc lớn hơn.
  ```go
  s := []int{} // empty slice, len=0, cap=0
  s := make([]int, 0) // empty slice, len=0, cap=0
  ```
Trong hầu hết các trường hợp, chúng hoạt động giống nhau (ví dụ `len(s)` đều là 0, `for range` không chạy). Sự khác biệt rõ nhất là khi encoding sang JSON, `nil` slice sẽ thành `null`, còn `empty` slice sẽ thành `[]`.

### 6. So sánh `map` và `struct`
- **Struct:**
    - Các trường (key) được định nghĩa tại thời điểm biên dịch.
    - Tên trường phải là một định danh hợp lệ.
    - Các giá trị có thể có các kiểu dữ liệu khác nhau.
    - Thường nhanh hơn và sử dụng ít bộ nhớ hơn.
    - An toàn về kiểu (type-safe).
- **Map:**
    - Các key được xác định tại thời điểm chạy.
    - Tất cả các key phải cùng kiểu, và tất cả các value cũng phải cùng kiểu.
    - Lý tưởng khi số lượng và tên của các key không được biết trước.
    - Chậm hơn một chút so với struct.

Chọn `struct` khi bạn có một cấu trúc dữ liệu cố định, và chọn `map` khi bạn cần một cấu trúc động.

### 7. Từ khóa `defer` dùng để làm gì?
`defer` trì hoãn việc thực thi một lệnh gọi hàm cho đến ngay trước khi hàm chứa `defer` kết thúc. Các lệnh `defer` được đẩy vào một stack, vì vậy chúng được thực thi theo thứ tự **Last-In-First-Out (LIFO)**.

Công dụng phổ biến nhất là để dọn dẹp tài nguyên, chẳng hạn như đóng file, mở khóa mutex, hoặc đóng kết nối mạng, đảm bảo chúng luôn được thực thi dù hàm kết thúc bình thường hay do `panic`.

### 8. Generics trong Golang là gì?
Generics, được giới thiệu từ Go 1.18, cho phép viết code có thể hoạt động trên nhiều kiểu dữ liệu khác nhau mà vẫn đảm bảo an toàn kiểu (type safety).
- **Type Parameters:** Cho phép định nghĩa các hàm và kiểu dữ liệu với các "tham số kiểu" placeholder.
- **Constraints:** Là các interface định nghĩa những yêu cầu mà một kiểu dữ liệu phải đáp ứng để có thể được sử dụng với hàm generic đó (ví dụ: kiểu đó phải hỗ trợ toán tử `+` và `>`).

Generics giúp giảm thiểu việc lặp lại code cho các kiểu dữ liệu khác nhau (ví dụ: viết một hàm `Max` cho cả `int` và `float64`).

### 9. Tại sao nên chọn Golang?
- **Đơn giản:** Cú pháp gọn gàng, ít từ khóa, dễ học.
- **Hiệu năng cao:** Được biên dịch ra mã máy, tốc độ thực thi nhanh, gần với C/C++.
- **Concurrency mạnh mẽ:** Goroutines và Channels là các tính năng được tích hợp sẵn trong ngôn ngữ, giúp việc viết các chương trình đồng thời trở nên cực kỳ đơn giản và hiệu quả.
- **Thư viện chuẩn phong phú:** Cung cấp nhiều công cụ mạnh mẽ cho việc xây dựng ứng dụng mạng, xử lý chuỗi, JSON...
- **Build nhanh:** Thời gian biên dịch rất nhanh, giúp tăng năng suất phát triển.
- **Hệ sinh thái tốt:** Cộng đồng lớn mạnh và có nhiều công cụ hỗ trợ (testing, profiling, static analysis).

### 10. Cách xử lý lỗi (Error Handling) trong Golang
Go có một cách tiếp cận rõ ràng và tường minh để xử lý lỗi. Thay vì dùng `try-catch`, các hàm có thể thất bại sẽ trả về một giá trị `error` là giá trị trả về cuối cùng.
- **Cơ chế:** Idiom phổ biến là `if err != nil`. Lập trình viên có trách nhiệm kiểm tra lỗi ngay sau khi gọi hàm.
- **Wrapping Errors:** Go 1.13 giới thiệu `fmt.Errorf` với `%w` để "bọc" lỗi, giữ lại thông tin của lỗi gốc.
- **Kiểm tra lỗi:** Sử dụng `errors.Is()` để kiểm tra một lỗi có phải là một lỗi cụ thể không, và `errors.As()` để trích xuất một lỗi có kiểu cụ thể từ chuỗi lỗi đã được bọc.

Cách tiếp cận này buộc lập trình viên phải suy nghĩ về các trường hợp lỗi, làm cho code trở nên tin cậy hơn.

---

## II. Concurrency (Lập trình đồng thời)

### 11. Concurrency là gì? Làm sao để dừng một goroutine?
- **Concurrency (tính đồng thời):** Là khả năng cấu trúc một chương trình để xử lý nhiều tác vụ một cách độc lập và không theo thứ tự. Các tác vụ này có thể chạy song song (parallelism) nếu có nhiều CPU, hoặc chạy xen kẽ trên một CPU. Goroutines là cơ chế để thực hiện concurrency trong Go.
- **Cách dừng một goroutine:** Phương pháp an toàn và phổ biến nhất là sử dụng channel.
    1.  Tạo một channel `done` (ví dụ `chan struct{}`).
    2.  Truyền channel này vào goroutine.
    3.  Trong goroutine, sử dụng `select` để lắng nghe trên cả channel công việc và channel `done`.
    4.  Khi muốn dừng, gửi một tín hiệu (hoặc đóng channel `done`) từ goroutine chính. Goroutine con sẽ nhận tín hiệu và thoát ra một cách an toàn.
    
    Một cách khác là sử dụng `context.Context`, đặc biệt hữu ích khi truyền tín hiệu hủy xuống một cây các goroutines.

### 12. Channel là gì? Phân biệt buffered và unbuffered channel
- **Channel:** Là một "đường ống" có kiểu dữ liệu, cho phép các goroutine giao tiếp và đồng bộ hóa với nhau một cách an toàn. Nguyên tắc của Go là: *"Don't communicate by sharing memory, share memory by communicating."*
- **Unbuffered Channel (kênh không đệm):**
    - `ch := make(chan int)`
    - Việc gửi dữ liệu vào channel sẽ **block** cho đến khi có một goroutine khác sẵn sàng nhận dữ liệu.
    - Việc nhận dữ liệu cũng sẽ **block** cho đến khi có goroutine khác gửi dữ liệu vào.
    - Đây là một cơ chế đồng bộ hóa mạnh mẽ.
- **Buffered Channel (kênh có đệm):**
    - `ch := make(chan int, 10)`
    - Việc gửi dữ liệu sẽ chỉ **block** khi buffer đã đầy.
    - Việc nhận dữ liệu sẽ chỉ **block** khi buffer đang rỗng.
    - Cho phép các goroutine giao tiếp một cách bất đồng bộ ở mức độ nhất định, giúp giảm bớt sự chờ đợi và tăng thông lượng.

### 13. Tại sao nên dùng channel để giao tiếp giữa các goroutine?
- **An toàn (Thread Safety):** Channel cung cấp cơ chế đồng bộ hóa tích hợp sẵn, loại bỏ nguy cơ data race khi nhiều goroutine truy cập cùng một dữ liệu. Bạn không cần phải tự quản lý mutex.
- **Code rõ ràng:** Việc truyền dữ liệu qua channel làm cho luồng dữ liệu và sự phụ thuộc giữa các goroutine trở nên tường minh.
- **Tách rời (Decoupling):** Goroutine gửi không cần biết về goroutine nhận và ngược lại. Chúng chỉ cần biết về channel, giúp thiết kế trở nên linh hoạt và dễ bảo trì hơn.

### 14. Nil channel dùng để làm gì?
Một `nil` channel (chưa được khởi tạo) có một đặc tính rất hữu ích:
- Cả việc gửi và nhận trên một `nil` channel sẽ **block mãi mãi**.

Công dụng chính của nó là trong câu lệnh `select`. Bằng cách gán một case trong `select` thành `nil`, bạn có thể **vô hiệu hóa tạm thời** case đó một cách linh hoạt tại thời điểm chạy, mà không cần thay đổi cấu trúc của `select`. Điều này rất hữu ích khi xây dựng các state machine hoặc khi bạn muốn ngừng nhận/gửi trên một kênh cụ thể.

### 15. `select` trong Golang hoạt động như thế nào?
`select` là một câu lệnh cho phép một goroutine chờ trên nhiều channel cùng một lúc.
- `select` sẽ block cho đến khi một trong các `case` (gửi hoặc nhận) có thể thực hiện được.
- Nếu nhiều case cùng sẵn sàng, `select` sẽ chọn một cách **ngẫu nhiên** để thực thi. Điều này đảm bảo sự công bằng và tránh tình trạng một channel luôn được ưu tiên.
- `case default`: Nếu có `default`, `select` sẽ không block. Nó sẽ thực thi `default` ngay lập tức nếu không có case nào khác sẵn sàng.

### 16. Race condition và data race là gì?
- **Data Race:** Xảy ra khi hai hoặc nhiều goroutine truy cập **cùng một vùng nhớ** một cách đồng thời, và ít nhất một trong các truy cập đó là **ghi (write)**, mà không có sự đồng bộ hóa (như mutex hay channel). Go cung cấp `go run -race` để phát hiện data race.
- **Race Condition:** Là một khái niệm rộng hơn. Nó là một lỗi trong chương trình mà kết quả cuối cùng phụ thuộc vào thứ tự hoặc thời gian thực thi không thể đoán trước của các sự kiện song song. Data race là một loại race condition, nhưng bạn vẫn có thể có race condition mà không có data race (ví dụ: lỗi logic trong việc sắp xếp thứ tự các thao tác trên channel).

### 17. Deadlock là gì?
Deadlock xảy ra khi một nhóm các goroutines bị kẹt, mỗi goroutine đang chờ một goroutine khác trong nhóm giải phóng tài nguyên mà nó cần. Kết quả là không goroutine nào có thể tiếp tục. Ví dụ phổ biến:
- Một goroutine gửi vào một unbuffered channel nhưng không có goroutine nào nhận.
- Hai goroutine cố gắng khóa hai mutex theo thứ tự trái ngược nhau.
- Các goroutine trong một vòng lặp chờ đợi lẫn nhau.

### 18. Thiết kế graceful shutdown cho ứng dụng Go
Graceful shutdown là quá trình đóng ứng dụng một cách có kiểm soát, cho phép các tác vụ đang thực hiện hoàn thành, đóng kết nối và giải phóng tài nguyên một cách an toàn.
1.  **Lắng nghe tín hiệu OS:** Sử dụng `os.Signal` và `signal.Notify` để bắt các tín hiệu như `SIGINT` (Ctrl+C) và `SIGTERM`.
2.  **Tạo một Context:** Khi nhận được tín hiệu, gọi `cancel()` trên một `context.Context` chính.
3.  **Truyền Context:** Truyền context này xuống các thành phần của ứng dụng (ví dụ: server HTTP, background workers).
4.  **Phản ứng với Context:** Các thành phần này sẽ lắng nghe `ctx.Done()`. Khi nhận được tín hiệu, chúng sẽ ngừng nhận request mới, hoàn thành các request đang xử lý, và dọn dẹp tài nguyên.
5.  **Sử dụng `WaitGroup`:** Dùng `sync.WaitGroup` để goroutine chính chờ tất cả các thành phần con hoàn thành việc shutdown trước khi thoát chương trình.

---

## III. Performance & Quản lý bộ nhớ (Performance & Memory)

### 19. Các nguyên nhân gây ra memory leak trong Go?
Memory leak trong Go xảy ra khi bộ nhớ được cấp phát nhưng không bao giờ được giải phóng vì vẫn còn tham chiếu đến nó, dù nó không còn được sử dụng nữa. Các nguyên nhân phổ biến:
- **Goroutine Leaks:** Một goroutine bị block vĩnh viễn (ví dụ: chờ trên một channel không bao giờ có dữ liệu) và không bao giờ kết thúc. Vùng nhớ mà goroutine này tham chiếu đến sẽ không được GC thu dọn.
- **Tham chiếu không cần thiết:** Các slice con tham chiếu đến một mảng (underlying array) rất lớn. Dù slice con nhỏ, nhưng toàn bộ mảng gốc vẫn không được giải phóng.
- **Global Variables:** Thêm dữ liệu vào một map hoặc slice toàn cục mà không bao giờ xóa đi.
- **Timer/Ticker không được dừng:** Tạo `time.Ticker` hoặc `time.Timer` nhưng quên gọi `Stop()` khi không cần nữa.

### 20. Khi nào nên sử dụng `panic`?
`panic` nên được sử dụng cho các **lỗi không thể phục hồi**, những tình huống mà chương trình không thể tiếp tục một cách hợp lý.
- **Lỗi của lập trình viên:** Ví dụ, truy cập mảng ngoài giới hạn, nil pointer dereference.
- **Trạng thái không thể xảy ra:** Khi một điều kiện mà bạn cho là "không thể xảy ra" lại xảy ra, cho thấy một bug nghiêm trọng.

Nó không nên được dùng cho các lỗi có thể dự đoán được (ví dụ: không tìm thấy file, lỗi kết nối mạng). Đối với những trường hợp đó, hãy trả về một `error`. Một `panic` không được `recover` sẽ làm sập toàn bộ chương trình.

### 21. Garbage Collection (GC) trong Go hoạt động như thế nào?
Go sử dụng một **concurrent, tri-color, mark-and-sweep garbage collector**.
- **Mục tiêu:** Tối ưu hóa cho độ trễ thấp, nghĩa là các khoảng tạm dừng (pause) của GC ("stop-the-world") là cực kỳ ngắn (thường dưới 1 mili giây).
- **Cơ chế:**
    1.  **Mark (Đánh dấu):** GC bắt đầu từ các gốc (global variables, stacks của goroutines) và duyệt qua cây đối tượng, đánh dấu tất cả các đối tượng có thể truy cập được. Giai đoạn này chạy đồng thời với chương trình chính.
    2.  **Sweep (Quét dọn):** Sau khi đánh dấu xong, GC quét qua heap và giải phóng tất cả các đối tượng không được đánh dấu.

Điều này cho phép chương trình tiếp tục chạy trong phần lớn thời gian GC hoạt động, giảm thiểu ảnh hưởng đến hiệu năng.

### 22. Giải thích về Go Runtime và Go Scheduler
- **Go Runtime:** Là môi trường thực thi quản lý các khía cạnh cốt lõi của một chương trình Go. Nó bao gồm GC, quản lý goroutine, và Go Scheduler.
- **Go Scheduler:** Chịu trách nhiệm phân phối các goroutine lên các luồng của hệ điều hành (OS threads) để thực thi. Nó sử dụng mô hình **M:N**, ánh xạ M goroutines lên N OS threads (N thường bằng số lõi CPU).
    - **G (Goroutine):** Tác vụ thực thi nhẹ.
    - **M (Machine):** Một luồng của OS.
    - **P (Processor):** Đại diện cho một tài nguyên thực thi (context), cần thiết để M chạy G. Số P mặc định bằng `GOMAXPROCS`.
Scheduler của Go có khả năng "trộm việc" (work-stealing), khi một P hết goroutine, nó sẽ "trộm" goroutine từ P khác để đảm bảo các lõi CPU luôn bận rộn.

### 23. Làm sao để giám sát việc sử dụng RAM của ứng dụng Go?
- **`runtime.ReadMemStats`:** Lấy các chỉ số bộ nhớ chi tiết từ runtime một cách programmatic, như `Alloc` (bộ nhớ heap đang được sử dụng), `TotalAlloc` (tổng bộ nhớ đã cấp phát), `Sys` (tổng bộ nhớ lấy từ OS), `NumGC`, v.v.
- **`pprof`:** Công cụ profiling mạnh mẽ nhất của Go. Gói `net/http/pprof` cho phép bạn xem thông tin heap profile (`/debug/pprof/heap`) qua HTTP, cho thấy các đối tượng nào đang chiếm nhiều bộ nhớ nhất và được cấp phát từ đâu.
- **`expvar`:** Một cách đơn giản để expose các biến tùy chỉnh (ví dụ: bộ đếm request, goroutine đang hoạt động) qua một endpoint HTTP JSON (`/debug/vars`).

### 24. Cách tối ưu việc sử dụng Heap trong Go?
Mục tiêu chính là **giảm số lần cấp phát bộ nhớ**, vì mỗi lần cấp phát đều tạo gánh nặng cho GC.
- **Sử dụng `sync.Pool`:** Tái sử dụng các đối tượng thường xuyên được cấp phát và giải phóng (ví dụ: buffer) để tránh cấp phát mới liên tục.
- **Giảm cấp phát trong hot path:** Trong các vòng lặp hoặc các hàm được gọi thường xuyên, hãy cố gắng cấp phát bộ nhớ trước và tái sử dụng nó.
- **Sử dụng đúng kích thước:** Khi tạo slice hoặc map với `make`, hãy cung cấp kích thước dự kiến để tránh việc cấp phát lại bộ nhớ khi chúng phát triển.
- **Sử dụng con trỏ một cách cẩn thận:** Truyền các struct lớn bằng con trỏ để tránh sao chép toàn bộ dữ liệu.

---

## IV. Hệ thống & Backend (System & Backend)

### 25. gRPC và gRPC Gateway là gì?
- **gRPC (gRPC Remote Procedure Call):** Là một framework RPC (Remote Procedure Call) hiệu năng cao, mã nguồn mở do Google phát triển.
    - **Đặc điểm:** Sử dụng Protocol Buffers (protobuf) để định nghĩa service và cấu trúc dữ liệu, giao tiếp qua HTTP/2, hỗ trợ streaming hai chiều.
    - **Ưu điểm:** Hiệu năng rất cao do sử dụng binary serialization, nhẹ hơn nhiều so với JSON/XML. Lý tưởng cho giao tiếp giữa các microservices (backend-to-backend).
    - **Nhược điểm:** Không thân thiện với trình duyệt, yêu cầu chia sẻ file `.proto` để tạo client, tạo ra sự phụ thuộc giữa các services.
- **gRPC-Gateway:** Là một plugin cho `protoc` (trình biên dịch protobuf) giúp tự động tạo một **reverse-proxy** server. Proxy này dịch một RESTful JSON API thành các lời gọi gRPC. Điều này cho phép bạn cung cấp cả API gRPC (cho internal services) và API RESTful (cho client như trình duyệt) từ một định nghĩa service duy nhất.

### 26. Cách viết Unit Test trong Golang
Go có hỗ trợ testing tích hợp sẵn trong bộ công cụ.
- **`testing` package:** Cung cấp các công cụ cơ bản để viết test.
- **Lệnh `go test`:** Tự động tìm và chạy các file test.
- **Quy ước:**
    - File test có tên kết thúc bằng `_test.go` (ví dụ: `app_test.go`).
    - Hàm test có tên bắt đầu bằng `Test` và nhận `*testing.T` làm tham số (ví dụ: `func TestMyFunction(t *testing.T)`).
    - Sử dụng các hàm như `t.Errorf()`, `t.Fatalf()`, `t.Logf()` để báo cáo kết quả.
- **Table-Driven Tests:** Một pattern phổ biến để chạy nhiều test case cho cùng một hàm, giúp code test ngắn gọn và dễ mở rộng.
- **Mocking & Stubs:** Đối với các phụ thuộc bên ngoài (database, API), sử dụng interfaces để định nghĩa hành vi và tạo các "mock" hoặc "stub" (phiên bản giả) trong test để cô lập thành phần đang được kiểm thử.

### 27. Các phương pháp tối ưu truy vấn cơ sở dữ liệu
(Câu trả lời gốc đã khá tốt, được tóm tắt và sắp xếp lại)
1.  **Giới hạn kết quả trả về:**
    - **`SELECT` cụ thể các cột:** Tránh `SELECT *`. Chỉ lấy những cột bạn thực sự cần.
    - **`LIMIT` số lượng hàng:** Sử dụng `LIMIT` (và `OFFSET` để phân trang) để hạn chế số bản ghi trả về.
2.  **Đánh chỉ mục (Indexing) thông minh:**
    - Đánh index cho các cột thường xuyên được dùng trong mệnh đề `WHERE`, `JOIN`, và `ORDER BY`.
    - Index giúp tăng tốc `SELECT` nhưng làm chậm `INSERT`, `UPDATE`, `DELETE`. Cần cân bằng.
    - Tránh các toán tử phủ định (`!=`, `NOT IN`, `NOT LIKE`) trên cột có index.
3.  **Sử dụng `LIKE` hợp lý:** Tránh bắt đầu pattern `LIKE` bằng ký tự đại diện (`%value`). Điều này ngăn cản DB sử dụng index.
4.  **`EXISTS` vs `IN` vs `COUNT`:**
    - Khi chỉ cần kiểm tra sự tồn tại, `EXISTS` thường hiệu quả hơn `IN` hoặc `COUNT(*)`.
    - Giữa `IN` và `BETWEEN` trên một khoảng số, `BETWEEN` thường nhanh hơn.
5.  **`UNION` vs `UNION ALL`:** Nếu bạn chắc chắn không có dữ liệu trùng lặp giữa các câu lệnh `SELECT`, hãy dùng `UNION ALL`. Nó nhanh hơn vì không cần thực hiện bước loại bỏ trùng lặp.
6.  **Sử dụng Replication:** Với các hệ thống có lượng đọc lớn, sử dụng mô hình Master-Slave. Gửi tất cả các truy vấn ghi (`INSERT`, `UPDATE`, `DELETE`) đến Master và các truy vấn đọc (`SELECT`) đến các Slave để phân tán tải.

---

## V. Kiến thức Khoa học Máy tính Tổng quát (General CS)

### 28. Giải thích các nguyên lý SOLID
SOLID là 5 nguyên tắc thiết kế hướng đối tượng giúp code dễ bảo trì, mở rộng và hiểu hơn.
- **S - Single Responsibility Principle (Nguyên lý Đơn trách nhiệm):**
    - Một struct hoặc package chỉ nên có một lý do duy nhất để thay đổi.
    - **Ví dụ Go:** Một struct `UserPersistence` chỉ nên chịu trách nhiệm lưu và đọc user từ DB, không nên chứa logic xác thực mật khẩu.
- **O - Open/Closed Principle (Nguyên lý Đóng/Mở):**
    - Có thể mở rộng hành vi của một thực thể (struct, function) mà không cần sửa đổi mã nguồn của nó.
    - **Ví dụ Go:** Sử dụng interface. Một hàm chấp nhận một interface `PaymentGateway` có thể làm việc với các phương thức thanh toán mới (Paypal, Stripe) mà không cần thay đổi hàm đó, miễn là các phương thức mới triển khai interface.
- **L - Liskov Substitution Principle (Nguyên lý Thay thế Liskov):**
    - Các đối tượng của một kiểu con có thể thay thế các đối tượng của kiểu cha mà không làm thay đổi tính đúng đắn của chương trình.
    - **Ví dụ Go:** Vì Go không có kế thừa, nguyên lý này áp dụng cho interface. Bất kỳ struct nào triển khai một interface đều phải tuân thủ "hợp đồng" của interface đó, đảm bảo hành vi nhất quán.
- **I - Interface Segregation Principle (Nguyên lý Phân tách Interface):**
    - Thay vì một interface lớn, cồng kềnh, hãy sử dụng nhiều interface nhỏ, chuyên biệt hơn. Client không nên bị buộc phải phụ thuộc vào các phương thức mà nó không sử dụng.
    - **Ví dụ Go:** Thay vì một interface `Worker` có cả `Work()` và `Sleep()`, hãy tách thành `Workable` và `Sleepable`.
- **D - Dependency Inversion Principle (Nguyên lý Đảo ngược Phụ thuộc):**
    - Các module cấp cao không nên phụ thuộc vào các module cấp thấp. Cả hai nên phụ thuộc vào abstraction (interface).
    - Abstraction không nên phụ thuộc vào chi tiết, mà chi tiết nên phụ thuộc vào abstraction.
    - **Ví dụ Go:** Một service `NotificationService` không nên phụ thuộc trực tiếp vào `SMSSender` (struct cụ thể), mà nên phụ thuộc vào interface `MessageSender`. Điều này cho phép dễ dàng thay thế `SMSSender` bằng `EmailSender`.

### 29. 7 cấu trúc dữ liệu cơ bản
(Câu trả lời gốc đã chính xác)
1.  **Mảng (Array):** Tập hợp các phần tử cùng kiểu, lưu trữ liên tiếp, truy cập bằng chỉ số. Kích thước cố định.
2.  **Danh sách liên kết (Linked List):** Chuỗi các nút, mỗi nút chứa dữ liệu và con trỏ đến nút tiếp theo. Kích thước động.
3.  **Ngăn xếp (Stack):** Cấu trúc LIFO (Last-In, First-Out). Thêm và xóa ở cùng một đầu.
4.  **Hàng đợi (Queue):** Cấu trúc FIFO (First-In, First-Out). Thêm ở cuối, xóa ở đầu.
5.  **Bảng băm (Hash Table):** Sử dụng hàm băm để ánh xạ key thành value. Truy cập trung bình O(1).
6.  **Cây (Tree):** Cấu trúc phân cấp cha-con (ví dụ: Cây nhị phân tìm kiếm).
7.  **Đồ thị (Graph):** Tập hợp các đỉnh và các cạnh nối chúng, biểu diễn các mối quan hệ.

### 30. 5 loại thuật toán cơ bản
(Câu trả lời gốc đã chính xác)
1.  **Thuật toán Sắp xếp (Sorting):** Sắp xếp các phần tử theo một thứ tự nhất định (Quick Sort, Merge Sort, Bubble Sort).
2.  **Thuật toán Tìm kiếm (Searching):** Tìm một phần tử trong một cấu trúc dữ liệu (Binary Search, Linear Search).
3.  **Thuật toán Đồ thị (Graph):** Duyệt hoặc tìm đường đi trên đồ thị (DFS, BFS, Dijkstra).
4.  **Thuật toán Quy hoạch động (Dynamic Programming):** Giải quyết các bài toán tối ưu bằng cách chia chúng thành các bài toán con và lưu trữ kết quả.
5.  **Thuật toán Cây (Tree):** Các thuật toán liên quan đến cấu trúc cây (duyệt cây, tìm kiếm, cân bằng cây).




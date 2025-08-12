## 🎯 **Hệ Thống Lending Tích Hợp Đa Đối Tác**

### **Mục đích và bối cảnh dự án**

Khi tôi tham gia dự án này, KiotViet đang đối mặt với một thách thức lớn: làm thế nào để kết nối hàng nghìn merchant với nhiều đối tác tài chính khác nhau một cách hiệu quả. Trước đây, mỗi khi merchant cần vay vốn, họ phải tự tìm hiểu và liên hệ từng ngân hàng riêng lẻ (công ty tài chính) - một quy trình phức tạp, tốn thời gian và không hiệu quả.

**Vấn đề cốt lõi được giải quyết:**
- Tiểu thương khó tiếp cận vốn ngân hàng do thủ tục phức tạp
- Mỗi ngân hàng có quy trình khác nhau, merchant phải làm nhiều lần
- Thiếu hệ thống đánh giá tín dụng thông minh dựa trên dữ liệu kinh doanh thực tế

### **🔥 Những thách thức kỹ thuật phức tạp đã giải quyết**

#### **1. Multi-Bank Integration Challenge**
Tích hợp với ngân hàng (VPBank, MBBank), công ty tài chính (tinvay) - mỗi bên có:
- API format khác nhau
- Quy trình approval khác nhau  
- Yêu cầu security khác nhau
Đồng bộ dữ liệu giữa nhiều systems:
- CMS → CRM/BI → Partner

#### **2. Bảo Mật Dữ Liệu Nhạy Cảm**
"Hệ thống xử lý thông tin tài chính cực kỳ nhạy cảm - CMND, thông tin thu nhập, tài sản. Một lỗ hổng bảo mật có thể gây thiệt hại nghiêm trọng"

### **💡 Những kỹ năng và công nghệ học được**

**Tech Stack hiện đại:**
- **Backend:** Gin Framework, MongoDB, Redis, Clean Architecture với Dependency Injection
- **Observability:** Grafana, Prometheus, Loki
- **Security:** JWT, RSA/AES, DES encryption, OAuth2
- **Integration:** REST APIs
- **Infrastructure:** Redis caching, SFTP, MinIO
- **Message Queue:** taskq PubSub
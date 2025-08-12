phỏng vấn senior backend dev
1. khác biệt giữa left-join, right-join, inner join
2. vì sao index làm tăng việc truy suất thông tin
3. nếu dùng nhiều index có ảnh hưởng gì không ?
4. khác biệt giữa delete, truncate, drop
5. Normalization trong db là gì ?
6. thiết kế bảng xyz, bắt giải thích chi tiết
7. giải thích về transactions
8. khác biệt chi tiết về oltp, olap
9. làm sao chuyển live database không gây downtime
10. cách handle số lượng lớn connection db
11. chiến lược scale db, khi nào dùng sharding, replication
12. xử lý deadlocks thế nào?
13. wal file là gì ?
14. nếu hệ thống có multiple write, làm sao đảm bảo consistency
15. làm sao theo dõi, fix các câu queries chậm trong live production
16. nguyên lý CAP trong distributed system, sau đó sẽ tự triển khai ý của ứng viên
ví dụ: db thì mọi người sử dụng cơ chế nào để scale db. Master slave hay sharding. Nếu master slave thì đang tập trung vào availability và partition tolerance. Với các cơ chế bầu master node, hãy tại sao tính consistency đôi khi không đảm bảo được ? Còn về sharding, cũng là AP và hi sinh tính consistency (tùy db). Tại sao sharding tốt như vậy, hỗ trợ cả read và write nặng nhưng ít khi sử dụng ở size công ty nhỏ. Vì sao vận hành nó phức tạp. nếu bạn thực sự làm hãy trả lời các câu hỏi trên ? Tương tự với message queue như Kafka cũng vậy, đáp ứng những tiêu chí CAP như thế nào ? hy sinh ở đâu và tại sao. Do đó hãy giải thích hiểu bản chất vận hành của hệ thống, sau đó giải thích công nghệ, lý thuyết công nghệ ? Hay đơn giản tại sao chúng ta phải distributed system ?
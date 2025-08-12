# PHÂN TÍCH SÂU CÁC CÂU HỎI PHỎNG VẤN SENIOR BACKEND DEVELOPER

## PHẦN I: DATABASE FUNDAMENTALS

### 1. Khác biệt giữa LEFT JOIN, RIGHT JOIN, INNER JOIN

**BẢN CHẤT VẬN HÀNH:**
- **INNER JOIN**: Chỉ trả về bản ghi khi có dữ liệu khớp ở cả 2 bảng
- **LEFT JOIN**: Trả về tất cả bản ghi từ bảng trái + bản ghi khớp từ bảng phải (NULL nếu không khớp)
- **RIGHT JOIN**: Trả về tất cả bản ghi từ bảng phải + bản ghi khớp từ bảng trái (NULL nếu không khớp)

**LÝ THUYẾT CÔNG NGHỆ:**
```sql
-- INNER JOIN: Giao của 2 tập hợp
SELECT u.name, o.amount 
FROM users u INNER JOIN orders o ON u.id = o.user_id;

-- LEFT JOIN: Tập hợp A + phần giao với B
SELECT u.name, o.amount 
FROM users u LEFT JOIN orders o ON u.id = o.user_id;
```

**IMPACT VẬN HÀNH HỆ THỐNG:**
- INNER JOIN nhanh nhất vì ít dữ liệu xử lý
- LEFT/RIGHT JOIN có thể tạo ra kết quả lớn hơn, ảnh hưởng memory
- Query planner sẽ chọn thuật toán join khác nhau (Hash Join, Nested Loop, Sort-Merge)

### 2. Vì sao Index làm tăng việc truy xuất thông tin

**BẢN CHẤT VẬN HÀNH:**
Index là cấu trúc dữ liệu bổ sung (thường là B-Tree) tạo "shortcut" đến dữ liệu thực tế.

**LÝ THUYẾT CÔNG NGHỆ:**
- **B-Tree Structure**: Độ phức tạp O(log n) thay vì O(n) của full table scan
- **Page-based Access**: Index giúp đọc ít pages từ disk
- **Clustered vs Non-clustered**: Clustered index sắp xếp dữ liệu vật lý theo key

**MECHANISM:**
```
Không có index: Scan toàn bộ table (1M records = 1M comparisons)
Có index B-Tree: log₂(1M) ≈ 20 comparisons
```

**VẬN HÀNH THỰC TẾ:**
- Index được load vào memory (buffer pool)
- Giảm I/O operations từ disk
- Query optimizer sử dụng statistics để chọn index tối ưu

### 3. Nếu dùng nhiều Index có ảnh hưởng gì không?

**BẢN CHẤT VẬN HÀNH:**
Mỗi index là overhead cho mọi INSERT/UPDATE/DELETE operation.

**TRADE-OFFS:**
**Pros:**
- Tăng tốc SELECT queries
- Hỗ trợ nhiều query patterns khác nhau

**Cons:**
- **Write Penalty**: Mỗi INSERT/UPDATE/DELETE phải update tất cả indexes
- **Storage Overhead**: Index chiếm thêm disk space
- **Memory Usage**: Nhiều index cạnh tranh buffer pool
- **Maintenance Cost**: Index fragmentation cần rebuild

**LÝ THUYẾT CÔNG NGHỆ:**
```
Write Amplification = 1 + số_lượng_index
Ví dụ: 5 indexes → mỗi INSERT thực tế là 6 operations
```

### 4. Khác biệt giữa DELETE, TRUNCATE, DROP

**BẢN CHẤT VẬN HÀNH:**

**DELETE:**
- Row-by-row removal
- Có thể rollback (transaction log)
- Triggers được execute
- Statistics không reset

**TRUNCATE:**
- Page-level removal
- Không thể rollback (minimal logging)
- Triggers KHÔNG execute
- Reset identity counter
- Nhanh hơn DELETE đáng kể

**DROP:**
- Remove entire table structure
- Metadata removal từ system catalog
- Không thể rollback structure

**VẬN HÀNH HỆ THỐNG:**
```sql
-- DELETE: Slow, logged, recoverable
DELETE FROM large_table WHERE condition;

-- TRUNCATE: Fast, minimal log, reset counters
TRUNCATE TABLE large_table;

-- DROP: Remove everything
DROP TABLE large_table;
```

### 5. Normalization trong Database

**BẢN CHẤT VẬN HÀNH:**
Normalization là quá trình tổ chức dữ liệu để giảm redundancy và improve data integrity.

**CÁC NORMAL FORMS:**

**1NF (First Normal Form):**
- Mỗi cell chứa atomic value
- Không có repeating groups

**2NF (Second Normal Form):**
- Đạt 1NF + mọi non-key attribute phụ thuộc hoàn toàn vào primary key
- Loại bỏ partial dependencies

**3NF (Third Normal Form):**
- Đạt 2NF + không có transitive dependencies
- Non-key attributes không phụ thuộc vào nhau

**BCNF (Boyce-Codd Normal Form):**
- Mọi determinant đều là candidate key

**VẬN HÀNH THỰC TẾ:**
```sql
-- Denormalized (0NF)
orders: id, customer_name, customer_email, product_name, product_price

-- Normalized (3NF)
customers: id, name, email
products: id, name, price
orders: id, customer_id, product_id
```

**TRADE-OFFS:**
- **Pros**: Giảm redundancy, improve consistency, save storage
- **Cons**: Nhiều JOINs, complex queries, performance impact

## PHẦN II: ADVANCED DATABASE CONCEPTS

### 6. Thiết kế bảng XYZ - Giải thích chi tiết

**BẢN CHẤT VẬN HÀNH:**
Database design phải cân bằng giữa normalization và performance.

**DESIGN PRINCIPLES:**
1. **Identify Entities**: Users, Products, Orders, etc.
2. **Define Relationships**: One-to-Many, Many-to-Many
3. **Choose Keys**: Primary, Foreign, Composite keys
4. **Apply Constraints**: NOT NULL, UNIQUE, CHECK
5. **Consider Indexes**: Query patterns analysis

**VÍ DỤ E-COMMERCE DESIGN:**
```sql
-- Users table
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Products table
CREATE TABLE products (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INT NOT NULL DEFAULT 0,
    category_id BIGINT,
    INDEX idx_category (category_id),
    INDEX idx_price (price)
);

-- Orders table
CREATE TABLE orders (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status ENUM('pending', 'confirmed', 'shipped', 'delivered') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_user_status (user_id, status),
    INDEX idx_created_at (created_at)
);
```

### 7. Giải thích về Transactions

**BẢN CHẤT VẬN HÀNH:**
Transaction là unit of work đảm bảo ACID properties.

**ACID PROPERTIES:**

**Atomicity (Tính nguyên tử):**
- All-or-nothing execution
- Rollback nếu có lỗi

**Consistency (Tính nhất quán):**
- Database từ consistent state này sang consistent state khác
- Constraints được maintain

**Isolation (Tính cô lập):**
- Concurrent transactions không ảnh hưởng lẫn nhau
- Isolation levels: READ UNCOMMITTED, READ COMMITTED, REPEATABLE READ, SERIALIZABLE

**Durability (Tính bền vững):**
- Committed changes persist ngay cả khi system crash
- Write-ahead logging (WAL)

**VẬN HÀNH HỆ THỐNG:**
```sql
BEGIN TRANSACTION;
    UPDATE accounts SET balance = balance - 100 WHERE id = 1;
    UPDATE accounts SET balance = balance + 100 WHERE id = 2;
    
    -- Nếu có lỗi, tất cả sẽ rollback
    IF @@ERROR <> 0
        ROLLBACK TRANSACTION;
    ELSE
        COMMIT TRANSACTION;
```

**ISOLATION LEVELS VÀ PROBLEMS:**
- **Dirty Read**: Đọc uncommitted data
- **Non-repeatable Read**: Cùng query khác kết quả trong transaction
- **Phantom Read**: Số lượng rows thay đổi

### 8. Khác biệt chi tiết OLTP vs OLAP

**BẢN CHẤT VẬN HÀNH:**

**OLTP (Online Transaction Processing):**
- **Purpose**: Operational transactions
- **Queries**: Short, frequent, simple
- **Data**: Current, detailed
- **Users**: Many concurrent users
- **Response Time**: Milliseconds
- **Database Design**: Normalized

**OLAP (Online Analytical Processing):**
- **Purpose**: Business intelligence, reporting
- **Queries**: Long, complex, aggregated
- **Data**: Historical, summarized
- **Users**: Fewer analytical users
- **Response Time**: Seconds to minutes
- **Database Design**: Denormalized (star/snowflake schema)

**VẬN HÀNH HỆ THỐNG:**
```sql
-- OLTP Example
SELECT * FROM orders WHERE user_id = 12345 AND status = 'pending';

-- OLAP Example
SELECT 
    YEAR(order_date),
    MONTH(order_date),
    SUM(total_amount),
    COUNT(*),
    AVG(total_amount)
FROM orders 
WHERE order_date >= '2023-01-01'
GROUP BY YEAR(order_date), MONTH(order_date);
```

**ARCHITECTURE PATTERNS:**
- **Lambda Architecture**: Batch + Speed layers
- **Kappa Architecture**: Stream processing only
- **Data Lake**: Raw data storage
- **Data Warehouse**: Structured analytical storage

## PHẦN III: PRODUCTION DATABASE OPERATIONS

### 9. Làm sao chuyển Live Database không gây Downtime

**BẢN CHẤT VẬN HÀNH:**
Zero-downtime migration yêu cầu chiến lược phức tạp với multiple phases.

**STRATEGIES:**

**1. Blue-Green Deployment:**
```
Blue (Current) ←→ Green (New)
Traffic switch: DNS/Load Balancer
```

**2. Rolling Migration:**
- Migrate từng phần nhỏ
- Verify integrity từng bước
- Rollback plan cho mỗi phase

**3. Database Replication:**
```
Master → Slave (New Schema)
1. Setup replication
2. Apply schema changes to slave
3. Switch master role
4. Update application connections
```

**4. Application-Level Migration:**
```python
# Dual-write pattern
def create_order(order_data):
    # Write to old table
    old_db.orders.insert(order_data)
    
    # Write to new table (with error handling)
    try:
        new_db.orders_v2.insert(transform_data(order_data))
    except Exception as e:
        log_error(e)  # Don't fail transaction
```

**TOOLS VÀ TECHNIQUES:**
- **pt-online-schema-change** (Percona)
- **gh-ost** (GitHub)
- **Database triggers** cho data sync
- **Feature flags** cho gradual rollout

### 10. Cách handle số lượng lớn Connection Database

**BẢN CHẤT VẬN HÀNH:**
Database có giới hạn concurrent connections. Cần optimize connection management.

**STRATEGIES:**

**1. Connection Pooling:**
```python
# Connection Pool Configuration
pool = ConnectionPool(
    host='localhost',
    database='mydb',
    user='user',
    password='password',
    minconn=1,
    maxconn=20,
    blocking=True,
    timeout=30
)
```

**2. Connection Multiplexing:**
- **PgBouncer** (PostgreSQL)
- **ProxySQL** (MySQL)
- **Connection Router** patterns

**3. Read Replicas:**
```
Write: Master DB (few connections)
Read: Multiple Slaves (distribute load)
```

**4. Application-Level Optimizations:**
```python
# Connection per request (BAD)
def get_user(user_id):
    conn = create_connection()
    result = conn.execute("SELECT * FROM users WHERE id = %s", user_id)
    conn.close()
    return result

# Connection pooling (GOOD)
def get_user(user_id):
    with connection_pool.get_connection() as conn:
        return conn.execute("SELECT * FROM users WHERE id = %s", user_id)
```

**MONITORING VÀ TUNING:**
```sql
-- PostgreSQL
SELECT * FROM pg_stat_activity;
SELECT count(*) FROM pg_stat_activity;

-- MySQL
SHOW PROCESSLIST;
SHOW STATUS LIKE 'Threads_connected';
```

### 11. Chiến lược Scale Database - Sharding vs Replication

**BẢN CHẤT VẬN HÀNH:**

**VERTICAL SCALING (Scale Up):**
- Tăng CPU, RAM, Storage
- Đơn giản nhưng có limit
- Single point of failure

**HORIZONTAL SCALING (Scale Out):**

**1. READ REPLICAS (Master-Slave):**
```
Master (Write) → Slave 1 (Read)
              → Slave 2 (Read)
              → Slave 3 (Read)
```

**Pros:**
- Đơn giản implement
- Tốt cho read-heavy workloads
- Backup tự động

**Cons:**
- Write bottleneck tại master
- Replication lag
- Eventual consistency

**2. SHARDING (Horizontal Partitioning):**
```
User ID 1-1000    → Shard 1
User ID 1001-2000 → Shard 2
User ID 2001-3000 → Shard 3
```

**SHARDING STRATEGIES:**
- **Range-based**: Chia theo range giá trị
- **Hash-based**: Hash function để distribute
- **Directory-based**: Lookup service

**Pros:**
- Scale cả read và write
- No single point of failure
- Linear scalability

**Cons:**
- Complex application logic
- Cross-shard queries khó khăn
- Rebalancing complexity
- Operational overhead

**KHI NÀO DÙNG GÌ:**
- **Replication**: Read-heavy, < 10K QPS writes
- **Sharding**: Write-heavy, > 10K QPS, need linear scale

**VÍ DỤ IMPLEMENTATION:**
```python
class ShardManager:
    def __init__(self, shards):
        self.shards = shards
    
    def get_shard(self, user_id):
        shard_key = hash(user_id) % len(self.shards)
        return self.shards[shard_key]
    
    def execute_query(self, user_id, query):
        shard = self.get_shard(user_id)
        return shard.execute(query)
```

### 12. Xử lý Deadlocks thế nào?

**BẢN CHẤT VẬN HÀNH:**
Deadlock xảy ra khi 2+ transactions chờ lẫn nhau, tạo circular dependency.

**DEADLOCK SCENARIO:**
```
Transaction 1: Lock A → Wait for Lock B
Transaction 2: Lock B → Wait for Lock A
Result: Circular wait → Deadlock
```

**DETECTION VÀ RESOLUTION:**

**1. Database Level:**
```sql
-- PostgreSQL: Deadlock detection
SHOW deadlock_timeout;  -- Default 1s

-- MySQL: Check deadlock info
SHOW ENGINE INNODB STATUS;
```

**2. Application Level:**
```python
def transfer_money(from_account, to_account, amount):
    max_retries = 3
    for attempt in range(max_retries):
        try:
            with transaction():
                # Always lock accounts in consistent order
                first_account = min(from_account, to_account, key=lambda x: x.id)
                second_account = max(from_account, to_account, key=lambda x: x.id)
                
                first_account.lock()
                second_account.lock()
                
                # Perform transfer
                from_account.balance -= amount
                to_account.balance += amount
                
            break  # Success
        except DeadlockException:
            if attempt == max_retries - 1:
                raise
            time.sleep(random.uniform(0.1, 0.5))  # Random backoff
```

**PREVENTION STRATEGIES:**
1. **Lock Ordering**: Luôn lock resources theo thứ tự consistent
2. **Timeout**: Set reasonable lock timeout
3. **Reduce Transaction Time**: Minimize lock duration
4. **Appropriate Isolation Level**: Sử dụng isolation level phù hợp

### 13. WAL File là gì?

**BẢN CHẤT VẬN HÀNH:**
WAL (Write-Ahead Logging) là mechanism đảm bảo durability và recovery.

**HOẠT ĐỘNG:**
1. Mọi changes được write vào log TRƯỚC khi write vào data files
2. Log entries có sequence number
3. Checkpoint process flush dirty pages từ memory ra disk
4. Recovery sử dụng WAL để replay changes

**WAL STRUCTURE:**
```
LSN (Log Sequence Number) | Transaction ID | Operation | Old Value | New Value
00000001                  | TXN_123       | UPDATE    | balance=100 | balance=150
00000002                  | TXN_124       | INSERT    | NULL       | user_data
```

**BENEFITS:**
- **Durability**: Changes persist ngay cả khi crash
- **Performance**: Sequential writes nhanh hơn random writes
- **Recovery**: Point-in-time recovery
- **Replication**: WAL shipping cho standby servers

**TUNING WAL:**
```sql
-- PostgreSQL
wal_level = replica
max_wal_size = 1GB
min_wal_size = 80MB
checkpoint_completion_target = 0.9
```

### 14. Multiple Write - Đảm bảo Consistency

**BẢN CHẤT VẬN HÀNH:**
Multiple writers có thể tạo ra race conditions và data inconsistency.

**PROBLEMS:**
- **Lost Updates**: T1 và T2 cùng update, một update bị mất
- **Dirty Reads**: T1 đọc uncommitted data từ T2
- **Inconsistent Reads**: T1 đọc data đang được T2 modify

**SOLUTIONS:**

**1. Database Level:**
```sql
-- Optimistic Locking
UPDATE accounts 
SET balance = balance - 100, version = version + 1
WHERE id = 123 AND version = @current_version;

-- Pessimistic Locking
SELECT balance FROM accounts WHERE id = 123 FOR UPDATE;
UPDATE accounts SET balance = balance - 100 WHERE id = 123;
```

**2. Application Level:**
```python
# Optimistic Locking Pattern
def update_account(account_id, amount):
    while True:
        account = get_account(account_id)
        current_version = account.version
        new_balance = account.balance + amount
        
        rows_affected = db.execute("""
            UPDATE accounts 
            SET balance = %s, version = version + 1
            WHERE id = %s AND version = %s
        """, [new_balance, account_id, current_version])
        
        if rows_affected > 0:
            break  # Success
        else:
            # Retry with exponential backoff
            time.sleep(random.uniform(0.01, 0.1))
```

**3. Distributed Consensus:**
- **2-Phase Commit (2PC)**
- **Raft Consensus**
- **SAGA Pattern** cho distributed transactions

### 15. Theo dõi và Fix Slow Queries trong Production

**BẢN CHẤT VẬN HÀNH:**
Slow queries là bottleneck chính trong production systems.

**MONITORING TOOLS:**

**1. Database Built-in:**
```sql
-- PostgreSQL: pg_stat_statements
SELECT query, calls, total_time, mean_time, rows
FROM pg_stat_statements
ORDER BY total_time DESC
LIMIT 10;

-- MySQL: Performance Schema
SELECT * FROM performance_schema.events_statements_summary_by_digest
ORDER BY SUM_TIMER_WAIT DESC
LIMIT 10;
```

**2. Application Level:**
```python
import time
from functools import wraps

def monitor_query(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start_time = time.time()
        try:
            result = func(*args, **kwargs)
            execution_time = time.time() - start_time
            
            if execution_time > 1.0:  # Log slow queries
                logger.warning(f"Slow query detected: {func.__name__}, "
                             f"execution_time: {execution_time:.2f}s")
            
            return result
        except Exception as e:
            logger.error(f"Query failed: {func.__name__}, error: {e}")
            raise
    return wrapper

@monitor_query
def get_user_orders(user_id):
    return db.execute("SELECT * FROM orders WHERE user_id = %s", [user_id])
```

**OPTIMIZATION STRATEGIES:**

**1. Query Analysis:**
```sql
-- PostgreSQL: EXPLAIN ANALYZE
EXPLAIN (ANALYZE, BUFFERS) 
SELECT * FROM orders o 
JOIN users u ON o.user_id = u.id 
WHERE o.created_at > '2024-01-01';
```

**2. Index Optimization:**
```sql
-- Missing index detection
SELECT schemaname, tablename, attname, n_distinct, correlation
FROM pg_stats
WHERE schemaname = 'public'
AND n_distinct > 100
AND correlation < 0.1;

-- Create appropriate indexes
CREATE INDEX CONCURRENTLY idx_orders_created_at_user_id 
ON orders(created_at, user_id);
```

**3. Query Rewriting:**
```sql
-- Bad: Function in WHERE clause
SELECT * FROM orders WHERE YEAR(created_at) = 2024;

-- Good: Range condition
SELECT * FROM orders 
WHERE created_at >= '2024-01-01' 
AND created_at < '2025-01-01';
```

## PHẦN IV: DISTRIBUTED SYSTEMS VÀ CAP THEOREM

### 16. Nguyên lý CAP trong Distributed System

**BẢN CHẤT VẬN HÀNH:**
CAP Theorem: Trong distributed system, chỉ có thể đảm bảo 2 trong 3 tính chất:
- **Consistency**: Mọi node thấy cùng data tại cùng thời điểm
- **Availability**: System luôn response (có thể là error)
- **Partition Tolerance**: System hoạt động dù có network failures

**LÝ THUYẾT CÔNG NGHỆ:**
```
Network Partition xảy ra → Phải chọn:
- CP: Từ chối requests để maintain consistency
- AP: Accept requests nhưng có thể inconsistent
```

**PRACTICAL EXAMPLES:**

**1. Database Systems:**

**Master-Slave Replication (CA → AP khi partition):**
```
Normal: Master ←→ Slaves (Consistent + Available)
Partition: Master | Slaves (Choose Availability → Inconsistent reads)
```

**Sharding (AP):**
```
Shard 1: Users 1-1000
Shard 2: Users 1001-2000
Network partition → Some shards unavailable → Partial availability
```

**2. Message Queue (Kafka):**

**Kafka Design (AP):**
```
Producer → Broker 1 (Leader)
        → Broker 2 (Follower)  
        → Broker 3 (Follower)

Network partition:
- Producer vẫn write được vào available brokers
- Consumers có thể miss messages từ partitioned brokers
- Eventually consistent khi partition heals
```

**Configuration Trade-offs:**
```properties
# Favor Consistency
acks=all
min.insync.replicas=2

# Favor Availability  
acks=1
min.insync.replicas=1
```

**3. Distributed Databases:**

**MongoDB (CP):**
```javascript
// Primary election khi partition
rs.status()
{
  "set": "myReplicaSet",
  "members": [
    {"_id": 0, "name": "mongo1:27017", "stateStr": "PRIMARY"},
    {"_id": 1, "name": "mongo2:27017", "stateStr": "SECONDARY"},
    {"_id": 2, "name": "mongo3:27017", "stateStr": "SECONDARY"}
  ]
}

// Network partition → Majority side becomes PRIMARY
// Minority side becomes SECONDARY (read-only)
```

**Cassandra (AP):**
```sql
-- Tunable consistency
SELECT * FROM users WHERE id = 123 
USING CONSISTENCY QUORUM;  -- CP behavior

SELECT * FROM users WHERE id = 123 
USING CONSISTENCY ONE;     -- AP behavior
```

**VÌ SAO SHARDING PHỨC TÁP VẬN HÀNH:**

**1. Operational Complexity:**
```python
# Cross-shard queries
def get_user_with_orders(user_id):
    # Query user shard
    user_shard = get_shard_for_user(user_id)
    user = user_shard.get_user(user_id)
    
    # Query all order shards (expensive!)
    orders = []
    for shard in all_order_shards:
        shard_orders = shard.get_orders_by_user(user_id)
        orders.extend(shard_orders)
    
    return user, orders
```

**2. Rebalancing Challenges:**
```
Initial: Shard1[1-1000], Shard2[1001-2000]
Growth: Need Shard3
Rebalance: Shard1[1-666], Shard2[667-1333], Shard3[1334-2000]

Problems:
- Data migration during rebalancing
- Application downtime
- Hotspot redistribution
```

**3. Transaction Complexity:**
```python
# Distributed transaction across shards
def transfer_between_users(from_user_id, to_user_id, amount):
    from_shard = get_shard_for_user(from_user_id)
    to_shard = get_shard_for_user(to_user_id)
    
    if from_shard == to_shard:
        # Same shard - simple transaction
        with from_shard.transaction():
            from_shard.debit(from_user_id, amount)
            from_shard.credit(to_user_id, amount)
    else:
        # Cross-shard - need 2PC or SAGA
        coordinator = DistributedTransactionCoordinator()
        coordinator.begin()
        try:
            coordinator.prepare(from_shard, "debit", from_user_id, amount)
            coordinator.prepare(to_shard, "credit", to_user_id, amount)
            coordinator.commit()
        except:
            coordinator.rollback()
```

**TẠI SAO CẦN DISTRIBUTED SYSTEMS:**

**1. Scale Requirements:**
- Single machine limits: CPU, Memory, Storage, Network
- Linear scalability needs multiple machines

**2. Availability Requirements:**
- Single point of failure unacceptable
- Geographic distribution for disaster recovery

**3. Performance Requirements:**
- Latency: Data closer to users
- Throughput: Parallel processing

**4. Business Requirements:**
- Global presence
- Regulatory compliance (data locality)
- Cost optimization

**TRADE-OFFS ANALYSIS:**

**Monolithic vs Distributed:**
```
Monolithic:
+ Simple deployment
+ ACID transactions
+ No network latency
- Single point of failure
- Vertical scaling limits

Distributed:
+ Horizontal scalability  
+ Fault tolerance
+ Geographic distribution
- Network complexity
- Eventual consistency
- Operational overhead
```

## KẾT LUẬN

**BẢN CHẤT VẬN HÀNH HỆ THỐNG:**
1. **Performance**: Mọi optimization đều có trade-offs
2. **Consistency vs Availability**: Fundamental trade-off trong distributed systems
3. **Complexity Management**: Tools và patterns để handle complexity
4. **Monitoring**: Critical để understand system behavior

**LÝ THUYẾT CÔNG NGHỆ CORE:**
1. **Data Structures**: B-Trees, Hash tables, LSM trees
2. **Algorithms**: Consensus (Raft, Paxos), Consistent hashing
3. **Protocols**: TCP/IP, HTTP, gRPC, database protocols
4. **Mathematical Foundations**: CAP theorem, ACID properties, Consistency models

**PRACTICAL WISDOM:**
- Start simple, scale when needed
- Measure before optimizing
- Understand your specific use case
- Plan for failure scenarios
- Invest in observability early

Đây là foundation knowledge mà mọi senior backend developer cần nắm vững để design và operate production systems hiệu quả.

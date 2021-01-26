use test;

set session autocommit=0;

start transaction;

insert into t values(0);

insert into t values(2);

insert into t values(12);

## inset 7 will block, because 7 is in(3, 10)
insert into t values(7);


show engine innodb status;

#
# =====================================
# 2020-05-18 07:46:01 7fe8f52be700 INNODB MONITOR OUTPUT
# =====================================
# Per second averages calculated from the last 0 seconds
# -----------------
# BACKGROUND THREAD
# -----------------
# srv_master_thread loops: 5 srv_active, 0 srv_shutdown, 672 srv_idle
# srv_master_thread log flush and writes: 677
#     ----------
#     SEMAPHORES
#     ----------
#     OS WAIT ARRAY INFO: reservation count 10
#     OS WAIT ARRAY INFO: signal count 10
#     Mutex spin waits 1, rounds 30, OS waits 1
#     RW-shared spins 9, rounds 270, OS waits 9
#     RW-excl spins 0, rounds 0, OS waits 0
#     Spin rounds per wait: 30.00 mutex, 30.00 RW-shared, 0.00 RW-excl
#     ------------
#     TRANSACTIONS
#     ------------
#     Trx id counter 2326
# Purge done for trx's n:o < 2319 undo n:o < 0 state: running but idle
# History list length 6
# LIST OF TRANSACTIONS FOR EACH SESSION:
# ---TRANSACTION 2317, ACTIVE 59 sec
# 1 lock struct(s), heap size 360, 0 row lock(s), undo log entries 3
# MySQL thread id 9, OS thread handle 0x7fe8f52be700, query id 2081 172.17.0.1 root init
# /* ApplicationName=IntelliJ IDEA 2019.3.4 */ show engine innodb status
# ---TRANSACTION 2316, ACTIVE 102 sec
# 2 lock struct(s), heap size 360, 1 row lock(s)
# MySQL thread id 8, OS thread handle 0x7fe8f5300700, query id 1933 172.17.0.1 root
# ---TRANSACTION 2315, ACTIVE 273 sec
# MySQL thread id 4, OS thread handle 0x7fe8f5342700, query id 1774 172.17.0.1 root
# Trx read view will not see trx with id >= 2316, sees < 2316
# --------
# FILE I/O
# --------
# I/O thread 0 state: waiting for completed aio requests (insert buffer thread)
# I/O thread 1 state: waiting for completed aio requests (log thread)
# I/O thread 2 state: waiting for completed aio requests (read thread)
# I/O thread 3 state: waiting for completed aio requests (read thread)
# I/O thread 4 state: waiting for completed aio requests (read thread)
# I/O thread 5 state: waiting for completed aio requests (read thread)
# I/O thread 6 state: waiting for completed aio requests (write thread)
# I/O thread 7 state: waiting for completed aio requests (write thread)
# I/O thread 8 state: waiting for completed aio requests (write thread)
# I/O thread 9 state: waiting for completed aio requests (write thread)
# Pending normal aio reads: 0 [0, 0, 0, 0] , aio writes: 0 [0, 0, 0, 0] ,
#  ibuf aio reads: 0, log i/o's: 0, sync i/o's: 0
# Pending flushes (fsync) log: 0; buffer pool: 0
# 183 OS file reads, 82 OS file writes, 51 OS fsyncs
# 0.00 reads/s, 0 avg bytes/read, 0.00 writes/s, 0.00 fsyncs/s
# -------------------------------------
# INSERT BUFFER AND ADAPTIVE HASH INDEX
# -------------------------------------
# Ibuf: size 1, free list len 0, seg size 2, 0 merges
# merged operations:
#  insert 0, delete mark 0, delete 0
# discarded operations:
#  insert 0, delete mark 0, delete 0
# Hash table size 276671, node heap has 0 buffer(s)
# 0.00 hash searches/s, 0.00 non-hash searches/s
# ---
# LOG
# ---
# Log sequence number 1639769
# Log flushed up to   1639769
# Pages flushed up to 1639769
# Last checkpoint at  1639769
# 0 pending log writes, 0 pending chkp writes
# 26 log i/o's done, 0.00 log i/o's/second
# ----------------------
# BUFFER POOL AND MEMORY
# ----------------------
# Total memory allocated 137363456; in additional pool allocated 0
# Dictionary memory allocated 63828
# Buffer pool size   8191
# Free buffers       8013
# Database pages     178
# Old database pages 0
# Modified db pages  0
# Pending reads 0
# Pending writes: LRU 0, flush list 0, single page 0
# Pages made young 0, not young 0
# 0.00 youngs/s, 0.00 non-youngs/s
# Pages read 167, created 11, written 51
# 0.00 reads/s, 0.00 creates/s, 0.00 writes/s
# No buffer pool page gets since the last printout
# Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
# LRU len: 178, unzip_LRU len: 0
# I/O sum[0]:cur[0], unzip sum[0]:cur[0]
# --------------
# ROW OPERATIONS
# --------------
# 0 queries inside InnoDB, 0 queries in queue
# 1 read views open inside InnoDB
# Main thread process no. 1, id 140638518425344, state: sleeping
# Number of rows inserted 6, updated 0, deleted 0, read 3
# 0.00 inserts/s, 0.00 updates/s, 0.00 deletes/s, 0.00 reads/s
# ----------------------------
# END OF INNODB MONITOR OUTPUT
# ============================

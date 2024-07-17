
**定时触发任务是如何实现的？：使用时间轮实现**

- xxl_job_info表是记录定时任务的db表，里面有个trigger_next_time（Long）字段，表示下一次触发的时间点
- 任务时间被修改 / 每一次任务触发后，可以根据cronb表达式计算下一次触发时间戳：Date nextValidTime = new CronExpression(jobInfo.getJobCron()).getNextValidTimeAfter(new Date())），更新trigger_next_time字段
- 定时执行任务逻辑：

- 定时任务scheduleThread：不断从db把5秒内要执行的任务读出，立即触发 / 放到时间轮等待触发，并更新trigger_next_time

- 获取当前时间now
- 轮询db，找出trigger_next_time在距now 5秒内的任务
- （0）对到达now时间后的任务（超出now 5秒外）：

- 直接跳过不执行；
- 重置trigger_next_time

- （1）对到达now时间后的任务（超出now 5秒内）：

- 开线程执行触发逻辑；
- 若任务下一次触发时间是在5秒内，则放到时间轮内（Map<Integer, List<Integer>> 秒数(1-60) => 任务id列表）；
- 重置trigger_next_time

- （2）对未到达now时间的任务：

- 直接放到时间轮内；
- 重置trigger_next_time

- 定时任务ringThread：时间轮实现到点触发任务

- 时间轮数据结构：Map<Integer, List<Integer>> key是秒数(1-60) ，value是任务id列表

- 获取当前时间秒数
- 从时间轮内移出当前秒数前2个秒数（避免处理耗时太长，跨过刻度，向前校验一个刻度）的任务列表id，一一触发任务；



**2、当xxl-job应用本身集群部署（实现高可用HA）时，如何避免集群中的多个服务器同时调度任务？：通过mysql悲观锁实现分布式锁（for update语句）**

- setAutoCommit(false)关闭隐式自动提交事务，启动事务
- select lock for update（显式排他锁，其他事务无法进入&无法实现for update）
- 读db任务信息 -> 拉任务到内存时间轮 -> 更新db任务信息
- commit提交事务，同时会释放for update的排他锁（悲观锁）

  

**3、任务执行器注册中心是如何实现的？**

- 使用db表xxl_job_group记录下执行器的信息：执行器AppName、执行器名称title、执行器地址列表address_list(多地址逗号分隔)


**4、如何实现任务执行器的路由？**

- 执行器集群部署时提供丰富的路由策略，包括：第一个、最后一个、轮询、随机、一致性HASH、最不经常使用、最近最久未使用、故障转移、忙碌转移等；

- **第一个、最后一个、轮询、随机**：都是简单读address_list即可
- **一致性HASH**：TreeSet实现一致性hash算法
- **最不经常使用、最近最久未使用**：HashMap、LinkedHashMap
- **故障转移**：遍历address_list获取address时，逐个检查该address的心跳（请求返回状态）；只有心跳正常的address才返回使用
- **忙碌转移**：遍历address_list获取address时，逐个检查该address是否忙碌（请求返回状态）；只有状态为idle的address才返回使用



**5、如何实现任务分片、并行执行？**

- 拉出任务的执行机器列表，逐个设置index / total，把index / total分发到任务执行器
- 任务执行器可根据index / total参数开发分片任务



参考 
https://zhuanlan.zhihu.com/p/670687010



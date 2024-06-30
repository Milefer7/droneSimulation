# 无人机战斗小组问题模拟仿真

## 设计要求

* 一个无人机战斗小组，其中侦察无人机M个，战斗无人机N个，战斗小组协调控制中心1个。侦察无人机负责侦察战场环境情况，并且获得的信息发往协调控制中心，战斗无人机从协调控制中心获取情报，根据信息内容，执行战斗任务。
* 允许多个战斗无人机同时从协调控制中心获取情报；
  * goroutine
* 每次只允许一个侦察无人机向协调控制中心提供情报；任何侦察无人机在完成情报提交之前，协调控制中心暂停服务；侦察无人机上传情报信息的任务优先级最高。根据题目描述利用进程相关知识模拟实现战斗过程。

## 项目设计

这个程序模拟了一个控制中心 (ControlCenter) 接收和分发情报给战斗无人机和侦察无人机的过程。控制中心实现了读者-写者问题，其中侦察无人机作为写者，战斗无人机作为读者。

### 三个模块：

1. **执行控制模块**：
   - 存储情报并处理任务编号。
   - 通过读写锁 (`sync.RWMutex`) 确保并发安全。
   - 使用条件变量 (`sync.Cond`) 来实现读写等待机制。
2. **侦察无人机 (Scout Drone)模块**：
   - 生成并提交情报给控制中心。
   - 每次提交情报后，会等待一段时间再提交下一次情报。
3. **战斗无人机 (Combat Drone)模块**：
   - 请求获取最新情报并进行处理。
   - 每次处理完情报后，会等待一段时间再请求下一次情报。

### 时序图

![image-20240622012030162](https://my-note-drawing-bed-1322822796.cos.ap-shanghai.myqcloud.com/picture/202406220120290.png)

**用户启动程序**：

- 用户（User）启动了主程序（Main），触发了整个流程。

**主程序初始化控制中心**：

- 主程序调用 `NewControlCenter()` 来初始化控制中心（ControlCenter）。

**启动侦察无人机和战斗无人机的循环**：

- 主程序开始侦察无人机（ScoutDrone）和战斗无人机（CombatDrone）的循环操作。

**侦察无人机的操作循环**：

- 侦察无人机进入循环，执行以下操作：
  - 向控制中心提交情报（SubmitIntelligence(intel, id)）。
  - 控制中心打印提交信息（Print submission message）。
  - 控制中心更新情报和任务计数器（Update intelligence and taskCounter）。
  - 控制中心广播新情报（Broadcast new intelligence）。
  - 侦察无人机进入休眠状态（Sleep）。

**战斗无人机的操作循环**：

- 战斗无人机进入循环，执行以下操作：
  - 请求情报（GetIntelligence()）。
  - 如果没有新情报，控制中心让战斗无人机等待（Wait if intelligence is nil）。
  - 控制中心返回情报（Return intelligence）。
  - 战斗无人机打印收到的情报信息（Print received message）。
  - 战斗无人机进入休眠状态（Sleep）。

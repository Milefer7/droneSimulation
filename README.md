# 无人机战斗小组问题模拟仿真

## 设计要求

一个无人机战斗小组，其中侦察无人机M个，战斗无人机N个，战斗小组协调控制中心1个。侦察无人机负责侦察战场环境情况，并且获得的信息发往协调控制中心，战斗无人机从协调控制中心获取情报，根据信息内容，执行战斗任务。允许多个战斗无人机同时从协调控制中心获取情报；每次只允许一个侦察无人机向协调控制中心提供情报；任何侦察无人机在完成情报提交之前，协调控制中心暂停服务；侦察无人机上传情报信息的任务优先级最高。根据题目描述利用进程相关知识模拟实现战斗过程。

## 项目设计
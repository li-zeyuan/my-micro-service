# my-micro-service

通过编写一个小项目，学习微服务，学习micro

参考：https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro

## 背景

- 已经从事了go一年左右web开发，觉得微服务很有价值，很神秘，是一种趋势
- 抱着学习微服务，学习micro的心态，参考GitHub上面的例子，编写这个项目

## 目标

- 掌握micro框架

- 对微服务有更深的理解

- 基于micro，能够编写一个属于自己项目

## 项目介绍

	### 项目架构

```sequence
用户端->web层: http请求
Note over web层: (1)登录、token颁发
Note over web层: (2)鉴权
web层->服务层: rpc调用

```



## 技术栈

## 业务模块

## 进度
- 加载db配置文件
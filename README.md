# Setu no!

## Definition

Setu服务器端v0.1

## Function

- [ ] 多线程处理客户端发来的请求（HTTP)
- [ ] 对每个请求，基于API要求返回数据或对setu数据库进行操作

## API

### 重定向

- `/setu/latest/(.*)` --> `/setu/v0.1/${0}`

### 基础请求

以下所有HTTP请求格式，均省略前面的`/setu/v0.1/`

- [ ] **GET** `/info` --> *JSON*
  - 返回当前服务端版本，其他信息（后续补充
- [ ] **GET** `/view?range=$range&sort=$sort` --> *JSON*
  - 返回所有可查看的图片ID，以JSON列表的格式
  - `$range` 
    - '' --> All pictures
    - ':10' --> Pic [0, 1, ..., 9]
    - '10:' --> Pic [10, 11, ..., end]
    - '10:20' --> Pic [10, 11, ..., 19]
  - `$sort`
    - '', 'D' --> 时间新的在前
    - 'A' --> 时间旧的在前
- [ ] **GET** `/view/$id` --> *JPEG/PNG/GIF*
  - 返回id为$id的图片
- [ ] **GET** `/view/$id/status` --> *JSON*
  - 返回$id图片信息
- [ ] **POST** `/upload`
  - 上传图片及其信息

## 后端返回状态码

**GET** 

- 200 (OK) 已经在响应中发出
- 204 (无内容) 返回空资源
- 301 (Moved Permanently) 资源的URI已经被更新
- 303 (See Other) 其他（如负载均衡）
- 304 (Not Modified) 资源未更改（客户端已经有缓存了）
- 400 (Bad Request) 参数错误之类
- 404 (Not Found) 你懂的
- 406 (Not Acceptable) 服务端不支持所需表示
- 500 (Internal Server Error) 通用服务器错误
- 503 (Service Unavailable) 服务端当前无法处理请求

**POST**

- 200 (OK) 现有资源已被修改
- 201 (Created) 如果新资源被创建
- 202 (Accepted)已接受请求但尚未完成（异步处理）
- 301 (Moved Permanently) 资源的URI已经被更新
- 303 (See Other) 其他（如负载均衡）
- 400 (Bad Request) 参数错误之类
- 404 (Not Found) 你懂的
- 406 (Not Acceptable) 服务端不支持所需表示
- 409 (Conflict) 通用冲突
- 412 (Precondition Failed) 前置条件失败（如执行条件更新时的冲突）
- 415 (Unsupported Media Type) 不支持的媒体种类
- 500 (Internal Server Error) 通用服务器错误
- 503 (Service Unavailable) 服务端当前无法处理请求

**PUT**

- 200 (OK) 现有资源已被修改
- 201 (Created) 如果新资源被创建
- 301 (Moved Permanently) 资源的URI已经被更新
- 303 (See Other) 其他（如负载均衡）
- 400 (Bad Request) 参数错误之类
- 404 (Not Found) 你懂的
- 406 (Not Acceptable) 服务端不支持所需表示
- 409 (Conflict) 通用冲突
- 412 (Precondition Failed) 前置条件失败（如执行条件更新时的冲突）
- 415 (Unsupported Media Type) 不支持的媒体种类
- 500 (Internal Server Error) 通用服务器错误
- 503 (Service Unavailable) 服务端当前无法处理请求

**DELETE**

- 200 (OK) 资源已被删除
- 301 (Moved Permanently) 资源的URI已经被更新
- 303 (See Other) 其他（如负载均衡）
- 400 (Bad Request) 参数错误之类
- 404 (Not Found) 你懂的
- 409 (Conflict) 通用冲突
- 500 (Internal Server Error) 通用服务器错误
- 503 (Service Unavailable) 服务端当前无法处理请求




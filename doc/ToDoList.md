---
title: ToDoList
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# ToDoList

Base URLs:

# Authentication

# 用户处理

## POST 注册接口

POST /todolist/user/signup

> Body 请求参数

```json
{
  "name": "水大改",
  "email": "ob5j6i.i37@qq.com",
  "password": "laborum quis"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 名称|名称|
|» email|body|string| 是 | 邮箱|邮箱|
|» password|body|string| 是 | 密码|密码|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

```json
{
  "message": "err",
  "content": "已有的邮箱"
}
```

```json
{
  "message": "err",
  "content": "已有的名字"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 登录接口

POST /todolist/user/signin

> Body 请求参数

```json
{
  "name": "",
  "email": "ob5j.i37@qq.com",
  "password": "laborum quis"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» name|body|string| 是 ||名称|
|» email|body|string| 是 ||none|
|» password|body|string| 是 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI1N2UyYTFmNy05YmQ1LTExZWYtODgzNy04NDVjZjNlY2RkMzYiLCJleHAiOjE3MzE0NTc0MjUsImlhdCI6MTczMDg1MjYyNX0.euHuxxwbfNNG0BChjNn1a2cejd1z6CjqTCjjd6DOeuQ"
  }
}
```

```json
{
  "message": "err",
  "content": "没有此用户"
}
```

```json
{
  "message": "err",
  "content": "密码错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||消息|
|» content|object|true|none||none|
|»» token|string|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

# 数据模型


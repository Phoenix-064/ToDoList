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
  "email": "3440480965@qq.com",
  "password": "123",
  "verification_code": "930028"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» email|body|string| 是 | 邮箱|邮箱|
|» password|body|string| 是 | 密码|密码|
|» verification_code|body|string| 是 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

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

# 注意验证时采用 Bearer Token 认证流程，后端只发送给前端 Token，但前端每次请求时要返回的是 Bearer Token，即：
## Bearer Token 格式:

标准的 Bearer Token 认证头部格式应该是：Authorization: Bearer `<token>`
例如：Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

> Body 请求参数

```json
{
  "email": "123",
  "password": "123"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
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
  "message": "ok",
  "content": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiOWY4N2U3YjEtOWMzYy0xMWVmLWIyMDItODQ1Y2YzZWNkZDM2IiwiaXNBZG1pbiI6ZmFsc2UsImV4cCI6MTczMTA1MTM1NSwiaWF0IjoxNzMwOTY0OTU1fQ.WMn9kt0mfZ47Y3NVyylnhjoqWza0qGQbgxmnsm3fto8"
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

## POST 发送验证码

POST /todolist/user/signup/send-code

还有几种异常，属于数据库或连接出错的类型。其中content会返回错误，message为err。

> Body 请求参数

```json
{
  "email": "3440480965@qq.com"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» email|body|string| 是 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": "发送验证码成功"
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
  "content": "设置收件人失败: 550 Invalid User: "
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|none|Inline|

### 返回数据结构

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 用户删除用户

POST /todolist/user/delete

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 是 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": "删除成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 修改密码

POST /todolist/user/change-password

> Body 请求参数

```json
{
  "FormerPassword": "123",
  "LaterPassword": "123"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» former_password|body|string| 是 ||none|
|» later_password|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "message": "string",
  "content": "string"
}
```

```json
{
  "message": "err",
  "content": "错误的原密码"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

# todoList

## GET 获取用户全部todos

GET /todolist

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||token|

> 返回示例

```json
{
  "message": "ok",
  "content": {
    "todos": []
  }
}
```

```json
{
  "message": "Excepteur fugiat enim ut irure",
  "content": {
    "todos": [
      {
        "id": "11",
        "event": "enim deserunt qui mollit ut",
        "completed": "do ut",
        "is_cycle": true,
        "description": "应种和观委。达动最。管必活去场东。少线加组物可。已存关改须目华县头样。等型飞干部并金。",
        "importance_level": 73
      },
      {
        "id": "82",
        "event": "occaecat deserunt reprehenderit in",
        "completed": "sint",
        "is_cycle": false,
        "description": "达那济通条义话式王。法斗况青知。",
        "importance_level": 68
      },
      {
        "id": "80",
        "event": "reprehenderit ut",
        "completed": "velit elit Ut irure sed",
        "is_cycle": false,
        "description": "运书总列却况办长消。最相新低一酸。及车节条干历据。利元北多整光义角油专。",
        "importance_level": 47
      }
    ]
  }
}
```

```json
{
  "message": "err",
  "content": "错误的请求格式"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||消息|
|» content|object|true|none||none|
|»» todos|[object]|true|none||none|
|»»» id|string|true|none||ID 编号|
|»»» event|string|true|none||none|
|»»» completed|boolean|true|none||none|
|»»» is_cycle|boolean|true|none||none|
|»»» description|string|true|none||none|
|»»» importance_level|integer|true|none||none|
|»»» completed_date|string|true|none||记录完成时间|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 添加一个todo

POST /todolist/add

> Body 请求参数

```json
{
  "id": "2",
  "event": "amet eiusmod in esse",
  "description": "使代证运自于。组容决车机包周。",
  "is_cycle": false,
  "importance_level": 41
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 ||ID 编号|
|» event|body|string| 是 ||none|
|» description|body|string| 是 ||none|
|» is_cycle|body|boolean| 是 ||none|
|» importance_level|body|integer| 是 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": "添加成功"
}
```

```json
{
  "message": "err",
  "content": "错误的请求格式"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|

### 返回数据结构

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 修改重要级（要提供所有 todos）

POST /todolist/updateImportanceLevel

> Body 请求参数

```json
{
  "todos": [
    {
      "id": "13",
      "event": "magna ut dolor eu",
      "completed": true,
      "is_cycle": false,
      "description": "志表得别细能华。京离即族龙参积。界头素断组现风。花王运加第行。",
      "importance_level": 87,
      "completed_date": "2025-05-30"
    },
    {
      "id": "13",
      "event": "magna ut dolor eu",
      "completed": true,
      "is_cycle": false,
      "description": "志表得别细能华。京离即族龙参积。界头素断组现风。花王运加第行。",
      "importance_level": 87,
      "completed_date": "2024-07-25"
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» todos|body|[object]| 是 ||none|
|»» id|body|string| 是 ||ID 编号|
|»» event|body|string| 是 ||none|
|»» completed|body|boolean| 是 ||none|
|»» is_cycle|body|boolean| 是 ||none|
|»» description|body|string| 是 ||none|
|»» importance_level|body|integer| 是 ||none|
|»» completed_date|body|string| 是 ||记录完成时间|

> 返回示例

```json
{
  "message": "ok",
  "content": "保存成功"
}
```

```json
{
  "message": "err",
  "content": "json: cannot unmarshal object into Go value of type []data.Todo"
}
```

```json
{
  "message": "err",
  "content": "token is expired by 24h58m53s"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|none|Inline|

### 返回数据结构

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 删除一个todo

POST /todolist/delete

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|query|string| 否 ||todo 的 ID 编号|
|Authorization|header|string| 否 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 修改一个 todo

POST /todolist/update

> Body 请求参数

```json
{
  "id": "2",
  "event": "Genshin",
  "completed": false,
  "is_cycle": false,
  "description": "Play Genshin Impact",
  "importance_level": 98,
  "completed_date": "2023-11-24"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 ||ID 编号|
|» event|body|string| 是 ||none|
|» completed|body|boolean| 是 ||none|
|» is_cycle|body|boolean| 是 ||none|
|» description|body|string| 是 ||none|
|» importance_level|body|integer| 是 ||none|
|» completed_date|body|string| 是 ||记录完成时间|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 记录完成时间

POST /todolist/record-finish-time

> Body 请求参数

```json
{
  "id": "2",
  "completed_date": "2024-10-10"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 ||ID 编号|
|» completed_date|body|string| 是 ||记录完成时间|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

# Wish

## GET 获取用户的随机一条 wish

GET /todolist/wish/random

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": {
    "id": "60",
    "event": "incididunt cillum dolor voluptate ullamco",
    "is_cycle": true,
    "description": "住被其心气思速难应层。见将持活质广。公示动色代。管他会料流火切。用中火面指学义识件。",
    "is_shared": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|object|true|none|发回什么都有可能|none|
|»» id|string|true|none||ID 编号|
|»» event|string|true|none||none|
|»» is_cycle|boolean|true|none||none|
|»» description|string|true|none||none|
|»» is_shared|boolean|true|none||none|

## GET 获取用户所有 wish

GET /todolist/wish

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": {
    "wishes": [
      {
        "id": "37",
        "event": "sed Ut cupidatat dolor",
        "is_cycle": true,
        "description": "用而自见。七却拉般起断亲京因。器管统今行。造以空化水机眼所。运动前极。",
        "is_shared": false
      },
      {
        "id": "59",
        "event": "fugiat adipisicing ut consequat cillum",
        "is_cycle": false,
        "description": "产论说关。人用与或除。对前列到快素气界路。",
        "is_shared": true
      },
      {
        "id": "13",
        "event": "non consectetur aliquip nulla",
        "is_cycle": true,
        "description": "都心队大因。看京量领去。相场众线因日叫。",
        "is_shared": false
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|object|true|none|发回什么都有可能|none|
|»» wishes|[object]|true|none||none|
|»»» id|string|true|none||ID 编号|
|»»» event|string|true|none||none|
|»»» is_cycle|boolean|true|none||none|
|»»» description|string|true|none||none|
|»»» is_shared|boolean|true|none||none|

## POST 删除一条 wish

POST /todolist/wish/delete

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|query|string| 否 ||wish 的 ID 编号|
|Authorization|header|string| 否 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 修改一个 wish

POST /todolist/wish/update

> Body 请求参数

```json
{
  "id": "86",
  "event": "voluptate et ut",
  "is_cycle": true,
  "description": "光入运值值基。",
  "is_shared": false
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 ||ID 编号|
|» event|body|string| 是 ||none|
|» is_cycle|body|boolean| 是 ||none|
|» description|body|string| 是 ||none|
|» is_shared|body|boolean| 是 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 添加一个 wish

POST /todolist/wish/add

> Body 请求参数

```json
{
  "id": "3",
  "event": "voluptate et ut",
  "is_cycle": true,
  "description": "产认清。只名质感新技备变。光入运值值基。",
  "is_shared": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 ||ID 编号|
|» event|body|string| 是 ||none|
|» is_cycle|body|boolean| 是 ||none|
|» description|body|string| 是 ||none|
|» is_shared|body|boolean| 是 ||none|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 将一个 wish 添加至待办

POST /todolist/wish/add-todo

> Body 请求参数

```json
{
  "id": "456"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 ||wish ID 编号|

> 返回示例

```json
{
  "message": "ok",
  "content": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

# 社区

## GET 获取所有社区里的 wish

GET /todolist/community

> 返回示例

```json
{
  "message": "cupidatat et",
  "content": [
    {
      "id": "26",
      "event": "adipisicing incididunt elit aute labore",
      "description": "场省组设都队。任真强利第。接则美。",
      "viewed": 91
    },
    {
      "id": "13",
      "event": "occaecat sed tempor sint ea",
      "description": "始部走单运好议。青导义。",
      "viewed": 50
    },
    {
      "id": "50",
      "event": "aliqua fugiat veniam elit",
      "description": "明查者。眼已深着说。收了离据适很消程。通儿难种想。",
      "viewed": 82
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|[[SharedWish](#schemasharedwish)]|true|none|发回什么都有可能|none|
|»» id|string|true|none||none|
|»» event|string|true|none||none|
|»» description|string|true|none||none|
|»» viewed|integer|true|none||none|

## POST 增加一次 viewed

POST /todolist/community/add-viewed

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|query|string| 否 ||心愿的 id|

> 返回示例

> 200 Response

```json
{
  "message": "string",
  "content": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

## POST 将其添加至心愿

POST /todolist/community/add-to-wish

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|query|string| 否 ||心愿 id|
|Authorization|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "message": "string",
  "content": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» message|string|true|none||none|
|» content|string|true|none||none|

# 数据模型

<h2 id="tocS_Todo">Todo</h2>

<a id="schematodo"></a>
<a id="schema_Todo"></a>
<a id="tocStodo"></a>
<a id="tocstodo"></a>

```json
{
  "id": "string",
  "event": "string",
  "completed": true,
  "is_cycle": true,
  "description": "string",
  "importance_level": 0,
  "user_uuid": "string",
  "completed_date": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||ID 编号|
|event|string|true|none||none|
|completed|boolean|true|none||none|
|is_cycle|boolean|true|none||none|
|description|string|true|none||none|
|importance_level|integer|true|none||none|
|user_uuid|string|true|none||只保存在后端|
|completed_date|string|true|none||记录完成时间|

<h2 id="tocS_User">User</h2>

<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "email": "string",
  "uuid": "string",
  "password": "string",
  "is_admin": true,
  "todos": {},
  "wishes": {}
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|email|string|true|none||none|
|uuid|string|true|none||后端生成|
|password|string|true|none||none|
|is_admin|boolean|true|none||从数据库直接修改，不经过程序|
|todos|object|true|none||用来关联 todo 数据库，不会发给前端|
|wishes|object|true|none||用来关联 wish 数据库，不会发给前端|

<h2 id="tocS_Response">Response</h2>

<a id="schemaresponse"></a>
<a id="schema_Response"></a>
<a id="tocSresponse"></a>
<a id="tocsresponse"></a>

```json
{
  "message": "string",
  "content": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|message|string|true|none||none|
|content|string|true|none||none|

<h2 id="tocS_Wish">Wish</h2>

<a id="schemawish"></a>
<a id="schema_Wish"></a>
<a id="tocSwish"></a>
<a id="tocswish"></a>

```json
{
  "id": "string",
  "user_uuid": "string",
  "event": "string",
  "is_cycle": true,
  "description": "string",
  "is_shared": true
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||ID 编号|
|user_uuid|string|true|none||只保存在后端|
|event|string|true|none||none|
|is_cycle|boolean|true|none||none|
|description|string|true|none||none|
|is_shared|boolean|true|none||none|

<h2 id="tocS_SharedWish">SharedWish</h2>

<a id="schemasharedwish"></a>
<a id="schema_SharedWish"></a>
<a id="tocSsharedwish"></a>
<a id="tocssharedwish"></a>

```json
{
  "id": "string",
  "event": "string",
  "description": "string",
  "viewed": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|event|string|true|none||none|
|description|string|true|none||none|
|viewed|integer|true|none||none|


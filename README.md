# gin框架
用来处理http

- [helloword](./gin/helloword.go)
  - **gin.New()**  创建gin引擎路由
  - **gin.Default()**  封装了gin.New(),会默认加载中间件
  - gin.New().GET/POST/PUT/DELETE(http路径, 处理函数(`*gin.Context`))
    - *gin.Context    封装了request和response
    - *gin.Context.String()    写response 字符串
  - gin.New().Run(:port)  监听端口启动路由
  
## 路由
- [URL注册分组管理](./bin/routesGroup.go)
  - gin.New().Group()  创建路由组
  - gin.New().Group().GET/POST/PUT/DELETE()  路由组添加路由


- TODO 路由文件拆分


## request 请求
- 获取http参数
    - *gin.Context.Param()		获取url
    - *gin.Context.Query()		获取param
    - *gin.Context.PostForm()		获取表单param
    - *gin.Context.FormFile()		获取multipart文件
    - *gin.Context.MultipartForm()	获取多个multipart表单参数

- body中的json解析
  - *gin.Context.ShouldBindJSON(结构体)  将body中的json字符串解析到结构体

## response 响应
- 常规body响应
  - *gin.Context.Json(响应码, gin.H{key:value})     json字符串
  - *gin.Context.Json(响应码, struct)      结构体
  - *gin.Context.XML(响应码, gin.H{key:value})      XML
  - *gin.Context.YAML(响应码, gin.H{key:value})      YAML
  - *gin.Context.ProtoBuf(响应码, protobuf格式数据)      protobuf格式

- HTML模板
  - *gin.Context.LoadHTMLGlob("模板路径")
  - *gin.Context.HTML(响应码, "模板文件名",gin.H{key:value})

- 重定向
  - *gin.Context.Redirect(响应码, "重定向URL路径")

- 同步异步处理

## 中间件 /过滤器拦截器
- gin.Default() 会默认加载框架内置的中间件 Logger(), Recovery()\
  logger可以很方便的进行调试，recovery可以使用panic中断的恢复。
- gin.New.Use(中间件)  加载中间件

## cookie session token
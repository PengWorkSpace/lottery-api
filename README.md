### 项目结构说明
* 项目是完全按照Gopher 公认的 [项目标准结构](https://github.com/golang-standards/project-layout)，很多组件的实现大量参考B站开源的Kratos设计理念。采用go module管理，基于go 1.14 版本开发
* cmd 程序主入口 main.go
* internal 程序集包 包括dao层、model层、server层、service层
* utils 公用程序包
* databasesql 整个项目的建表sql和初始化数据

### 部署步骤
* 创建lotteryrecord数据库，执行lottery.sql,建立相关表结构
* 进入cmd文件，执行 go run main.go ,启动应用程序
* 查看对应的路由地址，访问对应的地址信息测试


### 逻辑说明
* 发送验证码和校验验证码用于注册和报名时候校验
* 抽奖时候防止并发操作，采用redis的分布式锁
* 导出操作采用 encondig/csv 库


测试说明
* 报名注册 
请求地址：http://localhost:8080/lottery/join
![注册已存在](/images/lottery1.png)


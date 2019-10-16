
### 项目介绍

### 功能介绍
> 这一切都基于社区以及分类，先完善社区及分类

- [ ] 权限
- [x] 社区、分类
- [ ] 文章模块
    - [ ] CURD
    - [ ] 文章评论
    - [ ] 热门文章
    - [ ] 自动升级精华文章、手动升级精华文章
    - [ ] 点赞、关注
- [ ] 用户模块
    - [ ] CURD
    - [ ] 活跃用户
    - [ ] 用户间的关系（关注的人、粉丝）
    - [ ] 所点赞文章、收藏文章……
- [ ] 搜索模块
    - [ ] ES？
- [ ] 通知模块
    - [ ] 私信
    - [ ] 通知
    
### 安装

### 使用
1. cp .env.example.yaml .env.yaml
2. 修改配置参数

### 项目目录结构
<details>
<summary>展开查看</summary>
<pre><code>
├── cmd                         # 启动项目 main.go 以及可执行文件
├── config                      # 配置文件
├── handler                     # 数据库相关操作文件
├── http                        # 路由以及对应路由处理
├── middleware                  # 中间件
├── model                       # 模型的定义
├── pkg                         # 项目库
├── runtime                     # 运行时所产生的文件
├── service                     # 具体逻辑处理
└── tools                       # 工具函数
</code></pre>
</details>

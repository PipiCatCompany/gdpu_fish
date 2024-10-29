# Nunu — A CLI tool for building go aplication.
文档不一定好用
```
我的使用指南：
    参考： https://github.com/go-nunu/nunu/blob/main/docs/zh/tutorial.md
    开发流程如下：
        - nunu create all [modelName]   # 创建模型，并自动生成对应handler/repo/service基础文件
        - 完成接口编写
        - 写路由 -> 在internal/server/http.go 
        - 编辑 cmd/server/wire.go，将刚刚生成文件中的工厂函数添加到providerSet中
        - nunu wire all 生成依赖，wire_gen.go
        - 迁移数据库 -> internal/server/migration.go
        - 测试接口
```


Nunu是一个基于Golang的应用脚手架，它的名字来自于英雄联盟中的游戏角色，一个骑在雪怪肩膀上的小男孩。和努努一样，该项目也是站在巨人的肩膀上，它是由Golang生态中各种非常流行的库整合而成的，它们的组合可以帮助你快速构建一个高效、可靠的应用程序。

[英文介绍](https://github.com/go-nunu/nunu/blob/main/README.md)

![Nunu](https://github.com/go-nunu/nunu/blob/main/.github/assets/banner.png)

## 文档
* [使用指南](https://github.com/go-nunu/nunu/blob/main/docs/zh/guide.md)
* [分层架构](https://github.com/go-nunu/nunu/blob/main/docs/zh/architecture.md)
* [上手教程](https://github.com/go-nunu/nunu/blob/main/docs/zh/tutorial.md)
* [高效编写单元测试](https://github.com/go-nunu/nunu/blob/main/docs/zh/unit_testing.md)

## 许可证

Nunu是根据MIT许可证发布的。有关更多信息，请参见[LICENSE](LICENSE)文件。

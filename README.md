doxygen.io - Doxygen as Service
======

[![Build Status](https://travis-ci.org/qiniu/doxygen.io.svg?branch=master)](https://travis-ci.org/qiniu/doxygen.io)
[![doxygen.io](http://doxygen.io/github.com/qiniu/doxygen.io/?status.svg)](http://doxygen.io/github.com/qiniu/doxygen.io/)

[![Qiniu Logo](http://assets.qiniu.com/qiniu-white-195x105.png)](http://qiniu.com/)

# 使用方法

当前仅仅支持 github.com 上的项目。使用方式类似 [http://godoc.org/](http://godoc.org/)，只需要访问：

```
http://doxygen.io/github.com/<UserName>/<ProjectName>
```

比如，访问 [http://doxygen.io/github.com/qiniu/php-sdk](http://doxygen.io/github.com/qiniu/php-sdk) 即可访问七牛云存储的 PHP SDK 文档。

# 后续事项

- 支持一些定制性功能（要求项目根目录下有 .doxygen.io 这样的配置文件或文件夹），比如只根据项目的特定目录生成文档。
- 支持 $repo/.doxygen.io/README.dox, $repo/.doxygen.io/README.md 来作为 MainPage。
- 美化 Project admin tools 页面。
- 支持更多的 Source hosting 服务，比如 bitbucket.org 之类。
- 考虑支持多个分支维护独立的文档。`http://doxygen.io/github.com/<UserName>/<ProjectName>/` 表示 master 分支的 MainPage。其他分支用 `http://doxygen.io/github.com/<UserName>/<ProjectName>/<BranchName>/` 表示 MainPage。
- doxygen.io 网站的 favicon 支持。

# 参与改进

- 要提意见建议，请到 [Issues](https://github.com/qiniu/doxygen.io/issues)。
- 希望贡献源代码，请提 [Pull Request](https://github.com/qiniu/doxygen.io/compare)。

你可以到这里查看所有的[改进记录](https://github.com/qiniu/doxygen.io/releases)。

# 贡献者名单

- 你可以到这里查看[贡献者列表](https://github.com/qiniu/doxygen.io/graphs/contributors)。


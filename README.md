## About Project

最近在学Go和Gin，以此项目练手熟悉Gin相关API。

## Environment Build

PHPStudy + MySQL 8.0.12

## Command

```shell
# 初始化项目
mkdir -p ~/blog-service
cd blog-service
go mod init github.com/testercc/blog-service

# 安装gin
# 1.不指定版本，默认最新版
go get -u github.com/gin-gonic/gin

# 2.指定版本
go get -u github.com/gin-gonic/gin@v1.5.0
go get -u gopkg.in/natefinch/lumberjack.v2
```

## Dev Design

1. model
2. router
3. [todo](https://golang2.eddycjy.com/posts/ch2/02-project-design/)

## Blog Structure

- configs：配置文件。
- docs：文档集合。
- global：全局变量。
- internal：内部模块。
  - dao：数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL、ElasticSearch 等。
  - middleware：HTTP 中间件。
  - model：模型层，用于存放 model 对象。
  - routers：路由相关逻辑处理。
  - service：项目核心业务逻辑。
- pkg：项目相关的模块包。
- storage：项目生成的临时文件。
- scripts：各类构建，安装，分析等操作的脚本。
- third_party：第三方的资源工具，例如 Swagger UI。

## Test URL

测试url的命令
```shell
curl -v http://127.0.0.1:8000/api/v1/articles/1
```

## Database Initial

### Create database

```mysql
CREATE DATABASE
IF
	NOT EXISTS blog_service DEFAULT CHARACTER 
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
```

### Create blog_tag table
use blog_service;

```mysql
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';
```

### Create blog_article table

```mysql
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '文章简述',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';
```

### Create blog_article_tag table

```mysql
CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章 ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签 ID',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';
```

## REF

- [Gin官方文档](https://github.com/gin-gonic/gin)
- 《Go语言编程之旅：一起用Go做项目》  `Current Learn: P63`
- [Code](https://github.com/go-programming-tour-book)
- [online1](https://e.dangdang.com/pc/reader/index.html?id=1901224143)
- [online2](https://golang2.eddycjy.com/posts/ch2/02-project-design/)


## mongo驱动

用官方的mongo-driver
```
go get go.mongodb.org/mongo-driver/mongo
```

如果你使用的是不支持modules的go版本，你可以使用dep来安装：

```
dep ensure -add "go.mongodb.org/mongo-driver/mongo"
```

## Install 3rd-party Log Lib

```
go get -u github.com/spf13/viper
```

```
go get -u gopkg.in/natefinch/lumberjack.v2
```

## Install Swagger
安装 Go 对应的开源 Swagger 相关联的库，在项目 blog-service 根目录下执行安装命令，如下：
```
$ go get -u github.com/swaggo/swag/cmd/swag
$ go get -u github.com/swaggo/gin-swagger 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template
```
安装注意事项：

1.Windows 11
由于Windows下没有bin目录，所以需要先找到swag的下载位置。
在该目录下执行go install 会生成swag.exe到gopath的主目录下。
这样swag.exe就可以用来执行swag init操作了。
```
go env 查看pkg安装目录，找到swag的下载位置，一般为$GOPATH的路径下
C:~\go\pkg\mod\github.com\swaggo\swag@v1.8.3\cmd\swag

# 会生成swag.exe到gopath的主目录下
> go install
# 检验是否安装成功     
> swag -v
swag.exe version v1.8.3
```

在该目录下执行go install 会生成swag.exe到gopath的主目录下。
这个swag.exe就可以用来执行swag init操作了。

2.MacOS

MacOS zsh 没有swag命令的处理方式：
如果提示zsh: command not found: swag，先用`go env`看一下gopath的目录。
找到目录大概这样GOPATH="/Users/xx/go/bin"，ls一下，看到有swag这个文件，试一下`/Users/xx/go/bin/swag -v`命令。
执行后有类似"swag version v1.8.4"的回显。

则可确认是本地swag没有加到环境变量，我的终端是zsh，编辑一下配置：
vim ~/.zshrc
export PATH="/Users/xx/go/bin:$PATH"
保存退出后，重新打开Terminal即可正常使用swag命令。

## Generate Swagger Docs

在完成了所有的注解编写后，回到项目根目录下，执行如下命令：
```shell
$ swag init
```
在执行命令完毕后，会发现在 docs 文件夹生成 docs.go、swagger.json、swagger.yaml 三个文件。

如果需要重新生成文档，则在项目根目录下再次执行 swag init。

Swagger文档地址：`http://127.0.0.1:8000/swagger/index.html`

实质上在初始化 docs 包时，会默认执行 init 方法，而在 init 方法中，会注册相关方法，主体逻辑是 swag 会在生成时去检索项目下的注解信息，
然后将项目信息和接口路由信息按规范生成到包全局变量 doc 中去。 紧接着会在 ReadDoc() 方法中做一些 template 的模板映射等工作，完善 doc 的输出。


## Validator

使用开源项目 go-playground/validator 作为我们的本项目的基础库，它是一个基于标签来对结构体和字段进行值验证的一个验证器。
gin内部默认使用，如果要安装则执行如下命令：
`go get -u github.com/go-playground/validator/v10`

REF:
-[如何在windows下使用swaggo](https://blog.csdn.net/ran_Max/article/details/105718374)
# 💻ICP-Finder

ICP-Finder 是一个基于 Web 的服务，使用 Golang 编写。它的主要功能是查询指定的域名是否在中国大陆地区进行了网站备案，并提供相应的备案信息。该功能的实现是通过调用 https://guanjia.qq.com 的接口并加以封装实现的。

# 🌟特性

- 可以查询大部分常见的域名后缀，如 .com, .cn, .org 等等。
- 提供实时的备案信息查询，包括备案号、备案主体名称、备案主体类型、备案单位名称以及备案时间。
- 使用 Golang 编写，具有高效性和可扩展性，可以适应各种复杂的网络环境和使用场景。

# 安装🔧

## 从 release 页面下载使用

1. 打开项目的 release 页面。
2. 找到最新版本对应您服务器平台和架构的可执行文件，例如 icp-finder_0.1.0_linux_amd64.tar.gz。
3. 下载此文件到本地。
4. 在服务器上解压缩该文件。

```
tar -xzvf icp-finder_0.1.0_linux_amd64.tar.gz
```

5. 进入解压后的文件夹。

```
cd icp-finder_0.1.0_linux_amd64
```

6. 启动服务。

```
./icp-finder
```

*请注意，以上步骤仅适用于 Linux 平台和 AMD64 架构。如果您的服务器平台和架构不同，请下载相应的可执行文件并按照上述步骤进行操作。*

## 从源码安装

1. 克隆本项目到本地计算机。

```bash
git clone https://github.com/mengdiao/ICP-Finder.git
```

2. 进入项目目录，并运行 main.go 文件以启动服务器。

```bash
cd icp-finder
go run main.go
```

# 使用方法🚀

打开浏览器，输入 http://your-ip:8080/ 或者 http://localhost:8080/ 进入主页面。
在搜索框中输入待查询的域名，例如 `example.com`。
点击“查询”按钮，系统将会显示出该域名的备案信息。

此服务的接口是 `/geticpinfo`，您可以向该接口发送 GET 请求，使用 `url` 参数指定待查询的域名。例如：

```bash
GET /geticpinfo?url=example.com HTTP/1.1
Host: localhost:8080
```

系统将会返回该域名的备案信息。您还可以使用可选的 `feature` 参数控制返回结果的格式。如果 `feature` 参数为 `enabled`（默认值），系统将返回经过封装后的 JSON 格式数据；如果 `feature` 参数为 `disabled`，系统将返回腾讯接口源数据的 JSON 格式数据。

```bash
GET /geticpinfo?url=example.com&feature=disabled HTTP/1.1
Host: localhost:8080
```

例如，当您访问 `http://localhost:8080/geticpinfo?url=qq.com&feature=disabled` 时，系统将会返回如下 JSON 数据：

```
json{
  "data": {
    "results": {
      "ICPSerial": "粤B2-20090059-5",
      "Orgnization": "深圳市腾讯计算机系统有限公司",
      "Wording": "",
      "WordingTitle": "",
      "certify": 0,
      "detect_time": "1514189828",
      "eviltype": "0",
      "isDomainICPOk": 1,
      "url": "qq.com",
      "whitetype": 3
    },
    "retcode": 0
  },
  "reCode": 0
}
```

## 注意事项⚠️

1. 本项目仅限于查询中国大陆地区的网站备案信息，对于海外的网站或已经注销备案的网站，可能无法提供正确的结果。
2. 本项目所涉及的备案信息来源于 `https://guanjia.qq.com` 的接口，如果该接口出现故障或数据不准确，本项目将无法提供正确的结果。
3. 本项目仅用于学习和研究目的，任何人不得将其用于非法用途。

## 贡献

如果您对本项目感兴趣并希望为其贡献代码，请提交 pull request，并等待审核和合并。
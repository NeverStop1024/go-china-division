##

统计用区划代码

### ❗️ 前提

1. 务必遵守中国法律法规,在法律允许内使用本软件。不对社会、机构、个人、群体等等造成困扰。
2. 本软件只是学习参考使用，不作为任何商业软件的工具及插件。
3. 下载、使用本软件已代表同意上述观点，任何法律风险由使用者承担。
4. 本数据来源 [国家统计局](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/index.html)。


## 📦 安装

```bash
# 自动编译
go get https://github.com/lifegit/go-china-division

echo "go-china-division gain"

# 源码安装
git clone https://github.com/lifegit/go-china-division
go mod download

go install
echo "添加到 PATH环境变量"

echo "go-china-division gain"
```

### ⚠️ 命令行参数详解

```bash
[root@ericzhou felix]# go build && ./go-china-division gain --help
gain get zoning code and urban-rural division

Usage:
  china gain [flags]

Examples:
gain -o 2 -f c.json

Flags:
  -f, --fileName string   generate file filename (default "china.json")
  -h, --help              help for division
  -o, --option int        option plan (default 2)
  -p, --outPath string    generate file path (default "./")
```

### 👀 注意
1. 本软件依赖 [chrome](https://www.google.cn/chrome/) 浏览器，请先安装。
2. 使用时，可能会遇到人机验证。会弹出浏览器，要求输入验证码情况，正确输入即可。

#### ⭕️  QA

1. 
   Q：报错 "Failed to get the latest year" 或 "Failed to get the node"。

   A：获取数据太频繁，等几分钟后再试。
# 注意
## 版本
X.0 http库实现
x.1 gin框架实现
从5.X版本开始全面使用gin框架实现

## 使用
### apiServer
`apiServer`目录下`conf`目录是MasterNode的配置文件，格式是json(暂时只支持json，以后会考虑增加其他格式),启动脚本是`start.sh`

### dataServer
`dataServer`目录下`conf`目录是dataNode的配置文件，格式是json(暂时只支持json，以后会考虑增加其他格式),启动脚本在`test`目录

## 辅助工具
`test`目录中有一个golang写的测量文件大小的小程序


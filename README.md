# 静态文件更新工具

## 功能

1. 文件上传、zip压缩包上传（自动解压）
2. 项目管理
3. 服务器管理（支持密码、密钥两种认证方式）
4. 用户管理
5. 权限管理
6. 历史上传记录查询
7. 历史记录回滚

## 快速使用

1. 在Release中下载对应的压缩包
2. 解压后给予`tool`运行权限: `chmod 755 ./tool`
3. 使用`./tool start`启动服务，`./tool stop`停止服务
4. 启动后访问`http://localhost:8080`即可使用，修改端口，请到`resource/config/config.yml`中修改

## 注意事项

1. 首次启动会自动添加管理员信息，账号：`admin@qq.com` 密码：`123456`
2. 如果管理员用户丢失，可以找到sqlite(默认位置：`resource/db/update_tool.db`，可在`resource/config/config.yml`文件中设置db位置)文件，将用户`is_admin`设置为`1`，或者删除users表，再次启动之后会自动添加默认管理员，账号密码同上
3. 如果权限丢失，可以找到sqlite文件，删除`permissions`表，再次启动之后会自动添加默认权限

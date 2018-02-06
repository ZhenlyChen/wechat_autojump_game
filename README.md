# 全自动玩微信跳一跳
Golang实现的自动玩微信跳一跳

原来[sundy-li](https://github.com/sundy-li/wechat_autojump_game)的代码对于下一个跳跃点的判断在我本地测试里面很多时候不正确，于是用另外一个算法修改了一下，顺便修改了一下细节.

#### 准备条件
- 需要安装adb驱动, 这里有一篇国外作者的[教程](https://www.xda-developers.com/install-adb-windows-macos-linux/)
- 手机连接电脑后,进入设置-开发者选项-打开usb调试

准备就绪后,在终端输入`adb devices`, 如果可以看到对应设备,表示adb驱动已经安装配置完毕


#### 如何安装使用
- 安装方法一:

 一键下载安装,无需搭建环境,傻瓜化使用方法
请在[release](https://github.com/zhenlychen/wechat_autojump_game/releases) 页面下载对应操作系统的二进制压缩包,解压后,执行jump文件即可

```
 $  ./jump 
```

- 安装方法二: 

手动安装,开发调试
```
 $ go get -u github.com/zhenlychen/wechat_autojump_game
 $ cd $GOPATH/src/github.com/zhenlychen/wechat_autojump_game/cmd
 $ go run main.go
```


#### 参数
- `./jump -s 1.392`, 参数`s`表示距离速度值,不同分辨率可以适当调节
- `./jump -m 2000`, 参数 `m` 表示休眠时间毫秒值

#### 原理
- 利用adb shell截图游戏屏幕
- 读取截屏图片,获取当前位置,下一跳位置,计算跳动距离和触屏时间
- 利用adb shell发送input swipe事件来跳跃






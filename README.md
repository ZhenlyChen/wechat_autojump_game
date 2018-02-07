# 全自动玩微信跳一跳
Golang实现的自动玩微信跳一跳

原来[sundy-li](https://github.com/sundy-li/wechat_autojump_game)的代码对于下一个跳跃点的判断在我本地测试里面很多时候不正确，于是用另外一个算法修改了一下，顺便修改了一下细节.
另外实现了web操作界面

#### 准备条件
- 将adb套件放入当前目录lib文件夹下
- 手机连接电脑后,进入设置-开发者选项-打开usb调试

#### 原理
- 利用adb shell截图游戏屏幕
- 读取截屏图片,获取当前位置,下一跳位置,计算跳动距离和触屏时间
- 利用adb shell发送input swipe事件来跳跃






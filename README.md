# lcould

1. Linux

    推荐 inotify
   inotifywait 检测文件变化，rsync 同步
2. macos
   fswatch 检测文件变化，rsync 同步




win

它们分别是：FindFirstChangeNotification和ReadDirectoryChangesW。
其中FindFirstChangeNotification函数只能监控到某一目录下有文件发生改变，而不能监控到具体是哪一文件发生改变。
ReadDirectoryChangesW 能监控目录下某一文件发生改变。

　利用ReadDirectoryChangesW函数实现对一个目录进行监控的。
 具体的做法是：首先使用CreateFile获取要监控目录的句柄；然后在一个判断循环里面调用ReadDirectoryChangesW，
 并且把自己分配的用来存放目录变化通知的内存首地址、内存长度、目录句柄传给该函数。用户代码在该函数的调用中进行同步等待。
 当目录中有文件发生改变，控制函数把目录变化通知存放在指定的内存区域内，并把发生改变的文件名、文件所在目录和改变通知处理
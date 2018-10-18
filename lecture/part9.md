#并发
concurrency

涉及到的关键字 go  channel  

channel 是goroutine沟通的桥梁  大都是阻塞的
通过make创建，close关闭
channel是引用类型
可以使用for range来迭代不断操作channel
可以设置单向或双向通道
可以设置缓存大小 在未内填满前不会发生阻塞

学习笔记: 对于项目的工程化的结构中涉及概念还是比较模糊, 对各个层还没有完全的理解,所以这次的作业的结构在完成后,感觉还是十分十分的混乱.需要看reference逐渐的理解,后在重新设计一波 --- taco tuesday


api: 目录里面放了路由函数
cmd: main.go
internal:
	data: 数据库操作
	biz: 业务逻辑处理
	di: 使用wire构建的依赖文件
	service: 

1. go中没规定长度的数据类型跟着操作系统的位数来
例如64位系统int默认64位
2. rune为32位4字节 byte8位1字节
3. complex64，complex128是复数，前者实部虚部分别为32位均为float32，后者实部虚部分别是64位均为float64.
4. go类型转换是强制的没有隐式类型转化
5. go中指针不能运算
6. go中只有值传递一种方式，所以自定义数据类型定义时就需要考虑时当作指针还是具体值
7. 数组也是值类型
8. go一般不使用数组使用切片Slice
9. Slice内部实现：三个指针ptr len cap。ptr是Slice开头元素，len是Slice长度，cap是从ptr开始一直到数组结束。Slice都可以拓展可以超过len不能超过cap
10. map中处理slice,map,function的内建类型都可以作为key，struct类型不包括上述类型也可以做key
11. golang不支持继承多态，只有封装。面向接口编程
12. nil指针也可以调用方法
13. 方法中改变内容必须用指针接收者，结构过大也需要指针接收者。
14. 位结构定义的方法必须放在同一个包中，可以是不同的文件
15. 如何扩充系统类型或者别人的类型，1定义别名最简单2使用组合最常用3使用内嵌Embedding节省代码
16. go mod,很好的管理依赖是目前最好的依赖管理方式。
17. go test -cover 命令行测试代码覆盖率 go test -bench测试代码性能 go test 测试结果
18. 使用 go tool pprof cpu.out 工具查看cpu.out文件配合web工具发现代码慢的地方
19.  go test -bench . -cpuprofile cpu.out 性能测试输出到文件cpu.out
20. go tool pprof cpu.out  pprof工具查看文件 web命令用浏览器输出svg文件
21. go test -cover -coverprofile= c.out 测试结果输出到c.out文件
22. go test / go test -bench .  / go test -cover 命令进行测试/测试性能/代码覆盖率测试



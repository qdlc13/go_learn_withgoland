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
12. 
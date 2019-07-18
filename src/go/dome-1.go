package main

// 1. 左大括号 { 不能单独放一行
// 在其他大多数语言中，{ 的位置你自行决定。Go 比较特别，遵守分号注入规则（automatic semicolon injection）：编译器会在每行代码尾部特定分隔符后加 ; 来分隔多条语句，比如会在 ) 后加分号：


/*
// 错误示例
func main()
{
println("hello world")
}

// 等效于
func main();    // 无函数体
{
println("hello world")
}

./main.go: missing function body
./main.go: syntax error: unexpected semicolon or newline before {
*/

// 正确示例
func main() {
	println("hello world")
}     
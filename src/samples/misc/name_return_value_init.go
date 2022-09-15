// bufio.bufio中的 Reader.collectFragments方法的返回值fullBuffers未初始化
// 确认命名返回值已初始化
package main

import "fmt"

func test(msg ...string) (buf []byte) {
	for _, s := range msg {
		buf = append(buf, []byte(s)...)
	}
	return buf
}

func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	buf := test(proverbs...)
	fmt.Println(string(buf))
}

package frp

import (
	"bufio"
	"os"
)

type FrpLogListener interface {
	Log(log string)
}

func SetFrpLogListener(l FrpLogListener) {
	reader, writer, _ := os.Pipe()

	// 将多路复用器设置为标准输出
	os.Stdout = writer

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			log := scanner.Text()
			l.Log(log)
		}
	}()
}

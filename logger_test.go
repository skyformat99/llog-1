package llog

import (
	"fmt"
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/handler"
	"github.com/syyongx/llog/types"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	logger := NewLogger("test")
	h, err := handler.NewFile("/dev/null", types.WARNING, true, 0644)
	buf := handler.NewBuffer(h, 1, types.WARNING, true, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	f := formatter.NewLine("%Datetime% [%LevelName%] [%Channel%] %Message%\n", time.RFC3339)
	h.SetFormatter(f)
	logger.PushHandler(buf)
	logger.Warning("xxx")
	buf.Close()
}

func BenchmarkBasic(b *testing.B) {
	logger := NewLogger("test")
	h, err := handler.NewFile("/dev/null", types.WARNING, true, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	f := formatter.NewLine("%Datetime% [%LevelName%] [%Channel%] %Message%\n", time.RFC3339)
	h.SetFormatter(f)
	//buf := handler.NewBuffer(h, 1, types.WARNING, true, true)
	//logger.PushHandler(buf)
	logger.PushHandler(h)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Warning("xxx")
			//buf.Close()
		}
	})
}

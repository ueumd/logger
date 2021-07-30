package logger

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	err := Init("./test.log", true, 5)

	if err != nil {
		t.Fatal(err)
	}

	for i := 0;i < 1000; i++ {

		Infoln("this is test")

		InfoF("this is test %s", "xxx")

		time.Sleep(time.Second * 10)
	}

}

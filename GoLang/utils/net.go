package utils

import (
	"github.com/spf13/cast"
	"net"
	"sync"
	"time"
)

func NetTcping(host any, opts ...map[string]any) (ok bool, detail []map[string]any) {

	if len(opts) == 0 {
		opts = append(opts, map[string]any{
			"count":   4,
			"timeout": 5,
		})
	}

	opt := opts[0]
	count := cast.ToInt(opt["count"])
	timeout := cast.ToDuration(opt["timeout"]) * time.Second

	list := make([]map[string]any, 0)

	wg := sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {

		go func() {

			defer wg.Done()

			start := time.Now()
			_, err := net.DialTimeout("tcp", cast.ToString(host), timeout)

			if err != nil {
				list = append(list, map[string]any{
					"host":    host,
					"status":  false,
					"waist":   time.Now().Sub(start).Milliseconds(),
					"message": "Site unreachable, error: " + err.Error(),
				})
				return
			}

			list = append(list, map[string]any{
				"host":    host,
				"status":  true,
				"waist":   time.Now().Sub(start).Milliseconds(),
				"message": "tcp server is ok",
			})
		}()
	}

	wg.Wait()

	// 只要有一个 ping 成功就返回 true
	for _, val := range list {
		if val["status"].(bool) {
			return true, list
		}
	}

	return false, list
}

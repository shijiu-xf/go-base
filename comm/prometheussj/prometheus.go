package prometheussj

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)

func Monitoring(port int) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("[Monitoring] 启动失败 ", err)
		}
		log.Println("[Monitoring] 监控启动成功，端口：" + strconv.Itoa(port))
	}()
}

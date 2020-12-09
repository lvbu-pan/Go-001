package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		time.Sleep(time.Second * 60)
		fmt.Fprintln(w, "success")
	})

	s := http.Server{
		Addr:    "127.0.0.1:9091",
		Handler: mux,
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	done := make(chan struct{})
	quitDone := make(chan struct{})
	g, _ := errgroup.WithContext(context.Background())

	//启动监听端口
	g.Go(func() error {
		return s.ListenAndServe()
	})

	//停止http服务
	g.Go(func() error {
		select {
		case <-done:
			return s.Shutdown(context.Background())
		case <-quitDone:
			return s.Close()
		}
	})

	//监听信号量
	g.Go(func() error {
		for sig := range sigs {
			switch sig {
			case syscall.SIGINT, syscall.SIGQUIT:
				//不等待当前连接处理完成，立即关闭连接
				fmt.Println("立即关闭连接")
				close(quitDone)
				return nil
			case syscall.SIGTERM:
				//不在接收新的连接，等待已建立的连接处理完成后，关闭
				fmt.Println("优雅关闭连接")
				close(done)
				return nil
			case syscall.SIGHUP:
				// 重载配置
				fmt.Println("Reload")
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		close(sigs)
	}
}

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"week04/internal/di"
)

func main() {
	app, clean, err := di.InitApp()
	fmt.Println(err)
	if err != nil {
		clean()
		return
	}
	signs := make(chan os.Signal)
	signal.Notify(signs, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	done := make(chan struct{})
	quickDone := make(chan struct{})
	g, _ := errgroup.WithContext(context.Background())

	//启动监听端口
	g.Go(func() error {
		return app.Hs.ListenAndServe()
	})

	//停止http服务
	g.Go(func() error {
		select {
		case <-done:
			return app.Hs.Shutdown(context.Background())
		case <-quickDone:
			return app.Hs.Close()
		}
	})

	g.Go(func() error {
		for sig := range signs {
			switch sig {
			case syscall.SIGINT, syscall.SIGQUIT:
				//不等待当前连接处理完成，立即关闭连接
				fmt.Println("立即关闭连接")
				close(quickDone)
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
		close(done)
		os.Exit(0)
	}

}

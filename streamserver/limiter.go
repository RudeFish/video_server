package main

import "log"

// 流控机制

type ConnLimiter struct {
	// 连接数
	concurrentConn int
	// 用一个chan存储连接数
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		// 有buf的chan
		bucket: make(chan int, cc),
	}
}

// 获取一个token
func (cl ConnLimiter)GetConn()bool  {
	// 当bucker满的情况下
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}

	// 没有满，添加连接
	cl.bucket <- 1
	return true
}

// 释放token
func (cl ConnLimiter)ReleaseConn()  {
	c  :=<- cl.bucket
	log.Printf("New connction coming : %d", c)
}
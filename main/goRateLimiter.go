package main

//
//import (
//	"time"
//	"github.com/juju/ratelimit"
//)

///*
//*  @author liqiqiorz
//*  @data 2021/4/26 11:37
// */
//
////令牌桶
////令牌桶其实和漏桶的原理类似，令牌桶按固定的速率往桶里放入令牌，
////并且只要能从桶里取出令牌就能通过，令牌桶支持突发流量的快速处理。
//// 创建指定填充速率和容量大小的令牌桶
//func NewBucket(fillInterval time.Duration, capacity int64) *Bucket
//// 创建指定填充速率、容量大小和每次填充的令牌数的令牌桶
//func NewBucketWithQuantum(fillInterval time.Duration, capacity, quantum int64) *Bucket
//// 创建填充速度为指定速率和容量大小的令牌桶
//// NewBucketWithRate(0.1, 200) 表示每秒填充20个令牌
//func NewBucketWithRate(rate float64, capacity int64) *Bucket
//
//// 取token（非阻塞）
//func (tb *Bucket) Take(count int64) time.Duration
//func (tb *Bucket) TakeAvailable(count int64) int64
//
//// 最多等maxWait时间取token
//func (tb *Bucket) TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool)
//
//// 取token（阻塞）
//func (tb *Bucket) Wait(count int64)
//func (tb *Bucket) WaitMaxDuration(count int64, maxWait time.Duration) bool

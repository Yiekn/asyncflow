package util

import (
	"fmt"
	"sync"
	"time"
)

var SnowflakeGenerator *snowflake

func init() {
	var err error
	SnowflakeGenerator, err = newSnowflake(1, 1)
	if err != nil {
		Logger.Fatal("snowflake init failed, err: %s", err.Error())
	}
}

// 定义一些常量
const (
	// 工作节点ID的位数
	workerIdBits = 5
	// 数据中心ID的位数
	datacenterIdBits = 5
	// 序列号的位数
	sequenceBits = 12

	// 工作节点ID的左移位数
	workerIdShift = sequenceBits
	// 数据中心ID的左移位数
	datacenterIdShift = sequenceBits + workerIdBits
	// 时间戳的左移位数
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits
	// 序列号的掩码
	sequenceMask = -1 ^ (-1 << sequenceBits)

	// 2020-01-01 00:00:00 UTC 作为时间戳的起始点
	twepoch = 1577836800000
)

// snowflake 结构体
type snowflake struct {
	mu            sync.Mutex // 互斥锁，用于保护并发访问
	lastTimestamp int64      // 上次生成ID的时间戳
	workerId      int64      // 工作节点ID
	datacenterId  int64      // 数据中心ID
	sequence      int64      // 序列号
}

// NewSnowflake 创建一个新的 snowflake 实例
// 参数 workerId 和 datacenterId 分别表示工作节点ID和数据中心ID
// 返回值为 snowflake 实例和错误信息
func newSnowflake(workerId, datacenterId int64) (*snowflake, error) {
	// 检查 workerId 和 datacenterId 是否在有效范围内
	maxWorkerID := int64(-1) ^ (int64(-1) << workerIdBits)
	maxDatacenterID := int64(-1) ^ (int64(-1) << datacenterIdBits)
	if workerId < 0 || workerId > maxWorkerID {
		return nil, fmt.Errorf("worker ID must be between 0 and %d", maxWorkerID)
	}
	if datacenterId < 0 || datacenterId > maxDatacenterID {
		return nil, fmt.Errorf("datacenter ID must be between 0 and %d", maxDatacenterID)
	}

	return &snowflake{
		lastTimestamp: -1,
		workerId:      workerId,
		datacenterId:  datacenterId,
		sequence:      0,
	}, nil
}

// NextId 生成下一个唯一 ID
// 返回值为生成的唯一 ID 和错误信息
func (s *snowflake) NextId() (int64, error) {
	// 加锁，确保并发安全
	s.mu.Lock()
	// 函数结束时解锁
	defer s.mu.Unlock()

	// 获取当前时间戳（毫秒）
	timestamp := time.Now().UnixNano()/1e6 - twepoch

	// 处理时钟回拨问题
	if timestamp < s.lastTimestamp {
		return 0, fmt.Errorf("clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTimestamp-timestamp)
	}

	// 如果时间戳与上次相同，增加序列号
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		// 如果序列号溢出，等待下一毫秒
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano()/1e6 - twepoch
			}
		}
	} else {
		// 时间戳改变，重置序列号
		s.sequence = 0
	}

	// 更新上次时间戳
	s.lastTimestamp = timestamp

	// 生成 ID
	id := (timestamp << timestampLeftShift) |
		(s.datacenterId << datacenterIdShift) |
		(s.workerId << workerIdShift) |
		s.sequence

	return id, nil
}

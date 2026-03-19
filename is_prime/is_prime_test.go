package is_prime

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestIsPrime(t *testing.T) {
	// 第一个循环：测试真正的素数
	for i := 0; i < 1000; i++ {
		p, err := rand.Prime(rand.Reader, 20)
		if err != nil {
			t.Fatalf("生成素数失败: %v", err)
		}

		if !IsPrime(p) { // 期待返回 true
			t.Errorf("第 %d 次测试失败: IsPrime(%v) 期待 true，却返回了 false", i, p)
		}
	}

	// 第二个循环：测试合数（两个素数相乘）
	for i := 0; i < 1000; i++ {
		p1, _ := rand.Prime(rand.Reader, 10)
		p2, _ := rand.Prime(rand.Reader, 10)

		// 必须把两个大整数相乘，构造出一个合数
		comp := new(big.Int).Mul(p1, p2)

		if IsPrime(comp) { // 期待返回 false
			t.Errorf("第 %d 次测试失败: IsPrime(%v) 这是一个合数，却返回了 true", i, comp)
		}
	}
}

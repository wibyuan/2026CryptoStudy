package secret_share

import (
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

// Share 表示 Shamir 秘密共享的一个分片
type Share struct {
	X fr.Element // 在多项式中的采样点 x
	Y fr.Element // 多项式在 x 处的值 f(x)
} //114514

// SecretSplit 将秘密 secret 拆分为 n 个分片，并设置恢复门槛为 t。
//
// 数学原理：
// 1. 构造一个 t-1 次的多项式 f(x) = a0 + a1*x + ... + a_{t-1}*x^{t-1}。
// 2. 令常数项 a0 = secret。
// 3. 随机生成其余的系数 a1, ..., a_{t-1}。
// 4. 计算并返回 n 个不同的点 (x, f(x)) 作为分片，通常取 x = 1, 2, ..., n。
//
// 参数:
//
//	secret - 需要加密的原始秘密（有限域元素）
//	n - 生成分片的总数
//	t - 恢复秘密所需的最小分片数（门槛值）
//
// 返回值:
//
//	包含 n 个 Share 对象的切片
func SecretSplit(secret fr.Element, n int, t int) []*Share {
	// TODO: finish SecretSplit func
	//panic("No implement error")
	coeffs := make([]fr.Element, t)
	coeffs[0] = secret
	for i := 1; i < t; i++ {
		coeffs[i].SetRandom()
	}
	shares := make([]*Share, n)
	for i := 0; i < n; i++ {
		x := fr.NewElement(uint64(i + 1))
		y := fr.NewElement(0)
		for j := t - 1; j >= 0; j-- {
			y.Mul(&y, &x).Add(&y, &coeffs[j])
		}
		shares[i] = &Share{x, y}
	}
	return shares
}

// SecretCombine 使用拉格朗日插值法从给定的分片中恢复原始秘密 f(0)。
//
// 数学原理：
// 给定 k 个点 (x0, y0), ..., (xk-1, yk-1)，其中 k >= t，
// 原始秘密即为插值多项式在 x=0 处的值：
//
// 参数:
//
//	shares - 收集到的分片切片
//
// 返回值:
//
//	恢复出的原始秘密（有限域元素）
func SecretCombine(shares []*Share) fr.Element {
	// TODO: finish SecretCombine func
	//panic("No implement error")
	secret := fr.NewElement(0)
	t := len(shares)
	for i := 0; i < t; i++ {
		num := shares[i].Y
		den := fr.NewElement(1)
		for j := 0; j < t; j++ {
			if i != j {
				num.Mul(&num, &shares[j].X)
				sub := shares[j].X
				sub.Sub(&sub, &shares[i].X)
				den.Mul(&den, &sub)
			}
		}
		num.Div(&num, &den)
		secret.Add(&secret, &num)
	}
	return secret
}

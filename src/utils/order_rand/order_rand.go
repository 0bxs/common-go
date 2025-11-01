package order_rand

type OrderRand struct {
	min    int64 // 区间最小值
	length int64 // 区间长度
	prime  int64 // 与 length 互质的乘数，用来打乱顺序
	offest int64 // 偏移量，加在乘法结果上再取模，也是打乱作用
	pos    int64 // 当前进度（已经走到第几个数）
}

func New(min, max int64, seed int64) *OrderRand {
	n := max - min + 1
	a := (seed*2 + 1) % n
	if a == 0 {
		a = 1
	}
	for g := gcd(a, n); g != 1; g = gcd(a, n) {
		a += 2
		if a >= n {
			a %= n
			if a == 0 {
				a = 1
			}
		}
	}
	b := ((seed >> 17) ^ seed) % n
	if b < 0 {
		b += n
	}
	return &OrderRand{min: min, length: n, prime: a, offest: b, pos: 0}
}

func (r *OrderRand) Next() (int64, bool) {
	if r.pos >= r.length {
		return 0, false
	}
	val := r.min + (r.prime*r.pos+r.offest)%r.length
	r.pos++
	return val, true
}

func (r *OrderRand) GetPos() int64 { return r.pos }

func (r *OrderRand) LoadPos(pos int64) {
	if pos >= 0 && pos <= r.length {
		r.pos = pos
	}
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

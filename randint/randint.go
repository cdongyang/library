/*
为使测试结果不受随机数影响,这个包通过相同的Rand结构体来使产生的随机数相同
通过控制Mod和Add来使控制随机数是否重复,多少个数后出现重复
如果Add和Mod互质,则Rand会在产生Mod个数之后出现重复,Mod为最小循环长度
最小循环长度为 Mod/gcd(Add,Mod)
*/
package randint

type Rand struct {
	Add   int
	First int
	Mod   int
}

func (r *Rand) Int() int {
	r.First = (r.First + r.Add) % r.Mod
	return r.First
}

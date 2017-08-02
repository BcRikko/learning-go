package roman

type Arabic uint16

type Roman string

var nums = []struct {
	a Arabic
	r Roman
}{
	{1, "I"},
	{2, "II"},
}

// ToRoman は数値をローマ数字に変換する
func (in Arabic) ToRoman() (out Roman) {
	remainA := in
	for _, num := range nums {
		for num.a <= remainA {
			remainA -= num.a
			out += num.r
		}
	}
	return
}

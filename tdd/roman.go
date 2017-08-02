package roman

type Arabic uint16

type Roman string

// ToRoman は数値をローマ数字に変換する
func (in Arabic) ToRoman() Roman {
	return Roman("I")
}

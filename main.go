package main

import "fmt"

type Attribute int

const (
	Clear Attribute = iota
	Bold
	Light
	Italic
	Underline
	Blink
	FastBlink
	Reverse
	Hide
	Undo
)

const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

const FgClear Attribute = 39
const BgClear Attribute = 49
const esc = "\x1b"
const foreground = "38;5;"
const background = "48;5;"

func (a Attribute) String() string {
	attribute := 0
	ret := ""
	if a < -255 {
		// Background
		attribute = (int(a) + 255) * -1
		ret = fmt.Sprintf("%v%v", background, attribute)
	} else if a < 0 {
		// Forground
		attribute = int(a) * -1
		ret = fmt.Sprintf("%v%v", foreground, attribute)
	} else if a < 256 {
		ret = fmt.Sprintf("%v", attribute)
	}

	return ret
}

func GetMultiColorAttribute(colorCode int, isBackground bool) Attribute {
	if colorCode < 0 || 255 < colorCode {
		return 0
	} else {
		val := colorCode * -1
		if isBackground {
			val += -255
		}
		return Attribute(val)
	}
}

func ApplyAttribute(str string, attributes ...Attribute) string {
	ret := str
	codes := attributes[0].String()
	for i := range attributes[1:] {
		codes = fmt.Sprintf("%v;%v", codes, attributes[i+1].String())
	}
	ret = fmt.Sprintf("%v[%vm%v", esc, codes, ret)
	// color clear
	ret = fmt.Sprintf("%v%v[%vm", ret, esc, Clear)
	return ret
}

func main() {
	fgblack := GetMultiColorAttribute(232, false)
	fgwhite := GetMultiColorAttribute(255, false)
	fgoffwhite := GetMultiColorAttribute(231, false)
	systemcolor := make([][]int, 0)

	systemcolor = append(systemcolor, []int{0, 0, 0, 0})
	systemcolor = append(systemcolor, []int{1, 197, 15, 31})
	systemcolor = append(systemcolor, []int{2, 19, 161, 14})
	systemcolor = append(systemcolor, []int{3, 193, 156, 0})
	systemcolor = append(systemcolor, []int{4, 0, 55, 218})
	systemcolor = append(systemcolor, []int{5, 136, 23, 152})
	systemcolor = append(systemcolor, []int{6, 58, 150, 221})
	systemcolor = append(systemcolor, []int{7, 204, 204, 204})
	systemcolor = append(systemcolor, []int{8, 118, 118, 118})
	systemcolor = append(systemcolor, []int{9, 231, 72, 86})
	systemcolor = append(systemcolor, []int{10, 22, 198, 12})
	systemcolor = append(systemcolor, []int{11, 249, 241, 165})
	systemcolor = append(systemcolor, []int{12, 59, 120, 255})
	systemcolor = append(systemcolor, []int{13, 180, 0, 158})
	systemcolor = append(systemcolor, []int{14, 97, 214, 214})
	systemcolor = append(systemcolor, []int{15, 242, 242, 242})

	values := []int{0, 95, 135, 175, 215, 255}

	nocolor := make([][]int, 0)
	nocolor = append(nocolor, []int{16, 0x00})
	nocolor = append(nocolor, []int{232, 0x08})
	nocolor = append(nocolor, []int{233, 0x12})
	nocolor = append(nocolor, []int{234, 0x1C})
	nocolor = append(nocolor, []int{235, 0x26})
	nocolor = append(nocolor, []int{236, 0x30})
	nocolor = append(nocolor, []int{237, 0x3A})
	nocolor = append(nocolor, []int{238, 0x44})
	nocolor = append(nocolor, []int{239, 0x4E})
	nocolor = append(nocolor, []int{240, 0x58})
	nocolor = append(nocolor, []int{43, 0x5F})
	nocolor = append(nocolor, []int{241, 0x62})
	nocolor = append(nocolor, []int{242, 0x6C})
	nocolor = append(nocolor, []int{243, 0x76})
	nocolor = append(nocolor, []int{244, 0x80})
	nocolor = append(nocolor, []int{86, 0x87})
	nocolor = append(nocolor, []int{245, 0x8A})
	nocolor = append(nocolor, []int{246, 0x94})
	nocolor = append(nocolor, []int{247, 0x9E})
	nocolor = append(nocolor, []int{248, 0xA8})
	nocolor = append(nocolor, []int{129, 0xAF})
	nocolor = append(nocolor, []int{249, 0xB2})
	nocolor = append(nocolor, []int{250, 0xBC})
	nocolor = append(nocolor, []int{251, 0xC6})
	nocolor = append(nocolor, []int{252, 0xD0})
	nocolor = append(nocolor, []int{172, 0xD7})
	nocolor = append(nocolor, []int{253, 0xDA})
	nocolor = append(nocolor, []int{254, 0xE4})
	nocolor = append(nocolor, []int{255, 0xEE})
	nocolor = append(nocolor, []int{231, 0xFF})

	// system color
	for i, is := range systemcolor {
		c := GetMultiColorAttribute(is[0], true)
		str1 := fmt.Sprintf("%3d %02x", is[0], is[1])
		str2 := fmt.Sprintf("%02x%02x", is[2], is[3])
		fmt.Printf("%v", ApplyAttribute(str1, c, fgblack))
		fmt.Printf("%v", ApplyAttribute(str2, c, fgoffwhite))
		fmt.Printf(" ")
		if i%6 == 5 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
	fmt.Println()

	// 譛牙ｽｩ濶ｲ
	for i := 0; i < 36; i++ {
		for j := 0; j < 6; j++ {
			n := (i * 6) + j
			n2 := n + 16
			c := GetMultiColorAttribute(n2, true)
			str1 := fmt.Sprintf("%3d %02x", n2, values[n/36])
			str2 := fmt.Sprintf("%02x%02x", values[(n/6)%6], values[n%6])
			fmt.Printf("%v", ApplyAttribute(str1, c, fgblack))
			fmt.Printf("%v", ApplyAttribute(str2, c, fgwhite))
			fmt.Printf(" ")
		}
		fmt.Printf(" ")
		for j := 0; j < 6; j++ {
			n := j*36 + i
			n2 := n + 16
			c := GetMultiColorAttribute(n2, true)
			str1 := fmt.Sprintf("%3d %02x", n2, values[n/36])
			str2 := fmt.Sprintf("%02x%02x", values[(n/6)%6], values[n%6])
			fmt.Printf("%v", ApplyAttribute(str1, c, fgblack))
			fmt.Printf("%v", ApplyAttribute(str2, c, fgwhite))
			fmt.Printf(" ")
		}
		fmt.Printf(" ")
		for j := 0; j < 6; j++ {
			n := i/6*36 + i%6 + j*6
			n2 := n + 16
			c := GetMultiColorAttribute(n2, true)
			str1 := fmt.Sprintf("%3d %02x", n2, values[n/36])
			str2 := fmt.Sprintf("%02x%02x", values[(n/6)%6], values[n%6])
			fmt.Printf("%v", ApplyAttribute(str1, c, fgblack))
			fmt.Printf("%v", ApplyAttribute(str2, c, fgwhite))
			fmt.Printf(" ")
		}
		fmt.Println()
	}
	fmt.Println()

	// nocolor
	for i, val := range nocolor {
		hex := val[1]
		c := GetMultiColorAttribute(val[0], true)
		str1 := fmt.Sprintf("%3d %02x", val[0], hex)
		str2 := fmt.Sprintf("%02x%02x", hex, hex)
		fmt.Printf("%v", ApplyAttribute(str1, c, fgblack))
		fmt.Printf("%v", ApplyAttribute(str2, c, fgwhite))
		fmt.Printf(" ")
		if i%15%6 == 5 {
			fmt.Printf(" ")
		}
		if i == 14 {
			fmt.Println()
		}
	}
	fmt.Println()
}

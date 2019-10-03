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

func (a Attribute) String() string{
	attribute := 0
	ret:= ""
	if a < -255 {
		// Background
		attribute = (int(a) + 255 ) * -1
		ret = fmt.Sprintf("%v%v",background,attribute)
	} else if a < 0 {
		// Forground
		attribute = int(a) * -1
		ret = fmt.Sprintf("%v%v",foreground,attribute)
	} else if a < 256{
		ret = fmt.Sprintf("%v",attribute)
	}
	
		return ret
}

func GetMultiColorAttribute(colorCode int, isBackground bool) Attribute {
	if colorCode <0 || 255 < colorCode {
		return 0 
	} else {
		val := colorCode * -1
		if isBackground {
			val += -255 
		} 
		return Attribute(val)
	}
}

func ApplyAttribute(str string,attributes ...Attribute) string {
	ret := str
	codes := attributes[0].String()
	for i := range attributes[1:] {
		codes = fmt.Sprintf("%v;%v",codes,attributes[i+1].String())
	} 
	ret = fmt.Sprintf("%v[%vm%v",esc,codes, ret)
	// color clear
	ret = fmt.Sprintf("%v%v[%vm",ret,esc,Clear)
	return ret
}

func main() {
	fgblack := GetMultiColorAttribute(232,false)
	fgwhite := GetMultiColorAttribute(255,false)
	fgoffwhite := GetMultiColorAttribute(231,false)
	defaultColor := make ([][]int,0)

defaultColor = append (defaultColor,[]int {0,0,0,0})
defaultColor = append (defaultColor,[]int{1,197,15,31 })
defaultColor = append (defaultColor,[]int{2,19,161,14 })
defaultColor = append (defaultColor,[]int{3,193,156,0 })
defaultColor = append (defaultColor,[]int{4,0,55,218 })
defaultColor = append (defaultColor,[]int{5,136,23,152 })
defaultColor = append (defaultColor,[]int{6,58,150,221 })
defaultColor = append (defaultColor,[]int{7,204,204,204 })
defaultColor = append (defaultColor,[]int{8,118,118,118 })
defaultColor = append (defaultColor,[]int{9,231,72,86 })
defaultColor = append (defaultColor,[]int{10,22,198,12 })
defaultColor = append (defaultColor,[]int{11,249,241,165 })
defaultColor = append (defaultColor,[]int{12,59,120,255 })
defaultColor = append (defaultColor,[]int{13,180,0,158 })
defaultColor = append (defaultColor,[]int{14,97,214,214 })
defaultColor = append (defaultColor,[]int{15,242,242,242})


	values := []int{0,95,135,175,215,255}
	for i,is := range defaultColor {
			c := GetMultiColorAttribute(is[0],true) 
			str1 := fmt.Sprintf("%3d %02x",is[0],is[1])
			str2 := fmt.Sprintf("%02x%02x",is[2],is[3])
			fmt.Printf("%v",ApplyAttribute(str1,c,fgblack))
			fmt.Printf("%v",ApplyAttribute(str2,c,fgoffwhite))
			fmt.Printf(" ")
			if i % 6 == 5 {
			fmt.Printf(" ")
			}
	}
			fmt.Println()
			fmt.Println()
	for i := 0 ;i < 36; i++ {
		for j:= 0; j < 6; j++ {
			n := (i*6) + j 
			n2 := n + 16
			c := GetMultiColorAttribute(n2,true) 
			str1 := fmt.Sprintf("%3d %02x",n2,values[n/36])
			str2 := fmt.Sprintf("%02x%02x",values[(n/6)%6],values[n%6])
			fmt.Printf("%v",ApplyAttribute(str1,c,fgblack))
			fmt.Printf("%v",ApplyAttribute(str2,c,fgwhite))
			fmt.Printf(" ")
		}
		fmt.Printf(" ")
		for j:= 0; j < 6; j++ {
			n := j * 36 + i
			n2 := n + 16 
			c := GetMultiColorAttribute(n2,true) 
			str1 := fmt.Sprintf("%3d %02x",n2,values[n/36])
			str2 := fmt.Sprintf("%02x%02x",values[(n/6)%6],values[n%6])
			fmt.Printf("%v",ApplyAttribute(str1,c,fgblack))
			fmt.Printf("%v",ApplyAttribute(str2,c,fgwhite))
			fmt.Printf(" ")
		}
		fmt.Printf(" ")
		for j:= 0; j < 6; j++ {
			n := i / 6 * 36 + i % 6+ j * 6
			n2 := n + 16 
			c := GetMultiColorAttribute(n2,true) 
			str1 := fmt.Sprintf("%3d %02x",n2,values[n/36])
			str2 := fmt.Sprintf("%02x%02x",values[(n/6)%6],values[n%6])
			fmt.Printf("%v",ApplyAttribute(str1,c,fgblack))
			fmt.Printf("%v",ApplyAttribute(str2,c,fgwhite))
			fmt.Printf(" ")
		}
		fmt.Println()
	}
		fmt.Println()
	for i := 0 ;i < 24; i++ {
			n := 232 + i
			c := GetMultiColorAttribute(n,true) 
			hex := fmt.Sprintf("%02x%02x%02x",values[n/36],values[(n/6)%6],values[n%6])
			if i %2 == 0 {
				fmt.Printf("%v",ApplyAttribute(fmt.Sprintf("%3d %s",n,hex),c,fgblack))
			}else {
				fmt.Printf("%v",ApplyAttribute(fmt.Sprintf("%3d %s",n,hex),c,fgwhite)) 
			}
		if i ==  12 {
			fmt.Println()
		}
	}
	fmt.Println()
}

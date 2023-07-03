package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func numbers() {
	// 有符号整数，n-bit的值域为-2^(n-1)到2^(n-1)-1
	// int int8 int16 int32 int64
	// 无符号整数，n-bit的值域为0到2^(n)-1
	// uint uint8 uint16 uint32 uint64

	// rune === int32 === 一个Unicode码点
	// byte === uint8 === 一个字节

	// %取模运算符的符号和被取模数的符号总是一致
	// 除法运算符/的行为则依赖于操作数是否全为整数
	// 整数除法会向着0方向截断余数

	// 浮点数到整数的转换将丢失任何小数部分，然后向数轴零方向截断

	// math.MaxFloat32 为 3.5e38
	// math.MaxFloat64 为 1.8e308

	// 一个float32类型的浮点数可以提供大约6个十进制数的精度，
	// 而float64则可以提供约15个十进制数的精度
	// 当整数大于23bit能表达的范围时，float32的表示将出现误差
}

func pf() {
	// 通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，
	// 但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。
	// 第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	// 字符使用%c参数打印，或者是用%q参数打印带单引号的字符：
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	// 用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，
	// 但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	// 测试一个结果是否是非数NaN则是充满风险的，
	//因为NaN和任何数都不相等
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

	// 如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败
	// func compute()(value float64,ok bool){
	// 	// ...
	// 	if failed {
	// 		return 0, false
	// 	}
	// 	return result, true
	// }
}

func main() {
	// pf()
	// surface()
	// svgServer()
	// complexCalcServer()
	// strings()
	// packages()
	// fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println(notReComma("234235353412"))
}

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func svgServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "plain/text")
		// w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(`<svg version="1.0" xmlns="http://www.w3.org/2000/svg"
		width="1155.000000pt" height="1280.000000pt" viewBox="0 0 1155.000000 1280.000000"
		preserveAspectRatio="xMidYMid meet">
	 <metadata>
	 Created by potrace 1.15, written by Peter Selinger 2001-2017
	 </metadata>
	 <g transform="translate(0.000000,1280.000000) scale(0.100000,-0.100000)"
	 fill="#000000" stroke="none">
	 <path d="M5395 12788 c-32 -12 -5124 -2893 -5140 -2909 -6 -6 -4 -14 5 -23 8
	 -7 780 -377 1715 -821 935 -445 2229 -1060 2875 -1367 923 -440 1186 -561
	 1225 -565 43 -5 60 -1 113 26 50 25 4815 3030 5004 3155 59 39 71 55 59 74 -6
	 10 -408 182 -2396 1027 -159 68 -954 405 -1765 750 -811 345 -1501 636 -1533
	 646 -65 21 -117 23 -162 7z m98 -573 c388 -51 657 -235 657 -449 0 -197 -232
	 -382 -571 -455 -504 -108 -1044 30 -1187 304 -36 68 -37 169 -4 237 63 127
	 218 238 430 307 201 65 448 86 675 56z m1899 -804 c298 -52 507 -173 579 -337
	 17 -38 21 -61 17 -120 -9 -136 -101 -245 -284 -337 -101 -51 -199 -83 -326
	 -109 -161 -33 -464 -32 -616 0 -348 75 -554 239 -554 438 0 223 301 423 717
	 477 98 13 360 6 467 -12z m1835 -812 c225 -27 448 -121 551 -233 57 -62 73
	 -90 88 -153 53 -230 -199 -446 -617 -528 -117 -23 -363 -31 -485 -15 -232 29
	 -438 109 -554 216 -86 80 -112 129 -118 219 -4 63 -1 80 23 134 89 197 429
	 354 810 374 71 4 206 -2 302 -14z m-6505 -28 c367 -63 600 -235 600 -441 0
	 -222 -279 -416 -684 -476 -138 -20 -432 -14 -543 11 -285 64 -475 183 -540
	 339 -24 56 -19 152 10 213 79 170 329 307 651 357 137 21 370 20 506 -3z
	 m1781 -856 c396 -52 660 -232 661 -451 0 -337 -649 -576 -1239 -458 -101 21
	 -199 55 -296 105 -299 155 -334 420 -80 608 222 165 601 242 954 196z m1822
	 -845 c427 -38 726 -225 728 -453 2 -212 -240 -393 -625 -468 -134 -27 -413
	 -32 -538 -10 -306 53 -513 167 -591 326 -28 57 -31 72 -27 132 5 90 36 149
	 118 227 189 178 573 279 935 246z"/>
	 <path d="M11430 9941 c-192 -121 -4517 -2855 -4760 -3007 -151 -96 -296 -193
	 -322 -217 l-48 -43 0 -74 c0 -41 11 -829 25 -1750 14 -921 34 -2307 45 -3080
	 26 -1869 23 -1760 55 -1760 14 0 4991 3310 5058 3363 21 18 45 43 52 57 10 20
	 11 659 6 3290 -5 2891 -8 3265 -21 3267 -8 1 -49 -19 -90 -46z m-343 -1192
	 c156 -75 231 -245 220 -499 -22 -492 -399 -1014 -792 -1097 -90 -19 -162 -10
	 -242 28 -98 47 -170 150 -204 289 -18 75 -15 276 5 369 99 449 414 841 750
	 932 96 26 180 19 263 -22z m-1482 -3021 c97 -50 174 -158 206 -288 8 -33 13
	 -105 13 -180 -1 -140 -28 -273 -83 -413 -120 -303 -336 -558 -564 -666 -105
	 -50 -171 -65 -261 -59 -169 12 -290 132 -333 331 -22 103 -13 291 20 418 105
	 403 397 765 700 865 100 33 228 30 302 -8z m-1688 -3113 c112 -33 199 -123
	 243 -248 31 -88 39 -300 16 -418 -87 -440 -408 -849 -740 -944 -87 -25 -214
	 -17 -281 17 -86 44 -145 119 -187 237 -29 81 -36 286 -14 403 34 178 126 395
	 234 548 60 86 202 231 278 286 157 113 324 157 451 119z"/>
	 <path d="M0 9478 c0 -13 18 -576 40 -1253 79 -2418 100 -3083 129 -4005 40
	 -1259 37 -1198 70 -1230 29 -29 5689 -2990 5715 -2990 9 0 19 6 22 14 6 15
	 -81 6139 -92 6475 l-6 194 -32 33 c-24 26 -374 196 -1611 785 -869 413 -2170
	 1032 -2892 1375 -721 343 -1319 624 -1327 624 -10 0 -16 -9 -16 -22z m815
	 -1023 c342 -81 679 -480 775 -920 27 -125 30 -298 6 -399 -33 -138 -116 -250
	 -224 -303 -72 -35 -196 -43 -291 -18 -341 89 -670 488 -767 932 -23 106 -24
	 306 -1 392 40 150 143 273 260 310 74 24 160 26 242 6z m3621 -1776 c154 -27
	 300 -112 444 -258 194 -195 330 -453 381 -722 18 -95 15 -272 -5 -349 -65
	 -250 -251 -375 -485 -326 -304 64 -623 397 -754 788 -62 184 -80 402 -44 537
	 46 172 148 288 287 325 67 18 99 19 176 5z m-1709 -1198 c123 -43 223 -109
	 344 -230 342 -340 492 -864 341 -1185 -180 -380 -703 -259 -1035 240 -288 433
	 -304 959 -34 1142 71 49 115 61 212 61 74 1 105 -4 172 -28z m-1754 -1137 c29
	 -9 85 -33 125 -55 288 -154 536 -503 618 -871 23 -103 23 -312 0 -398 -91
	 -333 -404 -423 -730 -210 -171 112 -326 293 -431 505 -98 198 -135 344 -135
	 539 0 179 43 307 137 403 103 105 249 135 416 87z m3649 -1866 c332 -123 634
	 -517 713 -929 19 -103 19 -285 0 -367 -34 -145 -111 -249 -223 -304 -62 -31
	 -74 -33 -167 -33 -119 1 -191 22 -305 89 -282 166 -509 499 -586 856 -24 113
	 -25 301 -1 389 44 164 129 266 260 313 74 26 217 20 309 -14z"/>
	 </g>
	 </svg>`))
		// surface()
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calcImg(x, y, width, height int) image.Image {
	xmin, xmax, ymin, ymax := float64(-x), float64(x), float64(-y), float64(y)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	return img
}

func complexCalcServer() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	// const (
	// xmin, ymin, xmax, ymax = -2, -2, 2, 2
	// width, height          = 1024, 1024
	// )

	handler := func(w http.ResponseWriter, r *http.Request) {
		x, err := strconv.Atoi(r.URL.Query().Get("x"))
		y, err := strconv.Atoi(r.URL.Query().Get("y"))
		width, err := strconv.Atoi(r.URL.Query().Get("width"))
		height, err := strconv.Atoi(r.URL.Query().Get("height"))
		fmt.Println(x, y, width, height)

		if err != nil {
			fmt.Fprintf(w, "x is not a number")
		}
		img := calcImg(x, y, width, height)
		png.Encode(w, img)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func stringss() {
	// len(str)返回字节数目
	// str[i]返回第i个字节的字节值
	// 非ASCII字符的UTF-8编码会多两个或多个字节
	// 子字符串操作：str[0:5]

	// 转义
	// \a      响铃
	// \b      退格
	// \f      换页
	// \n      换行
	// \r      回车
	// \t      制表符
	// \v      垂直制表符
	// \'      单引号（只用在 '\'' 形式的rune符号面值中）
	// \"      双引号（只用在 "..." 形式的字符串面值中）
	// \\      反斜杠

	// 八/十六进制转义：\ooo \xhh

	// 原生字符串字面量：`...`

	// unicode码点：\uhhhh

	HasPrefix := func(s, prefix string) bool {
		return len(s) >= len(prefix) && s[:len(prefix)] == prefix
	}

	HasSuffix := func(s, suffix string) bool {
		return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
	}

	Contains := func(s, substr string) bool {
		for i := 0; i < len(s); i++ {
			if HasPrefix(s[i:], substr) {
				return true
			}
		}

		return false
	}

	fmt.Println(HasPrefix("abcdef", "abc1"))
	fmt.Println(HasSuffix("abcdefxyz1", "xyz1"))
	fmt.Println(Contains("abcqdqeqfqxyz1", "dqe"))

	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println("------------------------")

	// range隐式解码
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println("------------------------")

	l := "プログラム"
	fmt.Printf("% x\n", s)
	// 使用[]rune转换为utf-8
	r := []rune(l)
	fmt.Printf("% x\n", r)
	// 还原
	fmt.Println(string(r))
	// 将数字转换为字符串代表对应unicode码点的utf-8字符串
	fmt.Println(string(rune(65)))
	fmt.Println(string(rune(0x4eac)))
	// 无效码点使用\uFFFD代替
	fmt.Println(string(rune(1234567)))
}

func packages() {
	// strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能
	// strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换
	// unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类
	basename := func(s string) string {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '/' {
				s = s[i+1:]
				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '.' {
				s = s[:i]
				break
			}
		}

		return s
	}

	basename2 := func(s string) string {
		slash := strings.LastIndex(s, "/")
		s = s[slash+1:]

		if dot := strings.LastIndex(s, "."); dot >= 0 {
			s = s[:dot]
		}

		return s
	}

	fmt.Println(basename2("a/b/c.go"))
	fmt.Println(basename2("c.d.go"))
	fmt.Println(basename("abc"))
	fmt.Println(basename2("abc"))

}

// path和path/filepath包提供了关于文件路径名更一般的函数操作
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	// 当向bytes.Buffer添加任意字符的UTF8编码时，
	// 最好使用bytes.Buffer的WriteRune方法，
	// 但是WriteByte方法对于写入类似'['和']'等ASCII字符则会更加有效。
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func notReComma(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i > 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}

		buf.WriteByte(s[i])
	}

	return buf.String()
}

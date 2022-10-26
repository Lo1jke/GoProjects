package main

import "fmt"

func main() {
	fmt.Println("Введите число (≤1,000,000): ")
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err, "\n")
		return
	}
	if n > 1000000 || n < -1000000 {
		fmt.Printf("Error:|%d| > 1000000\n", n)
		return
	}

	fmt.Println(n, "Число:")
	fmt.Println(GetNumberName(n, false))
}

func GetNumberName(n int, thou bool) string {

	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	name := ""

	if exceptionsMap[n] != "" {
		name = exceptionsMap[n]
	} else if n > 0 && n < 10 {
		name = onesMap[n]
	} else if n >= 20 && n < 100 {
		tens := n / 10
		ones := n - tens*10
		name = tensMap[tens*10]
		if ones > 0 {
			if thou {
				name += "-" + onesMapT[ones]
			} else {
				name += "-" + onesMap[ones]
			}
		}
	} else if n >= 100 && n < 1000 {
		hundreds := n / 100
		tensAndLower := n - hundreds*100
		name = exceptionsHandredsMap[hundreds]
		if tensAndLower > 0 {
			name += " " + GetNumberName(tensAndLower, thou)
		}
	} else if n >= 1000 && n < 1000000 {
		thousands := n / 1000
		hundredsAndLower := n - thousands*1000
		name = GetNumberName(thousands, true) + " " + thousandsMap[thousands%10]
		if hundredsAndLower >= 100 {
			name += ", " + GetNumberName(hundredsAndLower, thou)
		} else if hundredsAndLower > 0 {
			name += " " + GetNumberName(hundredsAndLower, thou)
		}
	}

	if neg {
		name = "минус " + name
	}

	return name
}

var onesMap = map[int]string{
	1: "один",
	2: "два",
	3: "три",
	4: "четыре",
	5: "пять",
	6: "шесть",
	7: "семь",
	8: "восемь",
	9: "девять",
}

var onesMapT = map[int]string{
	1: "один",
	2: "две",
	3: "три",
	4: "четыре",
	5: "пять",
	6: "шесть",
	7: "семь",
	8: "восемь",
	9: "девять",
}

var thousandsMap = map[int]string{
	1: "тысяча",
	2: "тысячи",
	3: "тысячи",
	4: "тысячи",
	5: "тысяч",
	6: "тысяч",
	7: "тысяч",
	8: "тысяч",
	9: "тысяч",
}

var exceptionsMap = map[int]string{
	0:       "ноль",
	10:      "десять",
	11:      "одинадцать",
	12:      "двенадцать",
	13:      "тринадцать",
	14:      "четырнадцать",
	15:      "пятнадцать",
	16:      "честнадцать",
	17:      "семнадцать",
	18:      "восемьнадцать",
	19:      "девятнадцать",
	1000000: "один миллион",
}

var exceptionsHandredsMap = map[int]string{
	1: "сто",
	2: "двести",
	3: "триста",
	4: "четыреста",
	5: "пятьсот",
	6: "шестьсот",
	7: "семьсот",
	8: "восемьсот",
	9: "девятьсот",
}

var tensMap = map[int]string{
	20: "двадцать",
	30: "тридцать",
	40: "сорок",
	50: "пятьдесят",
	60: "шестьдесят",
	70: "семьдесят",
	80: "восемьдесят",
	90: "девяносто",
}

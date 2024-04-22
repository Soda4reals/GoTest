/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main
import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

type Calculator struct{}

func (c Calculator) calculate(op string, x, y int) int {  
switch op {
case "+":
return x + y
case "-":
return x - y
case "*":
return x * y
case "/":
if y != 0 {
return x / y
} else {
panic("невозможно деление на 0")
}
default:
panic("неверная операция")
}
}

func romanToArabic(roman string) int {             
romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10}
arabic := 0
prevValue := 0
for i := len(roman) - 1; i >= 0; i-- {
value := romanMap[rune(roman[i])]
if value < prevValue {
arabic -= value
} else {
arabic += value
}
prevValue = value
}
return arabic
}

func main() {                                   
calculator := Calculator{}
reader := bufio.NewReader(os.Stdin)

expr := strings.Split(input, " ")
if len(expr) != 3 {
fmt.Println("Неверный формат выражения")
continue
}

num1Str, num2Str := expr[0], expr[2]
num1, err1 := strconv.Atoi(num1Str)
num2, err2 := strconv.Atoi(num2Str)

if err1 != nil || err2 != nil {
num1 = romanToArabic(num1Str)
num2 = romanToArabic(num2Str)

if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
panic("Неверные числа")
}
} else {
if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
panic("Неверные числа")
}

}

result := calculator.calculate(expr[1], num1, num2)

if err1 != nil || err2 != nil {
if result < 0 {
panic("Результат меньше нуля")
}
fmt.Println("Результат:", arabicToRoman(result))
} else {
fmt.Println("Результат:", result)
}
}
}


func arabicToRoman(num int) string {   
vals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
romans := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var res strings.Builder
for i, v := range vals {
for num >= v {
num -= v
res.WriteString(romans[i])
}
}
return res.String()
}
package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	attempts := 0

	for sum := 0; sum != 200; {
		expression := generateExpression()
		result, err := evaluateExpression(expression)
		if err != nil {
			fmt.Printf("Ошибка при вычислении выражения: %s\n", err)
			break
		}

		sum = int(result)
		attempts++

		if sum == 200 {
			fmt.Printf("Найдено выражение: %s\n", expression)
			break
		}

		if attempts > 1000000 {
			fmt.Println("Не удалось найти выражение. Попробуйте увеличить количество попыток.")
			break
		}
	}
}

func generateExpression() string {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	operators := []string{"+", "-", "*", "/"}

	expression := ""

	for i := 0; i <= 8; i++ {
		expression += fmt.Sprintf("%d%s", digits[i], operators[rand.Intn(len(operators))])
	}

	expression += fmt.Sprintf("%d", 9)

	return expression
}

func evaluateExpression(expression string) (float64, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	return result.(float64), nil
}

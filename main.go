package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// --- Получение данных ---
	arguments := os.Args

	// Проверяем, что пользователь предоставил достаточно данных:
	// имя программы, операция и строка с числами (всего 3 элемента).

	if len(arguments) < 3 {
		fmt.Println("Пожалуйста, укажите операцию (AVG, SUM, MED) и строку чисел через запятую.")
		fmt.Println("Пример: go run main.go AVG 1,2,3")
		os.Exit(1) // Останавливаем программу, если данных не хватает
	}

	// Извлекаем операцию и строку с числами из среза аргументов.
	operation := arguments[1]     // Сохранение операции "AVG"
	numbersString := arguments[2] // Сохранение строки с числами

	// --- Подготовка данных ---
	numbersAsString := strings.Split(numbersString, ",")

	// Создаем пустой срез, куда будем складывать наши числа
	numbers := []int{}

	// Проходим в цикле по срезу строк, чтобы преобразовать каждую в число.
	for _, s := range numbersAsString {
		// strconv.Atoi (ASCII to integer) конвертирует строку в число.
		// Функция возвращает число и возможную ошибку.
		num, err := strconv.Atoi(s)

		//Проверяем ошибку сразу после вызова функции.
		if err != nil {
			fmt.Println("Ошибка! Не удалось преобразовать строку в число. Проверьте валидность.")
			os.Exit(1)
		}

		// Если ошибки не было, добавляем сконвертированное число в наш срез.
		numbers = append(numbers, num)
	}

	// --- Dвыбор и выполнение операции ---

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	quantityElements := len(numbers)

	//Выбора действия в зависимости от введенной операции.
	switch operation {

	// Для суммы нам нужно просто посчитать ее и вывести.
	case "SUM":

		fmt.Println(sum)

	// Для среднего считаем сумму и делим на количество элементов
	case "AVG":

		fmt.Println(float64(sum) / float64(quantityElements))

	// Для медианы первым делом нужно отсортировать срез.
	case "MED":

		sort.Ints(numbers)
		middle := quantityElements / 2

		// Проверяем, является ли количество элементов нечетным.
		if len(numbers)%2 != 0 {

			fmt.Println(numbers[middle])

			// Если четное, медиана - это среднее арифметическое двух центральных элементов.
		} else {
			fmt.Println((numbers[middle-1] + numbers[middle]) / 2)
		}

	// Если введена операция, которую мы не знаем, сообщаем об ошибке.
	default:
		fmt.Println("Ошибка! Неизвестная операция")
		os.Exit(1)
	}

}

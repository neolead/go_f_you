package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

// Генерация всех байтов для заданных кодировок, включаем мультибайтные символы.
func generateAllBytes(excludeAlphabet bool) []byte {
	var buf bytes.Buffer
	for i := 0; i <= 0x10FFFF; i++ {
		if utf8.ValidRune(rune(i)) {
			if excludeAlphabet {
				// Пропускаем буквы и цифры
				if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') || (i >= '0' && i <= '9') {
					continue
				}
			}
			buf.WriteRune(rune(i))
		}
	}
	return buf.Bytes()
}

// Функция для записи в текстовый файл с разделением символов через новую строку
func writeToTextFile(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, b := range data {
		_, err := file.WriteString(fmt.Sprintf("%c\n", b))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

// Функция для записи в бинарный файл
func writeToBinaryFile(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating binary file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to binary file:", err)
	}
}

func main() {
	// Генерация данных с алфавитом и без
	allBytes := generateAllBytes(false)
	allBytesWithoutAlphabet := generateAllBytes(true)

	// Запись в текстовые файлы
	writeToTextFile("fuzz_with_alphabet.txt", allBytes)
	writeToTextFile("fuzz_without_alphabet.txt", allBytesWithoutAlphabet)

	// Запись в бинарные файлы
	writeToBinaryFile("fuzz_with_alphabet.bin", allBytes)
	writeToBinaryFile("fuzz_without_alphabet.bin", allBytesWithoutAlphabet)

	fmt.Println("Файлы успешно сгенерированы.")
}

package main

import (
	"fmt"
	"os"

	"github.com/gagliardetto/solana-go"
)

func main() {
	// Запрашиваем у пользователя количество кошельков для генерации
	fmt.Print("Введите количество кошельков для генерации: ")
	var numWallets int
	_, err := fmt.Scanln(&numWallets)
	if err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число.")
		return
	}

	// Открываем файл для записи
	file, err := os.Create("wallets.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Генерируем кошельки и записываем в файл
	for i := 0; i < numWallets; i++ {
		account := solana.NewWallet()
		publicKey := account.PublicKey().String()
		privateKey := account.PrivateKey.String()

		// Записываем данные в файл
		_, err := file.WriteString(fmt.Sprintf("%s:%s\n", publicKey, privateKey))
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	fmt.Printf("Успешно сгенерировано и сохранено %d кошельков в файл wallets.txt\n", numWallets)

	// Добавляем паузу в конце, чтобы пользователь мог увидеть результат
	fmt.Println("Нажмите Enter, чтобы завершить программу...")
	fmt.Scanln() // Ждем, пока пользователь нажмет Enter
}

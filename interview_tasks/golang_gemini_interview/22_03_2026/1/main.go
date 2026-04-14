package main


var smallResult []byte

func storeResult() {
	bigData := make([]byte, 100*1024*1024) // 100 MB
	// ... заполняем bigData данными ...
	smallResult = bigData[:10]
}

func main() {
	storeResult()
	// Функция завершилась, но smallResult используется дальше.
}
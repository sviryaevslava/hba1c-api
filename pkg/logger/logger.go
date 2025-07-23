package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {
	// создаём лог-файл
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка при создании лог-файла: %v", err)
	}

	// создаём логгер
	Log = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	Log.Println("=== Логирование запущено ===")
}

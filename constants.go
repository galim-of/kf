package main

type userInput struct {
	
	server 	string // name of the server
	command string // key that will be executed
	param 	string // optional key
	files 	[]string // arguments for key
}

// Валидные ключи программы
const (
	server = "-s" 
	find = "-f"
	download = "-d"
	path = "-p"
	help = "-h"
)


const (
	maxdownload = 5
	minargs = 5
)

const helpString string =
				"kf -s Сервер -f | -d [-p Путь] Карта[[карта1, карта2, ... картаN]*]\n" +
				" -s - определяет сервер, на котором выполняется операция\n" +
				" -f - ищет все совпадения только для одной строки\n" +
				" -d - скачивает карту(-ы). По умолчанию, карты скачиваются в текущий каталог.\n" +
				" -p - необязательный параметр, определяет путь куда качать карты.\n" +
				" Пример: kf -s test.server.ru:7709 -f Biotics\n" +
				" программа покажет все совпадения для заданной строки\n" +
				" Еще пример: kf -s test.server.ru:7709 -f *\n" +
				" будут отображены все карты\n" +
				" Пример:\n" +
				" kf test.server.ru:7709 -d -p \"Killing Floor\\Maps\" KF-Bioticslab.rom\n" +
				" Если в пути имеются пробелы или символы кириллицы, тогда путь нужно взять в кавычкы.\n"
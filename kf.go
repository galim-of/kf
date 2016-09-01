package main

import (
	"fmt"
	"os"
	"strings"
	"net/http"
)
type userInput struct {
	server string
	command string
	param string
	files []string
}
const (
	SERVER = "-s"
	FIND = "-f"
	DOWNLOAD = "-d"
	PATH = "-p"
	HELP = "-h"

	MAXDOWNLOAD = 5
	MIN_ARGS = 5
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


func main() {
	switch len(os.Args) {
		case 1: promptHowToUse(helpString);
		case 2: {
					if os.Args[1] == HELP {
						promptHowToUse(helpString)
					} else {os.Exit(1)}
				}
	}
	if len(os.Args) < MIN_ARGS {os.Exit(1)}

	var keys = []string{SERVER, FIND, DOWNLOAD, PATH, HELP}
	var u userInput

	firstArgumentIsKey(keys)
	validKeys(keys, len(keys))
	keyAfterKey(keys)
	userKeys, count := countOfKeys(keys)
	switch count {
		// case 0: {fmt.Println("No keys"); os.Exit(1)}
		case -1: {fmt.Println("Finded dublicates"); os.Exit(1)}
		case 1: {fmt.Println("Not enough keys"); os.Exit(1)}
		default : {
			var userParams = make(map[string][]string)
			for _, key := range userKeys  {
				userParams[key] = countOfparametersForKey(key, keys)
			}
			for key, params := range userParams {
				switch key {
					case SERVER: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
					case FIND: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
					case DOWNLOAD: if len(params) > MAXDOWNLOAD || len(params) == 0 {fmt.Println(key, len(params));os.Exit(1)}
					case PATH: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
					case HELP: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
				}
			}
			keysAreCompatible(userKeys)
			u = fillStruct(userParams)
	}
	// fmt.Println("Good job you pass all checks")
	}
	connectToServer(u.server)
	if u.command == FIND {
		findMap(u.files[0])
	} else {
		for _, f := range u.files {
			downloadMap(f)
		}
	}

		
}

func fillStruct(params map[string][]string) (u userInput) {
	for key, v := range params {
		switch key {
			case SERVER: u.server = v[0]
			case FIND: u.command = FIND; u.files = v
			case DOWNLOAD: u.command = DOWNLOAD; u.files = v
			case PATH: u.param = v[0]
		} 
	}
	return
}
func connectToServer(s string) {
	conn, err := http.Get(s+"/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connecting to", s)
}
func findMap(m string) (n int) {
	var listOfFiles = make([]os.FileInfo, 0, 50)
	var temp = make([]string, 0, 25)
	fd, err := os.Open("Map")
	defer fd.Close()
	listOfFiles, err = fd.Readdir(50)
	if err != nil {fmt.Println(err)}
	if m == "*" {
		for _ , v := range listOfFiles {
			fmt.Println("#", v.Name())
		}
		return len(listOfFiles)
	}
	for j := 0; j < len(listOfFiles); j++ {
		if elem := listOfFiles[j]; strings.Contains(elem.Name(), m) == true {
			temp = append(temp, elem.Name())
		}
	}
	if len(temp) == 0 {
		fmt.Println("Совпадений не найдено")
		return 0
	}
	fmt.Printf("Найдены совпадения для %q в кол-ве %d: \n", m, len(temp))
	for _ , v := range temp {
		fmt.Println("#",v)
	}
	return len(temp)
	// fmt.Println()		
}
func downloadMap(m string) {
	fmt.Println("Downloading", m)
}
func keysAreCompatible(keys []string) {
	var sumPoints int
	// var points int
	var tempMap = make(map[string]int, len(keys))
	for _, key := range keys {
		switch key {
			case SERVER: tempMap[SERVER] = 0
			case DOWNLOAD: tempMap[DOWNLOAD] = 0
			case FIND: tempMap[FIND] = 0
			case PATH: tempMap[PATH] = 0
			case HELP: tempMap[HELP] = 0
		}
	}
	for key, points := range tempMap {
		switch key {
			case SERVER: points = 15; sumPoints += points; break
			case DOWNLOAD: points = 10; sumPoints += points; break
			case FIND: points = 5; sumPoints += points; break
			case PATH: points = 2; sumPoints += points; break
			case HELP: points = 1; sumPoints += points; break
		}
	}
	// fmt.Println(tempMap)
	// if sumPoints != 25 || sumPoints != 20 || sumPoints != 27 {
	switch sumPoints {
		case 25: return
		case 20: return
		case 27: return
		default : {fmt.Println("Non legal"); os.Exit(1)}
	}

}


func validKeys(keys []string, l int) {
// Функция ищет первый неправильный ключ, и если такой имеется, то завершает работу программы
	var count = 0 // количество несовпадений
	// var keys = [3]string{SERVER, FIND, DOWNLOAD}
	for i:= 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") == true {
			for _,key := range keys {// делаем перебор по всем значениеям keys
				if os.Args[i] == key {
					count = 0
					break
				} else {count++}
				
				if count == l {
					fmt.Println("invalid key", os.Args[i])
					os.Exit(1)
				}
			}
		}
	}
}
func countOfKeys(keys []string) (out []string, count int) {
// Функция возвращает количество ключей, в зависимости от ситуации count принимает следующие значения:
// 0: если не было найдено ни одного ключа 
// -1: если был найден дубликат ключа
// n: во всех остальных случаях
// Старое описание функции, не соотвествуюющее действительности
	var temp = make(map[string]int, len(keys))
	for i := 1; i < len(os.Args); i++ {
		for _,key := range keys {
			if os.Args[i] == key {
				out = append(out, key)
				temp[key]++
				count++
				// break
			}
			if temp[key] > 1 {
				count = -1
				out = out[:0]
				return 
			}
		}
	}
	return
}
func keyAfterKey(keys []string) {
	var end = len(os.Args) - 1
	// for i := 1; i < len(os.Args); i++ {
	for i := 1; i < end; i++ {
		// if i == end {return}
		for _, key := range keys {
			if os.Args[i] == key {
				for _, anotherKey := range keys {
					if os.Args[i+1] == anotherKey {
						fmt.Println("Key must not be after key")
						os.Exit(1)
					}
				}
			}
		}
	}
}

func countOfparametersForKey(k string, keys []string) (parameters []string){
	var end = len(os.Args) - 1
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == k {
			// if i == end {count = 0; return}
			if i == end {return}
			for j := i + 1; j < len(os.Args); j++ {
				for _, key := range keys {
					if os.Args[j] == key {
						return
					}
				}
				// count++
				parameters = append(parameters, os.Args[j])
			}
		}
	}
	return
}
func firstArgumentIsKey(keys []string) {
	var count int
	for _, key := range keys {
		if os.Args[1] == key {
			return
		} else {count++}
		if count == len(keys) {fmt.Println("First argument must be valid key");os.Exit(1)}
	}
}	
func promptHowToUse(s string) {
	fmt.Println(s)
	os.Exit(1)
}
func clearMapFromGarbage(m map[string][]string) (map[string][]string) {
// Функция удаляет элементы карты, которые равны 0
	for key, params := range m {
		if len(params) == 0 {
			delete(m, key)
		}
		// if count == -1 {
		// 	n[key] = 0
		// }
	}
	return m
}
// func findKey(key string, keys []string) (string) {
// 	for _, findedKey := range keys {
// 		if key == findedKey {
// 			return key
// 		}
// 	}
// }
// func dublicatesOfKeys(keys []string) {
// // Функция находит первый дубликат ключа и завершает программу
// 	var temp = make(map[string]int, len(keys))
// 	for i := 1; i < len(os.Args); i++ {
// 		for _,key := range keys {
// 			if os.Args[i] == key {temp[key]++; break}
// 			if temp[key] > 1 {os.Exit(1)}
// 		}
// 	}
// }
// func countOfparametersForKey(keys []string) (count int){
// 	var temp = make(map[string]int, len(keys))
// 	for i := 1; i < len(os.Args); i++ {
// 		for _, key := range keys	
// 			if os.Args[i] == key {
// 				for j := i + 1; j < len(os.Args); j++ {
// 					for _, key := range keys {
// 						if os.Args[j] == key {
// 							return
// 						}
// 					}	
// 					count++
// 				}
// 			}
// 		}
// 		return
// }
// func countOfKeys(keys []string) (count int) {
// // Функция возвращает количество ключей, в зависимости от ситуации count принимает следующие значения:
// // 0: если не было найдено ни одного ключа 
// // -1: если был найден дубликат ключа
// // n: во всех остальных случаях 
// 	var temp = make(map[string]int, len(keys))
// 	for i := 1; i < len(os.Args); i++ {
// 		for _,key := range keys {
// 			if os.Args[i] == key {temp[key]++;
// 				count++
// 				// break
// 			}
// 			if temp[key] > 1 {
// 				count = -1
// 				return 
// 			}
// 		}
// 	}
// 	return
// }
// func getMapsFromArgs(arg []string) ([]string) {
// 	arg = arg[2:]
// 	maps := make([]string, 10)
// 	for i , v := range arg {
// 		maps[i] = v
// 	}
// 	maps = maps[0:len(arg)]
// 	return maps
// }
// func getListOfArgs(args []string) (user userInput) {
// 	args = args[1:]
// 	var from int
// 	for i,_:= range args {
// 		if args[i] == SERVER {user.server = args[i+1]}
// 		if args[i] == FIND {user.command = FIND; from = i+1} 
// 		if args[i] == DOWNLOAD {user.command = DOWNLOAD; from = i+1} 
// 	}
// 	files := make([]string, 10)
// 	for j := 0; j < len(files); j++ {
// 		for i := from; i < len(args); i++ {
// 			files[j] = args[i]
// 		}
// 	}
// 	return
// }
// func anotherFunc() (u userInput) {
		// for i := 1; i < len(os.Args); i++ {
	// 	if os.Args[i] == SERVER {u.server = os.Args[i + 1]}
	// 	if os.Args[i] == FIND {u.command = FIND; u.files = append(u.files, os.Args[i + 1])}
	// 	if os.Args[i] == DOWNLOAD {
	// 		u.command = DOWNLOAD
	// 		// u.files = append(u.files, os.Args[i + 1])
	// 		for j := i+1; j < MAXDOWNLOAD; j++ {
	// 			u.files = append(u.files, os.Args[j])
	// 		} 
	// 	}
	// 	if os.Args[i] == PATH {u.param = os.Args[i + 1]}
	// }
	// return
// }
package main

import (
	"fmt"
	"os"
	"strings"
)

func countOfParameters(user map[string][]string) {
	for key, params := range user {
				switch key {
					case server: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
					case find: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
					case download: if len(params) > maxdownload || len(params) == 0 {fmt.Println(key, len(params));os.Exit(1)}
					case path: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
					case help: if len(params) != 1 {fmt.Println(key, len(params));os.Exit(1)}
				}
			}
}

func fillStruct(params map[string][]string) (u userInput) {
	for key, v := range params {
		switch key {
			case server: u.server = v[0]
			case find: u.command = find; u.files = v
			case download: u.command = download; u.files = v
			case path: u.param = v[0]
		} 
	}
	return
}

func keysAreCompatible(keys []string) {
	var sumPoints int
	// var points int
	var tempMap = make(map[string]int, len(keys))
	for _, key := range keys {
		switch key {
			case server: tempMap[server] = 0
			case download: tempMap[download] = 0
			case find: tempMap[find] = 0
			case path: tempMap[path] = 0
			case help: tempMap[help] = 0
		}
	}	
	for key, points := range tempMap {
		switch key {
			case server: points = 15; sumPoints += points; break
			case download: points = 10; sumPoints += points; break
			case find: points = 5; sumPoints += points; break
			case path: points = 2; sumPoints += points; break
			case help: points = 1; sumPoints += points; break
		}
	}
	// fmt.Println(tempMap)
	// if sumPoints != 25 || sumPoints != 20 || sumPoints != 27 {
	switch sumPoints {
		case 25: return
		case 20: return
		case 27: return
		default : {fmt.Println("Non legal", sumPoints); os.Exit(1)}
	}

}

func validKeys(keys []string, l int) {
// Функция ищет первый неправильный ключ, и если такой имеется, то завершает работу программы
	var count = 0 // количество несовпадений
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
		if os.Args[1] != key {
			count++
		} else {return}
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
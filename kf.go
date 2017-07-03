/*
Пакет kf реализует маенлькую программу, которая позволяет искать и скачивать карты с сервера
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {promptHowToUse(helpString)}
	if len(os.Args) == 2 {
		if os.Args[1] == help {
			promptHowToUse(helpString)
		} else {os.Exit(1)}
	}
	if len(os.Args) < minargs {os.Exit(1)}

	var keys = []string{server, find, download, path, help}
	var u userInput



	firstArgumentIsKey(keys)
	validKeys(keys, len(keys))
	keyAfterKey(keys)
	userKeys, count := countOfKeys(keys)
	switch count {
		case -1: {fmt.Println("Finded dublicates"); os.Exit(1)}
		case 1: {fmt.Println("Not enough keys"); os.Exit(1)}
		default : {
			var userParams = make(map[string][]string)
			for _, key := range userKeys  {
				userParams[key] = countOfparametersForKey(key, keys)
			}
			countOfParameters(userParams)
			keysAreCompatible(userKeys)
			u = fillStruct(userParams)
	}
	// Момент, когда пройдены все проверки
	}
	connectToServer(u.server)
	if u.command == find {
		findMap(u.files[0])
	} else {
		for _, f := range u.files {
			downloadMap(u.server, f, u.param)
			fmt.Println("Download for", f, "completed")
		}
	}

		
}

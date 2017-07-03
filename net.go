package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func connectToServer(s string) {
	_, err := http.Head("http://" + s)
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

func downloadMap(s, m, p string) {
	resp, err := http.Get("http://" + s + "/" + m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	if resp.StatusCode == 404 {
		fmt.Println("File not found")
		os.Exit(1)
	}	
	defer resp.Body.Close()
	
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(p + m, data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
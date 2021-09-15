package main

import "log"

func service1(bytes []byte) {
	log.Println("service 1 excuted")
	log.Println(string(bytes))
}

func service2(bytes []byte) {
	log.Println("service 2 excuted")
}

func service3(bytes []byte) {
	log.Println("service 3 excuted")
}

func service4(bytes []byte) {
	log.Println("service 4 excuted")
}

func service5(bytes []byte) {
	log.Println("service 5 excuted")
}

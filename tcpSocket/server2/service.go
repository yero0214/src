package main

import (
	"encoding/binary"
	"log"
)

func move(userId uint32, bytes []byte) {
	log.Println("move excuted")
	log.Println(bytes)

	for i, v := range users {
		if userId == v.userId {
			users[i].cx = binary.LittleEndian.Uint32(bytes[:4])
			users[i].cy = binary.LittleEndian.Uint32(bytes[4:8])
		}
	}
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

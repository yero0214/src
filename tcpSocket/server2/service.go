package main

import (
	"log"
)

func move(userId uint32, bytes []byte) {
	log.Println("move excuted")
	// log.Println(bytes)

	for i, v := range users {
		if userId == v.userId {
			position := byteSliceToFloat32Slice(bytes)
			log.Println(position)
			users[i].cx = position[0]
			users[i].cy = position[1]
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

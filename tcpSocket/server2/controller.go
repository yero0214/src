package main

import (
	"encoding/binary"
	"log"
)

func controller() {
	for {
		buffer := <-recvChan
		typeOfService := binary.LittleEndian.Uint32(buffer[:4])
		userId := binary.LittleEndian.Uint32(buffer[4:8])
		payloadLength := binary.LittleEndian.Uint32(buffer[8:12]) + 12

		log.Print("typeOfService: ")
		log.Println(typeOfService)

		switch typeOfService {
		case 1:
			// move
			move(userId, buffer[12:payloadLength])
		case 2:
			service2(buffer[12:payloadLength])
		case 3:
			service3(buffer[12:payloadLength])
		case 4:
			service4(buffer[12:payloadLength])
		case 5:
			service5(buffer[12:payloadLength])
		}
	}
}

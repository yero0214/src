package main

import (
	"encoding/binary"
	"log"
)

func controller() {
	for {
		buffer := <-recvChan
		typeOfService := binary.LittleEndian.Uint32(buffer[:4])
		payloadLength := binary.LittleEndian.Uint32(buffer[4:8]) + 8

		log.Print("typeOfService: ")
		log.Println(typeOfService)

		switch typeOfService {
		case 1:
			service1(buffer[8:payloadLength])
		case 2:
			service2(buffer[8:payloadLength])
		case 3:
			service3(buffer[8:payloadLength])
		case 4:
			service4(buffer[8:payloadLength])
		case 5:
			service5(buffer[8:payloadLength])
		}
	}
}

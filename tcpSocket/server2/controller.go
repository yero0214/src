package main

import (
	"encoding/binary"
	"log"
)

func controller() {
	for {
		recv := <-recvChan
		typeOfService := binary.LittleEndian.Uint32(recv.buffer[:4])
		// userId := binary.LittleEndian.Uint32(recv.buffer[4:8])
		payloadLength := binary.LittleEndian.Uint32(recv.buffer[8:12]) + 12

		log.Print("typeOfService: ")
		log.Println(typeOfService)

		switch typeOfService {
		case 1:
			// move
			move(recv.userId, recv.buffer[12:payloadLength])
		case 2:
			service2(recv.buffer[12:payloadLength])
		case 3:
			service3(recv.buffer[12:payloadLength])
		case 4:
			service4(recv.buffer[12:payloadLength])
		case 5:
			service5(recv.buffer[12:payloadLength])
		}
	}
}

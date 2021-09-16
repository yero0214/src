package main

import (
	"encoding/binary"
	"time"
)

func inGame() {
	for {
		time.Sleep(time.Second * 5)

		var buffer []byte
		var payload []byte
		typeOfService := make([]byte, 4)
		payloadLength := make([]byte, 4)
		for _, v := range users {
			userId := make([]byte, 4)
			x := make([]byte, 4)
			y := make([]byte, 4)

			binary.LittleEndian.PutUint32(userId, v.userId)
			binary.LittleEndian.PutUint32(x, v.x)
			binary.LittleEndian.PutUint32(y, v.y)

			payload = append(payload, userId...)
			payload = append(payload, x...)
			payload = append(payload, y...)
		}
		binary.LittleEndian.PutUint32(typeOfService, 1)
		binary.LittleEndian.PutUint32(payloadLength, uint32(len(payload)))

		packetHeader := append(typeOfService, payloadLength...)
		buffer = append(buffer, packetHeader...)
		buffer = append(buffer, payload...)

		broadCast(buffer)
	}
}

func xUpdate() {
	for {
		time.Sleep(time.Second)
		for i, _ := range users {
			if users[i].cx > users[i].x {
				users[i].x++
			} else if users[i].cx < users[i].x {
				users[i].x--
			} else {
				continue
			}
		}
	}
}

func yUpdate() {
	for {
		time.Sleep(time.Second)
		for i, _ := range users {
			if users[i].cy > users[i].y {
				users[i].y++
			} else if users[i].cy < users[i].y {
				users[i].y--
			} else {
				continue
			}
		}
	}
}

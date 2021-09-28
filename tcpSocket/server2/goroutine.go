package main

import (
	"encoding/binary"
	"math"
	"time"
)

func inGame() {
	for {
		time.Sleep(time.Millisecond * 16)

		var buffer []byte
		var payload []byte
		typeOfService := make([]byte, 4)
		payloadLength := make([]byte, 4)
		for _, v := range users {
			if v.x != v.cx || v.y != v.cy {
				userId := make([]byte, 4)
				x := make([]byte, 4)
				y := make([]byte, 4)
				math.Float32bits(v.x)
				binary.LittleEndian.PutUint32(userId, v.userId)
				binary.LittleEndian.PutUint32(x, math.Float32bits(v.x))
				binary.LittleEndian.PutUint32(y, math.Float32bits(v.y))
				payload = append(payload, userId...)
				payload = append(payload, x...)
				payload = append(payload, y...)
			}
		}
		binary.LittleEndian.PutUint32(typeOfService, 0)
		binary.LittleEndian.PutUint32(payloadLength, uint32(len(payload)))

		packetHeader := append(typeOfService, payloadLength...)
		buffer = append(buffer, packetHeader...)
		buffer = append(buffer, payload...)
		if len(payload) != 0 {
			broadCast(buffer)
		}
	}
}

func positionUpdate() {
	speed := float32(5)
	for {
		time.Sleep(time.Millisecond * 16)
		for i, _ := range users {
			if users[i].cx > users[i].x+speed {
				users[i].x += speed
			} else if users[i].cx < users[i].x-speed {
				users[i].x -= speed
			}
			if users[i].cy > users[i].y+speed {
				users[i].y += speed
			} else if users[i].cy < users[i].y-speed {
				users[i].y -= speed
			}
		}
	}
}

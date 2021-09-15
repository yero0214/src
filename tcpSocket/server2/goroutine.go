package main

import "time"

func inGame() {
	for {
		time.Sleep(time.Second * 5)
		var result string
		for _, v := range users {
			result += makeNo(v.no)
			result += makeNo(v.x)
			result += makeNo(v.y)
		}

		broadCast(result)
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

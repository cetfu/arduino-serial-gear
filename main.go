package main

import (
	"bufio"
	"github.com/micmonay/keybd_event"
	"go.bug.st/serial"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	gear1  = keybd_event.VK_1
	gear2  = keybd_event.VK_2
	gear3  = keybd_event.VK_3
	gear4  = keybd_event.VK_4
	gear5  = keybd_event.VK_5
	gear6  = keybd_event.VK_6
	gear0  = keybd_event.VK_7
	gearR1 = keybd_event.VK_8
)

func getGear(gear int) int {
	switch gear {
	case 1:
		return gear1
	case 2:
		return gear2
	case 3:
		return gear3
	case 4:
		return gear4
	case 5:
		return gear5
	case 6:
		return gear6
	default:
		return 0
	}
}

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	list, err := serial.GetPortsList()
	if err != nil {
		panic(err)
		return
	}

	for i, e := range list {
		println(i, e)
	}

	reader := bufio.NewReader(os.Stdin)
	print("Enter your device's port: ")
	read, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
		return
	}
	input := strings.TrimSuffix(read, "\n")
	number, _ := strconv.Atoi(input)
	portName := list[number]

	mode := &serial.Mode{
		BaudRate: 9600,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		panic(err)
		return
	} else {
		println("SERIAL CONNECTION OK.")
	}
	serialData := make([]byte, 100)
	for {
		n, err := port.Read(serialData)
		if err != nil {
			panic(err)
			return
		}
		if n == 0 {
			println("\nEOF")
			break
		}
		signal := string(serialData[:n])
		if ok := strings.HasPrefix(signal, "G_CH"); ok {
			d := strings.Split(signal, "_")
			number, err := strconv.Atoi(strings.TrimSpace(d[3]))

			if err != nil {
				panic(err)
				return
			}

			kb.SetKeys(getGear(number))
			if err := kb.Press(); err != nil {
				panic(err)
				return
			}
			time.Sleep(10 * time.Millisecond)
			if err := kb.Release(); err != nil {
				return
			}
		}
	}
}

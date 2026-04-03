package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func weatherName(id uint16) string {
	switch id {
	case 0:
		return "SUNSHINE"
	case 1:
		return "SUNSHINE"
	case 2:
		return "CLOUDS"
	case 3:
		return "FOG"
	case 4:
		return "HOT_SPELL"
	case 5:
		return "HEAT_WAVE"
	case 6:
		return "RAIN"
	case 7:
		return "SQUALL"
	case 8:
		return "DUST_STORM"
	case 9:
		return "SAND_STORM"
	case 10:
		return "WIND"
	case 11:
		return "GALES"
	case 12:
		return "SNOW"
	case 13:
		return "BLIZZARDS"
	case 14:
		return "THUNDER"
	case 15:
		return "THUNDERSTORMS"
	case 16:
		return "AURORAS"
	case 17:
		return "STELLAR_GLARE"
	case 18:
		return "GLOOM"
	case 19:
		return "DARKNESS"
	default:
		return "UNKNOWN"
	}
}

func decodePacked(v uint16) (uint16, uint16, uint16) {
	return v >> 10, (v >> 5) & 0x1F, v & 0x1F
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	eBlob := os.Getenv("BLOB")
	blob := strings.TrimPrefix(eBlob, "0x")

	raw, err := hex.DecodeString(blob)
	if err != nil {
		log.Fatal(err)
	}
	if len(raw)%2 != 0 {
		log.Fatal("blob length must be even")
	}

	startDay := 1144
	count := 8

	for day := startDay; day < startDay+count && day < len(raw)/2; day++ {
		v := binary.LittleEndian.Uint16(raw[day*2 : day*2+2])
		n, c, r := decodePacked(v)

		fmt.Printf(
			"day=%d raw=%04x v=%d => %s | %s | %s\n",
			day,
			v,
			v,
			weatherName(n),
			weatherName(c),
			weatherName(r),
		)
	}
}

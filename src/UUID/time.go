package uuid

import (
	"encoding/binary"
	"sync"
	"time"
)

type Time int64

const (
	lillian    = 2299160          // Julian day of 15 Oct 1582
	unix       = 2440587          // Julian day of 1 Jan 1970
	epoch      = unix - lillian   // Days between epochs
	g1582      = epoch * 86400    // seconds between epochs
	g1582ns100 = g1582 * 10000000 // 100s of a nanoseconds between epochs
)

var (
	timeMu   sync.Mutex
	lasttime uint64 // last time we returned
	clockSeq uint16 // clock sequence for this run

	timeNow = time.Now // for testing
)

func (t Time) UnixTime() (sec, nsec int64) {
	sec = int64(t - g1582ns100)
	nsec = (sec % 10000000) * 100
	sec /= 10000000
	return sec, nsec
}

func GetTime() (Time, uint16, error) {
	defer timeMu.Unlock()
	timeMu.Lock()
	return getTime()
}

func getTime() (Time, uint16, error) {
	t := timeNow()

	if clockSeq == 0 {
		setClockSequence(-1)
	}
	now := uint64(t.UnixNano()/100) + g1582ns100

	if now <= lasttime {
		clockSeq = ((clockSeq + 1) & 0x3fff) | 0x8000
	}
	lasttime = now
	return Time(now), clockSeq, nil
}

func ClockSequence() int {
	defer timeMu.Unlock()
	timeMu.Lock()
	return clockSequence()
}

func clockSequence() int {
	if clockSeq == 0 {
		setClockSequence(-1)
	}
	return int(clockSeq & 0x3fff)
}

func SetClockSequence(seq int) {
	defer timeMu.Unlock()
	timeMu.Lock()
	setClockSequence(seq)
}

func setClockSequence(seq int) {
	if seq == -1 {
		var b [2]byte
		randomBits(b[:]) // clock sequence
		seq = int(b[0])<<8 | int(b[1])
	}
	oldSeq := clockSeq
	clockSeq = uint16(seq&0x3fff) | 0x8000 // Set our variant
	if oldSeq != clockSeq {
		lasttime = 0
	}
}

func (uuid UUID) Time() Time {
	time := int64(binary.BigEndian.Uint32(uuid[0:4]))
	time |= int64(binary.BigEndian.Uint16(uuid[4:6])) << 32
	time |= int64(binary.BigEndian.Uint16(uuid[6:8])&0xfff) << 48
	return Time(time)
}

func (uuid UUID) ClockSequence() int {
	return int(binary.BigEndian.Uint16(uuid[8:10])) & 0x3fff
}
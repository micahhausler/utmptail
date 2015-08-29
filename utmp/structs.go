package utmp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Values for Utmp.Type field
type Utype int16

// Type for ut_exit, below
const (
	Empty        Utype = iota // Record does not contain valid info (formerly known as UT_UNKNOWN on Linux)
	RunLevel           = iota // Change in system run-level (see init(8))
	BootTime           = iota // Time of system boot (in ut_tv)
	NewTime            = iota // Time after system clock change (in ut_tv)
	OldTime            = iota // Time before system clock change (in ut_tv)
	InitProcess        = iota // Process spawned by init(8)
	LoginProcess       = iota // Session leader process for user login
	UserProcess        = iota // Normal process
	DeadProcess        = iota // Terminated process
	Accounting         = iota // Not implemented

	LineSize = 32
	NameSize = 32
	HostSize = 256
)

// MarshalJSON correctly marshals the constants for the utmp "type" field
func (u Utype) MarshalJSON() ([]byte, error) {
	switch u {
	case Empty:
		return json.Marshal("Empty")
	case RunLevel:
		return json.Marshal("RunLevel")
	case BootTime:
		return json.Marshal("BootTime")
	case NewTime:
		return json.Marshal("NewTime")
	case OldTime:
		return json.Marshal("OldTime")
	case InitProcess:
		return json.Marshal("InitProcess")
	case LoginProcess:
		return json.Marshal("LoginProcess")
	case UserProcess:
		return json.Marshal("UserProcess")
	case DeadProcess:
		return json.Marshal("DeadProcess")
	case Accounting:
		return json.Marshal("Accounting")
	default:
		return json.Marshal("")
	}
}

type exit_status struct {
	Termination int16 `json:"termination"` // Process termination status
	Exit        int16 `json:"exit"`        // Process exit status
}

type TimeVal struct {
	Sec  int32 `json:"seconds"`
	Usec int32 `json:"microseconds"`
}

// Correctly marshal unix times to human readable timestamps
func (t TimeVal) MarshalJSON() ([]byte, error) {
	ts := time.Unix(int64(t.Sec), int64(t.Usec))
	return json.Marshal(ts.Format(time.RFC1123Z))
}

type Utmp struct {
	Type    Utype          `json:"type"` // Type of record
	_       int16          // padding because Go doesn't 4-byte align
	Pid     int32          `json:"pid"`     // PID of login process
	Device  [LineSize]byte `json:"device"`  // Device name of tty - "/dev/"
	Id      [4]byte        `json:"id"`      // Terminal name suffix or inittab(5) ID
	User    [NameSize]byte `json:"user"`    // Username
	Host    [HostSize]byte `json:"host"`    // Hostname for remote login or kernel version for run-level messages
	Exit    exit_status    `json:"exit"`    // Exit status of a process marked as DeadProcess; not used by Linux init(1)
	Session int32          `json:"session"` // Session ID (getsid(2)), used for windowing
	Time    TimeVal        `json:"time"`    // Time entry was made
	Addr    [4]int32       `json:"address"` // Internet address of remote host; IPv4 address uses just Addr[0]
	Unused  [20]byte       `json:"-"`       // Reserved for future use
}

// AddrToString only handles IPv4 strings right now
func AddrToString(a [4]int32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(a[0]), byte(a[0]>>8), byte(a[0]>>16), byte(a[0]>>24))
}

// Overrides the default Marshal method for Utmp
//
// MarshalJSON correctly interprets the address field and byte arrays into
// properly formatted strings stripped of empty padding
func (u Utmp) MarshalJSON() ([]byte, error) {
	utmp := map[string]interface{}{}
	utmp["type"] = u.Type
	utmp["pid"] = u.Pid
	utmp["device"] = string(bytes.Trim(u.Device[:], "\u0000"))
	utmp["id"] = string(bytes.Trim(u.Id[:], "\u0000"))
	utmp["user"] = string(bytes.Trim(u.User[:], "\u0000"))
	utmp["host"] = string(bytes.Trim(u.Host[:], "\u0000"))
	utmp["exit"] = u.Exit
	utmp["session"] = u.Session
	utmp["time"] = u.Time
	utmp["address"] = AddrToString(u.Addr)
	return json.Marshal(utmp)
}

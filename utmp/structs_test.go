package utmp

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestUtypeMarshalJSON(t *testing.T) {
	cases := []struct {
		in   Utype
		want []byte
	}{
		{Empty, []byte("\"Empty\"")},
		{RunLevel, []byte("\"RunLevel\"")},
		{BootTime, []byte("\"BootTime\"")},
		{NewTime, []byte("\"NewTime\"")},
		{OldTime, []byte("\"OldTime\"")},
		{InitProcess, []byte("\"InitProcess\"")},
		{LoginProcess, []byte("\"LoginProcess\"")},
		{UserProcess, []byte("\"UserProcess\"")},
		{DeadProcess, []byte("\"DeadProcess\"")},
		{Accounting, []byte("\"Accounting\"")},
		{Utype(99), []byte("\"\"")},
	}
	for _, c := range cases {
		got, _ := json.Marshal(c.in)
		if !bytes.Equal(got, c.want) {
			t.Errorf("%q.MarshalJSON() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestTimeValMarshalJSON(t *testing.T) {
	cases := []struct {
		in   TimeVal
		want []byte
	}{
		{TimeVal{1257894000, 0}, []byte("\"Tue, 10 Nov 2009 23:00:00 +0000\"")},
	}
	for _, c := range cases {
		got, _ := json.Marshal(c.in)
		if !bytes.Equal(got, c.want) {
			t.Errorf("%d.MarshalJSON() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAddrToString(t *testing.T) {
	cases := []struct {
		in   [4]int32
		want string
	}{
		{[4]int32{603979786, 0, 0, 0}, "10.0.0.36"},
		{[4]int32{0, 0, 0, 0}, "0.0.0.0"},
	}
	for _, c := range cases {
		got := AddrToString(c.in)
		if got != c.want {
			t.Errorf("AddrToString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestUtmpMarshalJSON(t *testing.T) {
	var device [LineSize]byte
	copy(device[:], "pts/4")
	var id [4]byte
	copy(id[:], "ts/4")
	var user [NameSize]byte
	copy(user[:], "ubuntu")
	var host [HostSize]byte
	copy(host[:], "ip-10-0-0-36.local")

	cases := []struct {
		in   Utmp
		want []byte
	}{
		{
			Utmp{
				Type:    LoginProcess,
				Pid:     int32(13772),
				Device:  device,
				Id:      id,
				User:    user,
				Host:    host,
				Exit:    exit_status{0, 0},
				Session: int32(0),
				Time:    TimeVal{1257894000, 0},
				Addr:    [4]int32{603979786, 0, 0, 0},
				Unused:  [20]byte{},
			},
			[]byte("{\"address\":\"10.0.0.36\",\"device\":\"pts/4\",\"exit\":{\"termination\":0,\"exit\":0},\"host\":\"ip-10-0-0-36.local\",\"id\":\"ts/4\",\"pid\":13772,\"session\":0,\"time\":\"Tue, 10 Nov 2009 23:00:00 +0000\",\"type\":\"LoginProcess\",\"user\":\"ubuntu\"}"),
		},
	}
	for _, c := range cases {
		got, _ := json.Marshal(c.in)
		if !bytes.Equal(got, c.want) {
			t.Errorf("%d.MarshalJSON() == %q, want %q", c.in, got, c.want)
		}
	}
}

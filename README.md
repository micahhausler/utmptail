# utmptail
`utmptail` reads in a `/var/log/{b,u,w}tmp` file and emit each line as JSON.
This is useful for sending \*tmp logs to a logging service that parses
key/value logs

## Usage
<pre>
utmptail reads in a login record formatted fileand emits each record as a JSON to
a new line. On most systems these logs are located at /var/log/{btmp, utmp, wtmp}.
See <b>utmp(5)</b> for more information on login record files.

Usage of ./utmptail:

  &lt;filename&gt;
        The file to read in. (default "/dev/stdin")
  -version
        print version and exit

</pre>

## Output
The JSON output contains the following keys:

```
{
    "session": 0,  				# Session ID (getsid(2)), used for windowing
    "user": "ubuntu",			# Username
    "exit": {
        "termination": 0,		# Process termination status
        "exit": 0				# Process exit status
    },
    "type": "UserProcess",		# Type of record
    "Addr": "192.168.53.103",	# Internet address of remote host
    "host": "ip-192-168-53-103.ec2.internal", 	# Hostname for remote login
    "pid": 24203,				# PID of login process
    "time": "Fri, 28 Aug 2015 16:39:27 +0000",	# Login time, RFC 1123 formatted
    "id": "ts/5", 				# Terminal name suffix, or inittab(5) ID
    "device": "pts/5"			# Device name of tty - "/dev/" 
}
```

## Arch support
Currenlty this only supports 64 bit systems as 32 bit systems have different struct attributes 

## Wishlist

- [ ] Add tests/CI
- [ ] Add GoDocs
- [ ] Actually be able to follow the specified file, rather than having to pipe in a tail like so
	
	```
	tail -F --bytes 384 /var/log/wtmp | utmptail
	```
- [ ] Add different output formats besides json
- [ ] Add docker builds to CI

## License
MIT License. See [License](/LICENSE) for full text

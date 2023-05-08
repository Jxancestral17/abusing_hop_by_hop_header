# Abusing Hop By Hop Headers Attack 

```bash
git clone https://github.com/Jxancestral17/abusing_hop_by_hop_header.git
```

```bash
go run main.go [URL]
```
Example

```bash
go run main.go https://google.com
```

- Send 3/4 requests for second 
- If it detects a possible poisoning it creates a log file
- If you want to add headers add them ./headers/headers.dat (currently taken from seclists)

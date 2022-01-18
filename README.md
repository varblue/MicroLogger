# MicroLogger
📄 A tiny logger using the builtin log library and a color library, written in [Go](https://go.dev/)

![Version](https://img.shields.io/badge/Version-v1.0.0-success?style=for-the-badge) [![Go Reference](https://pkg.go.dev/badge//gitlab.com/tesseractge/MicroLogger.svg)](https://pkg.go.dev//gitlab.com/tesseractge/MicroLogger)

## Installation
You can install MicroLogger in your project using `go get gitlab.com/tesseractge/MicroLogger`

## Usage
You can use the library by calling the Log function
```go
package main

import l "gitlab.com/tesseractge/MicroLogger"

func main() {
	l.Log(l.OK, "This is a log entry with the severity level 'OK'")
	l.Log(l.INFO, "This is a log entry with the severity level 'INFO'")
	l.Log(l.EXTENSION, "This is a log entry with the severity level 'EXTENSION'")
	l.Log(l.DEBUG, "This is a log entry with the severity level 'DEBUG'")
	l.Log(l.WARN, "This is a log entry with the severity level 'WARN'")
	l.Log(l.FAILED, "This is a log entry with the severity level 'FAILED'")
	l.Log(l.FATAL, "This is a log entry with the severity level 'FATAL'")
}
```

## Support
You can get support by [emailing me](mailto:hello@quinnlane.dev?subject=%E2%9D%93%20Support%20-%20MicroLogger) or contacting me through [Discord](https://discord.com/users/422415365742788618) (Please make sure you let me know why you are there, so I don't mark you as spam! Thanks!)

## Authors and acknowledgment
This project was created by [Quinn Lane](https://gitlab.com/Identithree)

## License
This project is licensed under the [BSD 3-Clause "New" or "Revised" License](https://tldrlegal.com/license/bsd-3-clause-license-(revised)) [(fulltext)](LICENSE)

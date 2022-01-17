/*
The MicroLogger package is a Go library intended to provide a
simple and uniform way of logging data to the console for
debugging or status updates. The style is based off of Linux
boot sequences.

There are multiple different log severity levels.
OK, INFO, WARN, FAILED, FATAL, EXTENSION, and DEBUG.

	log.Log(log.OK, "This is a log entry with the severity level 'OK'")
	log.Log(log.INFO, "This is a log entry with the severity level 'INFO'")
	log.Log(log.WARN, "This is a log entry with the severity level 'WARN'")
	log.Log(log.FAILED, "This is a log entry with the severity level 'FAILED'")
	log.Log(log.FATAL, "This is a log entry with the severity level 'FATAL'")
	log.Log(log.EXTENSION, "This is a log entry with the severity level 'EXTENSION'")
	log.Log(log.DEBUG, "This is a log entry with the severity level 'DEBUG'")

*/
package MicroLogger

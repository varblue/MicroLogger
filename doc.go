/*
The MicroLogger package is a Go library intended to provide a
simple and uniform way of logging data to the console for
debugging or status updates. The style is based off of Linux
boot sequences.

There are multiple different log severity levels.
OK, INFO, EXTENSION, DEBUG, WARN, FAILED, and FATAL.
These are in a type called LoggerLevel.

	Log(OK, "This is a log entry with the severity level 'OK'")
	Log(INFO, "This is a log entry with the severity level 'INFO'")
	Log(EXTENSION, "This is a log entry with the severity level 'EXTENSION'")
	Log(DEBUG, "This is a log entry with the severity level 'DEBUG'")
	Log(WARN, "This is a log entry with the severity level 'WARN'")
	Log(FAILED, "This is a log entry with the severity level 'FAILED'")
	Log(FATAL, "This is a log entry with the severity level 'FATAL'")

The color for the FATAL logger level can be found in FatalColor.
*/
package main

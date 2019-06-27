package pipeline


// ForceLog should rarely be used. It forceable logs an entry to the
// Windows Event Log (on Windows) or to the SysLog (on Linux)
func ForceLog(level LogLevel, msg string) {
	if sanitizer != nil {
		msg = sanitizer.SanitizeLogLine(msg)
	}
	forceLog(level, msg)
}

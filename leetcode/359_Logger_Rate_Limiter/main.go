package main

type Logger struct {
	messages map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{messages: make(map[string]int)}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	currentTimestamp, ok := this.messages[message]
	if ok && timestamp-currentTimestamp < 10 {
		return false
	}

	this.messages[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

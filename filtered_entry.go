package logrus

import (
	"fmt"
	"os"
)

type FilteredLogEntry struct {
	*Entry
	Level Level
}

func (entry *FilteredLogEntry) GetLevel() Level {
	return entry.Level
}

func (entry *FilteredLogEntry) Debug(args ...interface{}) {
	if entry.Level >= DebugLevel {
		entry.log(DebugLevel, fmt.Sprint(args...))
	}
}

func (entry *FilteredLogEntry) Print(args ...interface{}) {
	entry.Info(args...)
}

func (entry *FilteredLogEntry) Info(args ...interface{}) {
	if entry.Level >= InfoLevel {
		entry.log(InfoLevel, fmt.Sprint(args...))
	}
}

func (entry *FilteredLogEntry) Warn(args ...interface{}) {
	if entry.Level >= WarnLevel {
		entry.log(WarnLevel, fmt.Sprint(args...))
	}
}

func (entry *FilteredLogEntry) Warning(args ...interface{}) {
	entry.Warn(args...)
}

func (entry *FilteredLogEntry) Error(args ...interface{}) {
	if entry.Level >= ErrorLevel {
		entry.log(ErrorLevel, fmt.Sprint(args...))
	}
}

func (entry *FilteredLogEntry) Fatal(args ...interface{}) {
	if entry.Level >= FatalLevel {
		entry.log(FatalLevel, fmt.Sprint(args...))
	}
	os.Exit(1)
}

func (entry *FilteredLogEntry) Panic(args ...interface{}) {
	if entry.Level >= PanicLevel {
		entry.log(PanicLevel, fmt.Sprint(args...))
	}
	panic(fmt.Sprint(args...))
}

// Entry Printf family functions

func (entry *FilteredLogEntry) Debugf(format string, args ...interface{}) {
	if entry.Level >= DebugLevel {
		entry.Debug(fmt.Sprintf(format, args...))
	}
}

func (entry *FilteredLogEntry) Infof(format string, args ...interface{}) {
	if entry.Level >= InfoLevel {
		entry.Info(fmt.Sprintf(format, args...))
	}
}

func (entry *FilteredLogEntry) Printf(format string, args ...interface{}) {
	entry.Infof(format, args...)
}

func (entry *FilteredLogEntry) Warnf(format string, args ...interface{}) {
	if entry.Level >= WarnLevel {
		entry.Warn(fmt.Sprintf(format, args...))
	}
}

func (entry *FilteredLogEntry) Warningf(format string, args ...interface{}) {
	entry.Warnf(format, args...)
}

func (entry *FilteredLogEntry) Errorf(format string, args ...interface{}) {
	if entry.Level >= ErrorLevel {
		entry.Error(fmt.Sprintf(format, args...))
	}
}

func (entry *FilteredLogEntry) Fatalf(format string, args ...interface{}) {
	if entry.Level >= FatalLevel {
		entry.Fatal(fmt.Sprintf(format, args...))
	}
	os.Exit(1)
}

func (entry *FilteredLogEntry) Panicf(format string, args ...interface{}) {
	if entry.Level >= PanicLevel {
		entry.Panic(fmt.Sprintf(format, args...))
	}
}

// Entry Println family functions

func (entry *FilteredLogEntry) Debugln(args ...interface{}) {
	if entry.Level >= DebugLevel {
		entry.Debug(entry.sprintlnn(args...))
	}
}

func (entry *FilteredLogEntry) Infoln(args ...interface{}) {
	if entry.Level >= InfoLevel {
		entry.Info(entry.sprintlnn(args...))
	}
}

func (entry *FilteredLogEntry) Println(args ...interface{}) {
	entry.Infoln(args...)
}

func (entry *FilteredLogEntry) Warnln(args ...interface{}) {
	if entry.Level >= WarnLevel {
		entry.Warn(entry.sprintlnn(args...))
	}
}

func (entry *FilteredLogEntry) Warningln(args ...interface{}) {
	entry.Warnln(args...)
}

func (entry *FilteredLogEntry) Errorln(args ...interface{}) {
	if entry.Level >= ErrorLevel {
		entry.Error(entry.sprintlnn(args...))
	}
}

func (entry *FilteredLogEntry) Fatalln(args ...interface{}) {
	if entry.Level >= FatalLevel {
		entry.Fatal(entry.sprintlnn(args...))
	}
	os.Exit(1)
}

func (entry *FilteredLogEntry) Panicln(args ...interface{}) {
	if entry.Level >= PanicLevel {
		entry.Panic(entry.sprintlnn(args...))
	}
}

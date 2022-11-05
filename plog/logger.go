/*
 * @Author: Charley
 * @Date: 2022-11-03 10:39:41
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-03 10:58:05
 * @FilePath: /mospanel/plog/logger.go
 * @Description: panel log
 */
package plog

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Level    string `yaml:"level"`
	File     string `yaml:"file"`
	ErrFile  string `yaml:"err_file"`
	InfoFile string `yaml:"info_file"`
}

var (
	stdout = zapcore.Lock(os.Stdout)
	stderr = zapcore.Lock(os.Stderr)

	lvl = zap.NewAtomicLevelAt(zap.InfoLevel)
	l   = newLogger(lvl, stdout, stderr)
	s   = l.Sugar()
)

func NewLogger(lc *LogConfig) (*zap.Logger, error) {
	var lvl zapcore.Level
	if len(lc.Level) > 0 {
		var ok bool
		lvl, ok = parseLogLevel(lc.Level)
		if !ok {
			return nil, fmt.Errorf("invalid log level [%s]", lc.Level)
		}
	}

	var infoWriter zapcore.WriteSyncer
	var errWriter zapcore.WriteSyncer

	if lf := lc.File; len(lf) > 0 {
		f, _, err := zap.Open(lf)
		if err != nil {
			return nil, fmt.Errorf("open log file: %w", err)
		}
		fLocked := zapcore.Lock(f)
		infoWriter = fLocked
		errWriter = fLocked
	} else {
		if lf := lc.InfoFile; len(lf) > 0 {
			f, _, err := zap.Open(lf)
			if err != nil {
				return nil, fmt.Errorf("open info log file: %w", err)
			}
			infoWriter = zapcore.Lock(f)
		}
		if lf := lc.ErrFile; len(lf) > 0 {
			f, _, err := zap.Open(lf)
			if err != nil {
				return nil, fmt.Errorf("open err log file: %w", err)
			}
			errWriter = zapcore.Lock(f)
		}
	}

	return newLogger(lvl, infoWriter, errWriter), nil
}

// newLogger creates a new zap.Logger, by default the infoWriter and errWriter are
// stdout and stderr.
func newLogger(lvl zapcore.LevelEnabler, infoWriter, errWriter zapcore.WriteSyncer) *zap.Logger {
	errLvl := zap.LevelEnablerFunc(func(l2 zapcore.Level) bool {
		return lvl.Enabled(l2) && l2 >= zapcore.ErrorLevel
	})
	infoLvl := zap.LevelEnablerFunc(func(l2 zapcore.Level) bool {
		return lvl.Enabled(l2) && l2 < zapcore.ErrorLevel
	})

	if infoWriter == nil {
		infoWriter = stdout
	}
	if errWriter == nil {
		errWriter = stderr
	}

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(defaultEncoderConfig()), infoWriter, infoLvl),
		zapcore.NewCore(zapcore.NewConsoleEncoder(defaultEncoderConfig()), errWriter, errLvl),
	)
	return zap.New(core)
}

func L() *zap.Logger {
	return l
}

func SetLevel(l zapcore.Level) {
	lvl.SetLevel(l)
}

func S() *zap.SugaredLogger {
	return s
}

func defaultEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func parseLogLevel(s string) (zapcore.Level, bool) {
	switch s {
	case "debug":
		return zap.DebugLevel, true
	case "", "info":
		return zap.InfoLevel, true
	case "warn":
		return zap.WarnLevel, true
	case "error":
		return zap.ErrorLevel, true
	default:
		return 0, false
	}
}

package log

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func ExampleFormatter_Format_default() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		NoColors:        true,
		TimestampFormat: "-",
	})

	l.Debug("test1")
	l.Info("test2")
	l.Warn("test3")
	l.Error("test4")

	// Output:
	// - [DEBU] test1
	// - [INFO] test2
	// - [WARN] test3
	// - [ERRO] test4
}

func ExampleFormatter_Format_full_level() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		ShowFullLevel:   true,
	})

	l.Debug("test1")
	l.Info("test2")
	l.Warn("test3")
	l.Error("   test4")

	// Output:
	// - [DEBUG] test1
	// - [INFO] test2
	// - [WARNING] test3
	// - [ERROR]    test4
}
func ExampleFormatter_Format_show_keys() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		HideKeys:        false,
	})

	ll := l.WithField("category", "rest")

	l.Info("test1")
	ll.Info("test2")

	// Output:
	// - [INFO] test1
	// - [INFO] test2 [category:rest]
}

func ExampleFormatter_Format_hide_keys() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		HideKeys:        true,
	})

	ll := l.WithField("category", "rest")

	l.Info("test1")
	ll.Info("test2")

	// Output:
	// - [INFO] test1
	// - [INFO] test2 [rest]
}

func ExampleFormatter_Format_trim_message() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		TrimMessages:    true,
		NoColors:        true,
		TimestampFormat: "-",
	})

	l.Debug(" test1 ")
	l.Info("test2 ")
	l.Warn(" test3")
	l.Error("   test4   ")

	// Output:
	// - [DEBU] test1
	// - [INFO] test2
	// - [WARN] test3
	// - [ERRO] test4
}

func TestFormatter_Format_with_report_caller(t *testing.T) {
	output := bytes.NewBuffer([]byte{})

	l := logrus.New()
	l.SetOutput(output)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		NoColors:        true,
		TimestampFormat: "-",
	})
	l.SetReportCaller(true)

	l.Debug("test1")

	line, err := output.ReadString('\n')
	if err != nil {
		t.Errorf("Cannot read log output: %v", err)
	}

	expectedRegExp := "- \\[DEBU\\] test1 \\(.+\\.go:[0-9]+ .+\\)\n$"
	match, err := regexp.MatchString(
		expectedRegExp,
		line,
	)
	if err != nil {
		t.Errorf("Cannot check regexp: %v", err)
	} else if !match {
		t.Errorf(
			"logger.SetReportCaller(true) output doesn't match, expected: %s to find in: '%s'",
			expectedRegExp,
			line,
		)
	}
}

func TestFormatter_Format_with_report_caller_and_CallerFirst_true(t *testing.T) {
	output := bytes.NewBuffer([]byte{})

	l := logrus.New()
	l.SetOutput(output)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		CallerFirst:     true,
	})
	l.SetReportCaller(true)

	l.Debug("test1")

	line, err := output.ReadString('\n')
	if err != nil {
		t.Errorf("Cannot read log output: %v", err)
	}

	expectedRegExp := "- \\(.+\\.go:[0-9]+ .+\\) \\[DEBU\\] test1\n$"
	match, err := regexp.MatchString(
		expectedRegExp,
		line,
	)

	if err != nil {
		t.Errorf("Cannot check regexp: %v", err)
	} else if !match {
		t.Errorf(
			"logger.SetReportCaller(true) output doesn't match, expected: %s to find in: '%s'",
			expectedRegExp,
			line,
		)
	}
}

func TestFormatter_Format_with_report_caller_and_CustomCallerFormatter(t *testing.T) {
	output := bytes.NewBuffer([]byte{})

	l := logrus.New()
	l.SetOutput(output)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		CallerFirst:     true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
	})
	l.SetReportCaller(true)

	l.Debug("test1")

	line, err := output.ReadString('\n')
	if err != nil {
		t.Errorf("Cannot read log output: %v", err)
	}

	expectedRegExp := "- \\[.+\\.go:[0-9]+\\]\\[.+\\(\\)\\] \\[DEBU\\] test1\n$"
	match, err := regexp.MatchString(
		expectedRegExp,
		line,
	)
	if err != nil {
		t.Errorf("Cannot check regexp: %v", err)
	} else if !match {
		t.Errorf(
			"logger.SetReportCaller(true) output doesn't match, expected: %s to find in: '%s'",
			expectedRegExp,
			line,
		)
	}
}

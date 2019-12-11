package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

var defaultFormatter = &logrus.TextFormatter{}

// FilePathMap 用于将不同级别的日志映射到不同的日志文件
// 多个级别的日志可以使用同一个文件，但是多个文件不能用于同一个级别的日志
type FilePathMap map[logrus.Level]string

// WriterMap 用于将不同级别的日志映射到不同的Writer
// 多个级别的日志可以使用同一个Writer，但是多个Writer不能用于同一个级别的日志
type WriterMap map[logrus.Level]io.Writer


type FileHook struct {
	lock      *sync.Mutex
	formatter logrus.Formatter

	//按日志级别映射到不同的target
	filePaths FilePathMap
	writers   WriterMap
	levels    []logrus.Level

	defaultPath      string
	defaultWriter    io.Writer
	hasDefaultPath   bool
	hasDefaultWriter bool
}

// NewHook returns new File hook
// Output can be a string , io.Writer
// If using io.Writer , user is responsible for closing the used io.Writer
func NewFileHook(output interface{}, formatter logrus.Formatter) *FileHook {
	hook := &FileHook{
		lock: new(sync.Mutex),
	}

	hook.SetFormatter(formatter)

	switch output.(type) {
	case string:
		hook.SetDefaultPath(output.(string))
		break
	case io.Writer:
		hook.SetDefaultWriter(output.(io.Writer))
		break
	case FilePathMap:
		hook.filePaths = output.(FilePathMap)
		for level := range output.(FilePathMap) {
			hook.levels = append(hook.levels, level)
		}
		break
	case WriterMap:
		hook.writers = output.(WriterMap)
		for level := range output.(WriterMap) {
			hook.levels = append(hook.levels, level)
		}
	default:
		panic(fmt.Sprintf("Unsupported level map type:%v", reflect.TypeOf(output)))
	}
	return hook
}

// SetFormatter set formatter for this hook
// If using text formatter , the method will disable color output to make the log file more readable
func (fh *FileHook) SetFormatter(formatter logrus.Formatter) {
	fh.lock.Lock()
	defer fh.lock.Unlock()

	if formatter == nil {
		formatter = defaultFormatter
	} else {
		switch formatter.(type) {
		case *logrus.TextFormatter:
			textFormatter := formatter.(*logrus.TextFormatter)
			textFormatter.DisableColors = true
		}
	}
	fh.formatter = formatter
}

// SetDefaultPath set the file path for all the levels
func (fh *FileHook) SetDefaultPath(defaultPath string) {
	fh.lock.Lock()
	defer fh.lock.Unlock()
	fh.defaultPath = defaultPath
	fh.hasDefaultPath = true
}

// SetDefaultWriter set the default writer for all the levels
func (fh *FileHook) SetDefaultWriter(defaultWriter io.Writer) {
	fh.lock.Lock()
	defer fh.lock.Unlock()

	fh.defaultWriter = defaultWriter
	fh.hasDefaultWriter = true
}

// Fire() writes the log file to defined path or using the defined writer.
// User who run this function needs write permissions to the file or directory if the file does not yet exist
func (fh *FileHook) Fire(entry *logrus.Entry) error {
	fh.lock.Lock()
	defer fh.lock.Unlock()

	if fh.writers != nil || fh.hasDefaultWriter {
		return fh.ioWriter(entry)
	}

	if fh.filePaths != nil || fh.hasDefaultPath {
		return fh.fileWriter(entry)
	}
	return nil
}

// Write a log line to an io.Writer.
func (fh *FileHook) ioWriter(entry *logrus.Entry) error {
	var (
		writer io.Writer
		msg    []byte
		err    error
		ok     bool
	)
	if writer, ok = fh.writers[entry.Level]; !ok {
		if !fh.hasDefaultWriter {
			return nil
		}
		writer = fh.defaultWriter
	}

	msg, err = fh.formatter.Format(entry)
	if err != nil {
		log.Println("failed to generate string for entry :", err)
		return err
	}
	_, err = writer.Write(msg)
	return err
}

// Write a log line directly to a file
// Create dir and file if file does not yet exist
func (fh *FileHook) fileWriter(entry *logrus.Entry) error {
	var (
		fd   *os.File
		path string
		msg  []byte
		err  error
		ok   bool
	)
	if path, ok = fh.filePaths[entry.Level]; !ok {
		if !fh.hasDefaultPath {
			return nil
		}
		path = fh.defaultPath
	}

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println("failed create Dir :", dir, err)
		return err
	}

	fd, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("failed to open file :", path, err)
		return err
	}

	defer fd.Close()

	msg, err = fh.formatter.Format(entry)
	if err != nil {
		log.Println("failed to generate string entry :", err)
		return err
	}
	_, err = fd.Write(msg)
	return err
}

func (fh *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

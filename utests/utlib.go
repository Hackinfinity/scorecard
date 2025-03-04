// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package utests defines util fns for Scorecard unit testing.
package utests

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/ossf/scorecard/v2/checker"
)

// TestReturn encapsulates expected CheckResult return values.
type TestReturn struct {
	Error         error
	Score         int
	NumberOfWarn  int
	NumberOfInfo  int
	NumberOfDebug int
}

// TestDetailLogger implements `checker.DetailLogger`.
type TestDetailLogger struct {
	messages []checker.CheckDetail
}

// Info implements DetailLogger.Info.
func (l *TestDetailLogger) Info(desc string, args ...interface{}) {
	cd := checker.CheckDetail{Type: checker.DetailInfo, Msg: checker.LogMessage{Text: fmt.Sprintf(desc, args...)}}
	l.messages = append(l.messages, cd)
}

// Warn implements DetailLogger.Warn.
func (l *TestDetailLogger) Warn(desc string, args ...interface{}) {
	cd := checker.CheckDetail{Type: checker.DetailWarn, Msg: checker.LogMessage{Text: fmt.Sprintf(desc, args...)}}
	l.messages = append(l.messages, cd)
}

// Debug implements DetailLogger.Debug.
func (l *TestDetailLogger) Debug(desc string, args ...interface{}) {
	cd := checker.CheckDetail{Type: checker.DetailDebug, Msg: checker.LogMessage{Text: fmt.Sprintf(desc, args...)}}
	l.messages = append(l.messages, cd)
}

// UPGRADEv3: to rename.

// Info3 implements DetailLogger.Info3.
func (l *TestDetailLogger) Info3(msg *checker.LogMessage) {
	cd := checker.CheckDetail{
		Type: checker.DetailInfo,
		Msg:  *msg,
	}
	cd.Msg.Version = 3
	l.messages = append(l.messages, cd)
}

// Warn3 implements DetailLogger.Warn3.
func (l *TestDetailLogger) Warn3(msg *checker.LogMessage) {
	cd := checker.CheckDetail{
		Type: checker.DetailWarn,
		Msg:  *msg,
	}
	cd.Msg.Version = 3
	l.messages = append(l.messages, cd)
}

// Debug3 implements DetailLogger.Debug3.
func (l *TestDetailLogger) Debug3(msg *checker.LogMessage) {
	cd := checker.CheckDetail{
		Type: checker.DetailDebug,
		Msg:  *msg,
	}
	cd.Msg.Version = 3
	l.messages = append(l.messages, cd)
}

func getTestReturn(cr *checker.CheckResult, logger *TestDetailLogger) (*TestReturn, error) {
	ret := new(TestReturn)
	for _, v := range logger.messages {
		switch v.Type {
		default:
			// nolint: goerr113
			return nil, fmt.Errorf("invalid type %v", v.Type)
		case checker.DetailInfo:
			ret.NumberOfInfo++
		case checker.DetailDebug:
			ret.NumberOfDebug++
		case checker.DetailWarn:
			ret.NumberOfWarn++
		}
	}
	ret.Score = cr.Score
	ret.Error = cr.Error
	return ret, nil
}

func errCmp(e1, e2 error) bool {
	return errors.Is(e1, e2) || errors.Is(e2, e1)
}

// ValidateTestReturn validates expected TestReturn with actual checker.CheckResult values.
// nolint: thelper
func ValidateTestReturn(t *testing.T, name string, expected *TestReturn,
	actual *checker.CheckResult, logger *TestDetailLogger) bool {
	actualTestReturn, err := getTestReturn(actual, logger)
	if err != nil {
		panic(err)
	}
	if !cmp.Equal(*actualTestReturn, *expected, cmp.Comparer(errCmp)) {
		log.Println(cmp.Diff(*actualTestReturn, *expected))
		return false
	}
	return true
}

// ValidateLogMessage tests that at least one log message returns true for isExpectedMessage.
func ValidateLogMessage(isExpectedMessage func(checker.LogMessage, checker.DetailType) bool,
	dl *TestDetailLogger) bool {
	for _, message := range dl.messages {
		if isExpectedMessage(message.Msg, message.Type) {
			return true
		}
	}
	return false
}

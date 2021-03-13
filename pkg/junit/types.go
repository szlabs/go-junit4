// Copyright szlabs authors Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package junit

import (
	"encoding/xml"
	"time"
)

const (
	CasePassed  = "passed"
	CaseError   = "error"
	CaseFailure = "failure"
	CaseSkipped = "skipped"
)

// Follow the JUint report format document listed in the link below:
// https://llg.cubic.org/docs/junit/

// TestSuites defines the root suite group.
type TestSuites struct {
	XMLName   xml.Name `xml:"testsuites"`
	Name      string   `xml:"name,attr,omitempty"`
	Tests     int64    `xml:"tests,attr,omitempty"`
	Disabled  int64    `xml:"disabled,attr,omitempty"`
	Errors    int64    `xml:"errors,attr,omitempty"`
	Failures  int64    `xml:"failures,attr,omitempty"`
	Time      int64    `xml:"time,attr,omitempty"`
	TestSuite []*TestSuite
}

// TestSuite defines the test suite model.
type TestSuite struct {
	XMLName    xml.Name    `xml:"testsuite"`
	Name       string      `xml:"name,attr"`
	Tests      int64       `xml:"tests,attr"`
	Disabled   int64       `xml:"disabled,attr,omitempty"`
	Errors     int64       `xml:"errors,attr,omitempty"`
	Failures   int64       `xml:"failures,attr,omitempty"`
	Skipped    int64       `xml:"skipped,attr,omitempty"`
	Hostname   string      `xml:"hostname,attr,omitempty"`
	ID         int         `xml:"id,attr,omitempty"`
	Package    string      `xml:"package,attr,omitempty"`
	Time       int64       `xml:"time,attr,omitempty"`
	Timestamp  time.Time   `xml:"timestamp,attr,omitempty"`
	Properties *Properties `xml:",omitempty"`
	TestCase   []*TestCase
	SystemOut  string `xml:"system-out,omitempty"`
	SystemErr  string `xml:"system-err,omitempty"`
}

// TestCase defines the test case model.
type TestCase struct {
	XMLName    xml.Name   `xml:"testcase"`
	Name       string     `xml:"name,attr"`
	Assertions int64      `xml:"assertions,attr,omitempty"`
	ClassName  string     `xml:"classname,attr,omitempty"`
	Status     string     `xml:"status,attr,omitempty"`
	Time       int64      `xml:"time,attr,omitempty"`
	Skipped    *Skipped   `xml:"skipped,omitempty"`
	Error      *Error     `xml:"error,omitempty"`
	Failure    *Failure   `xml:"failure,omitempty"`
	SystemOut  *SystemOut `xml:"system-out,omitempty"`
	SystemErr  *SystemErr `xml:"system-err,omitempty"`
}

// Skipped element.
type Skipped struct {
	XMLName xml.Name `xml:"skipped"`
	Message string   `xml:"message,attr,omitempty"`
}

// Error element.
type Error struct {
	XMLName     xml.Name `xml:"error"`
	Message     string   `xml:"message,attr,omitempty"`
	Type        string   `xml:"type,attr,omitempty"`
	Description string   `xml:",cdata"`
}

// Failure element.
type Failure struct {
	XMLName     xml.Name `xml:"failure"`
	Message     string   `xml:"message,attr,omitempty"`
	Type        string   `xml:"type,attr,omitempty"`
	Description string   `xml:",chardata"`
}

// Properties list element.
type Properties struct {
	XMLName    xml.Name `xml:"properties"`
	Properties []*Property
}

// Property element.
type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

// SystemOut element.
type SystemOut struct {
	XMLName xml.Name `xml:"system-out"`
	Data    string   `xml:",cdata"`
}

// SystemErr element.
type SystemErr struct {
	XMLName xml.Name `xml:"system-err"`
	Data    string   `xml:",cdata"`
}

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
	"io/ioutil"

	"github.com/pkg/errors"
)

// XMLWriter defines junit report exporting methods.
type XMLWriter interface {
	XML() ([]byte, error)
	XML2File(fp string) error
}

// TestSuitesSetter defines property setting methods for the test suites.
type TestSuitesSetter interface {
	SetTests(number int64) TestSuitesSetter
	SetDisabled(number int64) TestSuitesSetter
	SetErrors(number int64) TestSuitesSetter
	SetFailures(number int64) TestSuitesSetter
	// In seconds
	SetTime(number int64) TestSuitesSetter
	Append(suite *TestSuite) TestSuitesSetter
	Complete() XMLWriter
}

// XML writes junit data.
func (ts *TestSuites) XML() ([]byte, error) {
	return xml.MarshalIndent(ts, "  ", "    ")
}

// XML2File writes junit data to the specified file.
func (ts *TestSuites) XML2File(fp string) error {
	data, err := ts.XML()
	if err != nil {
		return errors.Wrap(err, "XML2File:TestSuites")
	}

	if err := ioutil.WriteFile(fp, data, 0644); err != nil {
		return errors.Wrap(err, "XML2File:TestSuites")
	}

	return nil
}

// SetTests sets the number of test cases in the whole test suites.
func (ts *TestSuites) SetTests(number int64) TestSuitesSetter {
	ts.Tests = number
	return ts
}

// SetDisabled sets the number of the disabled test cases in the whole test suites.
func (ts *TestSuites) SetDisabled(number int64) TestSuitesSetter {
	ts.Disabled = number
	return ts
}

// SetErrors sets the number of errored test cases in the whole test suites.
func (ts *TestSuites) SetErrors(number int64) TestSuitesSetter {
	ts.Errors = number
	return ts
}

// SetFailures sets the number of failed test cases in the whole test suites.
func (ts *TestSuites) SetFailures(number int64) TestSuitesSetter {
	ts.Failures = number
	return ts
}

// SetTime sets running duration of the test suites.
func (ts *TestSuites) SetTime(number int64) TestSuitesSetter {
	ts.Time = number
	return ts
}

// Append test suite to the test suites.
func (ts *TestSuites) Append(suite *TestSuite) TestSuitesSetter {
	if suite != nil {
		if ts.TestSuite == nil {
			ts.TestSuite = make([]*TestSuite, 0)
		}

		ts.TestSuite = append(ts.TestSuite, suite)
		// Update related metrics here
		ts.Tests += suite.Tests
		ts.Failures += suite.Failures
		ts.Errors += suite.Errors
		ts.Disabled += suite.Disabled
		ts.Time += suite.Time
	}

	return ts
}

// Complete the build process and return the xml writer.
func (ts *TestSuites) Complete() XMLWriter {
	return ts
}

// NewSuites is constructor method
func NewSuites(name string) TestSuitesSetter {
	return &TestSuites{
		Name: name,
	}
}

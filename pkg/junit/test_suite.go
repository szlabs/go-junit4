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

import "time"

// TestSuiteSetter defines methods to operator test suite element.
type TestSuiteSetter interface {
	SetTests(number int64) TestSuiteSetter
	SetDisabled(number int64) TestSuiteSetter
	SetErrors(number int64) TestSuiteSetter
	SetFailures(number int64) TestSuiteSetter
	SetSkipped(number int64) TestSuiteSetter
	SetHostname(host string) TestSuiteSetter
	SetPackage(p string) TestSuiteSetter
	SetTimestamp(t time.Time) TestSuiteSetter
	SetSystemOut(out string) TestSuiteSetter
	SetSystemErr(err string) TestSuiteSetter
	AddProperty(p *Property) TestSuiteSetter
	Append(tc *TestCase) TestSuiteSetter
	// In seconds
	SetTime(number int64) TestSuiteSetter
	TestSuite() *TestSuite
}

// NewSuite is constructor method
func NewSuite(id int, name string) TestSuiteSetter {
	return &TestSuite{
		Name: name,
		ID:   id,
	}
}

// SetTests sets the number of tests in the test suite.
func (ts *TestSuite) SetTests(number int64) TestSuiteSetter {
	ts.Tests = number
	return ts
}

// SetDisabled sets the number of disabled tests in the test suite.
func (ts *TestSuite) SetDisabled(number int64) TestSuiteSetter {
	ts.Disabled = number
	return ts
}

// SetErrors sets the number of errored tests in the test suite.
func (ts *TestSuite) SetErrors(number int64) TestSuiteSetter {
	ts.Errors = number
	return ts
}

// SetFailures sets the number of failed tests in the test suite.
func (ts *TestSuite) SetFailures(number int64) TestSuiteSetter {
	ts.Failures = number
	return ts
}

// SetSkipped sets the number of skipped tests in the test suite.
func (ts *TestSuite) SetSkipped(number int64) TestSuiteSetter {
	ts.Skipped = number
	return ts
}

// SetHostname sets the hostname of the test suite.
func (ts *TestSuite) SetHostname(host string) TestSuiteSetter {
	ts.Hostname = host
	return ts
}

// SetPackage sets the package of the test suite.
func (ts *TestSuite) SetPackage(p string) TestSuiteSetter {
	ts.Package = p
	return ts
}

// SetTimestamp sets the timestamp of the test suite.
func (ts *TestSuite) SetTimestamp(t time.Time) TestSuiteSetter {
	ts.Timestamp = t
	return ts
}

// SetTime sets the running duration of the test suite.
func (ts *TestSuite) SetTime(number int64) TestSuiteSetter {
	ts.Time = number
	return ts
}

// SetSystemOut sets the system out element of the test suite.
func (ts *TestSuite) SetSystemOut(out string) TestSuiteSetter {
	ts.SystemOut = out
	return ts
}

// SetSystemErr sets the system error element of the test suite.
func (ts *TestSuite) SetSystemErr(err string) TestSuiteSetter {
	ts.SystemErr = err
	return ts
}

// AddProperty adds property to the test suite.
func (ts *TestSuite) AddProperty(p *Property) TestSuiteSetter {
	if p != nil {
		if ts.Properties == nil {
			ts.Properties = &Properties{
				Properties: make([]*Property, 0),
			}
		}

		ts.Properties.Properties = append(ts.Properties.Properties, p)
	}

	return ts
}

// Append child test case to the test suite.
func (ts *TestSuite) Append(tc *TestCase) TestSuiteSetter {
	if tc != nil {
		if ts.TestCase == nil {
			ts.TestCase = make([]*TestCase, 0)
		}

		ts.TestCase = append(ts.TestCase, tc)
		// Update related metrics here
		ts.Time += tc.Time
		ts.Tests++
		if tc.Skipped != nil {
			ts.Skipped++
		}
		if tc.Error != nil {
			ts.Errors++
		}
		if tc.Failure != nil {
			ts.Failures++
		}

	}

	return ts
}

// TestSuite returns reference point.
func (ts *TestSuite) TestSuite() *TestSuite {
	return ts
}

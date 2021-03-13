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

// TestCaseSetter defines methods to operator test case element.
type TestCaseSetter interface {
	SetAssertions(number int64) TestCaseSetter
	SetClassName(c string) TestCaseSetter
	SetStatus(st string) TestCaseSetter
	SetError(msg, typ, desc string) TestCaseSetter
	SetFailure(msg, typ, desc string) TestCaseSetter
	SetSkipped(msg string) TestCaseSetter
	SetSystemOut(out string) TestCaseSetter
	SetSystemErr(err string) TestCaseSetter
	// In seconds
	SetTime(number int64) TestCaseSetter
	TestCase() *TestCase
}

// SetAssertions sets the number of assertions in the test case.
func (tc *TestCase) SetAssertions(number int64) TestCaseSetter {
	tc.Assertions = number
	return tc
}

// SetClassName sets the testing class name of the test case.
func (tc *TestCase) SetClassName(c string) TestCaseSetter {
	tc.ClassName = c
	return tc
}

// SetStatus sets the status info of the test case.
func (tc *TestCase) SetStatus(st string) TestCaseSetter {
	tc.Status = st
	return tc
}

// SetError sets the error info of the test case.
func (tc *TestCase) SetError(msg, typ, desc string) TestCaseSetter {
	e := &Error{
		Message:     msg,
		Type:        typ,
		Description: desc,
	}
	tc.Error = e
	return tc
}

// SetFailure sets the failure info of the test case.
func (tc *TestCase) SetFailure(msg, typ, desc string) TestCaseSetter {
	f := &Failure{
		Message:     msg,
		Type:        typ,
		Description: desc,
	}
	tc.Failure = f
	return tc
}

// SetSkipped sets skipped message of the test case.
func (tc *TestCase) SetSkipped(msg string) TestCaseSetter {
	s := &Skipped{
		Message: msg,
	}

	tc.Skipped = s
	return tc
}

// SetSystemOut sets system out of the test case.
func (tc *TestCase) SetSystemOut(out string) TestCaseSetter {
	tc.SystemOut = &SystemOut{
		Data: out,
	}
	return tc
}

// SetSystemErr sets system error of the test case.
func (tc *TestCase) SetSystemErr(err string) TestCaseSetter {
	tc.SystemErr = &SystemErr{
		Data: err,
	}
	return tc
}

// SetTime sets running duration of the test case.
func (tc *TestCase) SetTime(number int64) TestCaseSetter {
	tc.Time = number
	return tc
}

// TestCase returns test case reference pointer.
func (tc *TestCase) TestCase() *TestCase {
	return tc
}

// NewCase is constructor of test case
func NewCase(name string) TestCaseSetter {
	return &TestCase{
		Name: name,
	}
}

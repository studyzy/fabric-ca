/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package assert

import (
	http "github.com/studyzy/net/http"
	url "net/url"
	time "time"
)

// Condition uses a Comparison to assert a complex condition.
func (a *Assertions) Condition(comp Comparison, msgAndArgs ...interface{}) bool {
	return Condition(a.t, comp, msgAndArgs...)
}

// Conditionf uses a Comparison to assert a complex condition.
func (a *Assertions) Conditionf(comp Comparison, msg string, args ...interface{}) bool {
	return Conditionf(a.t, comp, msg, args...)
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    a.Contains("Hello World", "World")
//    a.Contains(["Hello", "World"], "World")
//    a.Contains({"Hello": "World"}, "Hello")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool {
	return Contains(a.t, s, contains, msgAndArgs...)
}

// Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    a.Containsf("Hello World", "World", "error message %s", "formatted")
//    a.Containsf(["Hello", "World"], "World", "error message %s", "formatted")
//    a.Containsf({"Hello": "World"}, "Hello", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool {
	return Containsf(a.t, s, contains, msg, args...)
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  a.Empty(obj)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) bool {
	return Empty(a.t, object, msgAndArgs...)
}

// Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  a.Emptyf(obj, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Emptyf(object interface{}, msg string, args ...interface{}) bool {
	return Emptyf(a.t, object, msg, args...)
}

// Equal asserts that two objects are equal.
//
//    a.Equal(123, 123)
//
// Returns whether the assertion was successful (true) or not (false).
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return Equal(a.t, expected, actual, msgAndArgs...)
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   a.EqualError(err,  expectedErrorString)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) bool {
	return EqualError(a.t, theError, errString, msgAndArgs...)
}

// EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   a.EqualErrorf(err,  expectedErrorString, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...interface{}) bool {
	return EqualErrorf(a.t, theError, errString, msg, args...)
}

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    a.EqualValues(uint32(123), int32(123))
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return EqualValues(a.t, expected, actual, msgAndArgs...)
}

// EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    a.EqualValuesf(uint32(123, "error message %s", "formatted"), int32(123))
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return EqualValuesf(a.t, expected, actual, msg, args...)
}

// Equalf asserts that two objects are equal.
//
//    a.Equalf(123, 123, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func (a *Assertions) Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return Equalf(a.t, expected, actual, msg, args...)
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.Error(err) {
// 	   assert.Equal(t, expectedError, err)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Error(err error, msgAndArgs ...interface{}) bool {
	return Error(a.t, err, msgAndArgs...)
}

// Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.Errorf(err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Errorf(err error, msg string, args ...interface{}) bool {
	return Errorf(a.t, err, msg, args...)
}

// Exactly asserts that two objects are equal is value and type.
//
//    a.Exactly(int32(123), int64(123))
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return Exactly(a.t, expected, actual, msgAndArgs...)
}

// Exactlyf asserts that two objects are equal is value and type.
//
//    a.Exactlyf(int32(123, "error message %s", "formatted"), int64(123))
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return Exactlyf(a.t, expected, actual, msg, args...)
}

// Fail reports a failure through
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) bool {
	return Fail(a.t, failureMessage, msgAndArgs...)
}

// FailNow fails test
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) bool {
	return FailNow(a.t, failureMessage, msgAndArgs...)
}

// FailNowf fails test
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...interface{}) bool {
	return FailNowf(a.t, failureMessage, msg, args...)
}

// Failf reports a failure through
func (a *Assertions) Failf(failureMessage string, msg string, args ...interface{}) bool {
	return Failf(a.t, failureMessage, msg, args...)
}

// False asserts that the specified value is false.
//
//    a.False(myBool)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) False(value bool, msgAndArgs ...interface{}) bool {
	return False(a.t, value, msgAndArgs...)
}

// Falsef asserts that the specified value is false.
//
//    a.Falsef(myBool, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Falsef(value bool, msg string, args ...interface{}) bool {
	return Falsef(a.t, value, msg, args...)
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  a.HTTPBodyContains(myHandler, "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}) bool {
	return HTTPBodyContains(a.t, handler, method, url, values, str)
}

// HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//  a.HTTPBodyContainsf(myHandler, "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}) bool {
	return HTTPBodyContainsf(a.t, handler, method, url, values, str)
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  a.HTTPBodyNotContains(myHandler, "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}) bool {
	return HTTPBodyNotContains(a.t, handler, method, url, values, str)
}

// HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//  a.HTTPBodyNotContainsf(myHandler, "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}) bool {
	return HTTPBodyNotContainsf(a.t, handler, method, url, values, str)
}

// HTTPError asserts that a specified handler returns an error status code.
//
//  a.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values) bool {
	return HTTPError(a.t, handler, method, url, values)
}

// HTTPErrorf asserts that a specified handler returns an error status code.
//
//  a.HTTPErrorf(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values) bool {
	return HTTPErrorf(a.t, handler, method, url, values)
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  a.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values) bool {
	return HTTPRedirect(a.t, handler, method, url, values)
}

// HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//  a.HTTPRedirectf(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values) bool {
	return HTTPRedirectf(a.t, handler, method, url, values)
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//  a.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values) bool {
	return HTTPSuccess(a.t, handler, method, url, values)
}

// HTTPSuccessf asserts that a specified handler returns a success status code.
//
//  a.HTTPSuccessf(myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values) bool {
	return HTTPSuccessf(a.t, handler, method, url, values)
}

// Implements asserts that an object is implemented by the specified interface.
//
//    a.Implements((*MyInterface)(nil), new(MyObject))
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	return Implements(a.t, interfaceObject, object, msgAndArgs...)
}

// Implementsf asserts that an object is implemented by the specified interface.
//
//    a.Implementsf((*MyInterface, "error message %s", "formatted")(nil), new(MyObject))
func (a *Assertions) Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool {
	return Implementsf(a.t, interfaceObject, object, msg, args...)
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 a.InDelta(math.Pi, (22 / 7.0), 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return InDelta(a.t, expected, actual, delta, msgAndArgs...)
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return InDeltaSlice(a.t, expected, actual, delta, msgAndArgs...)
}

// InDeltaSlicef is the same as InDelta, except it compares two slices.
func (a *Assertions) InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	return InDeltaSlicef(a.t, expected, actual, delta, msg, args...)
}

// InDeltaf asserts that the two numerals are within delta of each other.
//
// 	 a.InDeltaf(math.Pi, (22 / 7.0, "error message %s", "formatted"), 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	return InDeltaf(a.t, expected, actual, delta, msg, args...)
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	return InEpsilon(a.t, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	return InEpsilonSlice(a.t, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.
func (a *Assertions) InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {
	return InEpsilonSlicef(a.t, expected, actual, epsilon, msg, args...)
}

// InEpsilonf asserts that expected and actual have a relative error less than epsilon
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {
	return InEpsilonf(a.t, expected, actual, epsilon, msg, args...)
}

// IsType asserts that the specified objects are of the same type.
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	return IsType(a.t, expectedType, object, msgAndArgs...)
}

// IsTypef asserts that the specified objects are of the same type.
func (a *Assertions) IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) bool {
	return IsTypef(a.t, expectedType, object, msg, args...)
}

// JSONEq asserts that two JSON strings are equivalent.
//
//  a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) bool {
	return JSONEq(a.t, expected, actual, msgAndArgs...)
}

// JSONEqf asserts that two JSON strings are equivalent.
//
//  a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...interface{}) bool {
	return JSONEqf(a.t, expected, actual, msg, args...)
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    a.Len(mySlice, 3)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) bool {
	return Len(a.t, object, length, msgAndArgs...)
}

// Lenf asserts that the specified object has specific length.
// Lenf also fails if the object has a type that len() not accept.
//
//    a.Lenf(mySlice, 3, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Lenf(object interface{}, length int, msg string, args ...interface{}) bool {
	return Lenf(a.t, object, length, msg, args...)
}

// Nil asserts that the specified object is nil.
//
//    a.Nil(err)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) bool {
	return Nil(a.t, object, msgAndArgs...)
}

// Nilf asserts that the specified object is nil.
//
//    a.Nilf(err, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Nilf(object interface{}, msg string, args ...interface{}) bool {
	return Nilf(a.t, object, msg, args...)
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.NoError(err) {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) bool {
	return NoError(a.t, err, msgAndArgs...)
}

// NoErrorf asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.NoErrorf(err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NoErrorf(err error, msg string, args ...interface{}) bool {
	return NoErrorf(a.t, err, msg, args...)
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    a.NotContains("Hello World", "Earth")
//    a.NotContains(["Hello", "World"], "Earth")
//    a.NotContains({"Hello": "World"}, "Earth")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool {
	return NotContains(a.t, s, contains, msgAndArgs...)
}

// NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    a.NotContainsf("Hello World", "Earth", "error message %s", "formatted")
//    a.NotContainsf(["Hello", "World"], "Earth", "error message %s", "formatted")
//    a.NotContainsf({"Hello": "World"}, "Earth", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool {
	return NotContainsf(a.t, s, contains, msg, args...)
}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if a.NotEmpty(obj) {
//    assert.Equal(t, "two", obj[1])
//  }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) bool {
	return NotEmpty(a.t, object, msgAndArgs...)
}

// NotEmptyf asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if a.NotEmptyf(obj, "error message %s", "formatted") {
//    assert.Equal(t, "two", obj[1])
//  }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotEmptyf(object interface{}, msg string, args ...interface{}) bool {
	return NotEmptyf(a.t, object, msg, args...)
}

// NotEqual asserts that the specified values are NOT equal.
//
//    a.NotEqual(obj1, obj2)
//
// Returns whether the assertion was successful (true) or not (false).
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return NotEqual(a.t, expected, actual, msgAndArgs...)
}

// NotEqualf asserts that the specified values are NOT equal.
//
//    a.NotEqualf(obj1, obj2, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (a *Assertions) NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return NotEqualf(a.t, expected, actual, msg, args...)
}

// NotNil asserts that the specified object is not nil.
//
//    a.NotNil(err)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) bool {
	return NotNil(a.t, object, msgAndArgs...)
}

// NotNilf asserts that the specified object is not nil.
//
//    a.NotNilf(err, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotNilf(object interface{}, msg string, args ...interface{}) bool {
	return NotNilf(a.t, object, msg, args...)
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   a.NotPanics(func(){ RemainCalm() })
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) bool {
	return NotPanics(a.t, f, msgAndArgs...)
}

// NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   a.NotPanicsf(func(){ RemainCalm() }, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotPanicsf(f PanicTestFunc, msg string, args ...interface{}) bool {
	return NotPanicsf(a.t, f, msg, args...)
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//  a.NotRegexp("^start", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	return NotRegexp(a.t, rx, str, msgAndArgs...)
}

// NotRegexpf asserts that a specified regexp does not match a string.
//
//  a.NotRegexpf(regexp.MustCompile("starts", "error message %s", "formatted"), "it's starting")
//  a.NotRegexpf("^start", "it's not starting", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool {
	return NotRegexpf(a.t, rx, str, msg, args...)
}

// NotSubset asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    a.NotSubset([1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) bool {
	return NotSubset(a.t, list, subset, msgAndArgs...)
}

// NotSubsetf asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    a.NotSubsetf([1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool {
	return NotSubsetf(a.t, list, subset, msg, args...)
}

// NotZero asserts that i is not the zero value for its type and returns the truth.
func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) bool {
	return NotZero(a.t, i, msgAndArgs...)
}

// NotZerof asserts that i is not the zero value for its type and returns the truth.
func (a *Assertions) NotZerof(i interface{}, msg string, args ...interface{}) bool {
	return NotZerof(a.t, i, msg, args...)
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   a.Panics(func(){ GoCrazy() })
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Panics(f PanicTestFunc, msgAndArgs ...interface{}) bool {
	return Panics(a.t, f, msgAndArgs...)
}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   a.PanicsWithValue("crazy error", func(){ GoCrazy() })
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) PanicsWithValue(expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) bool {
	return PanicsWithValue(a.t, expected, f, msgAndArgs...)
}

// PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   a.PanicsWithValuef("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) PanicsWithValuef(expected interface{}, f PanicTestFunc, msg string, args ...interface{}) bool {
	return PanicsWithValuef(a.t, expected, f, msg, args...)
}

// Panicsf asserts that the code inside the specified PanicTestFunc panics.
//
//   a.Panicsf(func(){ GoCrazy() }, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Panicsf(f PanicTestFunc, msg string, args ...interface{}) bool {
	return Panicsf(a.t, f, msg, args...)
}

// Regexp asserts that a specified regexp matches a string.
//
//  a.Regexp(regexp.MustCompile("start"), "it's starting")
//  a.Regexp("start...$", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	return Regexp(a.t, rx, str, msgAndArgs...)
}

// Regexpf asserts that a specified regexp matches a string.
//
//  a.Regexpf(regexp.MustCompile("start", "error message %s", "formatted"), "it's starting")
//  a.Regexpf("start...$", "it's not starting", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool {
	return Regexpf(a.t, rx, str, msg, args...)
}

// Subset asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    a.Subset([1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) bool {
	return Subset(a.t, list, subset, msgAndArgs...)
}

// Subsetf asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    a.Subsetf([1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool {
	return Subsetf(a.t, list, subset, msg, args...)
}

// True asserts that the specified value is true.
//
//    a.True(myBool)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) True(value bool, msgAndArgs ...interface{}) bool {
	return True(a.t, value, msgAndArgs...)
}

// Truef asserts that the specified value is true.
//
//    a.Truef(myBool, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Truef(value bool, msg string, args ...interface{}) bool {
	return Truef(a.t, value, msg, args...)
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   a.WithinDuration(time.Now(), time.Now(), 10*time.Second)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	return WithinDuration(a.t, expected, actual, delta, msgAndArgs...)
}

// WithinDurationf asserts that the two times are within duration delta of each other.
//
//   a.WithinDurationf(time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) bool {
	return WithinDurationf(a.t, expected, actual, delta, msg, args...)
}

// Zero asserts that i is the zero value for its type and returns the truth.
func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) bool {
	return Zero(a.t, i, msgAndArgs...)
}

// Zerof asserts that i is the zero value for its type and returns the truth.
func (a *Assertions) Zerof(i interface{}, msg string, args ...interface{}) bool {
	return Zerof(a.t, i, msg, args...)
}

package script

//Error is a Go call error.
type Error struct {
	Q Ctx
	String
	Code Int
}

//NeedToLogin runs 'f' if the user is Unauthorized.
func (e Error) NeedToLogin(f func()) {
	var q = e.Q
	q.Javascript(`if (%v == 401) {`, e.Code)
	f()
	q.Javascript(`}`)
}

//Connection runs the `f` script if the error is a connection error.
func (e Error) Connection(f func()) {
	var q = e.Q
	q.Javascript(`if (rpc_result.status != 500) {`)
	f()
	q.Javascript(`}`)
}

//Go runs the `f` script if the error is a Go error.
func (e Error) Go(f func()) {
	var q = e.Q
	q.Javascript(`if (rpc_result.status == 500) {`)
	f()
	q.Javascript(`}`)
}

//Promise represents a future action that can either succeed or fail.
type Promise struct {
	Native
	q Ctx
}

//Raw returns the raw JS promise.
func (promise Promise) Raw() string {
	return promise.LanguageType().Raw()
}

//Wait waits for the promise to complete and returns the resulting value.
func (promise Promise) Wait() Dynamic {
	var a = Unique()
	promise.q.Javascript(`let %v = await %v;`, a, promise)
	return promise.q.Value(`%v`, a).Dynamic()
}

//Then executes the provided function when the promise succeeds.
func (promise Promise) Then(f func(value Dynamic)) Promise {
	promise.q.Javascript(promise.Raw() + ` = ` + promise.Raw() + ".then(async function(promise_result) {")
	if f != nil {
		f(promise.q.Value("promise_result").Dynamic())
	}
	promise.q.Javascript("});")
	return promise
}

//Catch executes the provided function when the promise fails.
func (promise Promise) Catch(f func(err Error)) Promise {
	promise.q.Javascript(promise.Raw() + ".catch(async function(rpc_result) {")
	if f != nil {
		f(Error{promise.q, promise.q.Value("rpc_result.response").String(), promise.q.Value("rpc_result.status").Int()})
	}
	promise.q.Javascript("});")
	return promise
}

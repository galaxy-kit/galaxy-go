package internal

// SafeRet 安全调用结果
type SafeRet struct {
	Err error // error
	Ret any   // 返回值
}

// SafeCall 安全调用
type SafeCall interface {
	// SafeCall 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，会阻塞。
	//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
	SafeCall(segment func() SafeRet) <-chan SafeRet

	// SafeCallNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，不会阻塞。
	//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
	SafeCallNoWait(segment func() SafeRet) <-chan SafeRet

	// SafeCallNoRet 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时会阻塞。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，会阻塞。
	SafeCallNoRet(segment func())

	// SafeCallNoRetNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时不会阻塞。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，不会阻塞。
	SafeCallNoRetNoWait(segment func())
}

// Callee 安全调用接收者
type Callee interface {
	// PushCall 将代码片段压入接收者的任务处理流水线，串行化的进行调用。
	PushCall(segment func())
}

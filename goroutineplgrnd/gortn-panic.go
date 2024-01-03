package gortntry

/** if flag is 1 then main panic recovered but new goroutine started in main can't be handled
If flag is 3 then goroutine started in TryPanic can handle the panic but main gorutine panic not recover.
for any other value both main goroutine and goroutine started in TryPanic not recovered  */

func panicFun(flag int) {
	if flag == 3 {
		defer func() {
			recover()
		}()
	}
	panic("panic in goroutine")
}

func TryPanic(flag int) {
	if flag == 1 {
		defer func() {
			recover()
		}()
	}
	go panicFun(flag)
	panic("panic in main")
}

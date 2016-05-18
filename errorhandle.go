//when program exit, print out the stack

func foo() {
	defer func() {
		if err := recover; err != nil {
			debug.PrintStack()
			log.Error("Error:", err)
		}
	}()
		//other code
}

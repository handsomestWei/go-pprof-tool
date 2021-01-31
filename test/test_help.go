package test

func BlockDefault() {
	var c chan int
	for {
		select {
		case <-c:
			// 不会进入该分支
		default:
			// 会持续占用cpu
		}
	}
}

func BlockNoneDefault() {
	var c chan int
	for {
		select {
		case <-c:
			// 不会进入该分支
		//default:
		}
	}
}
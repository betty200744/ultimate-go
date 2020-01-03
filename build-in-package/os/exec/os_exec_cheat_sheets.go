package main

import "os/exec"

func main() {
	/*
		➜  github open www.baidu.com
		The file /Users/betty/project/github/www.baidu.com does not exist.
		Perhaps you meant 'http://www.baidu.com'?
		➜  github open http://www.baidu.com

	*/
	exec.Command("open", "http://www.baidu.com").Output()
}

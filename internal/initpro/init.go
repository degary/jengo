package initpro

import (
	"fmt"
	"os"
	"os/exec"
)

func Makedir(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	//fmt.Printf("创建目录%s\n",dir)
	return nil
}

func StartDocker() {
	fmt.Println("start docker")
	cmd := exec.Command("/bin/bash", "-c", "docker ps")
	out, err := cmd.Output()
	if err != nil {
		fmt.Errorf("出错拉: %s", err)
	}
	fmt.Println(string(out))
}

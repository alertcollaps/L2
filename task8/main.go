package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"strconv"
	"strings"
)

func exists(path string) (bool, error) {
	val, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return val.IsDir(), nil
}

type currentPath struct {
	path string
}

func (c *currentPath) getPath() string {
	return c.path
}

func (c *currentPath) decreasePath() {
	cnt := strings.Count(c.path, "\\")
	if cnt > 1 {
		c.setPath(c.path[0:strings.LastIndex(c.path, "\\")])
		fmt.Println("Path now:", c.getPath())
		return
	}
	fmt.Println("Decrease path impossible")
	fmt.Println("Path now:", c.getPath())
}

func (c *currentPath) setPath(pth string) {
	//Проверка на существующий каталог
	if val, _ := exists(pth); val {
		c.path = pth
		return
	}
	if val, _ := exists(c.getPath() + "\\" + pth); val {
		c.path = c.getPath() + "\\" + pth
		return
	}
	//Если нет каталога, устанавливаем текущий
	fmt.Println("Path not found. Going to default path")
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.path = currDir
	defer fmt.Printf("Set path: %v\n\n", c.getPath())
}

func readConsole() string {
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(result, "\n")[0]
}

func cd(args []string, path *currentPath) {
	if len(args) == 0 {
		fmt.Println(path.getPath())
	} else if len(args) == 1 {
		switch args[0] {
		case "..":
			path.decreasePath()
		default:
			path.setPath(args[0])
		}
	}
}

func echo(args []string, path *currentPath) {
	switch len(args) {
	case 0:
		fmt.Println("Режим вывода команд на экран (ECHO) включен")
		return
	default:
		arg := strings.Join(args, " ")
		if strings.Contains(arg, ">") {
			ind := strings.Index(arg, ">")
			nameFile := strings.Split(strings.TrimSpace(arg[ind+1:]), " ")[0]
			file, err := os.Create(path.getPath() + "\\" + nameFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			writer := bufio.NewWriter(file)
			writer.WriteString(arg[:ind])
			writer.Flush()
			return
		}
		fmt.Println(arg)
	}
}

func whatever() {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

		// do os.* stuff on the pid
	}
}

func kill(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: kill [pid]")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln("Argument must be countable")
	}
	ps, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalln("Can't find process", err)
	}
	err = ps.Kill()
	if err != nil {
		log.Fatalln("Can't kill process", err)
	}
	fmt.Printf("Process (pid = %v) killed\n", pid)
}

func main() {
	Path := new(currentPath)
	Path.setPath("p")

	for {
		fmt.Printf("%v-> ", Path.getPath())
		command := readConsole()
		pipeStr := strings.Split(command, "|")
		for i := 0; i < len(pipeStr); i++ {
			cmSplit := strings.Split(strings.TrimSpace(pipeStr[i]), " ")
			switch cmSplit[0] {
			case "cd":
				cd(cmSplit[1:], Path)
			case "pwd":
				fmt.Println(Path.getPath())
			case "echo":
				echo(cmSplit[1:], Path)
			case "ps":
				whatever()
			case "kill":
				kill(cmSplit[1:])
			case "quit":
				return
			default:
				fmt.Println("Can't find command")
			}
		}
	}
}

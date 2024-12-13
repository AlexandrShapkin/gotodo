package main

import (
	"log"
	"time"
	"os"
	"strings"

	"github.com/AlexandrShapkin/cli"
)

func main() {
	taskManager := NewTaskManager(NewStorage("./tasks.json"))

	cliProcessor := cli.NewCli()
	cliProcessor.AddCmd(
		&cli.Command{
			Use: "create",
			Flags: []*cli.CommandFlag{
				{
					Type:  "name",
					Long:  "name",
					Short: "n",
				},
				{
					Type:  "desc",
					Long:  "desc",
					Short: "d",
				},
				{
					Type:  "stat",
					Long:  "status",
					Short: "s",
				},
				{
					Type:  "prior",
					Long:  "priority",
					Short: "p",
				},
			},
			Run: func(flags map[string]*cli.ParsedCommandFlags, args []string) {
				var (
					name        string   = "default name"
					description string   = "default description"
					status      string   = "default status"
					priority    Priority = LOW
				)

				nameFlag, ok := flags["name"]
				if ok {
					name = nameFlag.Args
				}

				descriptionFlag, ok := flags["desc"]
				if ok {
					description = descriptionFlag.Args
				}

				statusFlag, ok := flags["stat"]
				if ok {
					status = statusFlag.Args
				}

				priorityFlag, ok := flags["prior"]
				if ok {
					switch priorityFlag.Args {
					case "low":
						priority = LOW
					case "medium":
						priority = MEDIUM
					case "high":
						priority = HIGH
					case "highest":
						priority = HIGHEST
					}
				}

				task := &Task{
					Name:         name,
					Description:  description,
					Status:       status,
					Priority:     priority,
					CreationTime: time.Now(),
					Deadline:     time.Now().Add(time.Hour * 24),
				}
				taskManager.Cteate(task)
			},
		},
	)

	cmd := strings.Join(os.Args[1:], " ")
	err := cliProcessor.OneCmd(cmd)
	if err != nil {
		log.Println(err)
	}
}

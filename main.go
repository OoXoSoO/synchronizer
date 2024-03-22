package main

import "github.com/fieldryand/goflow/v2"

func main() {
	options := goflow.Options{
		UIPath:      "ui/",
		WithSeconds: true,
	}
	gf := goflow.New(options)
	gf.AddJob(personalJob)

	gf.Use(goflow.DefaultLogger())
	gf.Run(":8181")
}

func personalJob() *goflow.Job {
	j := &goflow.Job{
		Name:     "printer",
		Schedule: "1 * * * * *", //https://crontab.guru/
		Active:   true,
	}
	j.Add(&goflow.Task{
		Name:        "print",
		TriggerRule: "allDone",
		Operator:    &command{},
	})

	return j
}

type command struct{}

func (c *command) Run() (interface{}, error) {
	print("\nhello\n")
	return nil, nil
}

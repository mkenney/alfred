package alfred

import (
	"os"
	"strings"

	event "github.com/kcmerrill/hook"
)

// NewTask will execute a task
func NewTask(task string, context *Context, loadedTasks map[string]Task) {
	dir, t, tasks := FetchTask(task, context, loadedTasks)

	// switch the directory
	os.Chdir(dir)

	// copy our context
	c := *context

	// innocent until proven guilty
	c.Ok = true

	// set our taskname
	c.TaskFile, c.TaskName = TaskParser(task, ":default")

	/*for x := 0; x <= 256; x++ {
		color := strconv.Itoa(x)
		fmt.Println(ansi.ColorCode(color), "color#"+color)
	}*/

	components := []Component{
		Component{"log", log},
		Component{"defaults", defaults},
		Component{"summary", summary},
		Component{"config", configC},
		Component{"register", register},
		Component{"env", env},
		Component{"serve", serve},
		Component{"setup", setup},
		Component{"multitask", multitask},
		Component{"tasks", tasksC},
		Component{"watch", watch},
		Component{"for", forC},
		Component{"command", commandC},
		Component{"commands", commands},
		Component{"result", result},
		Component{"ok", ok},
		Component{"fail", fail},
		Component{"wait", wait},
		Component{"every", every},
	}

	// cycle through our components ...
	event.Trigger("task.started", t, &c, tasks)
	for _, component := range components {
		context.Component = component.Name
		event.Trigger("before."+component.Name, t, &c, tasks)
		component.F(t, &c, tasks)
		event.Trigger("after."+component.Name, t, &c, tasks)
	}
	event.Trigger("task.completed", t, &c, tasks)
}

// Task holds all of our task components
type Task struct {
	Aliases  string
	Summary  string
	Usage    string
	Args     []string
	Setup    string
	Defaults []string
	Dir      string
	For      struct {
		Tasks     string
		MultiTask string
		Args      string
	}
	Config    string
	Log       string
	Every     string
	Command   string
	Retry     int
	Register  map[string]string
	Env       map[string]string
	Commands  string
	Serve     string
	Script    string
	Tasks     string
	MultiTask string
	Ok        string
	Fail      string
	Wait      string
	Watch     string
	ExitCode  int `yaml:"exit"`
}

// Exit determins whether a task should exit or not
func (t *Task) Exit(context *Context, tasks map[string]Task) {
	context.Ok = false
	if t.ExitCode != 0 {
		outFail("["+strings.Join(context.Args, ", ")+"]", "{{ .Text.Failure }}{{ .Text.FailureIcon }} exiting ...", context)
		os.Exit(t.ExitCode)
	}
}

// IsPrivate determines if a task is private
func (t *Task) IsPrivate() bool {
	// I like the idea of not needing to put an astrick next to a task
	// ... Descriptions automagically qualify for "important tasks"
	// No descriptions means it's filler, or private
	if t.Summary != "" {
		return false
	}

	return true
}

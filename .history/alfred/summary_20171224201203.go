package alfred

import (
	"strings"
	"time"
)

func summary(task Task, context *Context, tasks map[string]Task) {
	if task.Summary != "" {
		outOK("["+strings.Join(context.Args, ", ")+"]", task.Summary+"{{ .Text.Grey }} | {{ .Text.Reset }}started ...", context)
	} else {
		outOK("["+strings.Join(context.Args, ", ")+"]", "started ...", context)
	}
}

func result(task Task, context *Context, tasks map[string]Task) {
	if context.Ok {
		outOK("{{ .Text.SuccessIcon }} ok ["+strings.Join(context.Args, ", ")+"]", "elapsed time {{ .Text.Grey }}'{{ .Text.Success }}"+time.Since(context.Started).Round(time.Second).String()+"{{ .Text.Grey }}'", context)
	} else {
		outFail("{{ .Text.FailureIcon }} failed ["+strings.Join(context.Args, ", ")+"]", "elapsed time {{ .Text.Grey }}'{{ .Text.Success }}"+time.Since(context.Started).Round(time.Second).String()+"{{ .Text.Grey }}'", context)
	}
}

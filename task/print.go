package task

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
)

func (manager *TaskManager) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	idColor := color.New(color.FgYellow).SprintFunc()
	contentsColor := color.New(color.FgWhite).SprintFunc()
	durationColor := color.New(color.FgMagenta).SprintFunc()

	// Header
	fmt.Fprintf(w, "%s\t%s\t%s\n", color.CyanString("ID"), color.CyanString("CONTENTS"), color.CyanString("CREATED"))

	for _, task := range manager.tasks {
		duration := time.Unix(task.created, 0)
		since := time.Since(duration).Truncate(time.Second).String()

		fmt.Fprintf(w, "%v\t%v\t%v\n",
			idColor(task.id),
			contentsColor(task.contents),
			durationColor(since),
		)
	}

	w.Flush()
}

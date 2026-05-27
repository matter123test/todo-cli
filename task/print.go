package task

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func (manager *TaskManager) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Format header with tabs separating columns
	fmt.Fprintln(w, "ID\tCONTENTS\tCREATED")

	// Iterate and print data
	for _, task := range manager.tasks {
		createdAt := task.created
		duration := time.Unix(createdAt, 0)
		since := time.Since(duration).Truncate(time.Second)

		fmt.Fprintf(w, "%v\t%v\t%v\n", task.id, task.contents, since)
	}

	w.Flush()
}

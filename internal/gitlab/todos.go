package gitlab

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/zephinzer/dev/internal/types"
	gl "github.com/zephinzer/dev/pkg/gitlab"
)

func GetTodos(hostname, accessToken string, since ...time.Time) (types.Notifications, error) {
	todos, getTodosError := gl.GetTodos(hostname, accessToken, since...)
	if getTodosError != nil {
		return nil, getTodosError
	}
	notifications := types.Notifications{}
	for _, todo := range *todos {
		notifications = append(notifications, TodoSerializer(todo))
	}
	return notifications, nil
}

type TodoSerializer gl.APIv4Todo

func (ts TodoSerializer) GetTitle() string {
	return fmt.Sprintf("%s - %s", ts.Project.NameWithNamespace, ts.Target.Title)
}

func (ts TodoSerializer) GetMessage() string {
	targetType := "an item"
	switch ts.TargetType {
	case "Issue":
		targetType = "an issue"
	case "MergeRequest":
		targetType = "a merge request"
	}
	createdAt := humanize.Time(ts.CreatedAt)
	switch ts.ActionName {
	case "assigned":
		return fmt.Sprintf("you were assigned to %s by @%s about %s", targetType, ts.Author.Username, createdAt)
	case "mentioned":
		if ts.Body == ts.Target.Title {
			return fmt.Sprintf(
				"You were mentioned in %s ('%s') by @%s about %s, check it out at: %s",
				targetType,
				ts.Target.Title,
				ts.Author.Username,
				createdAt,
				ts.TargetURL,
			)
		}
		return fmt.Sprintf(
			"You were mentioned in %s ('%s') by @%s about %s: \"%s\", check it out at: %s",
			targetType,
			ts.Target.Title,
			ts.Author.Username,
			createdAt,
			ts.Body,
			ts.TargetURL,
		)
	case "build_failed":
		return fmt.Sprintf("your build failed in %s about %s", targetType, createdAt)
	case "marked":
		return fmt.Sprintf("%s was marked about %s", targetType, createdAt)
	case "approval_required":
		return fmt.Sprintf("your approval was required in %s about %s", targetType, createdAt)
	case "unmergeable":
		return fmt.Sprintf("%s seems unmergeable", targetType)
	case "directly_addressed":
		return fmt.Sprintf("you were directly addressed in %s by @%s about %s: %s", targetType, ts.Author.Username, createdAt, ts.Body)
	}
	return "unknown notification was received"
}

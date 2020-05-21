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
	targetType := ts.getTargetType()
	switch ts.ActionName {
	case "assigned":
		return fmt.Sprintf("You were assigned to %s in `%s`", targetType, ts.Project.PathWithNamespace)
	case "mentioned":
		return fmt.Sprintf("You were mentioned in %s in `%s`", targetType, ts.Project.PathWithNamespace)
	case "directly_addressed":
		return fmt.Sprintf("You were directly addressed in %s in `%s`", targetType, ts.Project.PathWithNamespace)
	case "build_failed":
		return fmt.Sprintf("Your build has failed in %s in `%s`", targetType, ts.Project.PathWithNamespace)
	case "marked":
		return fmt.Sprintf("Something was marked in %s in `%s`", targetType, ts.Project.PathWithNamespace)
	case "approval_required":
		return fmt.Sprintf("Your approval is required in %s in `%s`", targetType, ts.Project.PathWithNamespace)
	case "unmergeable":
		return fmt.Sprintf("Your merge request is unmergeable in `%s`", ts.Project.PathWithNamespace)
	}
	return fmt.Sprintf("%s - %s", ts.Project.PathWithNamespace, ts.Target.Title)
}

func (ts TodoSerializer) GetMessage() string {
	targetType := ts.getTargetType()
	createdAt := humanize.Time(ts.CreatedAt)
	switch ts.ActionName {
	case "assigned":
		return fmt.Sprintf("You were assigned to %s by @%s about %s, check it out at: %s", targetType, ts.Author.Username, createdAt, ts.TargetURL)
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
		return fmt.Sprintf("Your build failed in %s ('%s') about %s, check it out at: %s", targetType, ts.Target.Title, createdAt, ts.TargetURL)
	case "marked":
		return fmt.Sprintf("%s was marked about %s, check it out at: %s", targetType, createdAt, ts.TargetURL)
	case "approval_required":
		return fmt.Sprintf("Your approval was required in %s ('%s') about %s, check it out at: %s", targetType, ts.Target.Title, createdAt, ts.TargetURL)
	case "unmergeable":
		return fmt.Sprintf("Your merge request seems unmergeable, check it out at: %s", ts.TargetURL)
	case "directly_addressed":
		if ts.Body == ts.Target.Title {
			return fmt.Sprintf(
				"You were directly addressed in %s ('%s') by @%s about %s, check it out at: %s",
				targetType,
				ts.Target.Title,
				ts.Author.Username,
				createdAt,
				ts.TargetURL,
			)
		}
		return fmt.Sprintf(
			"You were directly addressed in %s ('%s') by @%s about %s: \"%s\", check it out at: %s",
			targetType,
			ts.Target.Title,
			ts.Author.Username,
			createdAt,
			ts.Body,
			ts.TargetURL,
		)
	}
	return "unknown notification was received"
}

func (ts TodoSerializer) getTargetType() string {
	targetType := "an item"
	switch ts.TargetType {
	case "Issue":
		targetType = "an issue"
	case "MergeRequest":
		targetType = "a merge request"
	}
	return targetType
}

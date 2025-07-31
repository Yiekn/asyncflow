package util

import (
	"fmt"
	"strings"
)

func JointTaskId(snowflakeId int64, tableName string) string {
	return fmt.Sprintf("%d_%s", snowflakeId, tableName)
}

func ConvertToTableName(taskID string) string {
	s := strings.Split(taskID, "_")[1:]
	return strings.Join(s, "_")
}

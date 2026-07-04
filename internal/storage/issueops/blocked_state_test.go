package issueops

import (
	"strings"
	"testing"
)

func TestBlockedStateSQLUsesSQLiteCompatibleUpdateTargets(t *testing.T) {
	for name, query := range map[string]string{
		"mark issues":   markBlockedTemplateForIssues(),
		"unmark issues": unmarkBlockedTemplateForIssues(),
		"mark wisps":    markBlockedTemplateForWisps(),
		"unmark wisps":  unmarkBlockedTemplateForWisps(),
	} {
		query = strings.Join(strings.Fields(query), " ")
		if strings.Contains(query, "UPDATE issues i SET") || strings.Contains(query, "UPDATE wisps w SET") {
			t.Fatalf("%s uses a target-table alias in UPDATE, which DoltLite's SQLite-compatible parser rejects: %s", name, query)
		}
	}
}

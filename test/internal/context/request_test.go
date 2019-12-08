package context

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sepuka/chat/internal/context"
	"github.com/sepuka/chat/internal/domain"
)

func TestTooLongCommandsAreDisabled(t *testing.T) {
	var testCases = []struct {
		command  string
		expected string
	}{
		{
			command:  `some short command`,
			expected: `some short command`,
		},
		{
			command:  `there is too long command, longer than 32 symbols`,
			expected: `there is too long command, longe`,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, context.NewRequest(`login`, domain.Manual, testCase.command).GetCommand())
	}
}

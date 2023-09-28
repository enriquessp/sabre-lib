package sabre

import (
	"fmt"
	"os"
	"testing"

	agtest "github.com/agentguru/test-lib"
	"github.com/stretchr/testify/assert"
)

func Test_convertRawSessionCreateRS(t *testing.T) {
	scenariosDir := fmt.Sprintf("%s/docs/xmls/session/create/resp", Root)
	tests := agtest.ReadDinamicallyTestCases(scenariosDir, createConvertSessionCreateRSTestcase)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := convertRawSessionCreateRS(tt.data)
			tt.want(t, rs.(*SessionCreateRS))
		})
	}
}

type convertSessionCreateRSTestcase struct {
	name string
	data []byte
	want func(*testing.T, *SessionCreateRS)
}

func createConvertSessionCreateRSTestcase(file string) convertSessionCreateRSTestcase {
	fmt.Println("Reading test file", file)
	data, _ := os.ReadFile(file)

	return convertSessionCreateRSTestcase{
		name: agtest.ScenarioName(file),
		data: data,
		want: assertSessionCreateRSRequiredFields,
	}
}

func assertSessionCreateRSRequiredFields(t *testing.T, rs *SessionCreateRS) {
	assert.NotNil(t, rs)
}

func Test_convertRawSessionCloseRS(t *testing.T) {
	scenariosDir := fmt.Sprintf("%s/docs/xmls/session/close/resp", Root)
	tests := agtest.ReadDinamicallyTestCases(scenariosDir, createConvertSessionCloseRSTestcase)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := convertRawSessionCloseRS(tt.data)
			tt.want(t, rs.(*SessionCloseRS))
		})
	}
}

type convertSessionCloseRSTestcase struct {
	name string
	data []byte
	want func(*testing.T, *SessionCloseRS)
}

func createConvertSessionCloseRSTestcase(file string) convertSessionCloseRSTestcase {
	fmt.Println("Reading test file", file)
	data, _ := os.ReadFile(file)

	return convertSessionCloseRSTestcase{
		name: agtest.ScenarioName(file),
		data: data,
		want: assertSessionCloseRSRequiredFields,
	}
}

func assertSessionCloseRSRequiredFields(t *testing.T, rs *SessionCloseRS) {
	assert.NotNil(t, rs)
	assert.NotEmpty(t, rs.Status)
}

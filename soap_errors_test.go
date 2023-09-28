package sabre

import (
	"fmt"
	"testing"
)

func Test_IssueError_CallServer(t *testing.T) {
	t.Skip()
	tests := map[string]struct {
		xmlRQ_File   string
		xmlRS_File   string
		errorMessage string
	}{
		"GENERAL_ERROR": {"reqs/orderChangeRQ_CC.xml", "errors/orderViewRS_error_general_issue.xml", "rpc error: code = Unknown desc = GENERAL_ERROR"},
	}

	for scenario, c := range tests {
		t.Run(scenario, func(t *testing.T) {
			fmt.Println(scenario, c)
			// httpServer, close := agtest.CreateHttpServerTest(map[string]string{
			// 	"AA001HGPJB3A5": issueXmlsPath(c.xmlRS_File),
			// })
			// defer close()
			//
			// server := NewSabreServer(
			// 	WithClient(http.DefaultClient),
			// 	WithUrl(httpServer.URL),
			// 	WithSabreSession(encodedTc),
			// )
			//
			// _, err := CallServer(server, readMockFile(issueXmlsPath(c.xmlRQ_File), &OrderChangeRQ{}), &OrderViewRS{})
			// assert.NotNil(t, err)
			// assert.Equal(t, c.errorMessage, err.Error())
		})
	}
}

// Helpers
func issueXmlsPath(file string) string {
	return filePath(fmt.Sprintf("docs/xmls/ticketing/issue/%s", file))
}

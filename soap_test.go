package sabre

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	agtest "github.com/agentguru/test-lib"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func Test_GenericCall(t *testing.T) {
	t.Skip()

	url := "https://sws-crt.cert.havail.sabre.com/websvc"
	username := "420831"
	password := "Z39B0fA"
	complement := "N8BK"
	pcc := "6GIF"

	client := newCall("SessionCreateRQ")

	session := &SabreSession{ConversationID: xid.New().String()}

	server := NewSabreServer(
		WithClient(client),
		WithUrl(url),
		WithCredenntials(username, password, complement, pcc),
		WithSabreSession(session),
	)

	// CREATING A SESSION

	sessionCreateRS, err := CallServer(server, convertSessionCreate(pcc), &RSEnvelope{})
	assert.Nil(t, err)
	assert.NotNil(t, sessionCreateRS)
	assert.NotEmpty(t, sessionCreateRS.Header.Security.BinarySecurityToken.Value)
	assert.NotEmpty(t, sessionCreateRS.Body.SessionCreateRS.ConversationId)

	session.SecurityToken = sessionCreateRS.Header.Security.BinarySecurityToken.Value
	session.ConversationID = sessionCreateRS.Body.SessionCreateRS.ConversationId

	defer CallServer(server, convertSessionClose(pcc), &SessionCloseRS{})

	reservationRS, err := CallServer(server, convertGetReservationRequest("CWZTXI"), &GetReservationRS{})
	assert.Nil(t, err)
	assert.NotNil(t, reservationRS)
	assert.NotEmpty(t, reservationRS.Reservation.BookingDetails.RecordLocator)
}

// HELPERS
func newCall(service string) *http.Client {
	scenario := fmt.Sprintf("Service_%s", service)
	r, _ := agtest.NewHttpRecorder(context.Background(), scenario, recorder.ModeRecordOnly)

	return r.GetDefaultClient()
}

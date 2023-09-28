package sabre

import (
	"fmt"
	"os"
	"testing"

	"github.com/agentguru/collections-lib"
	agtest "github.com/agentguru/test-lib"
	"github.com/stretchr/testify/assert"
)

const required_fields = "required_fields"

func Test_convertRawGetReservationRS(t *testing.T) {
	scenariosDir := fmt.Sprintf("%s/docs/xmls/booking/retrieve/resp", Root)
	tests := agtest.ReadDinamicallyTestCases(scenariosDir, createConvertGetReservationRSTestcase)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := convertRawGetReservationRS(tt.data)
			tt.want[required_fields](t, rs)

			if want, ok := tt.want[tt.name]; ok {
				want(t, rs)
			}
		})
	}
}

type convertGetReservationRSTestcase struct {
	name string
	data []byte
	want map[string]func(*testing.T, interface{})
}

func createConvertGetReservationRSTestcase(file string) convertGetReservationRSTestcase {
	data, _ := os.ReadFile(file)

	assertions := map[string]func(*testing.T, interface{}){
		required_fields:                 assertGetReservationRSRequiredFields,
		"getreservationrs_chld_ssr.xml": assertGetReservationRSChldSsr,
	}

	return convertGetReservationRSTestcase{
		name: agtest.ScenarioName(file),
		data: data,
		want: assertions,
	}
}

func assertGetReservationRSChldSsr(t *testing.T, rs interface{}) {
	assertGetReservationRSRequiredFields(t, rs)

	childPaxId := "02.01"

	reservation := rs.(*GetReservationRS)
	ssrs := reservation.GetOpenReservationsByPax(childPaxId)
	assert.NotEmpty(t, ssrs)
	assert.Equal(t, 3, len(ssrs))

	chldSsrs := collections.FilterSlice(ssrs, func(o OpenReservationElement) bool { return o.ServiceRequest.Code == SSR_CHLD_CODE })
	assert.Equal(t, 2, len(chldSsrs))

	assert.True(t, reservation.HasChild())
	assert.Equal(t, 2, len(reservation.GetPassengers()))
	assert.Equal(t, 1, len(reservation.GetAdultPassengers()))
	assert.Equal(t, 1, len(reservation.GetChildPassengers()))
	assert.Equal(t, 0, len(reservation.GetInfantPassengers()))
}

func assertGetReservationRSRequiredFields(t *testing.T, rs interface{}) {
	assert.NotNil(t, rs)

	reservation := rs.(*GetReservationRS)
	assert.NotEmpty(t, reservation.Reservation.BookingDetails.RecordLocator)
}

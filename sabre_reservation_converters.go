package sabre

import (
	"encoding/xml"
)

func convertGetReservationRequest(loc string) *GetReservationRQ {
	// <ns4:GetReservationRQ xmlns:ns2="http://schemas.xmlsoap.org/ws/2002/12/secext" xmlns:ns4="http://webservices.sabre.com/pnrbuilder/v1_19" xmlns:ns5="http://services.sabre.com/res/or/v1_14" xmlns:ns6="http://www.ebxml.org/namespaces/messageHeader" xmlns:ns7="http://www.w3.org/1999/xlink" xmlns:ns8="http://www.w3.org/2000/09/xmldsig#" Version="1.19.0">
	// 	<ns4:Locator>%s</ns4:Locator>
	// 	<ns4:RequestType>Stateful</ns4:RequestType>
	// 	<ns4:ReturnOptions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" PriceQuoteServiceVersion="4.0.0" xsi:type="ns4:ReturnOptions">
	// 	<ns4:SubjectAreas>
	// 	<ns4:SubjectArea>ACTIVE</ns4:SubjectArea>
	// 	<ns4:SubjectArea>ANCILLARY</ns4:SubjectArea>
	// 	<ns4:SubjectArea>FARE_TYPE</ns4:SubjectArea>
	// 	<ns4:SubjectArea>PRICING_INFORMATION</ns4:SubjectArea>
	// 	<ns4:SubjectArea>PRICE_QUOTE</ns4:SubjectArea>
	// 	</ns4:SubjectAreas>
	// 	<ns4:ViewName>Full</ns4:ViewName>
	// 	<ns4:ResponseFormat>STL</ns4:ResponseFormat>
	// 	</ns4:ReturnOptions>
	// </ns4:GetReservationRQ>

	rq := &GetReservationRQ{
		Ns2:         "http://schemas.xmlsoap.org/ws/2002/12/secext",
		Ns4:         "http://webservices.sabre.com/pnrbuilder/v1_19",
		Ns5:         "http://services.sabre.com/res/or/v1_14",
		Ns6:         "http://www.ebxml.org/namespaces/messageHeader",
		Ns7:         "http://www.w3.org/1999/xlink",
		Ns8:         "http://www.w3.org/2000/09/xmldsig#",
		Version:     "1.19.0",
		Locator:     loc,
		RequestType: "Stateful",
		ReturnOptions: ReturnOptions{
			PriceQuoteServiceVersion: "4.0.0",
			Xsi:                      "http://www.w3.org/2001/XMLSchema-instance",
			Type:                     "ns4:ReturnOptions",
			SubjectAreas: SubjectAreas{
				SubjectArea: []string{
					"ACTIVE",
					"ANCILLARY",
					"FARE_TYPE",
					"PRICING_INFORMATION",
					"PRICE_QUOTE",
				},
			},
			ResponseFormat: "STL",
			ViewName:       "Full",
		},
	}

	return rq
}

func convertRawGetReservationRS(resp []byte) interface{} {
	var rs RSEnvelope
	xml.Unmarshal(resp, &rs)

	return rs.Body.GetReservationRS
}

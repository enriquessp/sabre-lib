package sabre

import (
	"encoding/xml"
)

func convertSessionCreate(pcc string) *SessionCreateRQ {
	return &SessionCreateRQ{
		Ns2: "http://www.opentravel.org/OTA/2002/11",
		Ns3: "http://schemas.xmlsoap.org/ws/2002/12/secext",
		Ns5: "http://www.ebxml.org/namespaces/messageHeader",
		Ns6: "http://www.w3.org/1999/xlink",
		Ns7: "http://www.w3.org/2000/09/xmldsig#",
		POS: POSRQ{
			Source: SourceRQ{
				PseudoCityCode: pcc,
			},
		},
	}
}

func convertSessionClose(pcc string) *SessionCloseRQ {
	return &SessionCloseRQ{
		Ns2: "http://www.opentravel.org/OTA/2002/11",
		Ns3: "http://schemas.xmlsoap.org/ws/2002/12/secext",
		Ns5: "http://www.ebxml.org/namespaces/messageHeader",
		Ns6: "http://www.w3.org/1999/xlink",
		Ns7: "http://www.w3.org/2000/09/xmldsig#",
		POS: POSRQ{
			Source: SourceRQ{
				PseudoCityCode: pcc,
			},
		},
	}
}

func convertRawSessionCreateRSEnvelope(data []byte) interface{} {
	var rs RSEnvelope
	xml.Unmarshal(data, &rs)

	return &rs
}

func convertRawSessionCreateRS(data []byte) interface{} {
	var rs RSEnvelope
	xml.Unmarshal(data, &rs)

	return rs.Body.SessionCreateRS
}

func convertRawSessionCloseRS(data []byte) interface{} {
	var rs RSEnvelope
	xml.Unmarshal(data, &rs)

	return rs.Body.SessionCloseRS
}

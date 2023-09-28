package sabre

import (
	"encoding/xml"
	"fmt"
	"time"
)

const (
	loginService = "SessionCreateRQ"
	kgEquivalent = 0.45359237

	SSR_CHLD_CODE = "CHLD"
	SSR_DOCS_CODE = "DOCS"
)

func createSabreSoapEnvelope(req interface{}, service, securityToken string, session *SabreSession) (string, error) {
	requestBody, err := xml.Marshal(&req)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(sabreRQEnvelope,
		session.Pcc,
		session.ConversationID,
		service,
		service,
		time.Now().Format(DateTimeFormat),
		securityToken,
		requestBody,
	), nil
}

const sabreUsernameTokenTemplate = `
         <ns3:UsernameToken>
            <ns3:Username>%s</ns3:Username>
            <ns3:Password>%s</ns3:Password>
            <Organization>%s</Organization>
            <Domain>DEFAULT</Domain>
         </ns3:UsernameToken>
`

const binarySecurityTokenTemplate = `
         <ns3:BinarySecurityToken>%s</ns3:BinarySecurityToken>
	`

const sabreRQEnvelope = `
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/" xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Header>
      <ns5:MessageHeader ns5:version="1.0" xmlns:ns2="http://www.opentravel.org/OTA/2002/11" xmlns:ns3="http://schemas.xmlsoap.org/ws/2002/12/secext" xmlns:ns5="http://www.ebxml.org/namespaces/messageHeader" xmlns:ns6="http://www.w3.org/1999/xlink" xmlns:ns7="http://www.w3.org/2000/09/xmldsig#">
         <ns5:From>
            <ns5:PartyId>agent.guru</ns5:PartyId>
         </ns5:From>
         <ns5:To>
            <ns5:PartyId>sabre.com</ns5:PartyId>
         </ns5:To>
         <ns5:CPAId>%s</ns5:CPAId>
         <ns5:ConversationId>%s</ns5:ConversationId>
         <ns5:Service>%s</ns5:Service>
         <ns5:Action>%s</ns5:Action>
         <ns5:MessageData>
            <ns5:MessageId>mid:20210313-2047-43478@agent.guru</ns5:MessageId>
            <ns5:Timestamp>%s</ns5:Timestamp>
         </ns5:MessageData>
      </ns5:MessageHeader>
      <ns3:Security xmlns:ns2="http://www.opentravel.org/OTA/2002/11" xmlns:ns3="http://schemas.xmlsoap.org/ws/2002/12/secext" xmlns:ns5="http://www.ebxml.org/namespaces/messageHeader" xmlns:ns6="http://www.w3.org/1999/xlink" xmlns:ns7="http://www.w3.org/2000/09/xmldsig#">
			%s
      </ns3:Security>
   </S:Header>
   <S:Body>
			%s
   </S:Body>
</S:Envelope>
`

type SabreSession struct {
	Url            string
	Pcc            string
	SecurityToken  string
	ConversationID string
}

type RSEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SoapEnv string   `xml:"soap-env,attr"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  struct {
		Text          string `xml:",chardata"`
		MessageHeader struct {
			Text           string `xml:",chardata"`
			Eb             string `xml:"eb,attr"`
			Version        string `xml:"version,attr"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
			From           struct {
				Text    string `xml:",chardata"`
				PartyId struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"PartyId"`
			} `xml:"From"`
			To struct {
				Text    string `xml:",chardata"`
				PartyId struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"PartyId"`
			} `xml:"To"`
			CPAId          string `xml:"CPAId"`
			ConversationId string `xml:"ConversationId"`
			Service        struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"Service"`
			Action      string `xml:"Action"`
			MessageData struct {
				Text           string `xml:",chardata"`
				MessageId      string `xml:"MessageId"`
				Timestamp      string `xml:"Timestamp"`
				RefToMessageId string `xml:"RefToMessageId"`
			} `xml:"MessageData"`
		} `xml:"MessageHeader"`
		Security struct {
			Text                string `xml:",chardata"`
			Wsse                string `xml:"wsse,attr"`
			BinarySecurityToken struct {
				Value        string `xml:",chardata"`
				EncodingType string `xml:"EncodingType,attr"`
				ValueType    string `xml:"valueType,attr"`
			} `xml:"BinarySecurityToken"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		SessionCreateRS  *SessionCreateRS  `xml:"SessionCreateRS"`
		SessionCloseRS   *SessionCloseRS   `xml:"SessionCloseRS"`
		GetReservationRS *GetReservationRS `xml:"GetReservationRS"`
	} `xml:"Body"`
}

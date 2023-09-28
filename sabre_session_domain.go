package sabre

import "encoding/xml"

// REQUEST

type SessionCloseRQ struct {
	XMLName xml.Name `xml:"SessionCloseRQ"`
	Ns2     string   `xml:"xmlns:ns2,attr"`
	Ns3     string   `xml:"xmlns:ns3,attr"`
	Ns5     string   `xml:"xmlns:ns5,attr"`
	Ns6     string   `xml:"xmlns:ns6,attr"`
	Ns7     string   `xml:"xmlns:ns7,attr"`
	POS     POSRQ    `xml:"ns2:POS"`
}

type SessionCreateRQ struct {
	Ns2 string `xml:"xmlns:ns2,attr"`
	Ns3 string `xml:"xmlns:ns3,attr"`
	Ns5 string `xml:"xmlns:ns5,attr"`
	Ns6 string `xml:"xmlns:ns6,attr"`
	Ns7 string `xml:"xmlns:ns7,attr"`
	POS POSRQ  `xml:"ns2:POS"`
}

type POSRQ struct {
	Source SourceRQ `xml:"ns2:Source"`
}

type SourceRQ struct {
	PseudoCityCode string `xml:"PseudoCityCode,attr"`
}

type SabreSessionCreateRQEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  struct {
		Text          string `xml:",chardata"`
		MessageHeader struct {
			Text    string `xml:",chardata"`
			Version string `xml:"version,attr"`
			Ns2     string `xml:"ns2,attr"`
			Ns3     string `xml:"ns3,attr"`
			Ns5     string `xml:"ns5,attr"`
			Ns6     string `xml:"ns6,attr"`
			Ns7     string `xml:"ns7,attr"`
			From    struct {
				Text    string `xml:",chardata"`
				PartyId string `xml:"PartyId"`
			} `xml:"From"`
			To struct {
				Text    string `xml:",chardata"`
				PartyId string `xml:"PartyId"`
			} `xml:"To"`
			CPAId          string `xml:"CPAId"`
			ConversationId string `xml:"ConversationId"`
			Service        string `xml:"Service"`
			Action         string `xml:"Action"`
			MessageData    struct {
				Text      string `xml:",chardata"`
				MessageId string `xml:"MessageId"`
				Timestamp string `xml:"Timestamp"`
			} `xml:"MessageData"`
		} `xml:"MessageHeader"`
		Security struct {
			Text          string `xml:",chardata"`
			Ns2           string `xml:"ns2,attr"`
			Ns3           string `xml:"ns3,attr"`
			Ns5           string `xml:"ns5,attr"`
			Ns6           string `xml:"ns6,attr"`
			Ns7           string `xml:"ns7,attr"`
			UsernameToken struct {
				Text         string `xml:",chardata"`
				Username     string `xml:"Username"`
				Password     string `xml:"Password"`
				Organization string `xml:"Organization"`
				Domain       string `xml:"Domain"`
			} `xml:"UsernameToken"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text            string `xml:",chardata"`
		SessionCreateRQ `xml:"SessionCreateRQ"`
	} `xml:"Body"`
}

// RESPONSE

type SessionCreateRS struct {
	Text          string `xml:",chardata"`
	EchoToken     string `xml:"EchoToken,attr"`
	TimeStamp     string `xml:"TimeStamp,attr"`
	Target        string `xml:"Target,attr"`
	Version       string `xml:"version,attr"`
	SequenceNmbr  string `xml:"SequenceNmbr,attr"`
	PrimaryLangID string `xml:"PrimaryLangID,attr"`
	AltLangID     string `xml:"AltLangID,attr"`
	Status        string `xml:"status,attr"`
	Success       string `xml:"Success"`
	Warnings      struct {
		Text    string `xml:",chardata"`
		Warning struct {
			Text      string `xml:",chardata"`
			Language  string `xml:"Language,attr"`
			ShortText string `xml:"ShortText,attr"`
			Type      string `xml:"Type,attr"`
			Code      string `xml:"Code,attr"`
			DocURL    string `xml:"DocURL,attr"`
			Status    string `xml:"Status,attr"`
			Tag       string `xml:"Tag,attr"`
			RecordId  string `xml:"RecordId,attr"`
		} `xml:"Warning"`
	} `xml:"Warnings"`
	ConversationId string `xml:"ConversationId"`
	Errors         struct {
		Text  string `xml:",chardata"`
		Error struct {
			Text         string `xml:",chardata"`
			ErrorCode    string `xml:"ErrorCode,attr"`
			Severity     string `xml:"Severity,attr"`
			ErrorMessage string `xml:"ErrorMessage,attr"`
			ErrorInfo    struct {
				Text    string `xml:",chardata"`
				Message string `xml:"Message"`
			} `xml:"ErrorInfo"`
		} `xml:"Error"`
	} `xml:"Errors"`
}

type SessionCloseRS struct {
	Text          string `xml:",chardata"`
	EchoToken     string `xml:"EchoToken,attr"`
	TimeStamp     string `xml:"TimeStamp,attr"`
	Target        string `xml:"Target,attr"`
	Version       string `xml:"version,attr"`
	SequenceNmbr  string `xml:"SequenceNmbr,attr"`
	PrimaryLangID string `xml:"PrimaryLangID,attr"`
	AltLangID     string `xml:"AltLangID,attr"`
	Status        string `xml:"status,attr"`
	Success       string `xml:"Success"`
	Warnings      struct {
		Text    string `xml:",chardata"`
		Warning struct {
			Text      string `xml:",chardata"`
			Language  string `xml:"Language,attr"`
			ShortText string `xml:"ShortText,attr"`
			Type      string `xml:"Type,attr"`
			Code      string `xml:"Code,attr"`
			DocURL    string `xml:"DocURL,attr"`
			Status    string `xml:"Status,attr"`
			Tag       string `xml:"Tag,attr"`
			RecordId  string `xml:"RecordId,attr"`
		} `xml:"Warning"`
	} `xml:"Warnings"`
	ConversationId string `xml:"ConversationId"`
	Errors         struct {
		Text  string `xml:",chardata"`
		Error struct {
			Text         string `xml:",chardata"`
			ErrorCode    string `xml:"ErrorCode,attr"`
			Severity     string `xml:"Severity,attr"`
			ErrorMessage string `xml:"ErrorMessage,attr"`
			ErrorInfo    struct {
				Text    string `xml:",chardata"`
				Message string `xml:"Message"`
			} `xml:"ErrorInfo"`
		} `xml:"Error"`
	} `xml:"Errors"`
}

package sabre

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type SabreServer struct {
	customConverters                    map[string]func([]byte) interface{}
	client                              *http.Client
	url                                 string
	username, password, complement, pcc string
	session                             *SabreSession
}

func (s *SabreServer) GetSecurityToken(svc string) string {
	if svc == loginService {
		return fmt.Sprintf(sabreUsernameTokenTemplate,
			s.username,
			s.password,
			s.complement,
		)
	}
	return fmt.Sprintf(binarySecurityTokenTemplate, s.session.SecurityToken)
}

type serverOption func(*SabreServer)

func WithCustomConverters(v map[string]func([]byte) interface{}) serverOption {
	return func(ns *SabreServer) {
		ns.customConverters = v
	}
}

func WithClient(c *http.Client) serverOption {
	return func(ns *SabreServer) {
		ns.client = c
	}
}

func WithCredenntials(username, password, complement, pcc string) serverOption {
	return func(ns *SabreServer) {
		ns.username = username
		ns.password = password
		ns.complement = complement
		ns.pcc = pcc
	}
}

func WithSabreSession(s *SabreSession) serverOption {
	return func(ns *SabreServer) {
		ns.session = s
	}
}

func WithUrl(v string) serverOption {
	return func(ns *SabreServer) {
		ns.url = v
	}
}

func NewSabreServer(options ...serverOption) *SabreServer {
	server := &SabreServer{}

	for _, apply := range options {
		apply(server)
	}

	return server
}

var converters = map[string]func([]byte) interface{}{
	"SessionCreateRQ:SessionCreateRS":   convertRawSessionCreateRS,
	"SessionCreateRQ:RSEnvelope":        convertRawSessionCreateRSEnvelope,
	"SessionCloseRQ:SessionCloseRS":     convertRawSessionCloseRS,
	"GetReservationRQ:GetReservationRS": convertRawGetReservationRS,
}

func CallServer[T, V any](server *SabreServer, req T, resp V) (V, error) {
	key := fmt.Sprint(reflect.TypeOf(req).Elem().Name(), ":", reflect.TypeOf(resp).Elem().Name())

	converter, ok := server.customConverters[key]
	if !ok {
		converter, ok = converters[key]
	}
	if !ok {
		return resp, errors.New(fmt.Sprintf("No converter found for %s ", key))
	}

	serviceName := reflect.TypeOf(req).Elem().Name()
	headers := []header{
		{"Content-Type", "text/xml;charset=UTF-8"},
		{"SOAPAction", "OTA"},
		{"Host", extractHost(server.url)},
		{"Connection", "Keep-Alive"},
	}
	envelope, _ := createSabreSoapEnvelope(&req, serviceName, server.GetSecurityToken(serviceName), server.session)
	fmt.Println(envelope)
	if serviceName != "SessionCloseRQ" {
		os.WriteFile("/tmp/rq", []byte(envelope), os.ModeAppend)
	}

	body, err := callSoap(server.client, server.url, string(envelope), headers...)
	fmt.Println(string(body))
	if serviceName != "SessionCloseRQ" {
		os.WriteFile("/tmp/resp", body, os.ModeAppend)
	}

	if err != nil {
		return resp, err
	}

	// var faultEnvelope FaultEnvelope
	// xml.Unmarshal(body, &faultEnvelope)
	// if len(faultEnvelope.Body.Fault.Faultcode) > 0 || len(faultEnvelope.Body.Fault.Faultstring) > 0 {
	// 	return resp, errors.New(fmt.Sprintf("%s - %s", faultEnvelope.Body.Fault.Faultcode, faultEnvelope.Body.Fault.Faultstring))
	// }

	var sabreEnvelope RSEnvelope
	xml.Unmarshal(body, &sabreEnvelope)
	// if ndcEnvelope.Error() != nil {
	// 	return resp, ndcEnvelope.Error()
	// }

	return converter(body).(V), nil
}

func callSoap(client *http.Client, url string, postBody string, headers ...header) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s", url), bytes.NewBufferString(postBody))

	for _, h := range headers {
		req.Header.Add(h.key, h.value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type converter func([]byte) (interface{}, error)

func extractHost(url string) string {
	return strings.Split(url, "://")[1]
}

type header struct {
	key   string
	value string
}

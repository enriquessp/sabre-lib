package sabre

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DateTimeFormat = "2006-01-02T15:04:05"
	DateFormat     = "2006-01-02"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "./")
)

// 2039-01'
func ParseExpiration(exp string) string {
	if len(exp) != 7 {
		return exp
	}
	return fmt.Sprintf("%s%s", exp[5:], exp[2:4])
}

func RoundFloat64(value float64) float64 {
	return math.Round(float64(value)*100) / 100
}

func EncodeBase64(raw []byte) string {
	return base64.StdEncoding.EncodeToString(raw)
}

func DecodeBase64(raw string) ([]byte, error) {
	rs, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func FormatAmount(raw float64) string {
	return fmt.Sprintf("%.2f", raw)
}

func ParseAmount(raw string) float64 {
	if len(raw) <= 2 {
		return 0.0
	}
	raw = fmt.Sprintf("%s.%s", raw[:len(raw)-2], raw[len(raw)-2:])
	rs, _ := strconv.ParseFloat(raw, 32)
	return RoundFloat64(rs)
}

func UnmarshalJson[T any](data []byte, t T) (T, error) {
	err := json.Unmarshal(data, &t)
	return t, err
}

func UnmarshalXml[T any](data []byte, t T) (T, error) {
	err := xml.Unmarshal(data, &t)
	return t, err
}

func ConvertLbToKg(lb int) int {
	converted := float64(lb) * kgEquivalent
	roundedToInt := int(math.Round(converted))
	return roundedToInt
}

func ParseInt32(value string) int32 {
	v, _ := strconv.Atoi(value)
	return int32(v)
}

func ProtoDateTimeParse(date, hour string) *timestamppb.Timestamp {
	return timestamppb.New(DateTimeParse(date, hour))
}

func DateTimeParse(date, hour string) time.Time {
	d, _ := time.Parse(DateTimeFormat, fmt.Sprintf("%sT%s:00", date, hour))
	return d
}

// PT07H40M
func ParseDuration(d string) int32 {
	hour := 0
	minutes := 0

	reHour := regexp.MustCompile(`.*(\d+)H.*`)
	reMinute := regexp.MustCompile(`.*H(\d+)M.*`)

	if reHour.Match([]byte(d)) {
		h := reHour.FindAllStringSubmatch(d, -1)
		if len(h) > 0 && len(h[0]) > 1 {
			hour, _ = strconv.Atoi(h[0][1])
		}
	}

	if reMinute.Match([]byte(d)) {
		m := reMinute.FindAllStringSubmatch(d, -1)
		if len(m) > 0 && len(m[0]) > 1 {
			minutes, _ = strconv.Atoi(m[0][1])
		}
	}

	return int32(hour)*60 + int32(minutes)
}

func filePath(file string) string {
	return fmt.Sprintf("%s/%s", Root, file)
}

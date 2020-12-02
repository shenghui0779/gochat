package wx

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
)

// X is a convenient alias for a map[string]interface{}.
type X map[string]interface{}

// CDATA XML CDATA section which is defined as blocks of text that are not parsed by the parser, but are otherwise recognized as markup.
type CDATA string

// MarshalXML encodes the receiver as zero or more XML elements.
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

var BufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 4<<10)) // 4kb
	},
}

// FormatMap2XML format map to xml
func FormatMap2XML(m WXML) (string, error) {
	buf := BufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	defer BufferPool.Put(buf)

	if _, err := io.WriteString(buf, "<xml>"); err != nil {
		return "", err
	}

	for k, v := range m {
		if _, err := io.WriteString(buf, fmt.Sprintf("<%s>", k)); err != nil {
			return "", err
		}

		if err := xml.EscapeText(buf, []byte(v)); err != nil {
			return "", err
		}

		if _, err := io.WriteString(buf, fmt.Sprintf("</%s>", k)); err != nil {
			return "", err
		}
	}

	if _, err := io.WriteString(buf, "</xml>"); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ParseXML2Map parse xml to map
func ParseXML2Map(b []byte) (WXML, error) {
	m := make(WXML)

	xmlReader := bytes.NewReader(b)

	var (
		d     = xml.NewDecoder(xmlReader)
		tk    xml.Token
		depth = 0 // current xml.Token depth
		key   string
		buf   bytes.Buffer
		err   error
	)

	for {
		tk, err = d.Token()

		if err != nil {
			if err == io.EOF {
				return m, nil
			}

			return nil, err
		}

		switch v := tk.(type) {
		case xml.StartElement:
			depth++

			switch depth {
			case 2:
				key = v.Name.Local
				buf.Reset()
			case 3:
				if err = d.Skip(); err != nil {
					return nil, err
				}

				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				buf.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = buf.String()
			}

			depth--
		}
	}
}

// SignWithMD5 生成MD5签名
func SignWithMD5(m WXML, apikey string, toUpper bool) string {
	h := md5.New()
	h.Write([]byte(buildSignStr(m, apikey)))

	sign := hex.EncodeToString(h.Sum(nil))

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

// SignWithHMacSHA256 生成HMAC-SHA256签名
func SignWithHMacSHA256(m WXML, apikey string, toUpper bool) string {
	h := hmac.New(sha256.New, []byte(apikey))
	h.Write([]byte(buildSignStr(m, apikey)))

	sign := hex.EncodeToString(h.Sum(nil))

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

// Sign 生成签名
func buildSignStr(m WXML, apikey string) string {
	l := len(m)

	ks := make([]string, 0, l)
	kvs := make([]string, 0, l)

	for k := range m {
		if k == "sign" {
			continue
		}

		ks = append(ks, k)
	}

	sort.Strings(ks)

	for _, k := range ks {
		if v, ok := m[k]; ok && v != "" {
			kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
		}
	}

	kvs = append(kvs, fmt.Sprintf("key=%s", apikey))

	return strings.Join(kvs, "&")
}

// EncodeUint32ToBytes 把整数 uint32 格式化成 4 字节的网络字节序
func EncodeUint32ToBytes(i uint32) []byte {
	b := make([]byte, 4)

	b[0] = byte(i >> 24)
	b[1] = byte(i >> 16)
	b[2] = byte(i >> 8)
	b[3] = byte(i)

	return b
}

// DecodeBytesToUint32 从 4 字节的网络字节序里解析出整数 uint32
func DecodeBytesToUint32(b []byte) uint32 {
	if len(b) != 4 {
		return 0
	}

	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

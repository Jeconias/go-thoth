package models

import (
	"fmt"
	"github.com/leodido/go-urn"
	"net"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const fieldErrMsg = "Error: Validation of field '%s' failed on tag '%s'"

// FieldError contains all functions to get error details
type FieldError interface {
	Tag() string
	Field() string
}

// ValidationErrors is an array of FieldError's for use in custom error
// messages post validation.
type ValidationErrors []FieldError

type fieldError struct {
	tag   string
	field string
}

// Tag returns the validation tag that failed.
func (f *fieldError) Tag() string {
	return f.tag
}

// Field returns the fields name with the tag name taking precedence over the
// fields actual name.
func (f *fieldError) Field() string {
	return f.field
}

// Error returns the fieldError's error message
func (f *fieldError) Error() string {
	return fmt.Sprintf(fieldErrMsg, f.Field(), f.Tag())
}

var (
	// NewError TODO
	NewError = func(field, tag string) *fieldError {
		return &fieldError{
			tag:   tag,
			field: field,
		}
	}
)

/*
	Validators
*/

// Empty TODO
func Empty(v int) bool {
	return v == 0
}

// IsValid TODO
func IsValid(v interface{}) bool {
	return v == nil
}

/*
	Numbers
*/

// IsUint TODO
func IsUint(v uint) bool {
	return v == 0
}

// IsUint8 TODO
func IsUint8(v uint8) bool {
	return v == 0
}

// IsUint16 TODO
func IsUint16(v uint16) bool {
	return v == 0
}

// IsUint32 TODO
func IsUint32(v uint32) bool {
	return v == 0
}

// IsUint64 TODO
func IsUint64(v uint64) bool {
	return v == 0
}

// IsUintptr TODO
func IsUintptr(v uintptr) bool {
	return v == 0
}

// IsInt TODO
func IsInt(v int) bool {
	return v == 0
}

// IsInt8 TODO
func IsInt8(v int8) bool {
	return v == 0
}

// IsInt16 TODO
func IsInt16(v int16) bool {
	return v == 0
}

// IsInt32 TODO
func IsInt32(v int32) bool {
	return v == 0
}

// IsInt64 TODO
func IsInt64(v int64) bool {
	return v == 0
}

// IsFloat32 TODO
func IsFloat32(v float32) bool {
	return v == 0
}

// IsFloat64 TODO
func IsFloat64(v float64) bool {
	return v == 0
}

// IsComplex64 TODO
func IsComplex64(v complex64) bool {
	return v == 0
}

// IsComplex128 TODO
func IsComplex128(v complex128) bool {
	return v == 0
}
func isURL(ref string) bool {
	var i int
	s := ref

	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i = strings.Index(s, "#"); i > -1 {
		s = s[:i]
	}

	if len(s) == 0 {
		return false
	}

	url, err := url.ParseRequestURI(s)

	if err != nil || url.Scheme == "" {
		return false
	}

	return err == nil
}
func isURI(ref string) bool {
	s := ref

	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i := strings.Index(s, "#"); i > -1 {
		s = s[:i]
	}

	if len(s) == 0 {
		return false
	}

	_, err := url.ParseRequestURI(s)

	return err == nil
}
func isUrnRFC2141(ref string) bool {
	s := ref
	_, match := urn.Parse([]byte(s))
	return match
}
func isFile(ref string) bool {
	fileInfo, err := os.Stat(ref)
	if err != nil {
		return false
	}

	return !fileInfo.IsDir()
}

// IsDir is the validation function for validating if the current field's value is a valid directory.
func isDir(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
func isISBN10(ref string) bool {
	s := strings.Replace(strings.Replace(ref, "-", "", 3), " ", "", 3)

	if !iSBN10Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	for i = 0; i < 9; i++ {
		checksum += (i + 1) * int32(s[i]-'0')
	}

	if s[9] == 'X' {
		checksum += 10 * 10
	} else {
		checksum += 10 * int32(s[9]-'0')
	}

	return checksum%11 == 0
}
func isISBN13(ref string) bool {
	s := strings.Replace(strings.Replace(ref, "-", "", 4), " ", "", 4)

	if !iSBN13Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	factor := []int32{1, 3}

	for i = 0; i < 12; i++ {
		checksum += factor[i%2] * int32(s[i]-'0')
	}

	return (int32(s[12]-'0'))-((10-(checksum%10))%10) == 0
}
func isDataURI(ref string) bool {
	uri := strings.SplitN(ref, ",", 2)

	if len(uri) != 2 {
		return false
	}

	if !dataURIRegex.MatchString(uri[0]) {
		return false
	}

	return base64Regex.MatchString(uri[1])
}
func isSSN(ref string) bool {
	if len(ref) != 11 {
		return false
	}

	return sSNRegex.MatchString(ref)
}

// IsIP is the validation function for validating if the field's value is a valid v4 or v6 IP address.
func isIP(ref string) bool {
	ip := net.ParseIP(ref)

	return ip != nil
}

// IsIPv4 is the validation function for validating if a value is a valid v4 IP address.
func isIPv4(ref string) bool {
	ip := net.ParseIP(ref)

	return ip != nil && ip.To4() == nil
}

// IsIPv6 is the validation function for validating if the field's value is a valid v6 IP address.
func isIPv6(ref string) bool {
	ip := net.ParseIP(ref)

	// TODO
	return ip != nil && ip.To16() == nil
}

// IsCIDRv4 is the validation function for validating if the field's value is a valid v4 CIDR address.
func isCIDRv4(ref string) bool {
	ip, _, err := net.ParseCIDR(ref)

	return err == nil && ip.To4() != nil
}

// IsCIDRv6 is the validation function for validating if the field's value is a valid v6 CIDR address.
func isCIDRv6(ref string) bool {
	ip, _, err := net.ParseCIDR(ref)

	// TODO
	return err == nil && ip.To4() == nil
}

// IsCIDR is the validation function for validating if the field's value is a valid v4 or v6 CIDR address.
func isCIDR(ref string) bool {
	_, _, err := net.ParseCIDR(ref)

	return err == nil
}

// IsMAC is the validation function for validating if the field's value is a valid MAC address.
func isMAC(ref string) bool {
	_, err := net.ParseMAC(ref)

	return err == nil
}

// IsTCP4AddrResolvable is the validation function for validating if the field's value is a resolvable tcp4 address.
func isTCP4AddrResolvable(ref string) bool {

	if !isIP4Addr(ref) {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp4", ref)
	return err == nil
}

// IsTCP6AddrResolvable is the validation function for validating if the field's value is a resolvable tcp6 address.
func isTCP6AddrResolvable(ref string) bool {

	if !isIP6Addr(ref) {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp6", ref)

	return err == nil
}

// IsTCPAddrResolvable is the validation function for validating if the field's value is a resolvable tcp address.
func isTCPAddrResolvable(ref string) bool {

	if !isIP4Addr(ref) && !isIP6Addr(ref) {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp", ref)

	return err == nil
}

// IsUDP4AddrResolvable is the validation function for validating if the field's value is a resolvable udp4 address.
func isUDP4AddrResolvable(ref string) bool {

	if !isIP4Addr(ref) {
		return false
	}

	_, err := net.ResolveUDPAddr("udp4", ref)

	return err == nil
}

// IsUDP6AddrResolvable is the validation function for validating if the field's value is a resolvable udp6 address.
func isUDP6AddrResolvable(ref string) bool {

	if !isIP6Addr(ref) {
		return false
	}

	_, err := net.ResolveUDPAddr("udp6", ref)

	return err == nil
}

// IsUDPAddrResolvable is the validation function for validating if the field's value is a resolvable udp address.
func isUDPAddrResolvable(ref string) bool {

	if !isIP4Addr(ref) && !isIP6Addr(ref) {
		return false
	}

	_, err := net.ResolveUDPAddr("udp", ref)

	return err == nil
}

// IsIP4AddrResolvable is the validation function for validating if the field's value is a resolvable ip4 address.
func isIP4AddrResolvable(ref string) bool {

	if !isIPv4(ref) {
		return false
	}

	_, err := net.ResolveIPAddr("ip4", ref)

	return err == nil
}

// IsIP6AddrResolvable is the validation function for validating if the field's value is a resolvable ip6 address.
func isIP6AddrResolvable(ref string) bool {

	if !isIPv6(ref) {
		return false
	}

	_, err := net.ResolveIPAddr("ip6", ref)

	return err == nil
}

// IsIPAddrResolvable is the validation function for validating if the field's value is a resolvable ip address.
func isIPAddrResolvable(ref string) bool {

	if !isIP(ref) {
		return false
	}

	_, err := net.ResolveIPAddr("ip", ref)

	return err == nil
}

// IsUnixAddrResolvable is the validation function for validating if the field's value is a resolvable unix address.
func isUnixAddrResolvable(ref string) bool {

	_, err := net.ResolveUnixAddr("unix", ref)

	return err == nil
}

func isIP4Addr(ref string) bool {

	val := ref

	if idx := strings.LastIndex(val, ":"); idx != -1 {
		val = val[0:idx]
	}

	ip := net.ParseIP(val)

	return ip != nil && ip.To4() != nil
}

func isIP6Addr(ref string) bool {

	val := ref

	if idx := strings.LastIndex(val, ":"); idx != -1 {
		if idx != 0 && val[idx-1:idx] == "]" {
			val = val[1 : idx-1]
		}
	}

	ip := net.ParseIP(val)

	return ip != nil && ip.To4() == nil
}

func isFQDN(ref string) bool {
	val := ref

	if val == "" {
		return false
	}

	if val[len(val)-1] == '.' {
		val = val[0 : len(val)-1]
	}

	return strings.ContainsAny(val, ".") && hostnameRegexRFC952.MatchString(val)
}

const (
	alphaRegexString                 = "^[a-zA-Z]+$"
	alphaNumericRegexString          = "^[a-zA-Z0-9]+$"
	alphaUnicodeRegexString          = "^[\\p{L}]+$"
	alphaUnicodeNumericRegexString   = "^[\\p{L}\\p{N}]+$"
	numericRegexString               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString                = "^[0-9]+$"
	hexadecimalRegexString           = "^[0-9a-fA-F]+$"
	hexcolorRegexString              = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	rgbRegexString                   = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaRegexString                  = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslRegexString                   = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegexString                  = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	emailRegexString                 = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	base64RegexString                = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLRegexString             = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	iSBN10RegexString                = "^(?:[0-9]{9}X|[0-9]{10})$"
	iSBN13RegexString                = "^(?:(?:97(?:8|9))[0-9]{10})$"
	uUID3RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID4RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID5RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUIDRegexString                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID3RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uUID4RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUID5RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUIDRFC4122RegexString           = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	aSCIIRegexString                 = "^[\x00-\x7F]*$"
	printableASCIIRegexString        = "^[\x20-\x7E]*$"
	multibyteRegexString             = "[^\x00-\x7F]"
	dataURIRegexString               = "^data:.+\\/(.+);base64$"
	latitudeRegexString              = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegexString             = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	sSNRegexString                   = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	hostnameRegexStringRFC952        = `^[a-zA-Z][a-zA-Z0-9\-\.]+[a-zA-Z0-9]$`    // https://tools.ietf.org/html/rfc952
	hostnameRegexStringRFC1123       = `^[a-zA-Z0-9][a-zA-Z0-9\-\.]+[a-zA-Z0-9]$` // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	btcAddressRegexString            = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`        // bitcoin address
	btcAddressUpperRegexStringBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`              // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	btcAddressLowerRegexStringBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`              // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ethAddressRegexString            = `^0x[0-9a-fA-F]{40}$`
	ethAddressUpperRegexString       = `^0x[0-9A-F]{40}$`
	ethAddressLowerRegexString       = `^0x[0-9a-f]{40}$`
	uRLEncodedRegexString            = `(%[A-Fa-f0-9]{2})`
	hTMLEncodedRegexString           = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	hTMLRegexString                  = `<[/]?([a-zA-Z]+).*?>`
)

var (
	alphaRegex                 = regexp.MustCompile(alphaRegexString)
	alphaNumericRegex          = regexp.MustCompile(alphaNumericRegexString)
	alphaUnicodeRegex          = regexp.MustCompile(alphaUnicodeRegexString)
	alphaUnicodeNumericRegex   = regexp.MustCompile(alphaUnicodeNumericRegexString)
	numericRegex               = regexp.MustCompile(numericRegexString)
	numberRegex                = regexp.MustCompile(numberRegexString)
	hexadecimalRegex           = regexp.MustCompile(hexadecimalRegexString)
	hexcolorRegex              = regexp.MustCompile(hexcolorRegexString)
	rgbRegex                   = regexp.MustCompile(rgbRegexString)
	rgbaRegex                  = regexp.MustCompile(rgbaRegexString)
	hslRegex                   = regexp.MustCompile(hslRegexString)
	hslaRegex                  = regexp.MustCompile(hslaRegexString)
	emailRegex                 = regexp.MustCompile(emailRegexString)
	base64Regex                = regexp.MustCompile(base64RegexString)
	base64URLRegex             = regexp.MustCompile(base64URLRegexString)
	iSBN10Regex                = regexp.MustCompile(iSBN10RegexString)
	iSBN13Regex                = regexp.MustCompile(iSBN13RegexString)
	uUID3Regex                 = regexp.MustCompile(uUID3RegexString)
	uUID4Regex                 = regexp.MustCompile(uUID4RegexString)
	uUID5Regex                 = regexp.MustCompile(uUID5RegexString)
	uUIDRegex                  = regexp.MustCompile(uUIDRegexString)
	uUID3RFC4122Regex          = regexp.MustCompile(uUID3RFC4122RegexString)
	uUID4RFC4122Regex          = regexp.MustCompile(uUID4RFC4122RegexString)
	uUID5RFC4122Regex          = regexp.MustCompile(uUID5RFC4122RegexString)
	uUIDRFC4122Regex           = regexp.MustCompile(uUIDRFC4122RegexString)
	aSCIIRegex                 = regexp.MustCompile(aSCIIRegexString)
	printableASCIIRegex        = regexp.MustCompile(printableASCIIRegexString)
	multibyteRegex             = regexp.MustCompile(multibyteRegexString)
	dataURIRegex               = regexp.MustCompile(dataURIRegexString)
	latitudeRegex              = regexp.MustCompile(latitudeRegexString)
	longitudeRegex             = regexp.MustCompile(longitudeRegexString)
	sSNRegex                   = regexp.MustCompile(sSNRegexString)
	hostnameRegexRFC952        = regexp.MustCompile(hostnameRegexStringRFC952)
	hostnameRegexRFC1123       = regexp.MustCompile(hostnameRegexStringRFC1123)
	btcAddressRegex            = regexp.MustCompile(btcAddressRegexString)
	btcUpperAddressRegexBech32 = regexp.MustCompile(btcAddressUpperRegexStringBech32)
	btcLowerAddressRegexBech32 = regexp.MustCompile(btcAddressLowerRegexStringBech32)
	ethAddressRegex            = regexp.MustCompile(ethAddressRegexString)
	ethaddressRegexUpper       = regexp.MustCompile(ethAddressUpperRegexString)
	ethAddressRegexLower       = regexp.MustCompile(ethAddressLowerRegexString)
	uRLEncodedRegex            = regexp.MustCompile(uRLEncodedRegexString)
	hTMLEncodedRegex           = regexp.MustCompile(hTMLEncodedRegexString)
	hTMLRegex                  = regexp.MustCompile(hTMLRegexString)
)

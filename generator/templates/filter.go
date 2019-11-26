package templates

import (
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/validation"
)

func hasTag(structsThoth []*myasthurts.Struct) bool {
	for _, s := range structsThoth {
		for _, field := range s.Fields {
			if len(field.Tag.Params) > 0 {
				return true
			}
		}
	}
	return false
}

// FilterInput TODO
type FilterInput struct {
	Struct    *myasthurts.Struct
	StructRef string
	Field     *myasthurts.Field
	Tag       myasthurts.TagParam
	Ref       string
}

func filterValidate(_buffer io.StringWriter, input *FilterInput, args ...string) {
	switch input.Tag.Value {
	case "-":
		// Skip field...
	case "required":
		validation.HasValue(_buffer, &validation.RequiredInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "required_with":
		validation.RequiredWith(_buffer, &validation.RequiredWithInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "required_with_all":
		validation.RequiredWithAll(_buffer, &validation.RequiredWithAllInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "required_without":
		validation.RequiredWithout(_buffer, &validation.RequiredWithoutInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "required_without_all":
		validation.RequiredWithoutAll(_buffer, &validation.RequiredWithoutAllInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "len":
		if len(args) == 1 {
			validation.HasLengthOf(_buffer, &validation.HasLengthOfInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			})
		}
	case "eq":
		if len(args) == 1 {
			validation.IsEq(_buffer, &validation.IsEqInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			})
		}
	case "min":
		if len(args) == 1 {
			validation.HasMinOf(_buffer, &validation.HasMinOfInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			})
		}
	case "max":
		if len(args) == 1 {
			validation.HasMaxOf(_buffer, &validation.HasMaxOfInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			})
		}
	case "ne":
		if len(args) == 1 {
			validation.IsNe(_buffer, &validation.IsNeInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			})
		}
	case "lt":
		if len(args) == 1 {
			validation.IsCompare(_buffer, &validation.IsCompareInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			}, ">", "lt")
		}
	case "lte":
		if len(args) == 1 {
			validation.IsCompare(_buffer, &validation.IsCompareInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			}, ">=", "lte")
		}
	case "gt":
		if len(args) == 1 {
			validation.IsCompare(_buffer, &validation.IsCompareInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			}, "<", "gt")
		}
	case "gte":
		if len(args) == 1 {
			validation.IsCompare(_buffer, &validation.IsCompareInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			}, "<=", "gte")
		}
	case "alpha":
		validation.IsAlpha(_buffer, &validation.IsAlphaInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "alphanum":
		validation.IsAlphanum(_buffer, &validation.IsAlphanumInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "alphaunicode":
		validation.IsAlphaunicode(_buffer, &validation.IsAlphaunicodeInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "alphanumunicode":
		validation.IsAlphanumunicode(_buffer, &validation.IsAlphanumunicodeInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "numeric":
		validation.IsNumeric(_buffer, &validation.IsNumericInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "number":
		validation.IsNumber(_buffer, &validation.IsNumberInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "hexadecimal":
		validation.IsHexadecimal(_buffer, &validation.IsHexadecimalInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "hexcolor":
		validation.IsHexcolor(_buffer, &validation.IsHexcolorInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "rgb":
		validation.IsRGB(_buffer, &validation.IsRGBInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "rgba":
		validation.IsRGBa(_buffer, &validation.IsRGBaInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "hsl":
		validation.IsHSL(_buffer, &validation.IsHSLInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "hsla":
		validation.IsHSLA(_buffer, &validation.IsHSLAInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "email":
		validation.IsEmail(_buffer, &validation.IsEmailInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "url":
		validation.IsURL(_buffer, &validation.IsURLInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uri":
		validation.IsURI(_buffer, &validation.IsURIInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "urn_rfc2141":
		validation.IsUrnRFC2141(_buffer, &validation.IsUrnRFC2141Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "file":
		validation.IsFile(_buffer, &validation.IsFileInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "base64":
		validation.IsBase64(_buffer, &validation.IsBase64Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "base64url":
		validation.IsBase64URL(_buffer, &validation.IsBase64URLInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "isbn":
		validation.IsISBN(_buffer, &validation.IsISBNInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "isbn10":
		validation.IsISBN10(_buffer, &validation.IsISBN10Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "isbn13":
		validation.IsISBN13(_buffer, &validation.IsISBN13Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "eth_addr":
		validation.IsEthereumAddress(_buffer, &validation.IsEthereumAddressInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "btc_addr":
		validation.IsBitcoinAddress(_buffer, &validation.IsBitcoinAddressInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "btc_addr_bech32":
		validation.IsBitcoinAddress(_buffer, &validation.IsBitcoinAddressInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid":
		validation.IsUUID(_buffer, &validation.IsUUIDInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid3":
		validation.IsUUID3(_buffer, &validation.IsUUID3Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid4":
		validation.IsUUID4(_buffer, &validation.IsUUID4Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid5":
		validation.IsUUID5(_buffer, &validation.IsUUID5Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid_rfc4122":
		validation.IsUUIDRFC4122(_buffer, &validation.IsUUIDRFC4122Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid3_rfc4122":
		validation.IsUUID3RFC4122(_buffer, &validation.IsUUID3RFC4122Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid4_rfc4122":
		validation.IsUUID4RFC4122(_buffer, &validation.IsUUID4RFC4122Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "uuid5_rfc4122":
		validation.IsUUID5RFC4122(_buffer, &validation.IsUUID5RFC4122Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "ascii":
		validation.IsASCII(_buffer, &validation.IsASCIIInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "printascii":
		validation.IsPrintableASCII(_buffer, &validation.IsPrintableASCIIInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "multibyte":
		validation.HasMultiByteCharacter(_buffer, &validation.HasMultiByteCharacterInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "datauri":
		validation.IsDataURI(_buffer, &validation.IsDataURIInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "ssn":
		validation.IsSSN(_buffer, &validation.IsSSNInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "ip":
		validation.IsIP(_buffer, &validation.IsIPInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "ipv4":
		validation.IsIP(_buffer, &validation.IsIPInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "v4")
	case "ipv6":
		validation.IsIP(_buffer, &validation.IsIPInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "v6")
	case "cidr":
		validation.IsCIDR(_buffer, &validation.IsCIDRInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "cidrv4":
		validation.IsCIDR(_buffer, &validation.IsCIDRInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "v4")
	case "cidrv6":
		validation.IsCIDR(_buffer, &validation.IsCIDRInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "v6")
	case "tcp4_addr":
		validation.IsTCPAddrResolvable(_buffer, &validation.IsTCPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "4")
	case "tcp6_addr":
		validation.IsTCPAddrResolvable(_buffer, &validation.IsTCPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "6")
	case "tcp_addr":
		validation.IsTCPAddrResolvable(_buffer, &validation.IsTCPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "udp4_addr":
		validation.IsUDPAddrResolvable(_buffer, &validation.IsUDPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "4")
	case "udp6_addr":
		validation.IsUDPAddrResolvable(_buffer, &validation.IsUDPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "6")
	case "udp_addr":
		validation.IsUDPAddrResolvable(_buffer, &validation.IsUDPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "ip4_addr":
		validation.IsIPAddrResolvable(_buffer, &validation.IsIPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "4")
	case "ip6_addr":
		validation.IsIPAddrResolvable(_buffer, &validation.IsIPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		}, "6")
	case "ip_addr":
		validation.IsIPAddrResolvable(_buffer, &validation.IsIPAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "unix_addr":
		validation.IsUnixAddrResolvable(_buffer, &validation.IsUnixAddrResolvableInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "mac":
		validation.IsMac(_buffer, &validation.IsMacInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "hostname":
		validation.IsHostnameRFC952(_buffer, &validation.IsHostnameRFC952Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "hostname_rfc1123":
		validation.IsHostnameRFC1123(_buffer, &validation.IsHostnameRFC1123Input{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "fqdn":
		validation.IsFQDN(_buffer, &validation.IsFQDNInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "html":
		validation.IsHTML(_buffer, &validation.IsHTMLInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "html_encoded":
		validation.IsHTMLEncoded(_buffer, &validation.IsHTMLEncodedInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "url_encoded":
		validation.IsURLEncoded(_buffer, &validation.IsURLEncodedInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "dir":
		validation.IsDir(_buffer, &validation.IsDirInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	default:
		// TODO: (@edumarcal) to fix not match rule
		k, v := splitArgs(input.Tag)
		input.Tag.Value = k
		filterValidate(_buffer, input, v...)
	}
}

func splitArgs(tag myasthurts.TagParam) (string, []string) {
	args := strings.Split(tag.Value, "=")
	switch len(args) {
	case 1:
		return args[0], args[1:]
	case 2:
		v := strings.Split(args[1], " ")
		return args[0], v
	default:
		return "", nil
	}
}

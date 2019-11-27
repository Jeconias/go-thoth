package models

// Validate TODO
func (a *ASCIIValidate) Validate() (errs ValidationErrors) {

	if !aSCIIRegex.MatchString(a.ASCII) {
		errs = append(errs, NewError("ASCII", "ascii"))
	}

	if a.Pointer == nil || !aSCIIRegex.MatchString(*a.Pointer) {
		errs = append(errs, NewError("Pointer", "ascii"))
	}

	for _, v := range a.Slice {
		if !aSCIIRegex.MatchString(v) {
			errs = append(errs, NewError("Slice", "ascii"))
		}
	}
	return errs
}

// Validate TODO
func (b *Base64Validate) Validate() (errs ValidationErrors) {

	if !base64Regex.MatchString(b.Base64) {
		errs = append(errs, NewError("Base64", "base64"))
	}

	if b.Pointer == nil || !base64Regex.MatchString(*b.Pointer) {
		errs = append(errs, NewError("Pointer", "base64"))
	}

	for _, v := range b.Slice {
		if !base64Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "base64"))
		}
	}
	return errs
}

// Validate TODO
func (b *Base64urlValidate) Validate() (errs ValidationErrors) {

	if !base64URLRegex.MatchString(b.Base64url) {
		errs = append(errs, NewError("Base64url", "base64url"))
	}

	if b.Pointer == nil || !base64URLRegex.MatchString(*b.Pointer) {
		errs = append(errs, NewError("Pointer", "base64url"))
	}

	for _, v := range b.Slice {
		if !base64URLRegex.MatchString(v) {
			errs = append(errs, NewError("Slice", "base64url"))
		}
	}
	return errs
}

// Validate TODO
func (c *CIDRValidate) Validate() (errs ValidationErrors) {

	if !isCIDR(c.CIDR) {
		errs = append(errs, NewError("CIDR", "cidr"))
	}

	if c.Pointer == nil || !isCIDR(*c.Pointer) {
		errs = append(errs, NewError("Pointer", "cidr"))
	}

	for _, v := range c.Slice {
		if !isCIDR(v) {
			errs = append(errs, NewError("Slice", "cidr"))
		}
	}
	return errs
}

// Validate TODO
func (c *CIDRv4Validate) Validate() (errs ValidationErrors) {

	if !isCIDRv4(c.CIDRv4) {
		errs = append(errs, NewError("CIDRv4", "cidrv4"))
	}

	if c.Pointer == nil || !isCIDRv4(*c.Pointer) {
		errs = append(errs, NewError("Pointer", "cidrv4"))
	}

	for _, v := range c.Slice {
		if !isCIDRv4(v) {
			errs = append(errs, NewError("Slice", "cidrv4"))
		}
	}
	return errs
}

// Validate TODO
func (c *CIDRv6Validate) Validate() (errs ValidationErrors) {

	if !isCIDRv6(c.CIDRv6) {
		errs = append(errs, NewError("CIDRv6", "cidrv6"))
	}

	if c.Pointer == nil || !isCIDRv6(*c.Pointer) {
		errs = append(errs, NewError("Pointer", "cidrv6"))
	}

	for _, v := range c.Slice {
		if !isCIDRv6(v) {
			errs = append(errs, NewError("Slice", "cidrv6"))
		}
	}
	return errs
}

// Validate TODO
func (d *DirValidate) Validate() (errs ValidationErrors) {

	if !isDir(d.Dir) {
		errs = append(errs, NewError("Dir", "dir"))
	}

	if d.Pointer == nil || !isDir(*d.Pointer) {
		errs = append(errs, NewError("Pointer", "dir"))
	}

	for _, v := range d.Slice {
		if !isDir(v) {
			errs = append(errs, NewError("Slice", "dir"))
		}
	}
	return errs
}

// Validate TODO
func (d *DiveValidate) Validate() (errs ValidationErrors) {
	for _, v := range d.SliceString {
		if Empty(len(v)) {
			errs = append(errs, NewError("SliceString", "dive"))
		}
	}

	for _, v := range d.SliceInt {
		if IsInt(v) {
			errs = append(errs, NewError("SliceInt", "dive"))
		}
	}

	for _, v := range d.MapStringToString {
		if Empty(len(v)) {
			errs = append(errs, NewError("MapStringToString", "dive"))
		}
	}
	return errs
}

// Validate TODO
func (e *EmailValidate) Validate() (errs ValidationErrors) {

	if !emailRegex.MatchString(e.Email) {
		errs = append(errs, NewError("Email", "email"))
	}

	if e.Pointer == nil || !emailRegex.MatchString(*e.Pointer) {
		errs = append(errs, NewError("Pointer", "email"))
	}
	return errs
}

// Validate TODO
func (e *EmailSliceValidate) Validate() (errs ValidationErrors) {
	for _, v := range e.Emails {
		if !emailRegex.MatchString(v) {
			errs = append(errs, NewError("Emails", "email"))
		}
	}
	return errs
}

// Validate TODO
func (e *EqNumber) Validate() (errs ValidationErrors) {

	if e.Uint != 1 {
		errs = append(errs, NewError("Uint", "eq"))
	}

	if e.UintPointer == nil || *e.UintPointer != 2 {
		errs = append(errs, NewError("UintPointer", "eq"))
	}

	if e.Uint8 != 3 {
		errs = append(errs, NewError("Uint8", "eq"))
	}

	if e.Uint8Pointer == nil || *e.Uint8Pointer != 4 {
		errs = append(errs, NewError("Uint8Pointer", "eq"))
	}

	if e.Uint16 != 5 {
		errs = append(errs, NewError("Uint16", "eq"))
	}

	if e.Uint16Pointer == nil || *e.Uint16Pointer != 6 {
		errs = append(errs, NewError("Uint16Pointer", "eq"))
	}

	if e.Uint32 != 7 {
		errs = append(errs, NewError("Uint32", "eq"))
	}

	if e.Uint32Pointer == nil || *e.Uint32Pointer != 8 {
		errs = append(errs, NewError("Uint32Pointer", "eq"))
	}

	if e.Uint64 != 9 {
		errs = append(errs, NewError("Uint64", "eq"))
	}

	if e.Uint64Pointer == nil || *e.Uint64Pointer != 10 {
		errs = append(errs, NewError("Uint64Pointer", "eq"))
	}

	if e.Uintptr != 11 {
		errs = append(errs, NewError("Uintptr", "eq"))
	}

	if e.UintptrPointer == nil || *e.UintptrPointer != 12 {
		errs = append(errs, NewError("UintptrPointer", "eq"))
	}

	if e.Int != 13 {
		errs = append(errs, NewError("Int", "eq"))
	}

	if e.IntPointer == nil || *e.IntPointer != 14 {
		errs = append(errs, NewError("IntPointer", "eq"))
	}

	if e.Int8 != 15 {
		errs = append(errs, NewError("Int8", "eq"))
	}

	if e.Int8Pointer == nil || *e.Int8Pointer != 16 {
		errs = append(errs, NewError("Int8Pointer", "eq"))
	}

	if e.Int16 != 17 {
		errs = append(errs, NewError("Int16", "eq"))
	}

	if e.Int16Pointer == nil || *e.Int16Pointer != 18 {
		errs = append(errs, NewError("Int16Pointer", "eq"))
	}

	if e.Int32 != 19 {
		errs = append(errs, NewError("Int32", "eq"))
	}

	if e.Int32Pointer == nil || *e.Int32Pointer != 20 {
		errs = append(errs, NewError("Int32Pointer", "eq"))
	}

	if e.Int64 != 21 {
		errs = append(errs, NewError("Int64", "eq"))
	}

	if e.Int64Pointer == nil || *e.Int64Pointer != 22 {
		errs = append(errs, NewError("Int64Pointer", "eq"))
	}

	if e.Float32 != 23 {
		errs = append(errs, NewError("Float32", "eq"))
	}

	if e.Float32Pointer == nil || *e.Float32Pointer != 24 {
		errs = append(errs, NewError("Float32Pointer", "eq"))
	}

	if e.Float64 != 25 {
		errs = append(errs, NewError("Float64", "eq"))
	}

	if e.Float64Pointer == nil || *e.Float64Pointer != 26 {
		errs = append(errs, NewError("Float64Pointer", "eq"))
	}

	if e.Complex64 != 27 {
		errs = append(errs, NewError("Complex64", "eq"))
	}

	if e.Complex64Pointer == nil || *e.Complex64Pointer != 28 {
		errs = append(errs, NewError("Complex64Pointer", "eq"))
	}

	if e.Complex128 != 29 {
		errs = append(errs, NewError("Complex128", "eq"))
	}

	if e.Complex128Pointer == nil || *e.Complex128Pointer != 30 {
		errs = append(errs, NewError("Complex128Pointer", "eq"))
	}
	return errs
}

// Validate TODO
func (e *EqString) Validate() (errs ValidationErrors) {

	if e.String != "chico" {
		errs = append(errs, NewError("String", "eq"))
	}

	if e.Pointer == nil || *e.Pointer != "bento" {
		errs = append(errs, NewError("Pointer", "eq"))
	}
	return errs
}

// Validate TODO
func (f *FileValidate) Validate() (errs ValidationErrors) {

	if !isFile(f.File) {
		errs = append(errs, NewError("File", "file"))
	}

	if f.Pointer == nil || !isFile(*f.Pointer) {
		errs = append(errs, NewError("Pointer", "file"))
	}

	for _, v := range f.Slice {
		if !isFile(v) {
			errs = append(errs, NewError("Slice", "file"))
		}
	}
	return errs
}

// Validate TODO
func (f *FQDNValidate) Validate() (errs ValidationErrors) {

	if !isFQDN(f.FQDN) {
		errs = append(errs, NewError("FQDN", "fqdn"))
	}

	if f.Pointer == nil || !isFQDN(*f.Pointer) {
		errs = append(errs, NewError("Pointer", "fqdn"))
	}

	for _, v := range f.Slice {
		if !isFQDN(v) {
			errs = append(errs, NewError("Slice", "fqdn"))
		}
	}
	return errs
}

// Validate TODO
func (g *GtValidate) Validate() (errs ValidationErrors) {

	if len(g.Name) < 5 {
		errs = append(errs, NewError("Name", "gt"))
	}

	if g.Password == nil || len(*g.Password) < 3 {
		errs = append(errs, NewError("Password", "gt"))
	}

	if g.Age < 22 {
		errs = append(errs, NewError("Age", "gt"))
	}

	if len(g.Contents) < 2 {
		errs = append(errs, NewError("Contents", "gt"))
	}
	return errs
}

// Validate TODO
func (g *GteValidate) Validate() (errs ValidationErrors) {

	if len(g.Name) <= 5 {
		errs = append(errs, NewError("Name", "gte"))
	}

	if g.Password == nil || len(*g.Password) <= 3 {
		errs = append(errs, NewError("Password", "gte"))
	}

	if g.Age <= 22 {
		errs = append(errs, NewError("Age", "gte"))
	}

	if len(g.Contents) <= 2 {
		errs = append(errs, NewError("Contents", "gte"))
	}
	return errs
}

// Validate TODO
func (h *HexadecimalValidate) Validate() (errs ValidationErrors) {

	if !hexadecimalRegex.MatchString(h.Hexadecimal) {
		errs = append(errs, NewError("Hexadecimal", "hexadecimal"))
	}

	if h.Pointer == nil || !hexadecimalRegex.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "hexadecimal"))
	}

	for _, v := range h.Slice {
		if !hexadecimalRegex.MatchString(v) {
			errs = append(errs, NewError("Slice", "hexadecimal"))
		}
	}
	return errs
}

// Validate TODO
func (h *HexcolorValidate) Validate() (errs ValidationErrors) {

	if !hexcolorRegex.MatchString(h.Hexcolor) {
		errs = append(errs, NewError("Hexcolor", "hexcolor"))
	}

	if h.Pointer == nil || !hexcolorRegex.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "hexcolor"))
	}
	return errs
}

// Validate TODO
func (h *HostnameRFC952Validate) Validate() (errs ValidationErrors) {

	if !hostnameRegexRFC952.MatchString(h.HostnameRFC952) {
		errs = append(errs, NewError("HostnameRFC952", "hostname"))
	}

	if h.Pointer == nil || !hostnameRegexRFC952.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "hostname"))
	}

	for _, v := range h.Slice {
		if !hostnameRegexRFC952.MatchString(v) {
			errs = append(errs, NewError("Slice", "hostname"))
		}
	}
	return errs
}

// Validate TODO
func (h *HostnameRFC1123Validate) Validate() (errs ValidationErrors) {

	if !hostnameRegexRFC1123.MatchString(h.HostnameRFC1123) {
		errs = append(errs, NewError("HostnameRFC1123", "hostname_rfc1123"))
	}

	if h.Pointer == nil || !hostnameRegexRFC1123.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "hostname_rfc1123"))
	}

	for _, v := range h.Slice {
		if !hostnameRegexRFC1123.MatchString(v) {
			errs = append(errs, NewError("Slice", "hostname_rfc1123"))
		}
	}
	return errs
}

// Validate TODO
func (h *HSLValidate) Validate() (errs ValidationErrors) {

	if !hslRegex.MatchString(h.HSL) {
		errs = append(errs, NewError("HSL", "hsl"))
	}

	if h.Pointer == nil || !hslRegex.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "hsl"))
	}
	return errs
}

// Validate TODO
func (h *HSLAValidate) Validate() (errs ValidationErrors) {

	if !hslaRegex.MatchString(h.HSLA) {
		errs = append(errs, NewError("HSLA", "hsla"))
	}

	if h.Pointer == nil || !hslaRegex.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "hsla"))
	}
	return errs
}

// Validate TODO
func (h *HTMLValidate) Validate() (errs ValidationErrors) {

	if !hTMLRegex.MatchString(h.HTML) {
		errs = append(errs, NewError("HTML", "html"))
	}

	if h.Pointer == nil || !hTMLRegex.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "html"))
	}

	for _, v := range h.Slice {
		if !hTMLRegex.MatchString(v) {
			errs = append(errs, NewError("Slice", "html"))
		}
	}
	return errs
}

// Validate TODO
func (h *HTMLEncodedValidate) Validate() (errs ValidationErrors) {

	if !hTMLEncodedRegex.MatchString(h.HTMLEncoded) {
		errs = append(errs, NewError("HTMLEncoded", "html_encoded"))
	}

	if h.Pointer == nil || !hTMLEncodedRegex.MatchString(*h.Pointer) {
		errs = append(errs, NewError("Pointer", "html_encoded"))
	}

	for _, v := range h.Slice {
		if !hTMLEncodedRegex.MatchString(v) {
			errs = append(errs, NewError("Slice", "html_encoded"))
		}
	}
	return errs
}

// Validate TODO
func (i *IPValidate) Validate() (errs ValidationErrors) {

	if !isIP(i.IP) {
		errs = append(errs, NewError("IP", "ip"))
	}

	if i.Pointer == nil || !isIP(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "ip"))
	}

	for _, v := range i.Slice {
		if !isIP(v) {
			errs = append(errs, NewError("Slice", "ip"))
		}
	}
	return errs
}

// Validate TODO
func (i *IPv4Validate) Validate() (errs ValidationErrors) {

	if !isIPv4(i.IPv4) {
		errs = append(errs, NewError("IPv4", "ipv4"))
	}

	if i.Pointer == nil || !isIPv4(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "ipv4"))
	}

	for _, v := range i.Slice {
		if !isIPv4(v) {
			errs = append(errs, NewError("Slice", "ipv4"))
		}
	}
	return errs
}

// Validate TODO
func (i *IPv6Validate) Validate() (errs ValidationErrors) {

	if !isIPv6(i.IPv6) {
		errs = append(errs, NewError("IPv6", "ipv6"))
	}

	if i.Pointer == nil || !isIPv6(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "ipv6"))
	}

	for _, v := range i.Slice {
		if !isIPv6(v) {
			errs = append(errs, NewError("Slice", "ipv6"))
		}
	}
	return errs
}

// Validate TODO
func (i *IPAddrValidate) Validate() (errs ValidationErrors) {

	if !isIPAddrResolvable(i.IPAddr) {
		errs = append(errs, NewError("IPAddr", "ip_addr"))
	}

	if i.Pointer == nil || !isIPAddrResolvable(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "ip_addr"))
	}

	for _, v := range i.Slice {
		if !isIPAddrResolvable(v) {
			errs = append(errs, NewError("Slice", "ip_addr"))
		}
	}
	return errs
}

// Validate TODO
func (i *IP4AddrValidate) Validate() (errs ValidationErrors) {

	if !isIP4AddrResolvable(i.IP4Addr) {
		errs = append(errs, NewError("IP4Addr", "ip4_addr"))
	}

	if i.Pointer == nil || !isIP4AddrResolvable(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "ip4_addr"))
	}

	for _, v := range i.Slice {
		if !isIP4AddrResolvable(v) {
			errs = append(errs, NewError("Slice", "ip4_addr"))
		}
	}
	return errs
}

// Validate TODO
func (i *IP6AddrValidate) Validate() (errs ValidationErrors) {

	if !isIP6AddrResolvable(i.IP6Addr) {
		errs = append(errs, NewError("IP6Addr", "ip6_addr"))
	}

	if i.Pointer == nil || !isIP6AddrResolvable(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "ip6_addr"))
	}

	for _, v := range i.Slice {
		if !isIP6AddrResolvable(v) {
			errs = append(errs, NewError("Slice", "ip6_addr"))
		}
	}
	return errs
}

// Validate TODO
func (i *ISBNValidate) Validate() (errs ValidationErrors) {

	if !(isISBN10(i.ISBN) || isISBN13(i.ISBN)) {
		errs = append(errs, NewError("ISBN", "isbn"))
	}

	if i.Pointer == nil || !(isISBN10(*i.Pointer) || isISBN13(*i.Pointer)) {
		errs = append(errs, NewError("Pointer", "isbn"))
	}

	for _, v := range i.Slice {
		if !(isISBN10(v) || isISBN13(v)) {
			errs = append(errs, NewError("Slice", "isbn"))
		}
	}
	return errs
}

// Validate TODO
func (i *ISBN10Validate) Validate() (errs ValidationErrors) {

	if !isISBN10(i.ISBN10) {
		errs = append(errs, NewError("ISBN10", "isbn10"))
	}

	if i.Pointer == nil || !isISBN10(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "isbn10"))
	}

	for _, v := range i.Slice {
		if !isISBN10(v) {
			errs = append(errs, NewError("Slice", "isbn10"))
		}
	}
	return errs
}

// Validate TODO
func (i *ISBN13Validate) Validate() (errs ValidationErrors) {

	if !isISBN13(i.ISBN13) {
		errs = append(errs, NewError("ISBN13", "isbn13"))
	}

	if i.Pointer == nil || !isISBN13(*i.Pointer) {
		errs = append(errs, NewError("Pointer", "isbn13"))
	}

	for _, v := range i.Slice {
		if !isISBN13(v) {
			errs = append(errs, NewError("Slice", "isbn13"))
		}
	}
	return errs
}

// Validate TODO
func (l *LenStringValidate) Validate() (errs ValidationErrors) {

	if len(l.String) != 5 {
		errs = append(errs, NewError("String", "len"))
	}

	if l.Pointer == nil || len(*l.Pointer) != 5 {
		errs = append(errs, NewError("Pointer", "len"))
	}
	return errs
}

// Validate TODO
func (l *LenChanValidate) Validate() (errs ValidationErrors) {

	if len(l.ChanString) != 2 {
		errs = append(errs, NewError("ChanString", "len"))
	}

	if len(l.ChanStringPointer) != 2 {
		errs = append(errs, NewError("ChanStringPointer", "len"))
	}

	if len(l.ChanUint) != 2 {
		errs = append(errs, NewError("ChanUint", "len"))
	}

	if len(l.ChanUintPointer) != 2 {
		errs = append(errs, NewError("ChanUintPointer", "len"))
	}
	return errs
}

// Validate TODO
func (l *LenSliceStringValidate) Validate() (errs ValidationErrors) {

	if len(l.String) != 5 {
		errs = append(errs, NewError("String", "len"))
	}

	if len(l.SlicePointer) != 5 {
		errs = append(errs, NewError("SlicePointer", "len"))
	}

	if l.PointerSlice == nil || len(*l.PointerSlice) != 5 {
		errs = append(errs, NewError("PointerSlice", "len"))
	}

	if l.PointerSlicePointer == nil || len(*l.PointerSlicePointer) != 5 {
		errs = append(errs, NewError("PointerSlicePointer", "len"))
	}
	return errs
}

// Validate TODO
func (l *LenSliceIntValidate) Validate() (errs ValidationErrors) {

	if len(l.Int) != 5 {
		errs = append(errs, NewError("Int", "len"))
	}

	if len(l.SlicePointer) != 5 {
		errs = append(errs, NewError("SlicePointer", "len"))
	}

	if l.PointerSlice == nil || len(*l.PointerSlice) != 5 {
		errs = append(errs, NewError("PointerSlice", "len"))
	}

	if l.PointerSlicePointer == nil || len(*l.PointerSlicePointer) != 5 {
		errs = append(errs, NewError("PointerSlicePointer", "len"))
	}
	return errs
}

// Validate TODO
func (l *LenSliceFloat64Validate) Validate() (errs ValidationErrors) {

	if len(l.Float64) != 5 {
		errs = append(errs, NewError("Float64", "len"))
	}

	if len(l.SlicePointer) != 5 {
		errs = append(errs, NewError("SlicePointer", "len"))
	}

	if l.PointerSlice == nil || len(*l.PointerSlice) != 5 {
		errs = append(errs, NewError("PointerSlice", "len"))
	}

	if l.PointerSlicePointer == nil || len(*l.PointerSlicePointer) != 5 {
		errs = append(errs, NewError("PointerSlicePointer", "len"))
	}
	return errs
}

// Validate TODO
func (l *LtValidate) Validate() (errs ValidationErrors) {

	if len(l.Name) > 12 {
		errs = append(errs, NewError("Name", "lt"))
	}

	if l.Password == nil || len(*l.Password) > 3 {
		errs = append(errs, NewError("Password", "lt"))
	}

	if l.Age > 22 {
		errs = append(errs, NewError("Age", "lt"))
	}

	if len(l.Contents) > 2 {
		errs = append(errs, NewError("Contents", "lt"))
	}
	return errs
}

// Validate TODO
func (l *LteValidate) Validate() (errs ValidationErrors) {

	if len(l.Name) >= 12 {
		errs = append(errs, NewError("Name", "lte"))
	}

	if l.Password == nil || len(*l.Password) >= 3 {
		errs = append(errs, NewError("Password", "lte"))
	}

	if l.Age >= 22 {
		errs = append(errs, NewError("Age", "lte"))
	}

	if len(l.Contents) >= 2 {
		errs = append(errs, NewError("Contents", "lte"))
	}
	return errs
}

// Validate TODO
func (m *MACValidate) Validate() (errs ValidationErrors) {

	if !isMAC(m.MAC) {
		errs = append(errs, NewError("MAC", "mac"))
	}

	if m.Pointer == nil || !isMAC(*m.Pointer) {
		errs = append(errs, NewError("Pointer", "mac"))
	}

	for _, v := range m.Slice {
		if !isMAC(v) {
			errs = append(errs, NewError("Slice", "mac"))
		}
	}
	return errs
}

// Validate TODO
func (m *MaxValidate) Validate() (errs ValidationErrors) {

	if len(m.Name) >= 10 {
		errs = append(errs, NewError("Name", "max"))
	}

	if m.Password == nil || len(*m.Password) >= 8 {
		errs = append(errs, NewError("Password", "max"))
	}

	if m.Age >= 50 {
		errs = append(errs, NewError("Age", "max"))
	}

	if len(m.Contents) >= 2 {
		errs = append(errs, NewError("Contents", "max"))
	}

	if len(m.Attributes) >= 3 {
		errs = append(errs, NewError("Attributes", "max"))
	}
	return errs
}

// Validate TODO
func (m *MinValidate) Validate() (errs ValidationErrors) {

	if len(m.Name) <= 3 {
		errs = append(errs, NewError("Name", "min"))
	}

	if m.Password == nil || len(*m.Password) <= 3 {
		errs = append(errs, NewError("Password", "min"))
	}

	if m.Age <= 22 {
		errs = append(errs, NewError("Age", "min"))
	}

	if len(m.Contents) <= 1 {
		errs = append(errs, NewError("Contents", "min"))
	}

	if len(m.Attributes) <= 2 {
		errs = append(errs, NewError("Attributes", "min"))
	}
	return errs
}

// Validate TODO
func (h *Home) Validate() (errs ValidationErrors) {
	if IsInt(h.ID) {
		errs = append(errs, NewError("ID", "required"))
	}
	if (h.Address == Address{}) {
		errs = append(errs, NewError("Address", "required"))
	}
	return errs
}

// Validate TODO
func (c *Client) Validate() (errs ValidationErrors) {
	if IsInt64(c.ID) {
		errs = append(errs, NewError("ID", "required"))
	}

	if Empty(len(c.Name)) {
		errs = append(errs, NewError("Name", "required"))
	}

	if c.LastName == nil {
		errs = append(errs, NewError("LastName", "required"))
	}

	if c.Address == nil {
		errs = append(errs, NewError("Address", "required"))
	}
	return errs
}

// Validate TODO
func (u *User) Validate() (errs ValidationErrors) {
	if Empty(len(u.Name)) {
		errs = append(errs, NewError("Name", "required"))
	}

	return errs
}

// Validate TODO
func (n *NeNumber) Validate() (errs ValidationErrors) {

	if n.Uint == 1 {
		errs = append(errs, NewError("Uint", "ne"))
	}

	if n.UintPointer == nil || *n.UintPointer == 2 {
		errs = append(errs, NewError("UintPointer", "ne"))
	}

	if n.Uint8 == 3 {
		errs = append(errs, NewError("Uint8", "ne"))
	}

	if n.Uint8Pointer == nil || *n.Uint8Pointer == 4 {
		errs = append(errs, NewError("Uint8Pointer", "ne"))
	}

	if n.Uint16 == 5 {
		errs = append(errs, NewError("Uint16", "ne"))
	}

	if n.Uint16Pointer == nil || *n.Uint16Pointer == 6 {
		errs = append(errs, NewError("Uint16Pointer", "ne"))
	}

	if n.Uint32 == 7 {
		errs = append(errs, NewError("Uint32", "ne"))
	}

	if n.Uint32Pointer == nil || *n.Uint32Pointer == 8 {
		errs = append(errs, NewError("Uint32Pointer", "ne"))
	}

	if n.Uint64 == 9 {
		errs = append(errs, NewError("Uint64", "ne"))
	}

	if n.Uint64Pointer == nil || *n.Uint64Pointer == 10 {
		errs = append(errs, NewError("Uint64Pointer", "ne"))
	}

	if n.Uintptr == 11 {
		errs = append(errs, NewError("Uintptr", "ne"))
	}

	if n.UintptrPointer == nil || *n.UintptrPointer == 12 {
		errs = append(errs, NewError("UintptrPointer", "ne"))
	}

	if n.Int == 13 {
		errs = append(errs, NewError("Int", "ne"))
	}

	if n.IntPointer == nil || *n.IntPointer == 14 {
		errs = append(errs, NewError("IntPointer", "ne"))
	}

	if n.Int8 == 15 {
		errs = append(errs, NewError("Int8", "ne"))
	}

	if n.Int8Pointer == nil || *n.Int8Pointer == 16 {
		errs = append(errs, NewError("Int8Pointer", "ne"))
	}

	if n.Int16 == 17 {
		errs = append(errs, NewError("Int16", "ne"))
	}

	if n.Int16Pointer == nil || *n.Int16Pointer == 18 {
		errs = append(errs, NewError("Int16Pointer", "ne"))
	}

	if n.Int32 == 19 {
		errs = append(errs, NewError("Int32", "ne"))
	}

	if n.Int32Pointer == nil || *n.Int32Pointer == 20 {
		errs = append(errs, NewError("Int32Pointer", "ne"))
	}

	if n.Int64 == 21 {
		errs = append(errs, NewError("Int64", "ne"))
	}

	if n.Int64Pointer == nil || *n.Int64Pointer == 22 {
		errs = append(errs, NewError("Int64Pointer", "ne"))
	}

	if n.Float32 == 23 {
		errs = append(errs, NewError("Float32", "ne"))
	}

	if n.Float32Pointer == nil || *n.Float32Pointer == 24 {
		errs = append(errs, NewError("Float32Pointer", "ne"))
	}

	if n.Float64 == 25 {
		errs = append(errs, NewError("Float64", "ne"))
	}

	if n.Float64Pointer == nil || *n.Float64Pointer == 26 {
		errs = append(errs, NewError("Float64Pointer", "ne"))
	}

	if n.Complex64 == 27 {
		errs = append(errs, NewError("Complex64", "ne"))
	}

	if n.Complex64Pointer == nil || *n.Complex64Pointer == 28 {
		errs = append(errs, NewError("Complex64Pointer", "ne"))
	}

	if n.Complex128 == 29 {
		errs = append(errs, NewError("Complex128", "ne"))
	}

	if n.Complex128Pointer == nil || *n.Complex128Pointer == 30 {
		errs = append(errs, NewError("Complex128Pointer", "ne"))
	}
	return errs
}

// Validate TODO
func (n *NeString) Validate() (errs ValidationErrors) {

	if n.String == "chico" {
		errs = append(errs, NewError("String", "ne"))
	}

	if n.Pointer == nil || *n.Pointer == "bento" {
		errs = append(errs, NewError("Pointer", "ne"))
	}
	return errs
}

// Validate TODO
func (o *OmitemptyValidate) Validate() (errs ValidationErrors) {

	if !Empty(len(o.OmitEmptyString)) && !(len(o.OmitEmptyString) <= 5 && len(o.OmitEmptyString) >= 5) {
		errs = append(errs, NewError("OmitEmptyString", "[min = 5] ~ [max = 5]"))
	}

	if !IsInt(o.OmitEmptyInt) && !(o.OmitEmptyInt <= 5 && o.OmitEmptyInt >= 5) {
		errs = append(errs, NewError("OmitEmptyInt", "[min = 5] ~ [max = 5]"))
	}

	if !Empty(len(o.OmitEmptySliceString)) && !(len(o.OmitEmptySliceString) <= 1) {
		errs = append(errs, NewError("OmitEmptySliceString", "[min = 1]"))
	}

	if !Empty(len(o.OmitEmptySliceInt)) {
		errs = append(errs, NewError("OmitEmptySliceInt", "omitempty"))
	}
	return errs
}

// Validate TODO
func (r *RequiredChanString) Validate() (errs ValidationErrors) {
	if Empty(len(r.ChanString)) {
		errs = append(errs, NewError("ChanString", "required"))
	}

	if Empty(len(r.PointerChanString)) {
		errs = append(errs, NewError("PointerChanString", "required"))
	}

	if Empty(len(r.ChanSliceString)) {
		errs = append(errs, NewError("ChanSliceString", "required"))
	}

	if Empty(len(r.ChanSlicePointerString)) {
		errs = append(errs, NewError("ChanSlicePointerString", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredChanUint) Validate() (errs ValidationErrors) {
	if Empty(len(r.ChanUint)) {
		errs = append(errs, NewError("ChanUint", "required"))
	}

	if Empty(len(r.PointerChanUint)) {
		errs = append(errs, NewError("PointerChanUint", "required"))
	}

	if Empty(len(r.ChanSliceUint)) {
		errs = append(errs, NewError("ChanSliceUint", "required"))
	}

	if Empty(len(r.ChanSlicePointerUint)) {
		errs = append(errs, NewError("ChanSlicePointerUint", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredString) Validate() (errs ValidationErrors) {
	if Empty(len(r.String)) {
		errs = append(errs, NewError("String", "required"))
	}

	if r.Pointer == nil {
		errs = append(errs, NewError("Pointer", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredInterface) Validate() (errs ValidationErrors) {
	if IsValid(r.Interface) {
		errs = append(errs, NewError("Interface", "required"))
	}

	if r.PointerInterface == nil {
		errs = append(errs, NewError("PointerInterface", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredStruct) Validate() (errs ValidationErrors) {
	if (r.Struct == RequiredStructA{}) {
		errs = append(errs, NewError("Struct", "required"))
	}

	if r.StructPointer == nil {
		errs = append(errs, NewError("StructPointer", "required"))
	}

	if Empty(len(r.SliceStruct)) {
		errs = append(errs, NewError("SliceStruct", "required"))
	}

	if Empty(len(r.SliceStructPointer)) {
		errs = append(errs, NewError("SliceStructPointer", "required"))
	}

	if r.SlicePointerStruct == nil {
		errs = append(errs, NewError("SlicePointerStruct", "required"))
	}

	if r.SlicePointerStructPointer == nil {
		errs = append(errs, NewError("SlicePointerStructPointer", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceString) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceString)) {
		errs = append(errs, NewError("SliceString", "required"))
	}

	if r.PointerSliceString == nil {
		errs = append(errs, NewError("PointerSliceString", "required"))
	}

	if Empty(len(r.SlicePointerString)) {
		errs = append(errs, NewError("SlicePointerString", "required"))
	}

	if r.PointerSlicePointerString == nil {
		errs = append(errs, NewError("PointerSlicePointerString", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceUint) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceUint)) {
		errs = append(errs, NewError("SliceUint", "required"))
	}

	if r.PointerSliceUint == nil {
		errs = append(errs, NewError("PointerSliceUint", "required"))
	}

	if Empty(len(r.SlicePointerUint)) {
		errs = append(errs, NewError("SlicePointerUint", "required"))
	}

	if r.PointerSlicePointerUint == nil {
		errs = append(errs, NewError("PointerSlicePointerUint", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceUint8) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceUint8)) {
		errs = append(errs, NewError("SliceUint8", "required"))
	}

	if r.PointerSliceUint8 == nil {
		errs = append(errs, NewError("PointerSliceUint8", "required"))
	}

	if Empty(len(r.SlicePointerUint8)) {
		errs = append(errs, NewError("SlicePointerUint8", "required"))
	}

	if r.PointerSlicePointerUint8 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint8", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceUint16) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceUint16)) {
		errs = append(errs, NewError("SliceUint16", "required"))
	}

	if r.PointerSliceUint16 == nil {
		errs = append(errs, NewError("PointerSliceUint16", "required"))
	}

	if Empty(len(r.SlicePointerUint16)) {
		errs = append(errs, NewError("SlicePointerUint16", "required"))
	}

	if r.PointerSlicePointerUint16 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint16", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceUint32) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceUint32)) {
		errs = append(errs, NewError("SliceUint32", "required"))
	}

	if r.PointerSliceUint32 == nil {
		errs = append(errs, NewError("PointerSliceUint32", "required"))
	}

	if Empty(len(r.SlicePointerUint32)) {
		errs = append(errs, NewError("SlicePointerUint32", "required"))
	}

	if r.PointerSlicePointerUint32 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint32", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceUint64) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceUint64)) {
		errs = append(errs, NewError("SliceUint64", "required"))
	}

	if r.PointerSliceUint64 == nil {
		errs = append(errs, NewError("PointerSliceUint64", "required"))
	}

	if Empty(len(r.SlicePointerUint64)) {
		errs = append(errs, NewError("SlicePointerUint64", "required"))
	}

	if r.PointerSlicePointerUint64 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint64", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceUintptr) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceUintptr)) {
		errs = append(errs, NewError("SliceUintptr", "required"))
	}

	if r.PointerSliceUintptr == nil {
		errs = append(errs, NewError("PointerSliceUintptr", "required"))
	}

	if Empty(len(r.SlicePointerUintptr)) {
		errs = append(errs, NewError("SlicePointerUintptr", "required"))
	}

	if r.PointerSlicePointerUintptr == nil {
		errs = append(errs, NewError("PointerSlicePointerUintptr", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceInt) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceInt)) {
		errs = append(errs, NewError("SliceInt", "required"))
	}

	if r.PointerSliceInt == nil {
		errs = append(errs, NewError("PointerSliceInt", "required"))
	}

	if Empty(len(r.SlicePointerInt)) {
		errs = append(errs, NewError("SlicePointerInt", "required"))
	}

	if r.PointerSlicePointerInt == nil {
		errs = append(errs, NewError("PointerSlicePointerInt", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceInt8) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceInt8)) {
		errs = append(errs, NewError("SliceInt8", "required"))
	}

	if r.PointerSliceInt8 == nil {
		errs = append(errs, NewError("PointerSliceInt8", "required"))
	}

	if Empty(len(r.SlicePointerInt8)) {
		errs = append(errs, NewError("SlicePointerInt8", "required"))
	}

	if r.PointerSlicePointerInt8 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt8", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceInt16) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceInt16)) {
		errs = append(errs, NewError("SliceInt16", "required"))
	}

	if r.PointerSliceInt16 == nil {
		errs = append(errs, NewError("PointerSliceInt16", "required"))
	}

	if Empty(len(r.SlicePointerInt16)) {
		errs = append(errs, NewError("SlicePointerInt16", "required"))
	}

	if r.PointerSlicePointerInt16 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt16", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceInt32) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceInt32)) {
		errs = append(errs, NewError("SliceInt32", "required"))
	}

	if r.PointerSliceInt32 == nil {
		errs = append(errs, NewError("PointerSliceInt32", "required"))
	}

	if Empty(len(r.SlicePointerInt32)) {
		errs = append(errs, NewError("SlicePointerInt32", "required"))
	}

	if r.PointerSlicePointerInt32 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt32", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceInt64) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceInt64)) {
		errs = append(errs, NewError("SliceInt64", "required"))
	}

	if r.PointerSliceInt64 == nil {
		errs = append(errs, NewError("PointerSliceInt64", "required"))
	}

	if Empty(len(r.SlicePointerInt64)) {
		errs = append(errs, NewError("SlicePointerInt64", "required"))
	}

	if r.PointerSlicePointerInt64 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt64", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceFloat32) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceFloat32)) {
		errs = append(errs, NewError("SliceFloat32", "required"))
	}

	if r.PointerSliceFloat32 == nil {
		errs = append(errs, NewError("PointerSliceFloat32", "required"))
	}

	if Empty(len(r.SlicePointerFloat32)) {
		errs = append(errs, NewError("SlicePointerFloat32", "required"))
	}

	if r.PointerSlicePointerFloat32 == nil {
		errs = append(errs, NewError("PointerSlicePointerFloat32", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceFloat64) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceFloat64)) {
		errs = append(errs, NewError("SliceFloat64", "required"))
	}

	if r.PointerSliceFloat64 == nil {
		errs = append(errs, NewError("PointerSliceFloat64", "required"))
	}

	if Empty(len(r.SlicePointerFloat64)) {
		errs = append(errs, NewError("SlicePointerFloat64", "required"))
	}

	if r.PointerSlicePointerFloat64 == nil {
		errs = append(errs, NewError("PointerSlicePointerFloat64", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceComplex64) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceComplex64)) {
		errs = append(errs, NewError("SliceComplex64", "required"))
	}

	if r.PointerSliceComplex64 == nil {
		errs = append(errs, NewError("PointerSliceComplex64", "required"))
	}

	if Empty(len(r.SlicePointerComplex64)) {
		errs = append(errs, NewError("SlicePointerComplex64", "required"))
	}

	if r.PointerSlicePointerComplex64 == nil {
		errs = append(errs, NewError("PointerSlicePointerComplex64", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredSliceComplex128) Validate() (errs ValidationErrors) {
	if Empty(len(r.SliceComplex128)) {
		errs = append(errs, NewError("SliceComplex128", "required"))
	}

	if r.PointerSliceComplex128 == nil {
		errs = append(errs, NewError("PointerSliceComplex128", "required"))
	}

	if Empty(len(r.SlicePointerComplex128)) {
		errs = append(errs, NewError("SlicePointerComplex128", "required"))
	}

	if r.PointerSlicePointerComplex128 == nil {
		errs = append(errs, NewError("PointerSlicePointerComplex128", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredNumber) Validate() (errs ValidationErrors) {
	if IsUint(r.Uint) {
		errs = append(errs, NewError("Uint", "required"))
	}

	if r.UintPointer == nil {
		errs = append(errs, NewError("UintPointer", "required"))
	}

	if IsUint8(r.Uint8) {
		errs = append(errs, NewError("Uint8", "required"))
	}

	if r.Uint8Pointer == nil {
		errs = append(errs, NewError("Uint8Pointer", "required"))
	}

	if IsUint16(r.Uint16) {
		errs = append(errs, NewError("Uint16", "required"))
	}

	if r.Uint16Pointer == nil {
		errs = append(errs, NewError("Uint16Pointer", "required"))
	}

	if IsUint32(r.Uint32) {
		errs = append(errs, NewError("Uint32", "required"))
	}

	if r.Uint32Pointer == nil {
		errs = append(errs, NewError("Uint32Pointer", "required"))
	}

	if IsUint64(r.Uint64) {
		errs = append(errs, NewError("Uint64", "required"))
	}

	if r.Uint64Pointer == nil {
		errs = append(errs, NewError("Uint64Pointer", "required"))
	}

	if IsUintptr(r.Uintptr) {
		errs = append(errs, NewError("Uintptr", "required"))
	}

	if r.UintptrPointer == nil {
		errs = append(errs, NewError("UintptrPointer", "required"))
	}

	if IsInt(r.Int) {
		errs = append(errs, NewError("Int", "required"))
	}

	if r.IntPointer == nil {
		errs = append(errs, NewError("IntPointer", "required"))
	}

	if IsInt8(r.Int8) {
		errs = append(errs, NewError("Int8", "required"))
	}

	if r.Int8Pointer == nil {
		errs = append(errs, NewError("Int8Pointer", "required"))
	}

	if IsInt16(r.Int16) {
		errs = append(errs, NewError("Int16", "required"))
	}

	if r.Int16Pointer == nil {
		errs = append(errs, NewError("Int16Pointer", "required"))
	}

	if IsInt32(r.Int32) {
		errs = append(errs, NewError("Int32", "required"))
	}

	if r.Int32Pointer == nil {
		errs = append(errs, NewError("Int32Pointer", "required"))
	}

	if IsInt64(r.Int64) {
		errs = append(errs, NewError("Int64", "required"))
	}

	if r.Int64Pointer == nil {
		errs = append(errs, NewError("Int64Pointer", "required"))
	}

	if IsFloat32(r.Float32) {
		errs = append(errs, NewError("Float32", "required"))
	}

	if r.Float32Pointer == nil {
		errs = append(errs, NewError("Float32Pointer", "required"))
	}

	if IsFloat64(r.Float64) {
		errs = append(errs, NewError("Float64", "required"))
	}

	if r.Float64Pointer == nil {
		errs = append(errs, NewError("Float64Pointer", "required"))
	}

	if IsComplex64(r.Complex64) {
		errs = append(errs, NewError("Complex64", "required"))
	}

	if r.Complex64Pointer == nil {
		errs = append(errs, NewError("Complex64Pointer", "required"))
	}

	if IsComplex128(r.Complex128) {
		errs = append(errs, NewError("Complex128", "required"))
	}

	if r.Complex128Pointer == nil {
		errs = append(errs, NewError("Complex128Pointer", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredMapStringToInterface) Validate() (errs ValidationErrors) {
	if Empty(len(r.MapStringToInterface)) {
		errs = append(errs, NewError("MapStringToInterface", "required"))
	}

	if r.PointerMapStringToInterface == nil {
		errs = append(errs, NewError("PointerMapStringToInterface", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredMapStringToString) Validate() (errs ValidationErrors) {
	if Empty(len(r.MapStringToString)) {
		errs = append(errs, NewError("MapStringToString", "required"))
	}

	if r.PointerMapStringToString == nil {
		errs = append(errs, NewError("PointerMapStringToString", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredMapIntToString) Validate() (errs ValidationErrors) {
	if Empty(len(r.MapIntToString)) {
		errs = append(errs, NewError("MapIntToString", "required"))
	}

	if r.PointerMapIntToString == nil {
		errs = append(errs, NewError("PointerMapIntToString", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredMapIntToBool) Validate() (errs ValidationErrors) {
	if Empty(len(r.MapIntToBool)) {
		errs = append(errs, NewError("MapIntToBool", "required"))
	}

	if r.PointerMapIntToBool == nil {
		errs = append(errs, NewError("PointerMapIntToBool", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredMapIntToStruct) Validate() (errs ValidationErrors) {
	if Empty(len(r.MapIntToStructA)) {
		errs = append(errs, NewError("MapIntToStructA", "required"))
	}

	if Empty(len(r.PointerMapIntToStructA)) {
		errs = append(errs, NewError("PointerMapIntToStructA", "required"))
	}

	if r.MapIntToStructB == nil {
		errs = append(errs, NewError("MapIntToStructB", "required"))
	}

	if r.PointerMapIntToStructPointerB == nil {
		errs = append(errs, NewError("PointerMapIntToStructPointerB", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.Status) && Empty(len(r.Name)) {
		errs = append(errs, NewError("Name", "required_with"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.Status) && r.Name == nil {
		errs = append(errs, NewError("Name", "required_with"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithFields) Validate() (errs ValidationErrors) {

	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.ID == 1 || r.Status) && r.Name == nil {
		errs = append(errs, NewError("Name", "required_with"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithConfirmation) Validate() (errs ValidationErrors) {

	if !isURL(r.URL) {
		errs = append(errs, NewError("URL", "url"))
	}

	if (isURL(r.URL)) && !r.NeedConfirmation {
		errs = append(errs, NewError("NeedConfirmation", "required_with"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithAllField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.Status) && Empty(len(r.Name)) {
		errs = append(errs, NewError("Name", "required_with_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithAllFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.Status) && r.Name == nil {
		errs = append(errs, NewError("Name", "required_with_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithAllFields) Validate() (errs ValidationErrors) {

	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.ID == 1 && r.Status) && r.Name == nil {
		errs = append(errs, NewError("Name", "required_with_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.Status) && Empty(len(r.Name)) {
		errs = append(errs, NewError("Name", "required_without"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.Status) && r.Name == nil {
		errs = append(errs, NewError("Name", "required_without"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutFields) Validate() (errs ValidationErrors) {

	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if (r.ID == 1 || r.Status) && r.Name == nil {
		errs = append(errs, NewError("Name", "required_without"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutConfirmation) Validate() (errs ValidationErrors) {

	if !isURL(r.URL) {
		errs = append(errs, NewError("URL", "url"))
	}

	if (isURL(r.URL)) && !r.NeedConfirmation {
		errs = append(errs, NewError("NeedConfirmation", "required_without"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutAllField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if Empty(len(r.Name)) && (r.Status) {
		errs = append(errs, NewError("Name", "required_without_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutAllFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.Status) {
		errs = append(errs, NewError("Name", "required_without_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutAllFields) Validate() (errs ValidationErrors) {

	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.ID == 1 && r.Status) {
		errs = append(errs, NewError("Name", "required_without_all"))
	}
	return errs
}

// Validate TODO
func (r *RGBValidate) Validate() (errs ValidationErrors) {

	if !rgbRegex.MatchString(r.RGB) {
		errs = append(errs, NewError("RGB", "rgb"))
	}

	if r.Pointer == nil || !rgbRegex.MatchString(*r.Pointer) {
		errs = append(errs, NewError("Pointer", "rgb"))
	}
	return errs
}

// Validate TODO
func (r *RGBAValidate) Validate() (errs ValidationErrors) {

	if !rgbaRegex.MatchString(r.RGBA) {
		errs = append(errs, NewError("RGBA", "rgba"))
	}

	if r.Pointer == nil || !rgbaRegex.MatchString(*r.Pointer) {
		errs = append(errs, NewError("Pointer", "rgba"))
	}
	return errs
}

// Validate TODO
func (s *SSNValidate) Validate() (errs ValidationErrors) {

	if !isSSN(s.SSN) {
		errs = append(errs, NewError("SSN", "ssn"))
	}

	if s.Pointer == nil || !isSSN(*s.Pointer) {
		errs = append(errs, NewError("Pointer", "ssn"))
	}

	for _, v := range s.Slice {
		if !isSSN(v) {
			errs = append(errs, NewError("Slice", "ssn"))
		}
	}
	return errs
}

// Validate TODO
func (t *TCPAddrValidate) Validate() (errs ValidationErrors) {

	if !isTCPAddrResolvable(t.TCPAddr) {
		errs = append(errs, NewError("TCPAddr", "tcp_addr"))
	}

	if t.Pointer == nil || !isTCPAddrResolvable(*t.Pointer) {
		errs = append(errs, NewError("Pointer", "tcp_addr"))
	}

	for _, v := range t.Slice {
		if !isTCPAddrResolvable(v) {
			errs = append(errs, NewError("Slice", "tcp_addr"))
		}
	}
	return errs
}

// Validate TODO
func (t *TCP4AddrValidate) Validate() (errs ValidationErrors) {

	if !isTCP4AddrResolvable(t.TCP4Addr) {
		errs = append(errs, NewError("TCP4Addr", "tcp4_addr"))
	}

	if t.Pointer == nil || !isTCP4AddrResolvable(*t.Pointer) {
		errs = append(errs, NewError("Pointer", "tcp4_addr"))
	}

	for _, v := range t.Slice {
		if !isTCP4AddrResolvable(v) {
			errs = append(errs, NewError("Slice", "tcp4_addr"))
		}
	}
	return errs
}

// Validate TODO
func (t *TCP6AddrValidate) Validate() (errs ValidationErrors) {

	if !isTCP6AddrResolvable(t.TCP6Addr) {
		errs = append(errs, NewError("TCP6Addr", "tcp6_addr"))
	}

	if t.Pointer == nil || !isTCP6AddrResolvable(*t.Pointer) {
		errs = append(errs, NewError("Pointer", "tcp6_addr"))
	}

	for _, v := range t.Slice {
		if !isTCP6AddrResolvable(v) {
			errs = append(errs, NewError("Slice", "tcp6_addr"))
		}
	}
	return errs
}

// Validate TODO
func (u *UDPAddrValidate) Validate() (errs ValidationErrors) {

	if !isUDPAddrResolvable(u.UDPAddr) {
		errs = append(errs, NewError("UDPAddr", "udp_addr"))
	}

	if u.Pointer == nil || !isUDPAddrResolvable(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "udp_addr"))
	}

	for _, v := range u.Slice {
		if !isUDPAddrResolvable(v) {
			errs = append(errs, NewError("Slice", "udp_addr"))
		}
	}
	return errs
}

// Validate TODO
func (u *UDP4AddrValidate) Validate() (errs ValidationErrors) {

	if !isUDP4AddrResolvable(u.UDP4Addr) {
		errs = append(errs, NewError("UDP4Addr", "udp4_addr"))
	}

	if u.Pointer == nil || !isUDP4AddrResolvable(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "udp4_addr"))
	}

	for _, v := range u.Slice {
		if !isUDP4AddrResolvable(v) {
			errs = append(errs, NewError("Slice", "udp4_addr"))
		}
	}
	return errs
}

// Validate TODO
func (u *UDP6AddrValidate) Validate() (errs ValidationErrors) {

	if !isUDP6AddrResolvable(u.UDP6Addr) {
		errs = append(errs, NewError("UDP6Addr", "udp6_addr"))
	}

	if u.Pointer == nil || !isUDP6AddrResolvable(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "udp6_addr"))
	}

	for _, v := range u.Slice {
		if !isUDP6AddrResolvable(v) {
			errs = append(errs, NewError("Slice", "udp6_addr"))
		}
	}
	return errs
}

// Validate TODO
func (u *UNIXAddrValidate) Validate() (errs ValidationErrors) {

	if !isUnixAddrResolvable(u.UNIXAddr) {
		errs = append(errs, NewError("UNIXAddr", "unix_addr"))
	}

	if u.Pointer == nil || !isUnixAddrResolvable(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "unix_addr"))
	}

	for _, v := range u.Slice {
		if !isUnixAddrResolvable(v) {
			errs = append(errs, NewError("Slice", "unix_addr"))
		}
	}
	return errs
}

// Validate TODO
func (u *URIValidate) Validate() (errs ValidationErrors) {

	if !isURI(u.URI) {
		errs = append(errs, NewError("URI", "uri"))
	}

	if u.Pointer == nil || !isURI(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uri"))
	}

	for _, v := range u.Slice {
		if !isURI(v) {
			errs = append(errs, NewError("Slice", "uri"))
		}
	}
	return errs
}

// Validate TODO
func (u *URLValidate) Validate() (errs ValidationErrors) {

	if !isURL(u.URL) {
		errs = append(errs, NewError("URL", "url"))
	}

	if u.Pointer == nil || !isURL(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "url"))
	}

	for _, v := range u.Slice {
		if !isURL(v) {
			errs = append(errs, NewError("Slice", "url"))
		}
	}
	return errs
}

// Validate TODO
func (u *URLEncodedValidate) Validate() (errs ValidationErrors) {

	if !uRLEncodedRegex.MatchString(u.URLEncoded) {
		errs = append(errs, NewError("URLEncoded", "url_encoded"))
	}

	if u.Pointer == nil || !uRLEncodedRegex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "url_encoded"))
	}

	for _, v := range u.Slice {
		if !uRLEncodedRegex.MatchString(v) {
			errs = append(errs, NewError("Slice", "url_encoded"))
		}
	}
	return errs
}

// Validate TODO
func (u *URNValidate) Validate() (errs ValidationErrors) {

	if !isUrnRFC2141(u.URN) {
		errs = append(errs, NewError("URN", "urn_rfc2141"))
	}

	if u.Pointer == nil || !isUrnRFC2141(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "urn_rfc2141"))
	}

	for _, v := range u.Slice {
		if !isUrnRFC2141(v) {
			errs = append(errs, NewError("Slice", "urn_rfc2141"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUIDValidate) Validate() (errs ValidationErrors) {

	if !uUIDRegex.MatchString(u.UUID) {
		errs = append(errs, NewError("UUID", "uuid"))
	}

	if u.Pointer == nil || !uUIDRegex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid"))
	}

	for _, v := range u.Slice {
		if !uUIDRegex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUID3Validate) Validate() (errs ValidationErrors) {

	if !uUID3Regex.MatchString(u.UUID3) {
		errs = append(errs, NewError("UUID3", "uuid3"))
	}

	if u.Pointer == nil || !uUID3Regex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid3"))
	}

	for _, v := range u.Slice {
		if !uUID3Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid3"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUID4Validate) Validate() (errs ValidationErrors) {

	if !uUID4Regex.MatchString(u.UUID4) {
		errs = append(errs, NewError("UUID4", "uuid4"))
	}

	if u.Pointer == nil || !uUID4Regex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid4"))
	}

	for _, v := range u.Slice {
		if !uUID4Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid4"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUID5Validate) Validate() (errs ValidationErrors) {

	if !uUID5Regex.MatchString(u.UUID5) {
		errs = append(errs, NewError("UUID5", "uuid5"))
	}

	if u.Pointer == nil || !uUID5Regex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid5"))
	}

	for _, v := range u.Slice {
		if !uUID5Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid5"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUIDrfc4122Validate) Validate() (errs ValidationErrors) {

	if !uUIDRFC4122Regex.MatchString(u.UUID) {
		errs = append(errs, NewError("UUID", "uuid_rfc4122"))
	}

	if u.Pointer == nil || !uUIDRFC4122Regex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid_rfc4122"))
	}

	for _, v := range u.Slice {
		if !uUIDRFC4122Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid_rfc4122"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUID3rfc4122Validate) Validate() (errs ValidationErrors) {

	if !uUID3RFC4122Regex.MatchString(u.UUID3rfc4122) {
		errs = append(errs, NewError("UUID3rfc4122", "uuid3_rfc4122"))
	}

	if u.Pointer == nil || !uUID3RFC4122Regex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid3_rfc4122"))
	}

	for _, v := range u.Slice {
		if !uUID3RFC4122Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid3_rfc4122"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUID4rfc4122Validate) Validate() (errs ValidationErrors) {

	if !uUID4RFC4122Regex.MatchString(u.UUID4rfc4122) {
		errs = append(errs, NewError("UUID4rfc4122", "uuid4_rfc4122"))
	}

	if u.Pointer == nil || !uUID4RFC4122Regex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid4_rfc4122"))
	}

	for _, v := range u.Slice {
		if !uUID4RFC4122Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid4_rfc4122"))
		}
	}
	return errs
}

// Validate TODO
func (u *UUID5rfc4122Validate) Validate() (errs ValidationErrors) {

	if !uUID5RFC4122Regex.MatchString(u.UUID5rfc4122) {
		errs = append(errs, NewError("UUID5rfc4122", "uuid5_rfc4122"))
	}

	if u.Pointer == nil || !uUID5RFC4122Regex.MatchString(*u.Pointer) {
		errs = append(errs, NewError("Pointer", "uuid5_rfc4122"))
	}

	for _, v := range u.Slice {
		if !uUID5RFC4122Regex.MatchString(v) {
			errs = append(errs, NewError("Slice", "uuid5_rfc4122"))
		}
	}
	return errs
}

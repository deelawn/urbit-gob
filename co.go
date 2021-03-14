package urbitgob

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

const (
	namePartitionPattern string = ".{1,3}"
	pre                  string = "dozmarbinwansamlitsighidfidlissogdirwacsabwissib" +
		"rigsoldopmodfoglidhopdardorlorhodfolrintogsilmir" +
		"holpaslacrovlivdalsatlibtabhanticpidtorbolfosdot" +
		"losdilforpilramtirwintadbicdifrocwidbisdasmidlop" +
		"rilnardapmolsanlocnovsitnidtipsicropwitnatpanmin" +
		"ritpodmottamtolsavposnapnopsomfinfonbanmorworsip" +
		"ronnorbotwicsocwatdolmagpicdavbidbaltimtasmallig" +
		"sivtagpadsaldivdactansidfabtarmonranniswolmispal" +
		"lasdismaprabtobrollatlonnodnavfignomnibpagsopral" +
		"bilhaddocridmocpacravripfaltodtiltinhapmicfanpat" +
		"taclabmogsimsonpinlomrictapfirhasbosbatpochactid" +
		"havsaplindibhosdabbitbarracparloddosbortochilmac" +
		"tomdigfilfasmithobharmighinradmashalraglagfadtop" +
		"mophabnilnosmilfopfamdatnoldinhatnacrisfotribhoc" +
		"nimlarfitwalrapsarnalmoslandondanladdovrivbacpol" +
		"laptalpitnambonrostonfodponsovnocsorlavmatmipfip"
	suf string = "zodnecbudwessevpersutletfulpensytdurwepserwylsun" +
		"rypsyxdyrnuphebpeglupdepdysputlughecryttyvsydnex" +
		"lunmeplutseppesdelsulpedtemledtulmetwenbynhexfeb" +
		"pyldulhetmevruttylwydtepbesdexsefwycburderneppur" +
		"rysrebdennutsubpetrulsynregtydsupsemwynrecmegnet" +
		"secmulnymtevwebsummutnyxrextebfushepbenmuswyxsym" +
		"selrucdecwexsyrwetdylmynmesdetbetbeltuxtugmyrpel" +
		"syptermebsetdutdegtexsurfeltudnuxruxrenwytnubmed" +
		"lytdusnebrumtynseglyxpunresredfunrevrefmectedrus" +
		"bexlebduxrynnumpyxrygryxfeptyrtustyclegnemfermer" +
		"tenlusnussyltecmexpubrymtucfyllepdebbermughuttun" +
		"bylsudpemdevlurdefbusbeprunmelpexdytbyttyplevmyl" +
		"wedducfurfexnulluclennerlexrupnedlecrydlydfenwel" +
		"nydhusrelrudneshesfetdesretdunlernyrsebhulryllud" +
		"remlysfynwerrycsugnysnyllyndyndemluxfedsedbecmun" +
		"lyrtesmudnytbyrsenwegfyrmurtelreptegpecnelnevfes"
	galaxy string = "galaxy"
	star   string = "star"
	planet string = "planet"
	moon   string = "moon"
	comet  string = "comet"
)

var (

	// Numbers
	zero  = big.NewInt(0)
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
	five  = big.NewInt(5)
	eight = big.NewInt(8)

	// Prefixes and suffixes
	prefixes      = regexp.MustCompile(namePartitionPattern).FindAllString(pre, -1)
	suffixes      = regexp.MustCompile(namePartitionPattern).FindAllString(suf, -1)
	prefixesIndex = map[string]int{}
	suffixesIndex = map[string]int{}

	// Error format strings
	errInvalidBin string = "invalid binary string: %s"
	errInvalidHex string = "invalid hexadecimal string: %s"
	errInvalidInt string = "invalid integer string: %s"
	errInvalidP   string = "invalid @p: %s"
	errInvalidQ   string = "invalid @q: %s"
)

func init() {

	// This assumes length of prefixes and suffixes are the same, which they should be.
	for i := 0; i < len(prefixes); i++ {
		prefixesIndex[prefixes[i]] = i
		suffixesIndex[suffixes[i]] = i
	}
}

func patp2syls(name string) []string {

	removeCharsPattern := regexp.MustCompile(`[\^~-]`)
	normalizedName := removeCharsPattern.ReplaceAllString(name, "")
	partitionPattern := regexp.MustCompile(namePartitionPattern)
	return partitionPattern.FindAllString(normalizedName, -1)
}

func bex(n *big.Int) *big.Int {

	return big.NewInt(0).Exp(two, n, nil)
}

func rsh(a, b, c *big.Int) *big.Int {

	return big.NewInt(0).Div(c, bex(big.NewInt(0).Mul(bex(a), b)))
}

func met(a, b, c *big.Int) *big.Int {

	if c == nil {
		c = big.NewInt(0)
	}

	if b.Cmp(zero) == 0 {
		return c
	}

	return met(a, rsh(a, one, b), big.NewInt(0).Add(c, one))
}

func end(a, b, c *big.Int) *big.Int {

	return big.NewInt(0).Mod(c, bex(big.NewInt(0).Mul(bex(a), b)))
}

// TODO: look everywhere that uses SetString and return errors on failure
func Hex2Patp(hex string) (string, error) {

	v, ok := big.NewInt(0).SetString(hex, 16)
	if !ok {
		return "", fmt.Errorf(errInvalidHex, hex)
	}

	return Patp(v.String())
}

func Patp2Hex(name string) (string, error) {

	if !IsValidPat(name) {
		return "", fmt.Errorf(errInvalidP, name)
	}

	syls := patp2syls(name)

	var addr string
	hasLengthOne := len(syls) == 1
	for i := 0; i < len(syls); i++ {
		if i%2 != 0 || hasLengthOne {
			addr += syl2bin(suffixesIndex[syls[i]])
		} else {
			addr += syl2bin(prefixesIndex[syls[i]])
		}
	}

	bigAddr, ok := big.NewInt(0).SetString(addr, 2)
	if !ok {
		return "", fmt.Errorf(errInvalidBin, addr)
	}

	hex := Fynd(bigAddr).Text(16)

	if len(hex)%2 != 0 {
		return "0" + hex, nil
	}

	return hex, nil
}

func syl2bin(idx int) string {

	binStr := strconv.FormatInt(int64(idx), 2)
	return strings.Repeat("0", 8-len(binStr)) + binStr // padStart
}

func patp2bn(name string) (*big.Int, error) {

	hexStr, err := Patp2Hex(name)
	if err != nil {
		return nil, err
	}

	hex, ok := big.NewInt(0).SetString(hexStr, 16)
	if !ok {
		return nil, fmt.Errorf(errInvalidHex, hexStr)
	}

	return hex, nil
}

func Patp2Dec(name string) (string, error) {

	dec, err := patp2bn(name)
	if err != nil {
		return "", err
	}

	return dec.String(), nil
}

func Patq(arg string) (string, error) {

	v, ok := big.NewInt(0).SetString(arg, 10)
	if !ok {
		return "", fmt.Errorf(errInvalidInt, arg)
	}

	return buf2patq(v.Bytes()), nil
}

func buf2patq(buf []byte) string {

	var chunked [][]byte
	if len(buf)%2 != 0 && len(buf) > 1 {
		chunked = append([][]byte{{buf[0]}}, chunk(buf[1:], 2)...)
	} else {
		chunked = chunk(buf, 2)
	}

	patq := "~"
	chunkedLen := len(chunked)
	for _, elem := range chunked {

		if patq != "~" {
			patq += "-"
		}

		patq += alg(elem, chunkedLen)
	}

	return patq
}

func prefixName(pair []byte) string {

	if len(pair) == 1 {
		return prefixes[0] + suffixes[pair[0]]
	}

	return prefixes[pair[0]] + suffixes[pair[1]]
}

func name(pair []byte) string {

	if len(pair) == 1 {
		return suffixes[pair[0]]
	}

	return prefixes[pair[0]] + suffixes[pair[1]]
}

func alg(pair []byte, chunkedLen int) string {

	if len(pair)%2 != 0 && chunkedLen > 1 {
		return prefixName(pair)
	}

	return name(pair)
}

func chunk(items []byte, size int) [][]byte {

	slices := [][]byte{}

	for _, item := range items {

		sliceLength := len(slices)
		if sliceLength == 0 || len(slices[sliceLength-1]) == size {
			slices = append(slices, []byte{})
		}

		slices[sliceLength-1] = append(slices[sliceLength-1], item)
	}

	return slices
}

func Hex2Patq(arg string) (string, error) {

	hexStr := arg
	if len(arg)%2 != 0 {
		hexStr = "0" + hexStr
	}

	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", fmt.Errorf(errInvalidHex, arg)
	}

	return buf2patq(buf), nil
}

func Patq2Hex(name string) (string, error) {

	if !IsValidPat(name) {
		return "", fmt.Errorf(errInvalidQ, name)
	}

	if len(name) == 0 {
		return "00", nil
	}

	chunks := strings.Split(name[1:], "-")
	return splat(chunks), nil
}

func dec2hex(dec int) string {

	decStr := strconv.FormatInt(int64(dec), 16)
	if l := len(decStr); l < 2 {
		padding := strings.Repeat("0", 2-l)
		decStr = padding + decStr
	}

	return decStr
}

func splat(chunks []string) string {

	var hexStr string
	for _, chunk := range chunks {

		syls := []string{chunk}
		if len(syls) > 3 {
			syls = []string{chunk[:3], chunk[3:]}
		}
		if len(syls) == 1 {
			hexStr += dec2hex(suffixesIndex[syls[0]])
		} else {
			hexStr += dec2hex(prefixesIndex[syls[0]]) + dec2hex(suffixesIndex[syls[1]])
		}
	}

	return hexStr
}

func patq2bn(name string) (*big.Int, error) {

	hexStr, err := Patq2Hex(name)
	if err != nil {
		return nil, err
	}

	v, ok := big.NewInt(0).SetString(hexStr, 16)
	if !ok {
		return nil, fmt.Errorf(errInvalidHex, name)
	}

	return v, nil
}

func Patq2Dec(name string) (string, error) {

	v, err := patq2bn(name)
	if err != nil {
		return "", err
	}

	return v.String(), nil
}

func Clan(who string) (string, error) {

	name, err := patp2bn(who)
	if err != nil {
		return "", err
	}

	wid := met(three, name, nil)

	if wid.Cmp(one) <= 0 {
		return galaxy, nil
	}
	if wid.Cmp(two) <= 0 {
		return star, nil
	}
	if wid.Cmp(four) <= 0 {
		return planet, nil
	}
	if wid.Cmp(eight) <= 0 {
		return moon, nil
	}

	return comet, nil
}

func Sein(name string) (string, error) {

	who, err := patp2bn(name)
	if err != nil {
		return "", err
	}

	mir, err := Clan(name)
	if err != nil {
		return "", err
	}

	var res *big.Int
	switch mir {
	case galaxy:
		res = who
	case star:
		res = end(three, one, who)
	case planet:
		res = end(four, one, who)
	case moon:
		res = end(five, one, who)
	default:
		res = zero
	}

	return Patp(res.String())
}

func IsValidPat(name string) bool {

	if len(name) < 4 || name[0] != '~' {
		return false
	}

	syls := patp2syls(name)

	sylsLen := len(syls)
	for i, syl := range syls {
		if i%2 != 0 || sylsLen == 1 {
			if _, ok := suffixesIndex[syl]; !ok {
				return false
			}
		} else if _, ok := prefixesIndex[syl]; !ok {
			return false
		}
	}

	return sylsLen%2 != 0 && sylsLen != 1
}

func IsValidPatp(str string) bool {

	dec, err := Patp2Dec(str)
	if err != nil {
		return false
	}

	p, err := Patp(dec)
	if err != nil {
		return false
	}

	return IsValidPat(str) && str == p
}

func IsValidPatq(str string) bool {

	dec, err := Patq2Dec(str)
	if err != nil {
		return false
	}

	q, err := Patq(dec)
	if err != nil {
		return false
	}

	isValid, err := EqPatq(str, q)
	if err != nil {
		return false
	}

	return IsValidPat(str) && isValid
}

func removeLeadingZeros(str string) string {

	for i, c := range str {
		if c != '0' {
			return str[i:]
		}
	}

	return ""
}

func eqModLeadingZeros(s, t string) bool {

	return removeLeadingZeros(s) == removeLeadingZeros(t)
}

func EqPatq(p, q string) (bool, error) {

	phex, err := Patq2Hex(p)
	if err != nil {
		return false, err
	}

	qhex, err := Patq2Hex(q)
	if err != nil {
		return false, err
	}

	return eqModLeadingZeros(phex, qhex), nil
}

func Patp(arg string) (string, error) {

	v, ok := big.NewInt(0).SetString(arg, 10)
	if !ok {
		return "", fmt.Errorf(errInvalidInt, arg)
	}

	sxz := Fein(v.String())
	dyy := met(four, sxz, nil)
	dyx := met(three, sxz, nil)

	p := "~"

	if dyx.Cmp(one) <= 0 {
		p += suffixes[int(sxz.Int64())]
	} else {
		p += patpLoop(dyy, sxz, zero, "")
	}

	return p, nil
}

func patpLoop(dyy, tsxz, timp *big.Int, trep string) string {

	log := end(four, one, tsxz)
	pre := prefixes[int(rsh(three, one, log).Int64())]
	suf := suffixes[int(end(three, one, log).Int64())]

	var etc string
	if big.NewInt(0).Mod(timp, four).Cmp(zero) == 0 {
		if timp.Cmp(zero) != 0 {
			etc = "--"
		}
	} else {
		etc = "-"
	}

	res := pre + suf + etc + trep

	if timp.Cmp(dyy) == 0 {
		return trep
	}

	return patpLoop(dyy, rsh(four, one, tsxz), big.NewInt(0).Add(timp, one), res)
}

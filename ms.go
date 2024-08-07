//
// Copyright (C) Philip Schlump, 2013-2015.
//
// My Strings (ms) package
//
// String and Os/String related support functions that work with Go (golang) templates.
// test change from new system.  Just a test.
//

package ms

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	strftime "github.com/hhkbp2/go-strftime" // ../strftime
	"github.com/pschlump/MiscLib"
	tr "github.com/pschlump/dbgo"
	words "github.com/pschlump/gowords"
	"github.com/pschlump/picfloat" // "../picfloat"
	"github.com/pschlump/pictime"  // "../pictime"
)

const ISO8601 = "2006-01-02T15:04:05.99999Z07:00"

// This is annoying but I had to include this function or the Go (verison 1.2) compiler
// hurled.  Oh well....  It is short - just leave it in.
// func dummy() {
// 	fmt.Printf("Make da compiler happy\n")
// }

const (
	PathSep = string(os.PathSeparator)
)

// Implement Dijkstra's L algorythm on an array of strings.
// Return -1 if not found, else index of item.
func FindCol(aName string, nameArray []string) int {
	for i, v := range nameArray {
		if aName == v {
			return i
		}
	}
	return -1
}

// Center the string 's' in the width 'n'
func CenterStr(n int, t interface{}) (r string) {
	s := fmt.Sprintf("%v", t)
	l := len(s)
	// fmt.Printf ( "l=%d\n", l )
	if l < n {
		blanks := (n - l) / 2
		// fmt.Printf ( "blanks=%d\n", blanks )
		r = PadStrRight(n, " ", PadStr(blanks+l, " ", s))
	} else {
		r = s
	}
	return
}

// PadStr pads a string with 'w' to a length of 'l' - pad on left
func PadStr(l int, w string, s string) string {
	if len(s) >= l {
		return s
	}
	k := l - len(s)
	t := ""
	for i := 0; i < k; i++ {
		t += w
	}
	return t + s
}

// PadStrRight pads a string with 'w' to a length of 'l' - pad on right
func PadStrRight(l int, w string, s string) string {
	if len(s) >= l {
		return s
	}
	k := l - len(s)
	t := ""
	for i := 0; i < k; i++ {
		t += w
	}
	return s + t
}

// ZeroPad left pad a string with "0" to the desired length 'l'
func ZeroPad(l int, s string) string {
	return PadStr(l, "0", s)
}

// ZeroPadRight pad on right a string with "0" to desired length 'l'
func ZeroPadRight(l int, s string) string {
	return PadStrRight(l, "0", s)
}

// Merge sets of data ----------------------------------------------------------------------------------------------------

// LowerCaseNames convers the keys on a map of strings to lower case
func LowerCaseNames(a map[string]interface{}) (rv map[string]interface{}) {
	rv = make(map[string]interface{})
	for i, v := range a {
		rv[strings.ToLower(i)] = v
	}
	return
}

func ExtendData(a map[string]interface{}, b map[string]interface{}) (rv map[string]interface{}) {
	rv = make(map[string]interface{})
	for i, v := range a {
		rv[i] = v
	}
	for i, v := range b {
		rv[i] = v
	}
	return
}

// Copy 'a', if same key in 'b', then copy data from b, prefering data from 'b'
func LeftData(a map[string]interface{}, b map[string]interface{}) (rv map[string]interface{}) {
	rv = make(map[string]interface{})
	for i, v := range a {
		rv[i] = v
	}
	for i, v := range b {
		if _, ok := a[i]; ok {
			rv[i] = v
		}
	}
	return
}

// Keep the data that has common keys between 'a' and 'b', prefering data from 'b'
// not used at the moment.
func IntersectData(a map[string]interface{}, b map[string]interface{}) (rv map[string]interface{}) {
	rv = make(map[string]interface{})
	for i, v := range a {
		if _, ok := b[i]; ok {
			rv[i] = v
		}
	}
	for i, v := range b {
		if _, ok := a[i]; ok {
			rv[i] = v
		}
	}
	return
}

func SplitOnWords(s string) (record []string) {

	// fmt.Printf ( "For ->%s<-\n", s )

	var reader *words.Reader
	reader = words.NewReader(strings.NewReader(s))

	record, err := reader.Read()

	if err == io.EOF {
		// return
	} else if err != nil {
		fmt.Printf("Error(12015): %v\n", err)
	}

	// fmt.Printf ( "SpitOnWords Gets ->%v<-\n", record )

	return

}

func PadOnLeft(n int, s interface{}) (r string) {
	r = fmt.Sprintf(fmt.Sprintf("%%%dv", n), s)
	return
}

func PadOnRight(n int, s interface{}) (r string) {
	r = fmt.Sprintf(fmt.Sprintf("%%-%dv", n), s)
	return
}

//func centerStr ( n int, s string ) ( r string ) {
//	l := len(s)
//	if l < n {
//		// fmt.Printf ( "Centering a narow field in a wide area, l=%d, n=%d\n", l, n )
//		blanks := (n-l)/2
//		// fmt.Printf ( "blanks=%d\n", blanks )
//		p1 := fmt.Sprintf ( "%%%ds%%s", blanks )
//		// fmt.Printf ( "format to right just/leading blanks string p1=->%s<-\n", p1 )
//		p2 := fmt.Sprintf ( p1, "", s )
//		// fmt.Printf ( "value with leading blanks >%s<-\n", p2 )
//		p4 := fmt.Sprintf ( "%%-%ds", n )
//		p3 := fmt.Sprintf ( p4, p2 )
//		// fmt.Printf ( "value with trailing blanks >%s<-\n", p3 )
//		r = p3
//	} else {
//		r = s
//	}
//	return
//}

func FmtDate(f string, t time.Time) (r string) {
	r = t.Format(f)
	return
}

func FmtDateTS(f string, ts string) (r string) {
	fff := "2006-01-02T15:04:05.00000+0000"
	t, err := time.Parse(fff, ts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sUnable to parse ->%s<- with format ->%s<- error=%s%s\n", MiscLib.ColorRed, ts, fff, err, MiscLib.ColorReset)
		return "2006-01-02T15:04:05.000Z"
	}
	r = t.Format(f)
	return
}

func IsEven(x int) (r bool) {
	if (x % 2) == 0 {
		r = true
	} else {
		r = false
	}
	return
}

func StrFTime(f string, t time.Time) (r string) {
	// older version of library returned errros if there was one
	//r, err := strftime.Format(f, t)
	//if err != nil {
	//	r = fmt.Sprintf("%v", err)
	//}
	r = strftime.Format(f, t)
	return
}

func PicFloat(format string, flt interface{}) (r string) {
	switch flt.(type) {
	case int:
		r = picfloat.Format(format, float64(flt.(int)))
	case int64:
		r = picfloat.Format(format, float64(flt.(int64)))
	case float32:
		r = picfloat.Format(format, float64(flt.(float32)))
	case float64:
		r = picfloat.Format(format, flt.(float64))
	case string:
		f, err := strconv.ParseFloat(flt.(string), 64)
		if err != nil {
			f = 0.0
		}
		r = picfloat.Format(format, f)
	default:
		fmt.Printf("Error(12026): invalid data type for PicFloat, got %T\n", flt)
	}
	return
}

func Nvl(show string, d string) string {
	if d == "" {
		return show
	} else {
		return d
	}
}

// idiotic format for dates ( just use ISO format YYYY-MM-DDTHH:mm:ss.nnn! )
//
//	   <field member="Arrival" columns="40">Arrived   : {0:hh:mm tt  ddd, MMM dd, yyyy}</field>
//			hh - 2digit hours
//			mm - 2digit minutes
//			tt - lower case am/pm
//			ddd - day of week
//			MMM - 3 char month name abrev.
//			dd - 2 digit day of month.
//			yyyy - year 4 digit
func PicTime(f string, t interface{}) (r string) {
	switch t.(type) {
	case time.Time:
		var err error
		r, err = pictime.Format(f, t.(time.Time))
		if err != nil {
			r = fmt.Sprintf("%v", err)
		}
	case int:
		r = fmt.Sprintf("Error(14026): invalid data type for PicTime, got %T\n", t)
		//r = picfloat.Format ( format, float64(flt.(int)) )
	case int64:
		r = fmt.Sprintf("Error(14027): invalid data type for PicTime, got %T\n", t)
		//r = picfloat.Format ( format, float64(flt.(int64)) )
	case float32:
		r = fmt.Sprintf("Error(14031): invalid data type for PicTime, got %T\n", t)
		//r = picfloat.Format ( format, float64(flt.(float32)) )
	case float64:
		r = fmt.Sprintf("Error(14032): invalid data type for PicTime, got %T\n", t)
		//r = picfloat.Format ( format, flt.(float64) )
	case string:
		// f, err := strconv.ParseFloat ( flt.(string), 64 )
		// if err != nil {
		// 	f = 0.0
		// }
		// r = picfloat.Format ( format, f )
		d, _, err := FuzzyDateTimeParse(t.(string), false)
		// r = t.(string)
		r, err = pictime.Format(f, d)
		if err != nil {
			r = fmt.Sprintf("%v", err)
		}
	default:
		r = fmt.Sprintf("Error(14033): invalid data type for PicFloat, got %T\n", t)
	}
	return
}

type datePat struct {
	fmt  *regexp.Regexp
	prep func(s string, t string) string
	pfmt string
}

var datePatTab []datePat

const db_fuzzy_date = true

func rmQuote(s string, t string) string {
	return strings.Trim(s, `"`)
}
func resetToOrig(s string, t string) string {
	return t
}
func rmTrailingZ(s string, t string) string {
	matched := strings.HasSuffix(s, "Z")
	if matched {
		return strings.TrimRight(s, "Z")
	}
	return s
}

func init() {
	// isIntStringRe = regexp.MustCompile("[0-9][0-9]*")
	// const ISO8601 = "2006-01-02T15:04:05.99999Z07:00"
	/* 0 */
	datePatTab = append(datePatTab, datePat{regexp.MustCompile(`"`), rmQuote, ""})
	/* 1 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`Z$`), rmTrailingZ, ""})
	/* 2 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2} [aApP][mM]$`), nil, "01/02/2006 03:04:05 PM"})
	/* 3 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{2} \d{2}:\d{2}:\d{2} [aApP][mM]$`), nil, "01/02/06 03:04:05 PM"})
	/* 2 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{4} \d{1}:\d{2}:\d{2} [aApP][mM]$`), nil, "01/02/2006 3:04:05 PM"})
	/* 3 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{2} \d{1}:\d{2}:\d{2} [aApP][mM]$`), nil, "01/02/06 3:04:05 PM"})
	/* 4 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{4} \d{2}:\d{2} [aApP][mM]$`), nil, "01/02/2006 03:04 PM"})
	/* 5 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{2} \d{2}:\d{2} [aApP][mM]$`), nil, "01/02/06 03:04 PM"})
	/* 4 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{4} \d{1}:\d{2} [aApP][mM]$`), nil, "01/02/2006 3:04 PM"})
	/* 5 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}/\d{2}/\d{2} \d{1}:\d{2} [aApP][mM]$`), nil, "01/02/06 3:04 PM"})
	/* 6 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d+$`), nil, "2006-01-02T15:04:05.99999"})
	/* 7 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`), nil, "2006-01-02T15:04:05"})
	/* 8 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`), nil, "2006-01-02"})
	/* 9 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{2}-\d{2}-\d{2}$`), nil, "06-01-02"})
	/*10 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`.`), resetToOrig, ""})
	/*11 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`"`), rmQuote, ""})
	/*12 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d+[-+]\d{2}:\d{2}$`), nil, "2006-01-02T15:04:05.99999-07:00"})
	/*13 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d+[-+]\d{4}$`), nil, "2006-01-02T15:04:05.99999-0700"})
	/*14 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$`), nil, "2006-01-02T15:04:05-07:00"})
	/*15 */ datePatTab = append(datePatTab, datePat{regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+]\d{4}$`), nil, "2006-01-02T15:04:05-0700"})

	// 2013-11-06T00:00:00+0000
}

func FuzzyDateTimeParse(s string, nullOk bool) (d time.Time, isNull bool, err error) {
	t := s
	isNull = false
	if len(s) == 0 && nullOk {
		isNull = true
		err = nil
		return
	}
	for i, v := range datePatTab {
		// _ = i
		if v.fmt.MatchString(s) {
			if db_fuzzy_date {
				fmt.Printf("got match at %d, s=%s, pfmt=%s\n", i, s, v.pfmt)
			}
			if v.prep != nil {
				s = v.prep(s, t)
				if db_fuzzy_date {
					fmt.Printf("	mod s to s=%s\n", s)
				}
			}
			if len(v.pfmt) > 0 {
				d, err = time.Parse(v.pfmt, s)
				if err == nil {
					if db_fuzzy_date {
						fmt.Printf("	success at %d\n", i)
					}
					return
				}
			}
		}
	}
	err = errors.New("Invalid date, unable to parse it.")
	return
}

func FmtPrintfStr(f string, s interface{}) (r string) {
	r = fmt.Sprintf(f, s)
	return
}

func findFixBindAndQuote(s string) (qryFixed string, bindVars []string, err error) {
	err = nil
	q := []byte(s)
	var w []byte
	var bv []byte
	var fDollar bool = false
	st := 0
	for i, v := range q {
		switch st {
		case 0:
			switch v {
			case '\'': // Quoted String
				st = 2
				w = append(w, v)
			case '"':
				st = 1
				w = append(w, '[')
			case '$':
				bv = bv[:0]
				fDollar = true
				st = 4
				w = append(w, '?')
			default:
				w = append(w, v)
			}
		case 1:
			switch v {
			case '"':
				if i+1 < len(q) && q[i+1] == '"' {
					st = 3
					w = append(w, v)
				} else {
					st = 0
					w = append(w, ']')
				}
			default:
				w = append(w, v)
			}
		case 2: // Inside ' string
			switch v {
			case '\'':
				st = 0
				w = append(w, v)
			default:
				w = append(w, v)
			}
		case 3: // Advance 1 char, ignore
			st = 0
		case 4: // Grab a bind variable
			switch v {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				bv = append(bv, v)
			default: // gone 1 too far [ mimic state 0 ]
				st = 0
				bindVars = append(bindVars, string(bv))
				bv = bv[:0]
				fDollar = false
				{
					switch v {
					case '\'': // Quoted String
						st = 2
						w = append(w, v)
					case '"':
						st = 1
						w = append(w, '[')
					case '$':
						bv = bv[:0]
						st = 4
						w = append(w, '?')
					default:
						w = append(w, v)
					}
				}
			}
		default:
			panic("UnreacableCode(12028): Invalid state")
		}
	}
	if st == 2 {
		err = errors.New("Ended inside quoted string")
	}
	if len(bv) > 0 || fDollar {
		bindVars = append(bindVars, string(bv))
	}
	qryFixed = string(w)
	return
}

func FixBindParams(qry string, data ...interface{}) (qryFixed string, retData []interface{}, err error) {
	var k int64
	err = nil
	qryFixed, bindVars, e2 := findFixBindAndQuote(qry) // Get array of bind params // R.E. to match $[0-9][0-9]*
	err = e2
	// fmt.Printf ( "bindVars = %v\n", bindVars )
	for _, v := range bindVars {
		k, err = strconv.ParseInt(v, 10, 32)
		if err != nil {
			return
		}
		// fmt.Printf ( "bind set k=%d\n", k )
		if k > 0 && int(k) <= len(data) {
			retData = append(retData, data[k-1])
		} else {
			err = errors.New(fmt.Sprintf("Error(12027): Query (%s) is invalid, insufficient bind parameters were passed.", qry))
		}
	}
	return
}

func Concat(args ...interface{}) string {
	//fmt.Fprintf ( xOut, "Concat: args=%v\n", args )
	if len(args) == 0 {
		return ""
	}
	s := ""
	for i, y := range args {
		switch y.(type) {
		case string:
			s += y.(string)
			// fmt.Fprintf ( xOut, "Concat: %dth arg ->%v<-\n", i, y )
		case int, int64, byte, float32, float64:
			s += fmt.Sprintf("%v", y)
			// fmt.Fprintf ( xOut, "Concat: %dth arg ->%v<-\n", i, y )
		case time.Time:
			s += (args[i+2].(time.Time)).Format(ISO8601)
		default:
			s += ""
			fmt.Printf("Concat: don't know what to do with: %dth arg ->%T<-\n", i, y)
		}
	}

	return s
}

// ===================================================================================================================================================
// New Mon Feb 23 16:37:31 MST 2015
// ===================================================================================================================================================
// {{ifDef . "Placeholder" "placeholder=\"" "$$" "\""}}
func IfDef(dataHash map[string]interface{}, it string, ss ...string) string {
	if x, ok := dataHash[it]; ok {
		rv := ""
		for _, v := range ss {
			if v == "$$" {
				rv += fmt.Sprintf("%s", x)
			} else {
				rv += v
			}
		}
		return rv
	}
	return ""
}

func IfIsDef(dataHash map[string]interface{}, it string) bool {
	fmt.Printf("it=%s\n", it)
	if x, ok := dataHash[it]; ok {
		fmt.Printf("Returning %v, v2=%v\n", ok, x)
		return ok
	}
	return false
}

/*
Sat May 11 20:30:16 MDT 2019
func IfIsLen0(dataHash map[string]interface{}, it string) bool {
	fmt.Printf("IfIsLen0: it=%s\n", it)
	if x, ok := dataHash[it]; ok {
		l := len(x)
		fmt.Printf("Returning %v, v2=%v\n", l == 0, x)
		return l == 0
	}
	return false
}
*/

func IfIsNotNull(dataHash map[string]interface{}, it string) bool {
	if x, ok := dataHash[it]; ok {
		if ok {
			return x != nil
		}
	}
	return false
}

// ===================================================================================================================================================
func DoPost(client *http.Client, Url string, s string) string {
	r1, e0 := client.PostForm(Url, url.Values{"auth_token": {s}})
	if e0 != nil {
		fmt.Printf("Error!!!!!!!!!!! %v, %s\n", e0, tr.LF())
		return "Error"
	}
	rv, e1 := ioutil.ReadAll(r1.Body)
	if e1 != nil {
		fmt.Printf("Error!!!!!!!!!!! %v, %s\n", e1, tr.LF())
		return "Error"
	}
	r1.Body.Close()
	if string(rv[0:6]) == ")]}',\n" {
		rv = rv[6:]
	}

	return string(rv)
}

// ===================================================================================================================================================
func DoGet(client *http.Client, url string) string {
	r1, e0 := client.Get(url)
	if e0 != nil {
		fmt.Printf("Error!!!!!!!!!!! %v, %s\n", e0, tr.LF())
		return "Error"
	}
	rv, e1 := ioutil.ReadAll(r1.Body)
	if e1 != nil {
		fmt.Printf("Error!!!!!!!!!!! %v, %s\n", e1, tr.LF())
		return "Error"
	}
	r1.Body.Close()
	if string(rv[0:6]) == ")]}',\n" {
		rv = rv[6:]
	}

	return string(rv)
}

// ===================================================================================================================================================
func Tr(line, pat, rep string) (rv string) {
	inPat := make(map[rune]int)
	for ii, rn := range pat {
		inPat[rn] = ii
	}
	var buffer bytes.Buffer
	for _, rn := range line {
		if pos, found := inPat[rn]; found {
			if pos < len(rep) {
				buffer.WriteByte(rep[pos])
			} else {
				buffer.WriteByte(rep[len(rep)-1])
			}
		} else {
			buffer.WriteRune(rn)
		}
	}
	rv = buffer.String()
	return
}

// ChopAt cuts the string 's' at the first ocurance of string 'c'.
func ChopAt(c string, s string) (t string) {
	t = s
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i:], c) {
			t = s[0:i]
		}
	}
	return
}

/*
func main() {
	fmt.Printf("abcTxyz ->%s<-\n", ChopAt("T", "abcTxyz"))
}
*/

/* vim: set noai ts=4 sw=4: */

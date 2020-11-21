// Several functions converting variables types

package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func ToString(v interface{}) string {
	return fmt.Sprint(v)
}

/*
 * INT to STRING
 */
// Its8: INT8 to STRING
func Its8(n int8) string {
	return fmt.Sprint(n)
}

// Uits8: UINT8 to STRING
func Uits8(n uint8) string {
	return fmt.Sprint(n)
}

// Its16: INT16 to STRING
func Its16(n int16) string {
	return fmt.Sprint(n)
}

// Uits16: UINT16 to STRING
func Uits16(n uint16) string {
	return fmt.Sprint(n)
}

// Its32: INT to STRING
func Its32(n int32) string {
	return fmt.Sprint(n)
}

// Uits32: UINT to STRING
func Uits32(n uint32) string {
	return fmt.Sprint(n)
}

// Its: INT to STRING
func Its(n int) string {
	return fmt.Sprint(n)
}

// Uits: UINT to STRING
func Uits(n uint) string {
	return fmt.Sprint(n)
}

// Its64: INT64 to STRING
func Its64(n int64) string {
	return fmt.Sprint(n)
}

// Uits64: INT64 to STRING
func Uits64(n uint64) string {
	return fmt.Sprint(n)
}

/*
 * INT to base 2, 8 e 16
 */
// IntToHex: INT to HEX
func IntToHex(n int64) string {
	return strconv.FormatInt(n, 16)
}

// IntToOct: INT to OCT
func IntToOct(n int64) string {
	return strconv.FormatInt(n, 8)
}

// IntToBin: INT to BIN
func IntToBin(n int64) string {
	return strconv.FormatInt(n, 2)
}

/*
 * INT to FLOAT
 */
// Itd: INT to FLOAT64 (double)
func Itd(n int) float64 {
	return float64(n)
}

// Uitd: UINT to FLOAT64 (double)
func Uitd(n uint) float64 {
	return float64(n)
}

// Itf: INT to FLOAT32 (full)
func Itf(n int) float32 {
	return float32(n)
}

// Uitf: UINT to FLOAT32 (full)
func Uitf(n uint) float32 {
	return float32(n)
}

// Itd: INT64 to FLOAT64 (double)
func I64td(n int64) float64 {
	return float64(n)
}

// Itf: INT64 to FLOAT32 (full)
func I64tf(n int64) float32 {
	return float32(n)
}

/*
 * STRING to NUMBER
 */
// Sti8: Tries to convert STRING to INT8; checks for error
func Sti8(s string) int8 {
	res, e := strconv.ParseInt(s, 10, 8)
	Check(e, "Error trying to convert STRING to INT; string: "+s)

	return int8(res)
}

// Stui8: Tries to convert STRING to UINT8; checks for error
func Stui8(s string) uint8 {
	res, e := strconv.ParseInt(s, 10, 8)
	Check(e, "Erro tentando converter STRING to INT")

	return uint8(res)
}

// Sti16: Tries to convert STRING to INT16; checks for error
func Sti16(s string) int16 {
	res, e := strconv.ParseInt(s, 10, 16)
	Check(e, "Erro tentando converter STRING to INT")

	return int16(res)
}

// Stui16: Tries to convert STRING to UINT16; checks for error
func Stui16(s string) uint16 {
	res, e := strconv.ParseInt(s, 10, 16)
	Check(e, "Erro tentando converter STRING to INT")

	return uint16(res)
}

// Sti: Tries to convert STRING to INT; checks for error
func Sti(s string) int {
	if EmptyString(s) {
		Print("Convertendo string vazia, retornando 0")

		return 0
	}
	res, e := strconv.Atoi(s)
	StayCalm(e)

	return res
}

// Stui: Tries to convert STRING to UINT ; checks for error
func Stui(s string) uint {
	if EmptyString(s) {
		return 0
	}
	res, e := strconv.ParseInt(s, 10, 32)
	Check(e, "Erro tentando converter STRING to INT")

	return uint(res)
}

// Sti32: Tries to convert STRING to INT32; checks for error
func Sti32(s string) int {
	res, e := strconv.Atoi(s)
	Check(e, "Erro tentando converter STRING to INT")

	return res
}

// Stui32: Tries to convert STRING to UINT32; checks for error
func Stui32(s string) uint {
	res, e := strconv.ParseInt(s, 10, 32)
	Check(e, "Erro tentando converter STRING to INT")

	return uint(res)
}

// Sti64: Tries to convert STRING to INT64; checks for error
func Sti64(s string) int64 {
	res, e := strconv.ParseInt(s, 10, 64)
	Check(e, "Erro tentando converter STRING to INT")

	return res
}

// Sti64: Tries to convert STRING to INT64; checks for error
func Stui64(s string) uint64 {
	res, e := strconv.ParseInt(s, 10, 64)
	Check(e, "Erro tentando converter STRING to INT")

	return uint64(res)
}

// Stf64: Tries to convert STRING to FLOAT64; checks for error
func Stf64(s string) float64 {
	f, e := strconv.ParseFloat(s, 64)
	Check(e, "Error converting STRING to FLOAT")

	return f
}

//Stb: STRING to []BYTE
func Stb(s string) []byte {
	return []byte(s)
}

// F32to64: Convert float32 to float64
func F32to64(f float32) float64 {
	return float64(f)
}

// F32to64: Convert float64 to float 32; lose precision
func F64to32(f float64) float32 {
	return float32(f)
}

/*
 * FLOAT to STRING
 */
// F64ts: FLOAT64 (Double) to STRING
func F64ts(f float64) string {
	return fmt.Sprintf("%f", f)
}

// F32ts: FLOAT32 (Full) to STRING
func F32ts(f float32) string {
	return fmt.Sprintf("%f", f)
}

/*
 * FLOAT to INT
 */
// Dti: FLOAT64 (Double) to INT; (May lost precision)
func F64ti(f float64) int {
	return int(f)
}

// Dti64: FLOAT64 (Double) to INT64; (May lost precision)
func F64ti64(f float64) int64 {
	return int64(f)
}

// Fti: FLOAT32 (Full) to INT; (May lost precision)
func F32ti(f float32) int {
	return int(f)
}

// F32ti64: FLOAT32 (Full) to INT64; (May lost precision)
func F32ti64(f float32) int64 {
	return int64(f)
}

/*
 * []BYTE to STRING
 */
// Bts: []BYTE to STRING
func Bts(b []byte) string {
	return string(b[:])
}

// Bts: BYTE to STRING
func BtsSingle(b byte) string {
	return string(b)
}

/*
 * DATE&TIME to STRING
 */
// Dts: Return the time.Duration as STRING
func Dts(d time.Duration) string {
	return ToString(d)
}

// Dts: Return time.Time as STRING
func Tts(t time.Time) string {
	return t.String()
}

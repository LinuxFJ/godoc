package main

import (
	"encoding/binary"
	"fmt"
	// "hanqu/logger"
	// "math"
	// "math/big"
	"math/rand"
	"strconv"
	"strings"
)

const (
    DMARKET_ENCODE_TOKEN            = "N15P7D3EFUATYCS698VWGZ4MQ2KHBX0R"
    DMARKET_ENCODE_TOKEN36          = "N5P7D3EFUATYCS698VWGZ4MQ2KHBXRJ01ILO"
    DMARKET_ENCODE_TOKEN62          = "NEWGNIJ4VEVOQYCNAF71KWPC63NMSPN15P7D3EFUATYCS698VWGZ4MQ2KHBX0R"
    DMARKET_ENCODE_KEY              = "OAru2Ab%2$,-~8hrKT+v+S`%"
    DMARKET_ENCODE_LEN              = 12
    ENC_SHORT_BASE = 31
)

var (
	enc32_map  = make(map[byte]int)
	enc36_map  = make([]byte, 256)
	enc36_mapr = make([]byte, 256)
	enc62_map  = make([]byte, 256)
	enc62_mapr = make([]byte, 256)
)

// 254084768964048310
// 787662783788549760

const (
	ENC_GIFT_BASE   = 25500
	ENC_GIFT_MAX_ID = 78000
	ENC_LONG_MOD    = 10000000000000
	ENC_LONG_BASE   = 31
)

func initEncoder() {
    //30进制
	token := []byte(DMARKET_ENCODE_TOKEN)
	for k, v := range token {
		enc32_map[v] = k
	}

    //36进制
	token36 := []byte(DMARKET_ENCODE_TOKEN36)
	for k, v := range token36 {
		if k <= 9 {
            //0-9 ascii 48-57
			enc36_map[48+k] = v
			enc36_mapr[v] = byte(48 + k)
		} else {
            //A-Z ascii 65-90(55+k就是从65开始算得)
			enc36_map[55+k] = v
			enc36_mapr[v] = byte(55 + k)
		}
	}

    //62进制
    token62 := []byte(DMARKET_ENCODE_TOKEN62)
	for k, v := range token62 {
		if k <= 9 {
            //0-9 ascii 48-57
			enc62_map[48+k] = v
			enc62_mapr[v] = byte(48 + k)
		} else if k > 9 && k <= 35 {
            //A-Z ascii 65-90(55+k就是从65开始算得)
			enc62_map[55+k] = v
			enc62_mapr[v] = byte(55 + k)
        } else {
            //a-z ascii 97-122(61+k就是从97开始算得)
            enc62_map[61+k] = v
            enc62_mapr[v] = byte(61 + k)
        }
    }
}

type CdkeyEncoder struct {
	rnd1  uint8  // 1-14 只用4位
	shift uint8  // rand
	gift  uint16 // gift编号
	idx   uint16 // 序号 offset + idx
	rnd2  uint16 // rand 0-65535
}

func (this *CdkeyEncoder) GetGiftId() int {
	return int(this.gift)
}

func (this *CdkeyEncoder) Encode(idx uint16) string {
	iCdkey := ENC_LONG_MOD * (int(this.gift) + ENC_GIFT_BASE)
	iCdkey = iCdkey + rand.Intn(ENC_LONG_MOD)

	strFmt := strconv.FormatUint(uint64(iCdkey), ENC_LONG_BASE)
	strFmt = strings.ToUpper(strFmt)

	strEnc := this.doEncode(strFmt)
	// strDec := this.doDecode(strEnc)
	//	logger.DEBUG("------------:%v,%v,%v", strFmt, strEnc, strDec)

	return strEnc
}

func (this *CdkeyEncoder) EnCode62(charid uint64) string {
    strFmt := strconv.FormatUint(charid, ENC_SHORT_BASE)
    strFmt = strings.ToUpper(strFmt)
    strEnc := this.doEncode62(strFmt)
    return strEnc
}

func (this *CdkeyEncoder) doEncode62(in string) string {
    fmt.Println(in)
	nlen := len(in)
	byteIn := []byte(in)
	byteStr := make([]byte, nlen)

	for i := 0; i < nlen; i++ {
		idx := int(byteIn[i])
        fmt.Println(idx)
		byteStr[i] = enc62_map[idx]
	}
	return string(byteStr)
}

func (this *CdkeyEncoder) Decode(cdkey string) error {
	if len(cdkey) != 12 {
		return fmt.Errorf("inv cdkey len:", len(cdkey))
	}

	cdkey = this.doDecode(cdkey)

	u64, err := strconv.ParseUint(cdkey, ENC_LONG_BASE, 64)
	if err != nil {
		return err
	}

	iCdkey := int(u64)
	this.gift = uint16(iCdkey/ENC_LONG_MOD - 25500)

	return nil
}

func (this *CdkeyEncoder) doEncode(in string) string {
	nlen := len(in)
	byteIn := []byte(in)
	byteStr := make([]byte, nlen)

	for i := 0; i < nlen; i++ {
		idx := int(byteIn[i])
		byteStr[i] = enc36_map[idx]
	}

	return string(byteStr)
}

func (this *CdkeyEncoder) doDecode(in string) string {
	nlen := len(in)
	byteIn := []byte(in)
	byteStr := make([]byte, nlen)

	for i := 0; i < nlen; i++ {
		idx := int(byteIn[i])
		byteStr[i] = enc36_mapr[idx]
	}

	return string(byteStr)
}

// func (this *CdkeyEncoder) Encode(idx uint16) string {
// 	this.rnd1 = uint8(rand.Intn(14) + 1)
// 	if this.rnd1%5 == 0 {
// 		this.rnd1 = this.rnd1 - 1
// 	}
// 	this.rnd2 = uint16(rand.Intn(65535))
// 	this.shift = uint8(math.Log(float64(this.rnd2+1))*20.17017 + float64(this.rnd1))
// 	this.idx = uint16(this.shift) + idx

// 	buf := this.tobytes()
// 	bi := big.NewInt(0)
// 	bi.SetBytes(buf)

// 	xx := uint64(bi.Int64())
// 	x0 := uint64(xx & 0XFFFFFFFFFFFFFF)
// 	x1 := x0 >> this.rnd1
// 	x2 := (x0 << (64 - this.rnd1)) >> 8
// 	x3 := x2 + x1 + (uint64(this.rnd1) << 56)

// 	cdkey := this.enc32(uint(x3))

// 	// fmt.Println(buf)
// 	// fmt.Printf("enc:%v,%v,%v,%v,%v\n", this.rnd1, this.rnd2, this.gift, this.idx, this.shift)
// 	// fmt.Printf("cdkey encode:%v,%v,%v\n", cdkey, xx, x3)

// 	return cdkey
// }

// func (this *CdkeyEncoder) Decode(cdkey string) error {
// 	if len(cdkey) != DMARKET_ENCODE_LEN {
// 		return fmt.Errorf("inv len")
// 	}

// 	var x3 uint64
// 	x3 = 0
// 	keybytes := []byte(cdkey)
// 	for i := 0; i < DMARKET_ENCODE_LEN; i++ {
// 		idx := DMARKET_ENCODE_LEN - 1 - i
// 		enc_val := keybytes[i]
// 		val, ok := enc32_map[enc_val]
// 		if ok == false {
// 			return fmt.Errorf("inv char")
// 		}
// 		x3 = x3 + uint64(val<<(uint(idx)*5))
// 	}

// 	rnd1 := (x3 & 0XF00000000000000) >> 56

// 	x0 := x3 & 0XFFFFFFFFFFFFFF
// 	x1 := (x0 << rnd1) & 0XFFFFFFFFFFFFFF
// 	x2 := (x0 >> (56 - rnd1))
// 	xx := (x3 & 0XF00000000000000) + x1 + x2

// 	bi := big.NewInt(int64(xx))
// 	buf := bi.Bytes()
// 	if len(buf) != 8 {
// 		return fmt.Errorf("inv buf")
// 	}

// 	this.frombytes(buf)

// 	// TODO 演算校验

// 	// fmt.Println(buf)
// 	// fmt.Printf("dec:%v,%v,%v,%v,%v\n", this.rnd1, this.rnd2, this.gift, this.idx, this.shift)
// 	// fmt.Printf("cdkey decode:%v, %v, %v\n", cdkey, xx, x3)
// 	return nil
// }

func (this *CdkeyEncoder) tobytes() []byte {
	buf := make([]byte, 8)
	buf[0] = byte(this.rnd1)
	buf[1] = byte(this.shift)
	binary.BigEndian.PutUint16(buf[2:], this.gift)
	binary.BigEndian.PutUint16(buf[4:], this.idx)
	binary.BigEndian.PutUint16(buf[6:], this.rnd2)

	return buf
}

func (this *CdkeyEncoder) frombytes(b []byte) {
	this.rnd1 = uint8(b[0])
	this.shift = uint8(b[1])
	this.gift = binary.BigEndian.Uint16(b[2:4])
	this.idx = binary.BigEndian.Uint16(b[4:6])
	this.rnd2 = binary.BigEndian.Uint16(b[6:])
}

func (this *CdkeyEncoder) enc32(n uint) string {
	buf := make([]byte, 12)
	base := uint(31)
	bidx := -1

	for i := 11; i >= 0; i-- {
		nshift := uint(i) * 5

		nand := (((base << nshift) & n) >> nshift)
		if nand > 0 || bidx >= 0 {
			if bidx == -1 {
				bidx = 0
			} else {
				bidx = bidx + 1
			}
			buf[bidx] = byte(DMARKET_ENCODE_TOKEN[nand])
		}
	}

	return string(buf)
}

func main() {
    initEncoder()
    key := &CdkeyEncoder{}
    code := key.EnCode62(uint64(316255622189298))
    fmt.Printf("1 %s\n", code)
    code = key.EnCode62(uint64(316255622188833))
    fmt.Printf("2 %s\n", code)
}

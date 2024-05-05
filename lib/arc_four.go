package ArcFour

type RC4 struct {
	i uint8
	j uint8
	k uint8
	s [255]uint8
}

func (rc *RC4) swaps() {
	var tmp1 uint8
	var tmp2 uint8

	tmp1 = rc.s[rc.i]
	tmp2 = rc.s[rc.j]
	rc.s[rc.i] = tmp2
	rc.s[rc.j] = tmp1
}

func (rc *RC4) Init(key string) {
	var tmp1 uint8
	var tmp2 uint8

	for i := 0; i < 255; i++ {
		rc.s[i] = uint8(i)
	}

	keyLength := len([]rune(key))
	for i := 0; i < 255; i++ {
		tmp1 = key[i%keyLength]
		tmp2 = rc.j + rc.s[i] + tmp1
		rc.j = tmp2 % 255
		rc.swaps()
	}

	rc.i = 0
	rc.j = 0
}

func (rc *RC4) Encrypt(source string) string {
	var cipherText = []byte(source)
	for i := 0; i < len(source); i++ {
		cipherText[i] = source[i] ^ rc.prgaKey()
	}
	rc.reset()
	return string(cipherText)
}

func (rc *RC4) Decrypt(source string) string {
	return rc.Encrypt(source)
}

func (rc *RC4) prgaKey() uint8 {
	rc.i = (rc.i + 1) % 255
	rc.j = (rc.j + rc.s[rc.i]) % 255
	rc.swaps()
	tmp := (rc.s[rc.i] + rc.s[rc.j]) % 255
	rc.k = rc.s[tmp]
	return rc.k
}

func (rc *RC4) reset() {
	rc.i = 0
	rc.j = 0
	rc.k = 0
	rc.s = [255]uint8{}
}

package packs

type sex string
type sexType []byte

const MALE = "MALE"
const FEMALE = "FEMALE"

const (
	male = 0 == 0
	female =0 != 0
)



type Hooman struct {
    name string
    age  int
    sex_  sex
}

func(*Hooman) NewHooman(nname string, nage int, male) *Hooman {
	newhoo := Hooman{nname,nage,sex(nsex)}
	return &newhoo
}
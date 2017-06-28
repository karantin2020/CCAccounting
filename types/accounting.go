package types

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CriminalCaseV1 struct {
	Id              bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Number          int64         `bson:"number,omitempty" json:"number,omitempty"`
	InitDate        time.Time     `bson:"initDate,omitempty" json:"initDate,omitempty"`
	Initiator       string        `bson:"initiator,omitempty" json:"initiator,omitempty"`
	Investigator    string        `bson:"investigator,omitempty" json:"investigator,omitempty"`
	Scene           string        `bson:"scene,omitempty" json:"scene,omitempty"` // fabula
	Victim          []string      `bson:"victim,omitempty" json:"victim,omitempty"`
	MinorVictim     []bool        `bson:"minorVictim,omitempty" json:"minorVictim,omitempty"`
	Damage          string        `bson:"damage,omitempty" json:"damage,omitempty"`
	Accused         []string      `bson:"accused,omitempty" json:"accused,omitempty"`           // obvinyaemiy
	MinorAccused    []string      `bson:"minorAccused,omitempty" json:"minorAccused,omitempty"` // obvinyaemiy
	Qualification   string        `bson:"qualification,omitempty" json:"qualification,omitempty"`
	ArrestDate      time.Time     `bson:"arrestDate,omitempty" json:"arrestDate,omitempty"`
	Evidence        string        `bson:"evidence,omitempty" json:"evidence,omitempty"`
	Termination     time.Time     `bson:"termination,omitempty" json:"termination,omitempty"`
	Suspend         time.Time     `bson:"suspend,omitempty" json:"suspend,omitempty"`
	ToCourt         time.Time     `bson:"toCourt,omitempty" json:"toCourt,omitempty"`
	PreСondDate     time.Time     `bson:"preСondDate,omitempty" json:"preСondDate,omitempty"` // predvarit slushanie
	СondDate        time.Time     `bson:"сondDate,omitempty" json:"сondDate,omitempty"`       // date of result
	СondResult      CondKind      `bson:"сondResult,omitempty" json:"сondResult,omitempty"`
	Sud             string        `bson:"sud,omitempty" json:"sud,omitempty"`
	SudName         string        `bson:"sudName,omitempty" json:"sudName,omitempty"`
	Judge           string        `bson:"judge,omitempty" json:"judge,omitempty"`
	Gosobvinitel    string        `bson:"gosobvinitel,omitempty" json:"gosobvinitel,omitempty"`
	SpecSubject     SubjectKind   `bson:"specSubject,omitempty" json:"specSubject,omitempty"`
	Сonviction      []bool        `bson:"сonviction,omitempty" json:"сonviction,omitempty"` // sudimost'
	CrimeKind       CrimeK        `bson:"crimeKind,omitempty" json:"crimeKind,omitempty"`
	CrimeCategory   CrimeCat      `bson:"crimeCategory,omitempty" json:"crimeCategory,omitempty"`
	ChangeCategory  bool          `bson:"changeCategory,omitempty" json:"changeCategory,omitempty"`
	GOQualif        []string      `bson:"goQualif,omitempty" json:"goQualif,omitempty"`
	GOPunishment    string        `bson:"goPunishment,omitempty" json:"goPunishment,omitempty"`
	GOAddPunish     string        `bson:"goAddPunish,omitempty" json:"goAddPunish,omitempty"`
	GORegim         Regim         `bson:"goRegim,omitempty" json:"goRegim,omitempty"`
	GOPosition      GoPos         `bson:"goPosition,omitempty" json:"goPosition,omitempty"`
	CourtOrder      OrderKind     `bson:"courtOrder,omitempty" json:"courtOrder,omitempty"`
	CourtFasten     bool          `bson:"courtFasten,omitempty" json:"courtFasten,omitempty"`
	CourtQualif     []string      `bson:"courtQualif,omitempty" json:"courtQualif,omitempty"`
	CourtPunishment string        `bson:"courtPunishment,omitempty" json:"courtPunishment,omitempty"`
	CourtAddPunish  string        `bson:"courtAddPunish,omitempty" json:"courtAddPunish,omitempty"`
	CourtRegim      Regim         `bson:"courtRegim,omitempty" json:"courtRegim,omitempty"`
	Convicted       int           `bson:"convicted,omitempty" json:"convicted,omitempty"`
	Justified       int           `bson:"justified,omitempty" json:"justified,omitempty"`
	RehabTotal      int           `bson:"rehabTotal,omitempty" json:"rehabTotal,omitempty"`
	RehabHard       int           `bson:"rehabHard,omitempty" json:"rehabHard,omitempty"`
	GORehabTotal    int           `bson:"goRehabTotal,omitempty" json:"goRehabTotal,omitempty"`
	GORehabHard     int           `bson:"goRehabHard,omitempty" json:"goRehabHard,omitempty"`
	PreRestraint    Restraint     `bson:"preRestraint,omitempty" json:"preRestraint,omitempty"`
	AfterRestraint  Restraint     `bson:"afterRestraint,omitempty" json:"afterRestraint,omitempty"`
	PrDecree        Decree        `bson:"prDecree,omitempty" json:"prDecree,omitempty"` // chastnoe postanovlenie
	PrDecreeProc    bool          `bson:"prDecreeProc,omitempty" json:"prDecreeProc,omitempty"`
	Claim           float64       `bson:"claim,omitempty" json:"claim,omitempty"`
	ClaimSatisf     float64       `bson:"claimSatisf,omitempty" json:"claimSatisf,omitempty"`
	RealEvidence    bool          `bson:"realEvidence,omitempty" json:"realEvidence,omitempty"`
	ReportCheck     time.Time     `bson:"reportCheck,omitempty" json:"reportCheck,omitempty"`
	AppealSubj      AppSubj       `bson:"appealSubj,omitempty" json:"appealSubj,omitempty"`
	AppealDate      time.Time     `bson:"appealDate,omitempty" json:"appealDate,omitempty"`
	ObjectionDate   time.Time     `bson:"objectionDate,omitempty" json:"objectionDate,omitempty"`
	PowerCondemn    bool          `bson:"powerCondemn,omitempty" json:"powerCondemn,omitempty"`
	Notes           string        `bson:"notes,omitempty" json:"notes,omitempty"`
}

type CondKind int

const (
	ObvinitPrigovor CondKind = iota
	OpravdatPrigovor
	Prekrashenie
)

var conds = []string{
	"Обвинительный приговор",
	"Оправдательный приговор",
	"Прекращение",
}

//go:generate $GOBIN/tmplr -t ../tmpls/SGBsonHead --subpkg=types -o SGBson
//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=CondKind --lst=conds

type SubjectKind int

const (
	Predprinimatel SubjectKind = iota
	Prokuror
	SledovatelSK
	SledovatelMVD
	ChlenSFFS
	ChlenGDFS
	DeputeeReg
	DeputeeOMS
	OtherElected
	Judge
	Lawer
)

var subjects = []string{
	"предприниматель",
	"прокурор",
	"следователь СК",
	"следователь МВД",
	"член Совета Федерации ФС РФ",
	"депутат ГосДумы РФ",
	"депутат субъекта РФ",
	"депутат ОМС",
	"иное выборное лицо ОМС",
	"судья",
	"адвокат",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=SubjectKind --lst=subjects

type CrimeK int

const (
	Corruption CrimeK = iota
	Ecology
	TEK
	Drugs
	NatProjects
	NDS
	Extrem
	Terror
	Commerc
	Legalisation
	ZKH
	CrimeBankroc
	Economy
	FinanPiram
	Budget
	PlayBussiness
	OrgPrestupn
	FsspMvdMcs
)

var crimes = []string{
	"коррупция",
	"экология",
	"ТЭК",
	"наркотики",
	"нацпроекты",
	"НДС",
	"экстремизм",
	"терроризм",
	"в отношении предпринимателя",
	"легализация",
	"ЖКХ",
	"кримин. банкрот.",
	"налоги/экономика",
	"фин.пирамиды",
	"бюджет",
	"игорный бизнес",
	"орг.прест., бандитизм",
	"в отношении ФССП,МВД,МЧС",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=CrimeK --lst=crimes

type CrimeCat int

const (
	Low CrimeCat = iota
	Middle
	High
	SuprHigh
)

var crimecats = []string{
	"небольшой",
	"средней",
	"тяжкое",
	"особо тяжкое",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=CrimeCat --lst=crimecats

type Regim int

const (
	VK Regim = iota
	Poselenie
	Obshego
	Strogogo
	Osobogo
	UK73
	UK82
	UK821
	UK83
	UK84
)

var regims = []string{
	"ВК",
	"поселение",
	"общий",
	"строгий",
	"особый",
	"73 УК",
	"81 УК",
	"82 УК",
	"82-1 УК",
	"83 УК",
	"84 УК",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=Regim --lst=regims

type GoPos int

const (
	Soglasie GoPos = iota
	PolnOtkaz
	ChastichnOtkazTyazhk
	ChastichnOtkaz
	IzmenenQualif
	HodatVozvrat
	HodatPrekrash
	IzmenPodsudnosti
)

var goposes = []string{
	"согласие с обвинением",
	"полный отказ",
	"частичный отказ по наиболее тяжкому",
	"частичный отказ",
	"изменение квалификации",
	"ход-во по 237",
	"ход-во о прекращении",
	"изменение подсудности",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=GoPos --lst=goposes

type OrderKind int

const (
	Obshiy OrderKind = iota
	Osobiy
	ObshiyPosleOsob
	DosudSoglashOsob
	DosudSoglashObshiy
	ZaochnoVOtsutst
	Sokrashen
)

var orders = []string{
	"общий",
	"особый",
	"общий после особого",
	"досудебное соглашение особый порядок",
	"досудебное соглашение общий порядок",
	"заочно в отсутствие подсудимого",
	"сокращенная форма дознания",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=OrderKind --lst=orders

type Restraint int

const (
	NoRestr Restraint = iota
	ObyazatYavka
	Podpiska
	LichnPoruchit
	PrismortNesoversh
	Zalog
	DomashnArrest
	Arrest
)

var restraints = []string{
	"не избиралась",
	"обязательство о явке",
	"подписка о невыезде",
	"личное поручительство",
	"присмотр за н/л обвиняемым",
	"залог",
	"домашний арест",
	"заключение под стражу",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=Restraint --lst=restraints

type Decree int

const (
	GO Decree = iota
	Sud
	Zashita
	Poterpevsh
	Inie
)

var decrees = []string{
	"г/о",
	"суд",
	"защита",
	"потерпевший",
	"иные",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=Decree --lst=decrees

type AppSubj int

const (
	GOSubj AppSubj = iota
	ZashitaSubj
	OsuzhdenSubj
	PoterpevshSubj
	GrIstec
	GrOtvetchik
)

var appsubjs = []string{
	"г/о",
	"защита",
	"осужденный",
	"потерпевший",
	"гр. истец",
	"гр. ответчик",
}

//go:generate $GOBIN/tmplr -t "../tmpls/SGBson" --subpkg=types --append -o SGBson --str=AppSubj --lst=appsubjs

// func createSetBson(raw bson.Raw, sl []string) (int, error) {
//     var res string
//     var reti int
//     var ok bool
//     if err := raw.Unmarshal(&res); err != nil {
//         return 0, err
//     } else {
//         if reti, ok = stringInSlice(res, sl); !ok {
//             return 0, errors.New("Out of range in SetBSON")
//         }
//     }
//     return reti, nil
// }

// func createGetBSON(i int, sl []string) (interface{}, error) {
//     if i < len(sl) && i > -1 {
//         return sl[i], nil
//     } else {
//         return nil, errors.New("Out of range in GetBSON")
//     }
// }

// func createString(i int, sl []string) string {
//     if i < len(sl) && i > -1 {
//         return sl[i]
//     } else {
//         return ""
//     }
// }

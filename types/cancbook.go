package types

import (
// "time"
)

type CancBookV1 struct {
	Pp            int        `bson:"pp,omitempty" json:"pp,omitempty"`
	Number        int        `bson:"number,omitempty" json:"number,omitempty"`
	Init          DateT      `bson:"init,omitempty" json:"init,omitempty"`
	Fact          string     `bson:"fact,omitempty" json:"fact,omitempty"` // fabula
	Victim        VictimLst  `bson:"victim,omitempty" json:"victim,omitempty"`
	MinorVictim   BoolT      `bson:"minorVictim,omitempty" json:"minorVictim,omitempty"`
	Damage        FloatT     `bson:"damage,omitempty" json:"damage,omitempty"`
	Accused       AccusedLst `bson:"accused,omitempty" json:"accused,omitempty"` // obvinyaemiy
	Qualification string     `bson:"qualification,omitempty" json:"qualification,omitempty"`
	ArrestDate    string     `bson:"arrestDate,omitempty" json:"arrestDate,omitempty"`
	Evidence      string     `bson:"evidence,omitempty" json:"evidence,omitempty"`
	Conditions    CondT      `bson:"conditions,omitempty" json:"conditions,omitempty"`
	MinorAccused  BoolT      `bson:"minorAccused,omitempty" json:"minorAccused,omitempty"`
	Detention     string     `bson:"detention,omitempty" json:"detention,omitempty"`
}

package types

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DateT struct {
	Date time.Time `json:"date,omitempty"`
	Name string    `json:"name,omitempty"`
}

func (r *DateT) UnmarshalCSV(in string) error {
	if len(in) < 15 {
		return nil
	}
	stl := [2]string{}
	stl[0] = in[:10]
	stl[1] = in[strings.Index(in, " ")+1:]
	if t, err := time.Parse("02.01.2006", stl[0]); err != nil {
		return nil
	} else {
		r.Date = t
		r.Name = stl[1]
	}
	return nil
}

type VictimLst []string

func (r *VictimLst) UnmarshalCSV(in string) error {
	*r = VictimLst(strings.Split(in, ", "))
	return nil
}

type AccusedLst []string

func (r *AccusedLst) UnmarshalCSV(in string) error {
	*r = AccusedLst(strings.Split(in, ", "))
	return nil
}

type CondT struct {
	Termination string    `json:"termination,omitempty"`
	Suspend     string    `json:"suspend,omitempty"`
	ToCourt     time.Time `json:"toCourt,omitempty"`
}

var revsu = regexp.MustCompile("в с/у[ ,0-9]+([0-9]{2}.[0-9]{2}.[0-9]{4})")
var revsud = regexp.MustCompile("в суд[ ,0-9]+([0-9]{2}.[0-9]{2}.[0-9]{4})")

func (r *CondT) UnmarshalCSV(in string) error {
	// fmt.Println(in)
	vsud := ""
	if lst := revsu.FindStringSubmatch(in); len(lst) > 0 {
		vsud = lst[1]
	} else if lst := revsud.FindStringSubmatch(in); len(lst) > 0 {
		vsud = lst[1]
	} else {
		return nil
	}
	if t, err := time.Parse("02.01.2006", vsud); err == nil {
		r.ToCourt = t
	}
	return nil
}

type BoolT bool

func (r *BoolT) UnmarshalCSV(in string) error {
	// fmt.Println("Parse BoolT:", in)
	in = strings.TrimSpace(in)
	if in == "" {
		*r = BoolT(false)
	} else if in == "+" {
		*r = BoolT(true)
	}
	return nil
}

type FloatT float64

func (r *FloatT) UnmarshalCSV(in string) error {
	// fmt.Println("Parse FloatT:", in)
	in = strings.TrimSpace(in)
	if in == "" {
		*r = FloatT(0.0)
	} else {
		if x, err := strconv.ParseFloat(in, 64); err == nil {
			*r = FloatT(x)
		} else if idx := strings.Index(in, "="); idx != -1 {
			v := strings.TrimSpace(in[idx+1:])
			v = strings.Replace(v, ",", ".", 1)
			if x, err = strconv.ParseFloat(v, 64); err == nil {
				*r = FloatT(x)
			} else {
				return err
			}
		}
	}
	return nil
}

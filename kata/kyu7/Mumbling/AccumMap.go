package mumbling

import "strings"

type Chain []string

func AccumMap(s string) string {
	return Split(s).Map(func(v string, i int) string {
		v = strings.ToLower(v)
		return strings.ToUpper(v) + strings.Repeat(v, i)
	}).Join(`-`)

}

func Split(s string) *Chain {
	p := new(Chain)
	*p = strings.Split(s, ``)
	return p
}

func (chain *Chain) Map(f func(string, int) string) *Chain {
	for i, v := range *chain {
		(*chain)[i] = f(v, i)
	}

	return chain
}

func (chain *Chain) Join(s string) string {
	return strings.Join(*chain, s)
}

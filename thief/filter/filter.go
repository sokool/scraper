package filter

import (
	"strings"
	"regexp"
	"fmt"
)

type filter struct {
	name   string
	params []string
}

func (self *filter) run(in string) (string, error) {
	f, ok := filters[self.name]
	if !ok {
		return "", fmt.Errorf("Given [%s] filter not exists", self.name)
	}

	return f(in, self.params), nil
}

var filters = map[string]func(string, []string) (string){
	"num" : num,
	"trim": trim,
	"upper": upper,
	"rspaces": rspaces,
	"regexp": regularExpression,
	"alpha": alpha,
	"title": title,
	"wrap": wrap,
	"cdata": cdata,
}

func parse(in string) []*filter {
	out := []*filter{}

	sF := strings.Split(in, ",")
	if len(sF) == 0 {
		return out
	}

	for _, f := range sF {
		params := strings.Split(f, ":")
		name := strings.TrimSpace(params[0])
		out = append(out, &filter{name, params[1:]})
	}

	return out
}

func regularExpression(in string, params []string) string {
	return strings.Join(regexp.MustCompile(params[0]).FindAllString(in, -1), "")
}

func rspaces(in string, params []string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(in, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")

	return final
}

func wrap(in string, params []string) string {
	return params[0] + in + params[1]
}

func cdata(in string, params []string) string {
	return wrap(in, []string{"<![CDATA[", "]]>"})
}

func num(in string, params []string) string {
	return regularExpression(in, []string{"[0-9]+"})
}

func alpha(in string, params []string) string {
	return regularExpression(in, []string{"[A-Za-z\\s-]+"})
}

func title(in string, params []string) string {
	return "---IMPLEMENTIT---"
}

func trim(in string, params []string) string {
	return strings.TrimSpace(in)
}

func upper(in string, params []string) string {
	return strings.ToUpper(in)
}

func Run(filter string, input string) (string, error) {
	if filter == "" {
		return "", fmt.Errorf("Given filter is empty")
	}

	var output string
	var err error
	for _, f := range parse(filter) {
		if output != "" {
			input = output
		}
		output, err = f.run(input)

		if err != nil {
			return "", err
		}
	}

	return output, nil

}
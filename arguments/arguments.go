package arguments

import "strings"

// Assemble 根据参数数组组合成一行命令
func Assemble(args []string) string {
	builder := commandLineBuilder{}
	for _, part := range args {
		builder.append(part)
	}
	return builder.make()
}

// Parse 把一行命令转换为参数数组
func Parse(line string) []string {
	reader := argumentReader{}
	reader.init(line)
	list := make([]string, 0)
	for {
		part, ok := reader.readNextPart()
		if ok {
			if part == "" {
				continue
			}
			list = append(list, part)
		} else {
			break
		}
	}
	return list
}

// UnwrapAll 把所有的参数解封装
func UnwrapAll(args []string) []string {
	aa := argumentTools{}
	src := args
	dst := make([]string, 0)
	for _, item := range src {
		item = aa.unwrapString(item)
		dst = append(dst, item)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type argumentReader struct {
	data []rune
	p    int
}

func (inst *argumentReader) init(line string) {
	inst.data = []rune(line)
	inst.p = 0
}

func (inst *argumentReader) skipSpace() {
	p0 := inst.p
	data := inst.data
	end := len(data)
	for p := p0; p < end; p++ {
		ch := data[p]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			continue
		} else {
			inst.p = p
			break
		}
	}
}

func (inst *argumentReader) readChar() rune {
	p := inst.p
	data := inst.data
	end := len(data)
	if p < end {
		inst.p++
		return data[p]
	}
	return 0
}

func (inst *argumentReader) getChar() rune {
	p := inst.p
	data := inst.data
	end := len(data)
	if p < end {
		return data[p]
	}
	return 0
}

func (inst *argumentReader) isTokenChar(ch rune) bool {
	if '0' <= ch && ch <= '9' {
		return true
	} else if 'a' <= ch && ch <= 'z' {
		return true
	} else if 'A' <= ch && ch <= 'Z' {
		return true
	} else if strings.ContainsRune("-+/*_.,", ch) {
		return true
	}
	return false
}

func (inst *argumentReader) readWrappedString(mark rune) (string, bool) {
	p0 := inst.p
	data := inst.data
	builder := strings.Builder{}
	end := len(data)
	p := p0
	countMark := 0
	for ; p < end; p++ {
		ch := data[p]
		if ch == mark {
			countMark++
			if countMark == 2 {
				p++
				break
			}
		} else {
			if countMark == 1 {
				builder.WriteRune(ch)
			}
		}
	}
	inst.p = p
	return builder.String(), true
}

func (inst *argumentReader) readToken() (string, bool) {
	p0 := inst.p
	data := inst.data
	builder := strings.Builder{}
	end := len(data)
	p := p0
	for ; p < end; p++ {
		ch := data[p]
		if inst.isTokenChar(ch) {
			builder.WriteRune(ch)
		} else {
			break
		}
	}
	inst.p = p
	return builder.String(), true
}

func (inst *argumentReader) readNextPart() (string, bool) {
	inst.skipSpace()
	ch := inst.getChar()
	if ch == '=' {
		inst.readChar()
		return "", true
	} else if ch == '\'' {
		return inst.readWrappedString(ch)
	} else if ch == '"' {
		return inst.readWrappedString(ch)
	} else if inst.isTokenChar(ch) {
		return inst.readToken()
	}
	return "", false
}

////////////////////////////////////////////////////////////////////////////////

type commandLineBuilder struct {
	builder strings.Builder
}

func (inst *commandLineBuilder) make() string {
	return inst.builder.String()
}

func (inst *commandLineBuilder) append(part string) {

	const sep = " "
	aa := argumentTools{}

	part = strings.TrimSpace(part)
	if part == "" {
		return
	}

	if inst.builder.Len() > 0 {
		inst.builder.WriteString(sep)
	}

	if aa.isNeedForWrapping(part) {
		part = aa.wrapString(part)
	}

	inst.builder.WriteString(part)
}

////////////////////////////////////////////////////////////////////////////////

type argumentTools struct {
}

func (inst *argumentTools) wrapString(str string) string {
	const mark1 = '\''
	const mark2 = '"'
	hasMark1 := strings.ContainsRune(str, mark1)
	hasMark2 := strings.ContainsRune(str, mark2)
	mark := mark2
	if hasMark1 && hasMark2 {
		return str
	} else if hasMark1 {
		mark = mark2
	} else if hasMark2 {
		mark = mark1
	}
	mkstr := string(mark)
	return mkstr + str + mkstr
}

func (inst *argumentTools) unwrapString(str string) string {

	str = strings.TrimSpace(str)
	length := len(str)
	if length < 2 {
		return ""
	}

	const mark1 = "'"
	if strings.HasPrefix(str, mark1) && strings.HasSuffix(str, mark1) {
		return str[1 : length-1]
	}

	const mark2 = "\""
	if strings.HasPrefix(str, mark2) && strings.HasSuffix(str, mark2) {
		return str[1 : length-1]
	}

	return str
}

func (inst *argumentTools) isStringWrapped(str string) bool {

	str = strings.TrimSpace(str)

	const mark1 = "'"
	if strings.HasPrefix(str, mark1) && strings.HasSuffix(str, mark1) {
		return true
	}

	const mark2 = "\""
	if strings.HasPrefix(str, mark2) && strings.HasSuffix(str, mark2) {
		return true
	}

	return false
}

func (inst *argumentTools) isNeedForWrapping(str string) bool {
	str = strings.TrimSpace(str)
	chs := []rune{'\t', ' ', '\r', '\n', '=', '\'', '"'}
	return strings.ContainsAny(str, string(chs))
}

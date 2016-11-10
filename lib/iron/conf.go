package iron

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	sectionRegex = regexp.MustCompile(`^.*\[(.*)\]$`)
	assignRegex  = regexp.MustCompile(`^([^=]+)=(.*)$`)
)

// ErrSyntax is returned when there is a syntax error in an INI file.
type ErrSyntax struct {
	Line   int
	Source string // The contents of the erroneous line, without leading or trailing whitespace
}

func (e ErrSyntax) Error() string {
	return fmt.Sprintf("invalid INI syntax on line %d: %s", e.Line, e.Source)
}

// A ConfFile represents a parsed INI file.
type ConfFile map[string]Section

// A Section represents a single section of an INI file.
type Section map[string]string

// Returns a named Section. A Section will be created if one does not already exist for the given name.
func (f ConfFile) Section(name string) Section {
	section := f[name]
	if section == nil {
		section = make(Section)
		f[name] = section
	}
	return section
}

// Looks up a value for a key in a section and returns that value, along with a boolean result similar to a map lookup.
func (f ConfFile) Get(section, key string) (value string, ok bool) {
	if s := f[section]; s != nil {
		value, ok = s[key]
	}
	return
}

func (f ConfFile) MustGet(section, key, defaultValue string) string {
	if confStr, ok := f.Get(section, key); false == ok {
		return defaultValue
	} else {
		return confStr
	}
}

// 获取某值，如果获取失败则设置新值
func (f ConfFile) GetOrSet(section, key, defaultValue string) {
	if _, ok := f.Get(section, key); false == ok {
		f.Set(section, key, defaultValue)
	}
}

func (f ConfFile) Set(section, key, value string) {
	s, ok := f[section]
	if false == ok {
		s = make(map[string]string)
	}
	s[key] = value
	f[section] = s
	return
}

/*
// Looks up a value for a key in a section and returns that value, along with a boolean result similar to a map lookup.
func (f ConfFile) GetInt32(section, key string) (value int, ok bool) {
	if s := f[section]; s != nil {
		value, ok = s[key]
		value, ok = strconv.ParseInt(value, 10, 32)
	}
	return
}

// Looks up a value for a key in a section and returns that value, along with a boolean result similar to a map lookup.
func (f ConfFile) GetFloat32(section, key string) (value float32, ok bool) {
	if s := f[section]; s != nil {
		value, ok = s[key]
	}
	return
}
*/

// Loads INI data from a reader and stores the data in the ConfFile.
func (f ConfFile) Load(in io.Reader) (err error) {
	bufin, ok := in.(*bufio.Reader)
	if !ok {
		bufin = bufio.NewReader(in)
	}
	return parseConfFile(bufin, f)
}

// Loads INI data from a named file and stores the data in the ConfFile.
func (f ConfFile) LoadConfFile(file string) (err error) {
	in, err := os.Open(file)
	if err != nil {
		return
	}
	defer in.Close()
	return f.Load(in)
}

func parseConfFile(in *bufio.Reader, file ConfFile) (err error) {
	section := ""
	lineNum := 0
	for done := false; !done; {
		var (
			line      string
			lineBytes []byte
		)
		if lineBytes, _, err = in.ReadLine(); err != nil {
			if err == io.EOF {
				done = true
			} else {
				return
			}
		}
		line = string(lineBytes)

		lineNum++
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			// Skip blank lines
			continue
		}
		if line[0] == ';' || line[0] == '#' {
			// Skip comments
			continue
		}

		if groups := assignRegex.FindStringSubmatch(line); groups != nil {
			key, val := groups[1], groups[2]
			key, val = strings.TrimSpace(key), strings.TrimSpace(val)
			file.Section(section)[key] = val
		} else if groups := sectionRegex.FindStringSubmatch(line); groups != nil {
			name := strings.TrimSpace(groups[1])
			section = name
			// Create the section if it does not exist
			file.Section(section)
		} else {
			return ErrSyntax{lineNum, line}
		}

	}
	return nil
}

// Loads and returns a ConfFile from a reader.
func Load(in io.Reader) (ConfFile, error) {
	file := make(ConfFile)
	err := file.Load(in)
	return file, err
}

// Loads and returns an INI ConfFile from a file on disk.
func LoadConfFile(filename string) (ConfFile, error) {
	file := make(ConfFile)
	err := file.LoadConfFile(filename)
	return file, err
}

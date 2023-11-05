package twoFA

import (
	"bytes"
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Keychain interface {
	GetAllNames() []string
	Add(name string, size int, key string, isHOTP bool) error
	GenerateCode(name string) string
}

type keychainImpl struct {
	filePath string
	data     []byte
	keys     map[string]Key
}

const counterLen = 20

func (c *keychainImpl) GetAllNames() []string {
	var names []string
	for name := range c.keys {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (c *keychainImpl) Add(name string, size int, key string, isHOTP bool) error {
	key += strings.Repeat("=", -len(key)&7) // pad to 8 bytes //TODO: understand why is this negative
	if _, err := decodeKey(key); err != nil {
		return err
	}
	line := fmt.Sprintf("%s %d %s", name, size, key)
	if isHOTP {
		line += " " + strings.Repeat("0", 20)
	}
	line += "\n"

	f, err := os.OpenFile(c.filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("opening keychain: %v", err)
	}
	err = f.Chmod(0600)
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(line)); err != nil {
		log.Fatalf("adding key: %v", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("adding key: %v", err)
	}
	return nil
}

func (c *keychainImpl) GenerateCode(name string) string {
	key, ok := c.keys[name]
	if !ok {
		log.Fatalf("no such key %q", name)
	}

	var code int
	code = totp(key.raw, time.Now(), key.digits)
	return fmt.Sprintf("%0*d", key.digits, code)
}
func decodeKey(key string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(strings.ToUpper(key))
}

func Init(filename string) Keychain {
	chain := &keychainImpl{
		filePath: filename,
		data:     make([]byte, 0),
		keys:     make(map[string]Key),
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return chain
		}
		log.Fatal(err)
	}

	chain.data = data

	lines := bytes.SplitAfter(data, []byte("\n"))
	offset := 0

	for i, line := range lines {
		lineNo := i + 1
		offset += len(line)

		words := bytes.Split(bytes.TrimSuffix(line, []byte("\n")), []byte(" "))

		if len(words) == 1 && len(words[0]) == 0 {
			//empty line
			continue
		}

		if len(words) >= 3 && len(words[1]) == 1 &&
			'6' <= words[1][0] && words[1][0] <= '8' {
			var key Key
			name := string(words[0])

			key.digits = int(words[1][0] - '0')

			rawKey, err := decodeKey(string(words[2]))
			if err != nil {
				log.Printf("%s:%d: malformed key, error : %+v", chain.filePath, lineNo, err)
			} else {
				key.raw = rawKey
				if len(words) == 3 {
					chain.keys[name] = key
				} else if len(words) == 4 && len(words[3]) == counterLen {
					_, err := strconv.ParseUint(string(words[3]), 10, 64)
					if err == nil {
						// Valid counter.
						key.offset = offset - counterLen

						if line[len(line)-1] == '\n' {
							key.offset--
						}
						chain.keys[name] = key
					}
				}
			}
		}

	}

	return chain
}

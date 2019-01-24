package seed

import "bufio"
import "os"
import "bytes"
import "fmt"
import "strings"

/*
	Possible names for this format?

	SEED
	SOON
*/

func openIML(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return make(map[string]string)
	}
	defer file.Close()

	var reader = bufio.NewReader(file)
	var result = make(map[string]string)
	var buffer bytes.Buffer
	
	var last byte
	for {
		var token, err = reader.ReadByte()
		if err != nil {
			break
		}
		
		if (last == '\n' || last == 0) && token == '\n' {
			continue
		} else if token == '\n' {

			var block, err = reader.ReadString('}')
			if err != nil {
				break
			}
			
			result[strings.TrimSpace(buffer.String())] = strings.TrimSpace(strings.Replace(block[:len(block)-1], "\t", "", -1))
			buffer.Reset()
			last = 0
			continue
		}

		if token == ':' {
			var line, err = reader.ReadString('\n')
			if err != nil {
				break
			}
		
			result[strings.TrimSpace(buffer.String())] = strings.TrimSpace(line[:len(line)-1])
			buffer.Reset()
			last = 0
			continue
		}
		
		buffer.WriteByte(token)
		last = token
	}

	return result
}
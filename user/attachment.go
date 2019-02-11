package user

import (
	"strconv"
	"mime/multipart"
	"os"
	"path/filepath"
	"io"
	"io/ioutil"
	"encoding/base64"
	"math/big"
)

var AttachmentDirectory = filepath.Dir(os.Args[0])+"/attachments"
var AttachmentDirectoryLength = 1

type Attachment struct {
	paths []string

	heads []multipart.FileHeader
	files []multipart.File
}

//Retrieve the either the first attachment or if index is provided, the attachment at the specified index.
//This can be used to recieve files from the user.
func (user User) Attachment(index ...int) (attachment Attachment) {
	err := user.Request.ParseForm()
	if err != nil {
		println(err.Error())
		return
	}

	var i = 1
	if len(index) > 0 {
		i = index[0]
	}

	var j = 1
	for {
		file, header, err := user.Request.FormFile("attachment-"+strconv.Itoa(i)+"-"+strconv.Itoa(j))
		if err != nil {
			println("attachment-"+strconv.Itoa(i)+"-"+strconv.Itoa(j), err.Error())
			return
		}
	
		attachment.heads = append(attachment.heads, *header)
		attachment.files = append(attachment.files, file)

		j++
	}
}

func init() {
	os.MkdirAll(AttachmentDirectory, 0755)
	files, _ := ioutil.ReadDir(AttachmentDirectory)
	AttachmentDirectoryLength = len(files)+1
}

//Return the attachment as file paths, the files will be written to disk if not already there.
func (a Attachment) Paths() []string {
	if a.files == nil  {
		return nil
	}

	var result []string

	for i := range a.files {

		var filename = base64.RawURLEncoding.EncodeToString(big.NewInt(int64(AttachmentDirectoryLength)).Bytes())
		var extension = filepath.Ext(a.heads[i].Filename)
		
		output, err := os.Create(AttachmentDirectory+"/"+filename+extension)
		if err != nil {
			println(AttachmentDirectory+"/"+filename+extension, err.Error())
			continue
		}
		
		io.Copy(output, a.files[i])

		output.Close()

		result = append(result, AttachmentDirectory+"/"+filename+extension)
	}

	return result
}

//Return the attachment as a file path, the file will be written to disk if not already there.
func (a Attachment) Path() string {
	if len(a.files) == 0  {
		return ""
	}

	var i = 0

	var filename = base64.RawURLEncoding.EncodeToString(big.NewInt(int64(AttachmentDirectoryLength)).Bytes())
	var extension = filepath.Ext(a.heads[i].Filename)
	
	output, err := os.Create(AttachmentDirectory+"/"+filename+extension)
	if err != nil {
		println(AttachmentDirectory+"/"+filename+extension, err.Error())
		return ""
	}
	
	io.Copy(output, a.files[i])

	output.Close()

	return AttachmentDirectory+"/"+filename+extension
}
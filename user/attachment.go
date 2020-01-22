package user

import (
	"io"
	"mime/multipart"
	"strconv"
)

//Attachment is file that a user has attached to a request.
type Attachment struct {
	paths []string

	heads []multipart.FileHeader
	files []multipart.File
}

//Name returns the name of the attachment.
func (a Attachment) Name() string {
	if len(a.files) == 0 {
		return ""
	}

	return a.heads[0].Filename
}

//Open returns a reader to the attachment.
func (a Attachment) Open() io.ReadCloser {
	if len(a.files) == 0 {
		return nil
	}

	return a.files[0]
}

//Size returns the size of the attachment.
func (a Attachment) Size() int64 {
	if len(a.files) == 0 {
		return 0
	}

	return a.heads[0].Size
}

//Attachment retrieve either the first attachment or if index is provided, the attachment at the specified index.
//This can be used to recieve files from the user.
func (u Ctx) Attachment(index ...int) (attachment Attachment) {

	var i = 1
	if len(index) > 0 {
		i = index[0]
	}

	var j = 1
	for {
		file, header, err := u.r.FormFile("attachment-" + strconv.Itoa(i) + "-" + strconv.Itoa(j))
		if err != nil {
			println("attachment-"+strconv.Itoa(i)+"-"+strconv.Itoa(j), err.Error())
			return
		}

		attachment.heads = append(attachment.heads, *header)
		attachment.files = append(attachment.files, file)

		j++
	}
}

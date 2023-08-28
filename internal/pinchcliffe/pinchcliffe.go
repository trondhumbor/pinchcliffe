package pinchcliffe

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
)

const (
	SIZE_MAGIC             = 0x8
	SIZE_FILEHEADER_STRUCT = 0x108
)

func ExtractArchive(archive, outfolder string) {
	file, err := os.Open(archive)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Seek(SIZE_MAGIC, io.SeekStart)
	var filecount uint32
	binary.Read(file, binary.LittleEndian, &filecount)
	for i := 0; i < int(filecount); i++ {
		header := make([]byte, SIZE_FILEHEADER_STRUCT)
		binary.Read(file, binary.LittleEndian, &header)

		filename := ""
		for _, b := range header {
			if b == 0x00 {
				break
			}
			filename += string(b)
		}

		// The last 4 bytes (uint32) in the header indicates the filelength
		filelength := binary.LittleEndian.Uint32(header[len(header)-4:])

		content := make([]byte, filelength)
		binary.Read(file, binary.LittleEndian, &content)

		outpath := filepath.Join(outfolder, filename)
		os.MkdirAll(filepath.Dir(outpath), os.ModeDir)

		outfile, ferr := os.Create(outpath)
		if ferr != nil {
			panic(ferr)
		}
		defer outfile.Close()
		binary.Write(outfile, binary.LittleEndian, &content)
	}
}

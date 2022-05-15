package pinchcliffe

import (
	"encoding/binary"
	"os"
	"path/filepath"
)

func ExtractArchive(archive, outfolder string) {
	file, err := os.Open(archive)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Seek(0x8, 0)
	var filecount uint32
	binary.Read(file, binary.LittleEndian, &filecount)
	for i := 0; i < int(filecount); i++ {
		filename := ""
		for {
			var tmp byte
			binary.Read(file, binary.LittleEndian, &tmp)
			if tmp == 0x00 {
				break
			}
			filename += string(tmp)
		}

		file.Seek(0x104-(int64(len(filename))+1), 1)
		var filelength uint32
		binary.Read(file, binary.LittleEndian, &filelength)

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

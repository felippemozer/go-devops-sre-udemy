package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func EncryptLargeFiles(inFile, outPath string, key []byte) error {
	buf := make([]byte, 4096)
	in, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer out.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)

	for {
		n, err := in.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			out.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	out.Write(iv)
	return nil
}

func DecryptLargeFiles(inFile, outPath string, key []byte) error {
	in, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer in.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	fi, err := in.Stat()
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	msgLen := fi.Size() - int64(len(iv))

	_, err = in.ReadAt(iv, msgLen)
	if err != nil {
		return err
	}

	out, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer out.Close()
	buf := make([]byte, 4096)
	stream := cipher.NewCTR(block, iv)

	for {
		n, err := in.Read(buf)
		if n > 0 {
			if n > int(msgLen) {
				n = int(msgLen)
			}
			msgLen -= int64(n)
			stream.XORKeyStream(buf, buf[:n])
			out.Write(buf[:n])
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

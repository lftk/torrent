package torrent

import (
	"crypto/sha1"
	"errors"

	"github.com/zeebo/bencode"
)

// FileInfo represent file info
type FileInfo struct {
	MD5    string   `bencode:"md5sum"`
	Length int64    `bencode:"length"`
	Path   []string `bencode:"path"`
	PathU8 []string `bencode:"path.utf-8"`
}

// MetaData represent a metadata of torrent
type MetaData struct {
	// Single file
	Name   string `bencode:"name"`
	NameU8 string `bencode:"name.utf-8"`
	Length int64  `bencode:"length"`
	MD5    string `bencode:"md5sum"`
	// Multiple files
	Files       []FileInfo `bencode:"files"`
	PieceLength int64      `bencode:"piece length"`
	Pieces      string     `bencode:"pieces"`
	Private     int64      `bencode:"private"`
}

// MetaInfo represent a torrent
type MetaInfo struct {
	Data         MetaData   `bencode:"info"`
	Hash         []byte     `bencode:"info hash"`
	Announce     string     `bencode:"announce"`
	AnnounceList [][]string `bencode:"announce-list"`
	CreationDate int64      `bencode:"creation date"`
	Comment      string     `bencode:"comment"`
	CreatedBy    string     `bencode:"created by"`
	Encoding     string     `bencode:"encoding"`
}

// Encode MetaInfo to bytes
func Encode(m *MetaInfo) ([]byte, error) {
	return bencode.EncodeBytes(m)
}

// Decode bytes to MetaInfo
func Decode(b []byte) (m *MetaInfo, err error) {
	m = new(MetaInfo)
	err = decodeBytes(b, m)
	if err != nil {
		return
	}
	data, err := bencode.EncodeBytes(&m.Data)
	if err != nil {
		return
	}
	m.Hash, err = infoHash(data)
	return
}

// EncodeMetadata encode MetaData to bytes
func EncodeMetadata(m *MetaData) ([]byte, error) {
	return bencode.EncodeBytes(m)
}

// DecodeMetadata decode bytes to MetaData
func DecodeMetadata(b []byte) (m *MetaData, hash []byte, err error) {
	m = new(MetaData)
	err = decodeBytes(b, m)
	if err != nil {
		return
	}
	hash, err = infoHash(b)
	return
}

func infoHash(b []byte) ([]byte, error) {
	h := sha1.New()
	_, err := h.Write(b)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func decodeBytes(b []byte, val interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New("happen painc when decode bytes")
		}
	}()
	err = bencode.DecodeBytes(b, val)
	return
}

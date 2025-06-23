package main

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	"io"
	"net/http"
	"os"
)

func decompressionBomb() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)
	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
	r.Close()
}

func sqlInject(db *sql.DB, req *http.Request) {
	id := req.URL.Query().Get("id")
	db.Exec("DELETE FROM table WHERE Id = " + id)
} 
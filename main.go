// 2019, Georg Sauthoff <mail@gms.tf>
// SPDX-License-Identifier: GPL-3.0-or-later

package gonzofilter

import (
	"log"
	"strings"
)

type args struct {
	read_size       int
	in_filename     string
	out_filename    string
	header_filename string
	body_filename   string
	mime_filename   string
	db_filename     string
	dump_text       bool
	dump_words      bool
	dump_mark       bool
	dump_db         bool
	ham             bool
	spam            bool
	check           bool
	passthrough     bool
	tmpdir          string
	sandbox         bool
	sandbox_debug   bool
	h               bool
	help            bool
	verbose         bool
}

func (a *args) SetDbFilename(value string) {
	a.db_filename = value
}

// Classify message as spam or ham
func ClassifyMessage(message string, db_filename string) string {
	a := new(args)
	a.SetDbFilename(db_filename)
	in := strings.NewReader(message)
	is_ham, err := classify_file(in, a)
	if err != nil {
		log.Fatal(err)
	}
	if is_ham {
		return "HAM"
	}
	return "SPAM"
}

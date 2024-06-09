package decoder

import (
	"encoding/json"
	"os"

	ownadventure "github.com/gogojson/own_adventure/own_adventure"
)

func JsonDecoder(fileName string) (ownadventure.Book, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var book ownadventure.Book
	d := json.NewDecoder(f)
	if err := d.Decode(&book); err != nil {
		return nil, err
	}
	return book, nil
}

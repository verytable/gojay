package basic_struct

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/francoispqt/gojay"
	"github.com/viant/assertly"
	"bytes"
)

func TestMessage_Unmarshal(t *testing.T) {

	input := `{
  "Id": 1022,
  "Name": "name acc",
  "Price": 13.3,
  "Ints": [
    1,
    2,
    5
  ],
  "Floats": [
    2.3,
    4.6,
    7.4
  ],
  "SubMessageX": {
    "Id": 102,
    "Description": "abcd"
  },
  "MessagesX": [
    {
      "Id": 2102,
      "Description": "abce"
    }
  ],
  "SubMessageY": {
    "Id": 3102,
    "Description": "abcf"
  },
  "MessagesY": [
    {
      "Id": 5102,
      "Description": "abcg"
    },
    {
      "Id": 5106,
      "Description": "abcgg"
    }
  ],
  "IsTrue": true,
  "Payload": ""
}`

	var err error
	var data = []byte(input)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	assert.Nil(t, err)
	assertly.AssertValues(t, input, message)
}



func TestMessage_Marshal(t *testing.T) {

	input := `{
  "Id": 1022,
  "Name": "name acc",
  "Price": 13.3,
  "Ints": [
    1,
    2,
    5
  ],
  "Floats": [
    2.3,
    4.6,
    7.4
  ],
  "SubMessageX": {
    "Id": 102,
    "Description": "abcd"
  },
  "MessagesX": [
    {
      "Id": 2102,
      "Description": "abce"
    }
  ],
  "SubMessageY": {
    "Id": 3102,
    "Description": "abcf"
  },
  "MessagesY": [
    {
      "Id": 5102,
      "Description": "abcg"
    },
    {
      "Id": 5106,
      "Description": "abcgg"
    }
  ],
  "IsTrue": true,
  "Payload": ""
}`

	var err error
	var data = []byte(input)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	assert.Nil(t, err)

	var writer = new(bytes.Buffer)

	encoder :=  gojay.NewEncoder(writer)
	err = encoder.Encode(message)
	assert.Nil(t, err)
	var JSON = writer.String()
	assertly.AssertValues(t, input,  JSON)
}

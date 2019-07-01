package utils

import (
  "encoding/json"
)

func RespondWithJson(payload interface{}) (string, error) {
  response, err := json.Marshal(payload)

  return string(response), err
}

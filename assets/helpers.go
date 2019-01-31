package assets

import(
  "crypto/sha1"
  "encoding/hex"
)

func hash_function(s string) string{
  h := sha1.New()
  h.Write([]byte(s))
  result := hex.EncodeToString(h.Sum(nil))
  return result
}

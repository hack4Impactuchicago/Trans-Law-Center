package assets

import(
  "hash/fnv"
)

func hash_function(s string) int{
        h := fnv.New32a()
        result = h.Write([]byte(s))
        return result
}

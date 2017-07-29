package fulltextreplace

import (
  "regexp"
  "strings"
  "unicode"
)

var alphabetConv = unicode.SpecialCase{
  // uppercase letters
  unicode.CaseRange{
    Lo: 0xff21, // 'Ａ'
    Hi: 0xff3a, // 'Ｚ'
    Delta: [unicode.MaxCase]rune{
      0,
      0x0041 - 0xff21, // 'A' - 'Ａ'
      0,
    },
  },
  // lowercase letters
  unicode.CaseRange{
    Lo: 0xff41, // 'ａ'
    Hi: 0xff5a, // 'ｚ'
    Delta: [unicode.MaxCase]rune{
      0,
      0x0061 - 0xff41, // 'a' - 'ａ'
      0,
    },
  },
}

func StrReplace(strObject string) (string, error) {
  rep := regexp.MustCompile(`[0-9]|[０-９]|□|　|\x20|〔|〕|・|-|´|／|．|:|：|\(|\)|（|）|「|」`)
  str := rep.ReplaceAllString(strObject, "")
  str = strings.ToLowerSpecial(alphabetConv, str)
  str = strings.ToUpper(str)
  return str, nil
}
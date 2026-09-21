package piscine

func CamelToSnakeCase(s string) string {
  if s == "" {
   return ""
 }

  for i, r := range s {
  if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
     return s
  }

   if i == len(s)-1 && r >= 'A' && r <= 'Z' {
return s
  }

   if i > 0 && r >= 'A' && r <= 'Z' {
     prev := rune(s[i-1])
     if prev >= 'A' && prev <= 'Z' {
    return s
    }
  }
}

 result := ""
  for i, r := range s {
    if r >= 'A' && r <= 'Z' {
   if i != 0 {
   result += "_"
  }
   result += string(r)
  } else {
result += string(r)
}
}

return result
}

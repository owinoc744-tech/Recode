package piscine

func FirstWord(s string )string{
start := 0
for start < len (s) && (s[start] == ' ' ||s[start] == '\t'){
       start++
    }
end := start
for end <len(s) && s[end] != ' '&&s[end] != '\t'{
   end++
   }

return s[start:end]+"\n"
}

{{ $top10 := dbTopEntries "taco_bank" 10 10 }}
{{ if not $top10 }}
Chưa có ai có taco_bank.
{{ else }}
Top 10 người có taco_bank nhiều nhất:
{{ range $top10 }}
- User: <@{{ .UserID }}> - Taco_bank: {{ .Value }}
{{ end }}
{{ end }}
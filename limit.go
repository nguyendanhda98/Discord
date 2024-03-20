{{ $args := parseArgs 1 "<amount>" (carg "int" "amount of the limit taco")}}
{{ $taco_limit := ($args.Get 0) }}
{{ if gt $taco_limit 0}}
{{ dbSet 0 "taco_limit" ($args.Get 0) }}
Bạn đã set thành công {{ $args.Get 0 }} mỗi ngày.
{{ else }}
Taco mỗi ngày phải lớn hơn 0.
{{ end }}
{{ $args := parseArgs 3 "<user> <amount> <message>" (carg "user" "target user") (carg "int" "amount of taco") (carg "string" "message") }}
{{ $user_target := ($args.Get 0) }}
{{ $amount := ($args.Get 1) }}
{{ $message := ($args.Get 2) }}
 
{{ if eq $user_target.ID .User.ID }}
Bạn không thể tự tặng chính mình
{{ return }}
{{end}}
 
{{ if le $amount 0 }}
Số lượng taco phải lớn hơn 0.
{{ return }}
{{end}}
 
{{/* This is current time in UTC + 7 */}}
{{ $currentTime := currentTime.Add 25200000000000 }}
{{/* This is how many seconds until the next day*/}}
{{ $timeUntilNextDay := sub 86400 (mod $currentTime.Unix 86400) }}
 
{{ $taco_daily := (dbGet .User.ID "taco_daily") }}
{{ $taco_limit := (dbGet 0 "taco_limit") }}
 
{{ $taco_user_limit := (dbGet .User.ID "taco_limit" ) }}
{{ if $taco_user_limit}}
{{ $taco_limit = $taco_user_limit }}
{{ end }}
 
{{ if $taco_daily }}
 
{{ if lt (toInt $taco_daily.Value) $amount }}
Bạn không đủ taco để tặng. Bạn có {{$taco_daily.Value}}/{{ $taco_limit.Value}} taco ngày hôm nay.
{{ return}}
{{ end }}
 
{{ else }}
{{ $taco_daily = $taco_limit }}
{{ end }}
 
{{ dbSetExpire .User.ID "taco_daily" (sub $taco_daily.Value $amount) $timeUntilNextDay}}
{{ $user_target_taco_bank := dbIncr $user_target.ID "taco_bank" $amount }}
{{ .User.Mention }} đã tặng {{ $user_target.Mention }} **{{ $amount }}** :taco: với lời nhắn: {{ $message }}
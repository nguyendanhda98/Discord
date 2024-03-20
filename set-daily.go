{{ $args := parseArgs 2 "<user> <amount>" (carg "user" "target user") (carg "int" "amount of taco") }}
{{ $user_target :=  ($args.Get 0) }}
{{ $amount := ($args.Get 1) }}
 
{{ if lt $amount 0 }}
Taco daily phải lớn hơn hoặc bằng 0.
{{ return }}
{{ end }}
 
{{ dbSet $user_target.ID "taco_limit" $amount }}
{{ $user_target.Mention }} đã sẽ nhận được {{ $amount }} taco mỗi ngày.
{{ $args := parseArgs 3 "<user> <amount>" (carg "user" "target user") (carg "int" "amount of taco") }}
{{ $user_target :=  ($args.Get 0) }}
{{ $amount := ($args.Get 1) }}
 
{{ $user_target_taco_bank := dbIncr $user_target.ID "taco_bank" $amount }}
Bạn đã thêm thành công {{ $amount }} taco bank cho {{ $user_target.Mention }}.
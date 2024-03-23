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


{{ sendMessage nil ( print  .User.Mention " đã tặng " $user_target.Mention "**" $amount "** :taco: với lời nhắn: " $message ) }}

{{/* ============================= taco check ========================== */}}

{{ $user_target := ($args.Get 0) }}
{{ $taco_daily := dbGet $user_target.ID "taco_daily" }}
{{ $taco_bank := dbGet $user_target.ID "taco_bank" }}
{{ $taco_daily_value := $taco_daily.Value }}
{{ $taco_bank_value := $taco_bank.Value }}
{{/* 25200000000000 = 7 giờ */}}
{{ $currentTime := currentTime.Add 25200000000000 }}
{{ $timeUntilNextDay := sub 86400 (mod $currentTime.Unix 86400) }}
{{ $taco_limit := (dbGet 0 "taco_limit") }}

{{ if not $taco_daily }}
{{ $taco_daily_value = $taco_limit.Value }}
{{ dbSetExpire $user_target.ID "taco_daily" $taco_limit.Value $timeUntilNextDay}}
{{ end }}

{{ if not $taco_bank }}
{{ $taco_bank_value = 0 }}
{{ dbSet $user_target.ID "taco_bank" 0}}
{{ end }}

{{ $title := print "Ví Taco :taco:" }}
{{ $MessageDiscription := print ":gift: **Daily**: " $taco_daily_value "/" $taco_limit.Value " :taco:\n:bank: **Bank**: " $taco_bank_value " :taco:"}}
{{ $serverIcon := (joinStr "" "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".png") }}
{{ $avatar := print "https://cdn.discordapp.com/avatars/" $user_target.ID "/" $user_target.Avatar ".png" }}

{{$Msg := cembed
    "author" 
	(sdict 
		"name" $user_target.Globalname
		"url" ($user_target.AvatarURL "512") 
		"icon_url" ($avatar)  )
    "title" $title
    "thumbnail" (sdict "url" $avatar)
    "description" $MessageDiscription
        "footer" 
	(sdict 
		"text" .Guild.Name
		"icon_url" $serverIcon
	)
    "timestamp" currentTime
}}

{{ sendMessage nil $Msg }}
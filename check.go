{{ $args := parseArgs 1 "<user>" (carg "user" "target user")}}
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
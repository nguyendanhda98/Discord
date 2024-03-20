{{ $taco_daily := dbGet .User.ID "taco_daily" }}
{{ $taco_bank := dbGet .User.ID "taco_bank" }}
{{ $taco_daily_value := $taco_daily.Value }}
{{ $taco_bank_value := $taco_bank.Value }}
{{/* 25200000000000 = 7 giờ */}}
{{ $currentTime := currentTime.Add 25200000000000 }}
{{ $timeUntilNextDay := sub 86400 (mod $currentTime.Unix 86400) }}
{{ $taco_limit := (dbGet 0 "taco_limit") }}
 
{{ $taco_user_limit := (dbGet .User.ID "taco_limit" ) }}
{{ if $taco_user_limit}}
{{ $taco_limit = $taco_user_limit }}
{{ end }}
 
{{ if not $taco_daily }}
{{ $taco_daily_value = $taco_limit.Value }}
{{ dbSetExpire .User.ID "taco_daily" $taco_limit.Value $timeUntilNextDay}}
{{ end }}
 
{{ if not $taco_bank }}
{{ $taco_bank_value = 0 }}
{{ dbSet .User.ID "taco_bank" 0}}
{{ end }}
 
{{ $title := print "Ví Taco :taco:" }}
{{ $MessageDiscription := print ":gift: **Daily**: " $taco_daily_value "/" $taco_limit.Value " :taco:\n:bank: **Bank**: " $taco_bank_value " :taco:"}}
{{ $serverIcon := (joinStr "" "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".png") }}
{{ $avatar := print "https://cdn.discordapp.com/avatars/" .User.ID "/" .User.Avatar ".png" }}
 
{{$Msg := cembed
    "author" 
	(sdict 
		"name" .User.Globalname
		"url" (.User.AvatarURL "512") 
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
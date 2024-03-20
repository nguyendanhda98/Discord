{{ $args := parseArgs 0 "<page>" (carg "int" "page of top") }}
{{ $top_page := 1 }}
{{ $users_per_page := 5 }}
{{ if $args.IsSet 0 }}
	{{ $top_page = $args.Get 0 }}
{{ end }}
{{ $users_total := dbTopEntries "taco_bank" 100 0 }}
{{ $users_length := len $users_total }}
{{ $max_page := 1}}
{{ if gt $users_length $users_per_page }}
{{ $max_page = roundCeil (fdiv $users_length $users_per_page) }}
{{ end }}
{{ if gt $top_page (toInt $max_page) }}
	{{ $top_page = $max_page }}
{{ end }}
{{ $users := dbTopEntries "taco_bank" $users_per_page ( mult (sub $top_page 1) $users_per_page )}}
{{ $MessageDiscription := ""}}
{{ if $users }}
	{{ $rank := add (mult (sub $top_page 1) $users_per_page) 1}}
	{{ range $users }}
		{{ $user := userArg .UserID }}
		{{ $MessageDiscription = print $MessageDiscription (print $rank ". " $user.Mention ": " .Value " :taco:\n") }}
		{{ $rank = add $rank 1 }}
	{{ end }}
{{ end }}
{{ $title := print "Top Taco Bank :taco:" }}
{{ $serverIcon := (joinStr "" "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".png") }}
{{ $avatar := print "https://cdn.discordapp.com/avatars/" .User.ID "/" .User.Avatar ".png" }}

{{$Msg := cembed
		"author"
	(sdict
		"name" .User.Globalname
		"url" (.User.AvatarURL "512")
		"icon_url" ($avatar)  )
		"title" $title
		"description" $MessageDiscription
		"footer"
	(sdict
		"text" (print "Trang " $top_page "/" $max_page " • " .Guild.Name)
		"icon_url" $serverIcon
	)
		"timestamp" currentTime
}}

{{ sendMessage nil $Msg }}
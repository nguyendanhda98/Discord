{{ $args := parseArgs 0 "<page>" (carg "int" "page of shop") }}
{{ $shop_page := 1 }}
{{ $items_per_page := 5 }}
{{ if $args.IsSet 0 }}
	{{ $shop_page = $args.Get 0 }}
{{ end }}
{{ $shop_ID := 1 }}
{{ $items_total := dbGetPattern $shop_ID "%" 100 0 }}
{{ $items_length := len $items_total }}
{{ $max_page := 1}}
{{ if gt $items_length $items_per_page }}
{{ $max_page = roundCeil (fdiv $items_length $items_per_page) }}
{{ end }}
{{ if gt $shop_page (toInt $max_page) }}
	{{ $shop_page = $max_page }}
{{ end }}
{{ $items := dbGetPattern $shop_ID "%" $items_per_page ( mult (sub $shop_page 1) $items_per_page )}}
{{ $fields := cslice 
	(sdict "name" "Cửa hàng trống" "value" "Cửa hàng hiện đang trống, hãy quay lại sau." "inline" false)
}}

{{ if $items }}
	{{ $fields = cslice }}
	{{ range $items }}
		{{ $fields = $fields.Append (sdict "name" (print "ID: " .Key " - " .Value.content) "value" (print "- " .Value.price " :taco: - Còn lại: " .Value.amount) "inline" false) }}
	{{ end }}
{{ end }}

{{ $title := print "Shop Taco :taco:" }}
{{ $MessageDiscription := "Để mua, hãy dùng lệnh `taco shop-buy <id>`"}}
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
		"fields" $fields
		"footer" 
	(sdict 
		"text" (print "Trang " $shop_page "/" $max_page " • " .Guild.Name)
		"icon_url" $serverIcon
	)
		"timestamp" currentTime
}}

{{ sendMessage nil $Msg }}
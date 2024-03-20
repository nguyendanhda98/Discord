{{ $shop_log := dbGet 0 "shop_log" }}
{{ if not $shop_log }}
	{{ sendMessage nil "Chưa cài đặt kênh và vai trò để nhận shop log." }}
	{{ return }}
{{ end }}
{{ $log_channel := $shop_log.Value.channel }}
{{ $log_role := $shop_log.Value.role }}
{{ sendMessageNoEscape nil (print "Kênh: <#" $log_channel.ID "> - Vai trò: <@&" $log_role.ID ">") }}
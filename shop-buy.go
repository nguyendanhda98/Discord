{{ $shop_log := dbGet 0 "shop_log" }}
{{ if not $shop_log }}
	{{ sendMessage nil "Chưa cài đặt kênh và vai trò để nhận shop log. Hãy bảo admin cài đặt bằng lệnh ` taco shop-log-set`." }}
	{{ return }}
{{ end }}

{{ $args := parseArgs 0 "<id>" (carg "int" "ID của vật phẩm") }}
{{ $shop_ID := 1 }}
{{ $item_ID := $args.Get 0 }}
{{ $item := dbGet $shop_ID (joinStr "" $item_ID) }}
{{ if not $item }}
	{{ sendMessage nil (print "Không tìm thấy vật phẩm có ID: " $item_ID) }}
	{{ return }}
{{ end }}
	{{ $item_content := $item.Value.content }}
	{{ $item_price := $item.Value.price }}
	{{ $item_amount := $item.Value.amount }}
{{ if lt $item_amount 1 }}
	{{ sendMessage nil (print "Vật phẩm có ID: " $item_ID " đã hết hàng.") }}
	{{ return }}
{{ end }}
{{ $user_taco_bank := dbGet .User.ID "taco_bank" }}
{{ if lt (toInt $user_taco_bank.Value) $item_price }}
	{{ sendMessage nil (print "Bạn không đủ Taco để mua vật phẩm này. Bạn cần thêm " (sub $item_price $user_taco_bank.Value) " :taco:") }}
	{{ return }}
{{ end }}
{{ dbSet .User.ID "taco_bank" (sub (toInt $user_taco_bank.Value) $item_price) }}
{{ dbSet $shop_ID (joinStr "" $item_ID) (sdict "content" $item_content "price" $item_price "amount" (sub $item_amount 1)) }}
{{ sendMessage nil (print "Bạn đã mua thành công vật phẩm: " $item_content " với giá " $item_price " :taco:") }}

{{ $log_channel := $shop_log.Value.channel }}
{{ $log_role := $shop_log.Value.role }}
{{ $log_embed := cembed
	"author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "512"))	
	"description" (print "Mua vật phẩm: " $item_content " với giá " $item_price " :taco:")
	"footer" (sdict "text" (print "ID: " .User.ID))
	"timestamp" currentTime
}}
{{ sendMessage $log_channel.ID $log_embed }}
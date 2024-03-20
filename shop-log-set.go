{{ $args := parseArgs 2 "<channel> <role>" (carg "channel" "Kênh để gửi log") (carg "role" "Vai trò để gửi log") }}
{{ $log_channel := $args.Get 0 }}
{{ $log_role := $args.Get 1 }}
{{ $value := sdict "channel" $log_channel "role" $log_role }}
{{ dbSet 0 "shop_log" $value }}
{{ sendMessage nil (print "Đã cài đặt kênh " $log_channel.Mention " và vai trò " $log_role.Mention " để nhận shop log.") }}
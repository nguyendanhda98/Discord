{{ $args := parseArgs 1 "<id>" (carg "int" "id") }}
{{ $id := $args.Get 0 }}
{{ $shop_ID := 1 }}
 
{{ if not (dbGet $shop_ID $id) }}
Không thể xoá item id {{$id}} vào shop vì item không tồn tại.
{{ return }}
{{ end }}
 
{{ dbDel $shop_ID $id }}
Đã xoá item id {{$id}} khỏi shop.
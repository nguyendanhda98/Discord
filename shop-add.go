{{ $args := parseArgs 4 "<id> <price> <amount> <content>" (carg "int" "id") (carg "int" "price") (carg "int" "amount") (carg "string" "content") }}
{{ $id := $args.Get 0 }}
{{ $price := $args.Get 1 }}
{{ $amount := $args.Get 2 }}
{{ $content := $args.Get 3 }}
{{ $item := sdict "price" $price "amount" $amount "content" $content }}
{{ $shop_ID := 1 }}

{{ if (dbGet $shop_ID $id) }}
Không thể thêm item id {{$id}} vào shop vì item đã tồn tại.
{{ return }}
{{ end }}

{{ dbSet $shop_ID $id $item }}
Thêm thành công item id {{$id}} vào shop với giá {{$price}}, số lượng {{$amount}}, và nội dung {{$content}}.
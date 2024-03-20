{{ $title := print "Taco Help" }}
{{ $description := print "Taco là loại tiền tệ ảo trong server. Bạn có thể kiếm được taco hàng ngày, mua vật phẩm từ shop, tặng taco cho người khác và kiểm tra số lượng taco bạn có. Dưới đây là danh sách các lệnh taco có sẵn:\n\n" }}
{{ $fields := cslice
	(sdict "name" "Member" "value" "`taco help`: trợ giúp về các lệnh taco :taco:\n`taco me`: Kiểm tra số taco :taco: mà bạn có\n`taco give <user> <amount> <message>`: Tặng ai đó taco :taco: để cảm ơn họ\n`taco shop [page]`: Mở shop taco\n`taco shop-buy <id>`: Mua sản phẩm <id> bằng taco bank\n`taco top [page]`: Top thành viên có nhiều taco bank nhất" "inline" true)
	(sdict "name" "Admin" "value" "`taco limit <number>`: Giới hạn số taco được làm mới mỗi ngày\n`taco check <user>`: Kiểm tra taco của thành viên\n`taco daily-set <user> <amount>`: Set giới hạn taco daily cho thành viên\n`taco bank-edit <user> <amount>`: Thêm taco bank cho thành viên\n`taco shop-add <id> <price> <amount> <content>`: Thêm sản phẩm vào shop với id, giá, số lượng, nội dung\n`taco shop-edit <id> <price> <amount> <content>`: Sửa sản phẩm <id> với giá, số lượng, nội dung\n`taco shop-remove <id>`: Xoá sản phẩm <id> trong shop\n`taco shop-log` : Xem kênh và role nhận log khi người dùng mua hàng\n`taco shop-log-set <channel> <role>`: Set kênh và role nhận log người dùng khi mua hàng" "inline" true)
}}
{{ $author := sdict "name" .User.Globalname "icon_url" (.User.AvatarURL "512") }}
{{ $footer := sdict "text" "<> - bắt buộc, [] - tuỳ chọn" }}
{{ $timestamp := currentTime }}
{{ $msg := cembed
	"title" $title
	"description" $description
	"fields" $fields
	"author" $author
	"footer" $footer
	"timestamp" $timestamp
}}
{{ sendMessage nil $msg }}
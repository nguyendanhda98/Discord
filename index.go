{{ $taco_everyday := (dbGet .User.ID "taco_everyday").Value }}
{{ if $taco_everyday }}
{{ print "taco everyday is " $taco_everyday }}
{{ else }}
{{ dbSet .User.ID "taco_everyday" 10 }}
{{ end }}

{{$test := (dbGet 0 "cooldown_replace with name here").Value}}
{{ print $test}}
# digitoy_assignment

Not: Direkt olarak bitmiş el ile denemeyin, serileri hesaplarken K1-M1-S1 K1-K2-K3-K4
gibi hem aynı renk hem aynı sayı farklı renk gruplarını bir araya getirip olabilecek bütün subsetler
üzerinden hesaplama yaptım. 1-2-3-4-5-6-7-8-9-10-11-12-13-1 şeklinde bitmiş bir elin 2^100 den fazla subset kombinasyonu var.
([[1,2,3], [2,3,4], [5,6,7]], [[5,6,7,8], [12,13,1]]...) Cevabı 10 yıl sonra görebilirsiniz.
Zaman yetmediğinden dolayı orayı optimize etmeden bıraktım (ayrıca optimize nasıl edeceğimi konusunda bir fikrim yok henüz).
Düz bir şekilde çalıştırmakta sıkıntı yok, rastgele dağıtırken çok yüksek ihtimalle birkaç tane valid serisi olan
ıstaka oluşuyor, onları anında hesaplayabilir. 2 okey çıktığında ıstakanın geri kalanının ne kadar iyi geldiğine göre 3-4 saniye 
sürebiliyor.

Çalıştırmak için 
zip olarak indirip bir directorye çıkartın
"go build" komutunu kullandıktan sonra oluşan executable'ı çalıştırın.

dipnot: ANSI renk kodları ile printledim konsola. kullandığınız terminalin ANSI color code support'u olması gerekiyor.
Öteki türlü [97m5 şeklinde sayılar görebilirsiniz. 
# digitoy_assignment

ANSI renk kodları ile printledim konsola. kullandığınız terminalin ANSI color code support'u olması gerekiyor.
Öteki türlü [97m5 şeklinde sayılar görebilirsiniz. Direkt olarak bitmiş el ile denemeyin, serileri hesaplarken K1-M1-S1 K1-K2-K3-K4
gibi hem aynı renk hem aynı sayı farklı renk gruplarını bir araya getirip olabilecek bütün subsetler
üzerinden hesaplama yaptım. 1-2-3-4-5-6-7-8-9-10-11-12-13-1 şeklinde bitmiş bir elin 2^100 den fazla subset kombinasyonu var.
([[1,2,3], [2,3,4], [5,6,7]], [[5,6,7,8], [12,13,1]]...) Cevabı 10 yıl sonra görebilirsiniz.
Düz bir şekilde çalıştırmakta sıkıntı yok, rastgele dağıtırken çok yüksek ihtimalle birkaç tane valid serisi olan
ıstaka oluşuyor, onları anında hesaplayabilir. 2 okey çıktığında ıstakanın geri kalanının ne kadar iyi geldiğine göre 3-4 saniye 
sürebiliyor. 

not: İç içe 15 tane if else ve alt alta 10dan fazla if else ile bütün olabilecek caseleri handle etmek yerine bu çözümü daha güzel bulduğumdan
bu şekilde yaptım ama şu anki versiyonu 1-2 tane edge case'i kaçırıyor ve puan hesaplaması kısmının optimize edilmesi gerekiyor. Zaman kısıtlamasından ötürü bu şekliyle bıraktım.
Ayrıca puanları hesaplarken nasıl optimize edeceğim konusunda da henüz bir fikrim yok, sonuçta gidip en büyük serileri almak mutlak sonucu vermiyor (küçük 3 tane serinin toplamı direkt olarak eli bitiredebilir).

Çalıştırmak için 
zip olarak indirip bir directorye çıkartın
"go build" komutunu kullandıktan sonra oluşan executable'ı çalıştırın.


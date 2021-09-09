USE sample;
CREATE TABLE players(
player_id INTEGER,
player_name VARCHAR(30),
player_rank VARCHAR(3),
average float,
HR INTEGER,
OPS float,
SB INTEGER,
PRIMARY KEY (player_id));

INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('0','Yuki Yanagita','SR','0.342','29','1.071','7');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('1','Munetaka Murakami','SR','0.307','28','1.012','11');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('2','Norichika Aoki','SR','0.317','18','0.981','2');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('3','Hideto Asamura','SR','0.28','32','0.969','1');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('4','Masataka Yoshida','SR','0.35','14','0.966','8');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('5','Seiya Suzuki','SR','0.3','25','0.953','6');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('6','Kensuke Kondo','SR','0.34','5','0.934','4');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('7','Yoshihiro Maru','SR','0.284','27','0.928','8');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('8','Keita Sano','SR','0.328','20','0.927','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('9','Yusuke Ohyama','SR','0.288','28','0.918','1');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('10','Takayuki Kajitani','SR','0.323','19','0.913','14');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('11','Kazuma Okamoto','SR','0.275','31','0.907','2');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('12','Stefen Romero','R','0.272','24','0.893','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('13','Hayato Sakamoto','R','0.289','19','0.879','4');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('14','Leonys Martin','R','0.234','25','0.866','7');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('15','Haruki Nishikawa','R','0.306','5','0.825','42');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('16','Jerry Sands','R','0.257','19','0.814','2');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('17','Sho Nakata','R','0.239','31','0.811','1');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('18','Cory Spangenberg','R','0.268','15','0.807','12');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('19','Hotaka Yamakawa','R','0.205','24','0.807','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('20','Toshiro Miyazaki','R','0.301','14','0.805','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('21','T-Okada','R','0.256','16','0.797','5');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('22','Shuhei Takahashi','R','0.305','7','0.794','1');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('23','Neftali Soto','R','0.252','25','0.793','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('24','Takumi Kuriyama','R','0.272','12','0.79','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('25','Shota Dohbayashi','R','0.279','14','0.787','17');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('26','Dayan Viciedo','R','0.267','17','0.776','3');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('27','Tetsuto Yamada','R','0.254','12','0.766','8');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('28','Yohei Ohshima','R','0.316','1','0.763','16');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('29','Justin Bour','R','0.243','17','0.76','1');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('30','Koji Chikamoto','R','0.293','9','0.759','31');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('31','Ryosuke Kikuchi','R','0.271','10','0.757','3');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('32','Hiroaki Shimauchi','R','0.281','8','0.755','9');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('33','Seiya Inoue','N','0.245','15','0.749','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('34','Hiroto Kobukata','N','0.288','3','0.745','17');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('35','Daichi Suzuki','N','0.295','4','0.744','1');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('36','Naoki Yoshikawa','N','0.274','8','0.734','11');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('37','Kurihara Ryoya','N','0.243','17','0.727','5');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('38','Ryo Watanabe','N','0.283','6','0.725','4');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('49','Ryuhei Matsuyama','N','0.277','9','0.722','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('40','Kosuke Tanaka','N','0.251','8','0.721','8');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('41','Taishi Ohta','N','0.275','14','0.721','3');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('42','Toshiki Abe','N','0.257','13','0.715','2');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('43','Akira Nakamura','N','0.271','6','0.709','0');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('44','Shogo Nakamura','N','0.249','8','0.706','6');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('45','Tomoya Mori','N','0.251','9','0.705','4');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('46','Tomotaka Sakaguchi','N','0.246','9','0.688','4');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('47','Shuta Tonosaki','N','0.247','8','0.688','21');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('48','Nobuhiro Matsuda','N','0.228','13','0.668','1');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('59','Sosuke Genda','N','0.27','1','0.656','18');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('50','Hisanori Yasuda','N','0.221','6','0.647','2');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('51','Yota Kyoda','N','0.244','5','0.642','8');
INSERT INTO players(player_id,player_name,player_rank,average,HR,OPS,SB) VALUES('52','Alcides Escobar','N','0.273','1','0.641','6');

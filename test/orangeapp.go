package test

import (
	"fmt"
	"strings"

	"durian/comic"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	err := ComicParse()
	if err != nil {
		return
	}
}

func ComicParse() error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(testcomic))
	if err != nil {
		fmt.Println(err)
		return err
	}

	comic.OrangeAppCrach(doc)
	return err
}

const testcomic = `<!doctype html>
<html lang="en">
<head>
<meta name="google" content="notranslate">
<meta property="og:site_name" content="禁漫天堂">
<meta property="og:title" content="我會對著妳汪汪的叫著 [Type-G (イシガキタカシ)] 僕はあなたにワンと鳴く Comics - 禁漫天堂">
<meta property="og:url" content="https://18comic.vip/album/419411">
<meta property="og:type" content="album">
<meta property="og:image" content="https://cdn-msp.jmcomic.me/media/logo/new_logo.png">
<meta property="og:description" content="https://www.pixiv.net/users/254423
https://www.dlsite.com/books/work/=/product_id/BJ056134.html免費成人H漫線上看">
<title>我會對著妳汪汪的叫著 [Type-G (イシガキタカシ)] 僕はあなたにワンと鳴く Comics - 禁漫天堂</title>
<meta name="theme-color" content="#FFFFFF">
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="robots" content="index, follow" />
<meta name="revisit-after" content="1 days" />
<meta name="keywords" content="個人整合,劇情向,純愛,扶他,百合,黑肉,校服,馬尾,浴衣,巨乳,泳裝,吊帶襪,女性支配,強暴,野砲,調教,群交,口交,乳交,透視,中出,中文-成人,免費A漫,H漫,同人誌,成人H漫,成人漫畫,生肉漫畫,漫畫線上看,中文A漫,H圖,H本,本本,禁漫漢化組,漢化" />
<meta name="description" content="https://www.pixiv.net/users/254423
https://www.dlsite.com/books/work/=/product_id/BJ056134.html免費成人H漫線上看" />
<meta name="juicyads-site-verification" content="9dc45d0c989efd2d562bc0bfa0ae960b">
<meta name="exoclick-site-verification" content="02a42f90b0f54accd3b5a8b44e3c2c54">
<meta name="trafficjunky-site-verification" content="vtt10qy00" />
<meta name="ero_verify" content="c55bde9f8536c9020bbbdee53a9ceee9" />
<link rel="Shortcut Icon" type="image/ico" href="/favicon.ico" />
<link rel="apple-touch-icon" href="/favicon.ico" /> <link rel="canonical" href="https://18comic.vip/album/419411" />
<script type="text/javascript">
        var base_url = "https://orangeapp.cc";
        var max_thumb_folders = "32000";
        var tpl_url = "/templates/frontend/airav";
                var lang_deleting = "删除...";
        var lang_flaging = "标记中...";
        var lang_loading = "加载中...";
        var lang_sending = "发送...";
        var lang_share_name_empty = "请输入您的名字!";
        var lang_share_rec_empty = "请至少输入一个收件人的电子邮件!";
        var fb_signin = "0";
        var fb_appid = "";
        var g_signin = "0";
        var g_cid = "";
        var signup_section = false;
        var relative = "";
        
        if (base_url === "" && atob) {
            location.href = atob("aHR0cHM6Ly8xOGNvbWljLm9yZw==") + "/album/419411";
        }
        
    </script>
<script src="/templates/frontend/airav/js/jquery.min.js"></script>
<script src="/templates/frontend/airav/js/jquery.cookie.min.js"></script>
<script src="/templates/frontend/airav/js/td_js/td_ga_tracker.js"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.lazyload-2.0.js?v=2023081601"></script>
<script src="/templates/frontend/airav/js/header.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.4_46b.js"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.avs-0.2.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/owlcarousel/owl.carousel.js"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/toastr.min.js"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/search-img.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.uploadfile.min.js"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.voting-album-0.1.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/style_phone.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.forum.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/emoji/emojionearea.js?v=1"></script>
<script async src="https://www.googletagmanager.com/gtag/js?id=G-VW05C6PGN3"></script>
<script>
            window.dataLayer = window.dataLayer || [];

            function gtag() {
                dataLayer.push(arguments);
            }

            gtag('js', new Date());
            gtag('config', 'G-VW05C6PGN3');
        </script>
<script type="text/javascript">
        (function ($) {
                        var ga_res_nums = ["UA-144553158-17", "UA-144553158-18"];
            var dimension1 = "Guest";
            var dimension2 = "簡體";
            td_ga_tracker.init(ga_res_nums );
            
            var ga_track_nums = td_ga_tracker.count();
            var ga_args = {
                'hitCallback': function () {
                    if (--ga_track_nums === 0) {

                    }
                }
            };
            td_ga_tracker.doga('require', 'linkid', 'linkid.js');
            td_ga_tracker.doga('require', 'displayfeatures');
            td_ga_tracker.doga('require', 'ec');
            td_ga_tracker.doga('set', 'dimension1', dimension1);
            td_ga_tracker.doga('set', 'dimension2', dimension2);
            td_ga_tracker.doga('send', 'pageview', ga_args);

            if ($.cookie('cover') != '1') {//聯播網錯誤
                window.addEventListener('error', function (e) {
                    if (e.target.nodeName === 'SCRIPT') {
                        td_ga_tracker.doga(
                            'send',
                            'exception',
                            {
                                'exDescription': e.target.src,
                                'exFatal': true
                            }
                        );
                    }
                }, true);
            }
            
            })(jQuery);
    </script>
<link href="/templates/frontend/airav/css/bootstrap.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/owlcarousel/assets/owl.carousel.min.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/owlcarousel/assets/owl.theme.default.min.css" rel="stylesheet">
<link href="/templates/frontend/airav/css/style.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/responsive.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/all.min.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/fontawesome.min.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/colors.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/style_phone.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/2021header.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/toastr.min.css" rel="stylesheet">
<link href="/templates/frontend/airav/css/search-img.css" rel="stylesheet">
<link href="/templates/frontend/airav/css/uploadfile.css" rel="stylesheet">
<style>
        .e8c78e-4_b {
            text-align: center;
            vertical-align: middle;
            padding: 9px !important;
            overflow: hidden;
            display: block !important;
        }

        .e8c78e-4_b img {
            height: auto;
        }

        @media screen and (max-width: 767px) {
            .e8c78e-4_b img {
                margin: 0px auto !important;
            }

            .col-lg-3.col-md-3.col-sm-3.col-xs-6:nth-child(2n+1) {
                clear: left;
            }
        }

        @media screen and (min-width: 768px) {
            .e8c78e-4_b {
                max-height: 250px !important;
            }
        }
    </style>
</head>
<body>
<input type="hidden" id="nb_path" value="e8c78e-4_b" />
<input type="hidden" id="ipcountry" value="US" />
<div id="search-board" class="modal fade in">
<div class="modal-dialog search-board">
<div class="modal-content text-center">
<div class="modal-header">
<div class="modal-title ">
搜寻的最佳姿势？
</div>
</div>
<div class="modal-body">
<div>
<img style="max-width: 265px;" class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/templates/frontend/airav/img/phicon/search-board.png">
</div>
<div class="m-t-5">【包含搜寻】<br>搜寻[+]全彩[空格][+]人妻,仅显示全彩且是人妻的本本<br>范例:+全彩 +人妻<br><br>
【排除搜寻】<br>搜寻全彩[空格][-]人妻,显示全彩并排除人妻的本本<br>範例:全彩 -人妻<br><br>
【我都要搜寻】<br>搜寻全彩[空格]人妻,会显示所有包含全彩及人妻的本本<br>范例:全彩 人妻<br></div>
</div>
<div class="modal-footer">
<button type="button" class="btn btn-primary" data-dismiss="modal" style="display: block;margin: auto;">我都学会了！
</button>
</div>
</div>
</div>
</div>
<div id="billboard-modal" class="modal fade in" data-backdrop="static">
<div class="modal-dialog billboard-modal">
<div class="modal-content text-center">
<div class="modal-header">
<div class="modal-title ">
<a href="/blog/1153">
<img style="max-width: 100%;height: auto;" src="/static/resources/files/20230705_jm_1600800_cn.jpg" alt="禁漫天堂">
</a>
</div>
</div>
<div class="modal-body">
<ul class="pop-list">
<li class="pop-list-item_img">
<a href="https://jm365.work/xPD8Un" target="_blank" rel="nofollow noopener">
<img src="/static/resources/images/games/350x70-V2.png">
</a>
</li>
<li class="pop-list-item_img">
<a href="https://m.cowrp.xyz/?attributionId=126" target="_blank" rel="nofollow noopener">
<img src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/RP_15_700_140_v6.jpg">
</a>
</li>
<li class="pop-list-item_img">
<a href="https://hhyae7.cc/8850.html" target="_blank" rel="nofollow noopener">
<img src="/static/resources/files/700x140.gif">
</a>
</li>
<li class="pop-list-item_img">
<a href="https://zno290.com/nydkg/" target="_blank" rel="nofollow noopener">
<img src="/static/resources/images/AD/700x140-20230608.gif">
</a>
</li>
<li class="pop-list-item_img">
<a href="https://98pro.cc/sbctPk" target="_blank" rel="nofollow noopener">
<img src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/20230721_QC_700140_TW.gif">
</a>
</li>
<li class="pop-list-item_img">
<a href="https://r.morexpbb.net/game/xr/?attributionId=126" target="_blank" rel="nofollow noopener">
<img src="/static/resources/images/%E5%BB%A3%E5%91%8A202308/XR_random_700_140_v43.jpg">
</a>
</li>
</ul>
</div>
<div class="modal-footer">
<div class="m-b-10 center" style="color: red;font-weight:bold;">
使用ADBLOCK将会使网站部份功能(下载、换页等)可能无法正常使用请自行注意!!
</div>
<div class="btn-group">
<button id="chk_cover" type="button" class="btn btn-primary col-xs-7" data-dismiss="modal">我保证我已18岁!
</button>
<a href="https://www.google.com.tw/search?q=禁漫天堂" class="btn col-xs-5" style="background-color: #ababab;color: #323232;margin: 0;flex:1;">未滿18歲，離開</a>
</div>
</div>
</div>
<div class="center m-t-10">
<a href="/static/_2257" style="color: #fff;">18 USC 合規免責聲明</a>
</div>
</div>
</div>
<div id="guide-modal" class="modal fade in" data-backdrop="static">
<div class="modal-dialog ">
<div class="modal-content text-center">
<div class="modal-head">
</div>
<div class="modal-body">
<div class="owl-carousel owl-pop2_img" style="background:initial;">
<div class="partial-item">
<a href="https://youtu.be/OUMeIfNq_2M" target="_self">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="/static/resources/files/20230717_%E5%90%8C%E4%BA%BA%E9%9B%80%E5%A3%AB_500250.jpg" class="owl-lazy" itemprop="image" />
</a>
</div>
<div class="partial-item">
<a href="https://98pro.cc/5JYpMv" target="_self">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/20230711_QC_500250_CN.gif" class="owl-lazy" itemprop="image" />
</a>
</div>
</div>
<ul style="margin-top:1rem">
<li>
<div class="news_card ">
<h4>新遊戲上架~ 星神少女</h4>
<a href="https://m.cowrp.xyz/?attributionId=126" rel="nofollow noopener" target="_blank">
<img alt="新遊戲上架~ 星神少女" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/RP_15_500_250_v2(1).gif" />
</a>
<a href="https://m.cowrp.xyz/?attributionId=126" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;">眾神之父對獵物伸出淫爪，擒服少女強制交合，巨根猛捅百合花唇，一晚七次灌注濃稠精液，堂堂阿開地亞公主淪為神的泄慾容器♡！！！</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong>兌換碼：<span style="color:#e74c3c;">COMICXOP0815</span></strong></p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong><span style="color:#e74c3c;">(日冕石&times;18. 星之鑰&times;2. 隨機禮物盒&times;2)</span></strong></p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>放置傳說</h4>
<a href="https://r.morexpbb.net/game/xr/?attributionId=126" rel="nofollow noopener" target="_blank">
<img alt="放置傳說" style="width: 100%;" src="/static/resources/images/%E5%BB%A3%E5%91%8A202308/XR_random_600_200_v42.gif" />
</a>
<a href="https://r.morexpbb.net/game/xr/?attributionId=126" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<div style="text-align: center;">「報告老大，成功生擒三上悠亞，等候你的發落！」<br/>
悠亞眼神迷糊，白滑美腿主動撩動褲檔，遵從體內的獸性，拋開理智，以粗壯的手臂撲上去將她反噬！</div>
<div style="text-align: center;">&nbsp;</div>
<div style="text-align: center;">兌換碼 ：<span style="color:#e74c3c;"><strong><span style="background-color:#dddddd;">HFU8R14JHR</span></strong></span><br/>
<strong><span style="color:#e74c3c;">&nbsp;-－Gem x100－召喚券 x3－一袋蒼玉 (2小時) x3</span></strong></div>
<div style="text-align: center;"><strong>有效至：8月31日23:59</strong></div>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>女皇調教手冊</h4>
<a href="https://s.aplagoon.xyz/?attributionId=126" rel="nofollow noopener" target="_blank">
<img alt="女皇調教手冊" style="width: 100%;" src="/static/resources/images/%E5%BB%A3%E5%91%8A202308/queen_random_600_200_v44.jpg" />
</a>
<a href="https://s.aplagoon.xyz/?attributionId=126" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<div style="text-align: center;">
<style type="text/css"><!--td {border: 1px solid #ccc;}br {mso-data-placement:same-cell;}-->
</style>
植物園裡迷路，催情香氣飄來，只見妖精迷糊張腿對著花蕊上下跩動！？</div>
<div style="text-align: center;">&nbsp;</div>
<div style="text-align: center;">&nbsp;</div>
<p style="text-align: center;"><strong>兌換碼<span style="background-color:#ffffff;">:&nbsp;&nbsp;</span><span style="color:#e74c3c;"><span style="background-color:#dddddd;">AUGIXDPICA</span></span><span style="background-color:#dddddd;">&nbsp;</span><span style="background-color:#ffffff;"> </span></strong></p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong><span style="background-color:#ffffff;"><span style="color:#e74c3c;">(</span></span><span style="color:#e74c3c;">鑽石*200 .五星戰姬石*30 .招募卷*5</span><span style="background-color:#ffffff;"><span style="color:#e74c3c;">)</span></span></strong><br/>
<span style="color:#95a5a6;"><strong><span style="background-color:#ffffff;">有效至：8月31日23:59</span></strong></span></p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>新遊戲上架~三國志侵略版</h4>
<a href="https://sgzqlb.com/fk/index.html?ag=18comic" rel="nofollow noopener" target="_blank">
<img alt="新遊戲上架~三國志侵略版" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/FK_2_500X250.gif" />
</a>
<a href="https://sgzqlb.com/fk/index.html?ag=18comic" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;">高颜值国创卡牌王者手游,再度开启魏蜀吴群英荟萃的将星时代!</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">全新玩法,百变搭配,零损换将,出奇制胜争夺无上荣耀。</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">佛系放置,不肝不氪,轻松扫荡,护肝设计打造超强躺赢阵容。</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">天梯争霸,联盟对抗,燃情竞技,一策奇谋逆风翻盘主宰战场。</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">铜雀春深,浮生一梦,万种风流,美人侍寝助力主公逐鹿九州。</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">这就是三国,也许你就是下一位天下之主!</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong>兌換碼:&nbsp;<span style="color:#e74c3c;">JSFK666&nbsp; (元寶*350、5星隨機碎片*50)</span></strong></p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>新遊戲~~三國艷遊錄</h4>
<a href="https://to.js-hban.site/tjtz/jump/15/189/1/jump.html" rel="nofollow noopener" target="_blank">
<img alt="新遊戲~~三國艷遊錄" style="width: 100%;" src="/static/resources/images/%E5%96%AE%E6%A9%9F%E9%81%8A%E6%88%B2/4-500x250(1).gif" />
</a>
<a href="https://to.js-hban.site/tjtz/jump/15/189/1/jump.html" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;">《三国艷遊錄》是一款以三國戰爭史為題材的回合制養成遊戲，劇情新穎，畫面精致，玩法眾多，超多著名的戰將等你來收集，炫目特技、破陣殺敵！調戲女神，招募戰將、培養舉世無雙的絕世神將，與怪物大軍、副本BOSS、甚至與玩家展開對決。 遊戲以主線劇情玩法為核心，結合主角征伐路上的激情艷遇，人妻！妖精！劍魂！貂蟬&hellip;統統一槍收服，金槍不倒，橫掃天下，女神服侍，稱霸三國，建立後宮！</p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>【新遊戲上架】慾神幻想</h4>
<a href="https://l.hyenadata.com/s/0Ac1Kc" rel="nofollow noopener" target="_blank">
<img alt="【新遊戲上架】慾神幻想" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/1779_25__500X250_CN.gif" />
</a>
<a href="https://l.hyenadata.com/s/0Ac1Kc" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">許久以前，人類最原始的慾望造就了兩名原初之神：女神「愛芙黛蒂」與男神「艾洛斯」。而後，隨著文明發展，增生的慾望導致越來越多的神祇現身，</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">並演變成神與神之間的大戰。儘管艾洛斯</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">獲得勝利，卻因身邊的愛芙黛蒂背叛而遭到封印，世界也隨之進入了無神時代。</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong>兌換碼 ：<span style="color:#e74c3c;">75qR4PafSMe</span></strong><br/>
<strong>&nbsp;-</strong>夢之晶x100+ 精力爆炸藥水 x1+ 3萬金幣</p>
<p style="text-align: center;">&nbsp;</p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>新遊戲上架~櫻境物語</h4>
<a href="https://l.hyenadata.com/s/17hZk2" rel="nofollow noopener" target="_blank">
<img alt="新遊戲上架~櫻境物語" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/1893_25_CherryTale_500x250.gif" />
</a>
<a href="https://l.hyenadata.com/s/17hZk2" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;">擁有從神而來的智慧與魔力的人類,與障天使們的戰爭</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">持續了上百年之在漫長的戰爭中,天使的力量日益苗</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">壯,人類逐漸無法與之抗衡,他們迫不得已,以自身的魔</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">力作為籌碼與天神交易。在天神的幫助下,人類戰勝了墮</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">天使,這場數百年的戰爭才終於畫下句號。</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong>兌換碼 ：</strong><strong><span style="color:#e74c3c;">VIPYX290</span>&nbsp;(<span style="color:#e74c3c;">裝備召喚券x3</span>)</strong></p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>新遊戲上架~禁漫斗羅</h4>
<a href="https://l.tyrantdb.com/QWtYYfRs?subc1=18" rel="nofollow noopener" target="_blank">
<img alt="新遊戲上架~禁漫斗羅" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/20230712500X250.gif" />
</a>
<a href="https://l.tyrantdb.com/QWtYYfRs?subc1=18" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;">華人銷量TOP1實體系列小說改編的H版同人遊戲</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">22年度國民級P戰鬥手游鉅作,《禁漫斗羅》,</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">邀您一同進入和動畫一模一樣的真實H版斗羅世界!</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong>兌換碼 ：<span style="color:#e74c3c;">GDFA6AG5</span></strong><br/>
<strong><span style="color:#e74c3c;"><span style="background-color:#ffffff;">&nbsp;(魂殿教皇令*10、金磚*20、奧斯卡特級香腸*1、奧斯卡香腸*1)</span></span></strong></p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong>有效至：12月31日23:59</strong></p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>翻牆神器 ~ Shark Cloud</h4>
<a href="https://1.moegura.com/#/register?code=wToZwsWa" rel="nofollow noopener" target="_blank">
<img alt="翻牆神器 ~ Shark Cloud" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/20230517%20600x200.jpg" />
</a>
<a href="https://1.moegura.com/#/register?code=wToZwsWa" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;"><strong>纯隧道SS机场 价格低廉 <span style="color:#e74c3c;">最低套餐低至5.99&nbsp;</span><span style="color:#e67e22;"><span style="background-color:#ffffff;">&nbsp;</span>&nbsp;</span><span style="color:#000000;">热门流媒体</span><span style="color:#e74c3c;">解锁</span></strong></p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong>客服工单<span style="color:#e74c3c;">1小时</span>内回应 节点故障维护迅速</strong></p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><br/>
<strong>网络互联网冲浪尽在&nbsp;<span style="color:#999999;"><span style="background-color:#ffffff;">Shark Cloud</span></span></strong></p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>戰國天堂~~七夕活動</h4>
<a href="https://url365.club/pzx2Yv" rel="nofollow noopener" target="_blank">
<img alt="戰國天堂~~七夕活動" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/20221102_war_500250.gif" />
</a>
<a href="https://url365.club/pzx2Yv" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;"><span style="color:#000000;"><font face="微軟雅黑"><font style="overflow-wrap:break-word; padding:0px; margin:0px; color:#596090; font-size:14px; font-style:normal; font-variant-ligatures:normal; font-weight:400; text-align:start; white-space:normal; background-color:#000000; text-decoration-thickness:initial; text-decoration-style:initial; text-decoration-color:initial"><font size="3"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff;">活動時間內於寶區、練區、戰場、江戶、肯特，怪物隨機掉落【紅粉知心】</span></strong></font></font></font></font></font></span></p>
<p style="text-align: center;"><br microsoft style="overflow-wrap:break-word; padding:0px; color:#596090; font-family:" yahei />
<span style="color:#000000;"><font face="微軟雅黑"><font style="overflow-wrap:break-word; padding:0px; margin:0px; color:#596090; font-size:14px; font-style:normal; font-variant-ligatures:normal; font-weight:400; text-align:start; white-space:normal; background-color:#000000; text-decoration-thickness:initial; text-decoration-style:initial; text-decoration-color:initial"><font size="3"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff;">【紅粉知心】</span></strong><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff;">&nbsp;- 使用後可回復HP800(無冷卻時間)</span></strong></font></font></font></font></font></span></p>
<p style="text-align: center;"><br microsoft style="overflow-wrap:break-word; padding:0px; color:#596090; font-family:" yahei />
<span style="color:#000000;"><font face="微軟雅黑"><font style="overflow-wrap:break-word; padding:0px; margin:0px; color:#596090; font-size:14px; font-style:normal; font-variant-ligatures:normal; font-weight:400; text-align:start; white-space:normal; background-color:#000000; text-decoration-thickness:initial; text-decoration-style:initial; text-decoration-color:initial"><font size="3"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff;">收集99個合成【紅粉禮包】開啟隨機獲得商城道具</span></strong></font></font></font></font></font></span></p>
<p style="text-align: center;"><br microsoft style="overflow-wrap:break-word; padding:0px; color:#596090; font-family:" yahei />
<span style="color:#000000;"><font face="微軟雅黑"><font style="overflow-wrap:break-word; padding:0px; margin:0px; color:#596090; font-size:14px; font-style:normal; font-variant-ligatures:normal; font-weight:400; text-align:start; white-space:normal; background-color:#000000; text-decoration-thickness:initial; text-decoration-style:initial; text-decoration-color:initial"><font size="3"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff;">收集999個＋贊助五萬就送-紅粉西施(永久</span></strong></font></font></font></font><font style="overflow-wrap:break-word; padding:0px; color:#596090; font-size:14px; font-style:normal; font-variant-ligatures:normal; font-weight:400; text-align:start; white-space:normal; background-color:#000000; text-decoration-thickness:initial; text-decoration-style:initial; text-decoration-color:initial"><font size="3"><font style="overflow-wrap:break-word; padding:0px"><font style="overflow-wrap:break-word; padding:0px"><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff">稀有金變</span></strong></font></font></font></font><font style="overflow-wrap:break-word; padding:0px; margin:0px; color:#596090; font-size:14px; font-style:normal; font-variant-ligatures:normal; font-weight:400; text-align:start; white-space:normal; background-color:#000000; text-decoration-thickness:initial; text-decoration-style:initial; text-decoration-color:initial"><font size="3"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><font style="overflow-wrap:break-word; padding:0px; margin:0px"><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff;">)</span></strong></font></font></font></font></font></span></p>
<p style="text-align: center;"><br/>
<span style="color:#000000;"><font face="微軟雅黑"><font style="overflow-wrap:break-word; padding:0px; color:#596090; font-size:14px; font-style:normal; font-variant-ligatures:normal; font-weight:400; text-align:start; white-space:normal; background-color:#000000; text-decoration-thickness:initial; text-decoration-style:initial; text-decoration-color:initial"><font size="3"><font style="overflow-wrap:break-word; padding:0px"><font style="overflow-wrap:break-word; padding:0px"><strong style="overflow-wrap:break-word; padding:0px; font-weight:700"><span style="background-color:#ffffff;">活動時間：7/7 ~ 8/31</span></strong></font></font></font></font></font></span></p>
</div>
</div>
</li>
<li>
<div class="news_card ">
<h4>[Boki Boki] 淫習的幽世村</h4>
<a href="https://www.boki2.fun/r18/proitem.php?obj_id=G0000221&18comic=1" rel="nofollow noopener" target="_blank">
<img alt="[Boki Boki] 淫習的幽世村" style="width: 100%;" src="/static/resources/images/%E4%BA%8C%E6%AC%A1%E8%93%8B%E6%9D%BF/500x250%20(1).jpg" />
</a>
<a href="https://www.boki2.fun/r18/proitem.php?obj_id=G0000221&18comic=1" rel="nofollow noopener" target="_blank">
<div>立刻<br/>遊玩</div>
</a>
<div>
<p style="text-align: center;"><strong><span style="font-size:14px;">那是個淫亂猖獗，女人對男人絕對服從的村子。奇諾理因為孩子氣的外表而被村民誤認為是男孩。<br/>
兩人試圖隱瞞著奇諾理是女孩的事實，直到他們找到逃離村莊的辦法......</span></strong></p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong><span style="font-size:18px;"><span style="color:#e74c3c;">禁漫天堂專屬邀請碼: R18COMIC </span></span></strong></p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;">&nbsp;</p>
<p style="text-align: center;"><strong><span style="font-size:18px;"><span style="color:#e74c3c;">(註冊 獲得25點，30天內消費滿200元 再獲得30點HPoint)</span></span></strong></p>
</div>
</div>
</li>
</ul>
<div><strong>禁漫官方社群</strong>
<ul>
<li><a href="https://www.instagram.com/18comicorg/" rel="nofollow noopener" target="_blank"><img alt="instagram" src="/templates/frontend/airav/img/social_ig.svg" /> </a></li>
<li><a href="https://twitter.com/18comic2" rel="nofollow noopener" target="_blank"><img alt="twitter" src="/templates/frontend/airav/img/social_twitter.svg" /> </a></li>
<li><a href="https://discord.gg/V74p7HM" rel="nofollow noopener" target="_blank"><img alt="發布頁" src="/templates/frontend/airav/img/social_ds.svg" />
</a></li>
<li><a href="https://jmcomic.bet/" rel="nofollow noopener" target="_blank"><img alt src="/templates/frontend/airav/img/release.svg" />
</a></li>
<li><a href="https://www.facebook.com/senshimanga/" rel="nofollow noopener" target="_blank"><img alt src="/templates/frontend/airav/img/social_fb.svg" /> </a></li>
</ul>
</div>
</div>
<div class="modal-footer">
<button id="chk_guide" type="button" class="btn btn-primary" data-dismiss="modal" style="display: block;margin: auto;">确定进入!
</button>
</div>
</div>
</div>
</div>
<div class="modal fade in" id="login-modal">
<div class="modal-dialog login-modal">
<div class="modal-content">
<form name="login_form" method="post" action="/login">
<div class="modal-header">
<button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
<h4 class="modal-title">会员登录</h4>
</div>
<div class="modal-body">
<center>
<div class="m-b-5"></div>
</center>
<div class="form-group">
<label for="login_username" class="control-label">用户名:</label>
<input name="username" type="text" value id="login_username" class="form-control" />
</div>
<div class="form-group">
<label for="login_password" class="control-label">密码:</label>
<input name="password" type="password" value id="login_password" autocomplete="on" class="form-control" />
</div>
<div class="form-group">
<div class="checkbox">
<label>
<input name="id_remember" type="checkbox" /> 在这台装置记住我的帐号密码
</label>
</div>
<div class="checkbox">
<label>
<input name="login_remember" type="checkbox" id="login_remember" /> 登陸180天，期限內將自動登入會員
</label>
</div>
</div>
<a href="/lost" id="lost_password">忘记用户名或密码？</a><br/>
<a href="/confirm" id="confirmation_email">没有收到确认邮件？</a>
</div>
<div class="modal-footer">
<button name="submit_login" id="login_submit" type="submit" class="btn btn-primary">登录</button>
<a href="/signup" class="btn btn-secondary">注册</a>
</div>
</form>
</div>
</div>
</div>
<div class="modal fade in" id="language-modal">
<div class="modal-dialog language-modal">
<div class="modal-content">
<div class="modal-header">
<button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
<h4 class="modal-title">Select Language</h4>
</div>
<div class="modal-body">
<div class="row">
<div class="col-xs-6 col-sm-4">
<a href="#" id="cn_CT" class="change-language">中文繁體</a>
</div>
<div class="col-xs-6 col-sm-4">
<span class="change-language language-active">中文简体</span>
</div>
</div>
<form name="languageSelect" id="languageSelect" method="post" action>
<input name="language" id="language" type="hidden" value />
</form>
</div>
</div>
</div>
</div>
<div class="modal fade in" id="shunt-modal">
<div class="modal-dialog shunt-modal">
<div class="modal-content">
<div class="modal-header">
<button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
<h4 class="modal-title">切換分流，加速你的網站體驗</h4>
</div>
<div class="modal-body">
<div class="row">
<div class="col-xs-6 col-sm-4">
<a href="?shunt=1" class="change-shunt language-active">分流一</a>
</div>
<div class="col-xs-6 col-sm-4">
<a href="?shunt=2" class="change-shunt ">分流二</a>
</div>
<div class="col-xs-6 col-sm-4">
<a href="?shunt=3" class="change-shunt ">分流三</a>
</div>
</div>
</div>
<div class="modal-footer">
<button type="button" class="btn btn-primary" data-dismiss="modal" style=" display: block;margin: auto;">我覺得現在很快了
</button>
</div>
</div>
</div>
</div>
<div class="modal fade in" id="headelikes_modal">
<div class="modal-dialog headelikes_modal">
<div class="modal-content" style="background-color: #202020">
<div class="modal-header" style="border-bottom: none;">
<button type="button" class="close" data-dismiss="modal" aria-hidden="true" style="color: #ffffff;">
&times;
</button>
<h4 class="modal-title" style="color: #ffffff">友站连结</h4>
<div class="modal-body" style="padding: 10px;">
<a data-label="menu-top-link" href="https://url365.club/xZwttR" target="_blank" rel="nofollow noopener" style="color: #ec79ac">戰國天堂</a>
</div>
<div class="modal-body" style="padding: 10px;">
<a data-label="menu-top-link" href="https://moefuns.info/?mfa=18comic" target="_blank" rel="nofollow noopener" style="color: #ec79ac">萌翻导航</a>
</div>
<div class="modal-body" style="padding: 10px;">
<a data-label="menu-top-link" href="https://theporndude.com/zh" target="_blank" rel="nofollow noopener" style="color: #ec79ac">成人Dude</a>
</div>
<div class="modal-body" style="padding: 10px;">
<a data-label="menu-top-link" href="https://69run.work/5mHfCW" target="_blank" rel="nofollow noopener" style="color: #ec79ac">束縛美學</a>
</div>
</div>
</div>
</div>
</div>
<div class="modal fade" id="search-Modalimg" tabindex="-1" role="dialog" aria-hidden="true" style="display: none;">
<div class="modal-dialog modal-search-img" role="document">
<div class="modal-content search-img-block">
<div class="modal-header border-0">
<button type="button" class="close" data-dismiss="modal" aria-label="Close">
<span aria-hidden="true">×</span>
</button>
<h4 class="modal-title user-active m-b-5">以圖搜圖</h4>
<p>不需輸入關鍵字，直接使用圖片搜尋。嘗試將圖片拖曳到這裡。
<br>※請上傳整張完整圖片，剪裁過的容易照成搜尋失敗
</p>
</div>
<form action="/search/image" method="post" enctype="multipart/form-data">
<div class="middle d-flex">
<div class="middle_bt">
<a href="#web">貼上圖片網址</a>
</div>
<div class="middle_bt">
<a href="#wab">上傳圖片</a>
</div>
</div>
<div id="web" class="modal-body scr">
<input type="hidden" name="imgBase64" id="imgBase64">
<input class="scr_a" type="url" name="urlImage" placeholder="請在這貼上圖片網址">
<input class="scr_b border-0 m-l-10" type="submit" value="以圖搜尋">
</div>
<div id="wab" class="modal-body scr">
<div id="fileuploader" name="fileuploader"></div>
</div>
</form>
</div>
</div>
</div>
<div class="black-back"></div>
<div class="top-nav">
<div class="container">
<ul class="top-menu">
<div class="center visible-xs visible-sm" style="white-space: nowrap;">
<li class="top-menu-m">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" href="https://9527.rocks/EwtPpR">
焕儿全集
</a>
</li>
<li class="top-menu-m">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" href="https://69run.work/8gbbn4">
免費片
</a>
</li>
<li class="top-menu-m">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" href="https://69run.work/qEPQSt">
好友裸拍视频
</a>
</li>
<li class="top-menu-m">
<a data-toggle="modal" href="#headelikes_modal" style="color: #fff;">友站</a></li>
</div>
<div class="pull-left">
<li class="dropdown hidden-xs hidden-sm">
<a href="#" class="dropdown-toggle" data-toggle="dropdown" style="color: #fff;">
热门排行
<b class="caret"></b>
</a>
<ul class="dropdown-menu">
<li class="top-menu-link "><a href="https://orangeapp.cc/albums?o=mv">热门</a>
</li>
<li class="top-menu-link "><a href="https://orangeapp.cc/albums?t=m&o=mv">本月热门</a></li>
<li class="top-menu-link "><a href="https://orangeapp.cc/albums?o=mv&t=w">本周热门</a></li>
</ul>
</li>
<li class="top-menu-link hidden-xs hidden-sm">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" style="color: #ec79ac" href="https://9527.rocks/EwtPpR">焕儿</a>
</li>
<li class="top-menu-link hidden-xs hidden-sm">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" style="color: #ec79ac" href="https://69run.work/8gbbn4">AIRAV免費A片</a>
</li>
<li class="top-menu-link hidden-xs hidden-sm">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" style="color: #ec79ac" href="https://69run.work/qEPQSt">好友裸拍</a>
</li>
<li class="top-menu-link hidden-xs hidden-sm">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" style="color: #ec79ac" href="https://url365.club/xZwttR">戰國天堂</a>
</li>
<li class="top-menu-link hidden-xs hidden-sm">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" style="color: #ec79ac" href="https://moefuns.info/?mfa=18comic">萌翻导航</a>
</li>
<li class="top-menu-link hidden-xs hidden-sm">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" style="color: #ec79ac" href="https://theporndude.com/zh">成人Dude</a>
</li>
<li class="top-menu-link hidden-xs hidden-sm">
<a data-label="menu-top-link" target="_blank" rel="nofollow noopener" style="color: #ec79ac" href="https://69run.work/5mHfCW">束縛美學</a>
</li>
</div>
<div class="pull-right hidden-xs hidden-sm">
<li><a data-toggle="modal" href="#language-modal">簡體 <span class="caret"></span></a></li>
</div>
<div class="pull-right hidden-xs hidden-sm">
<li><a data-toggle="modal" href="#shunt-modal">分流<span class="caret"></span></a></li>
</div>
<div class="pull-right hidden-xs hidden-sm" style="margin-left: 10px">
<a href="https://jm365.work/dFTN82" target="_blank" rel="nofollow noopener">打赏JM</a>
</div>
<div class="pull-right hidden-xs hidden-sm" style="margin-left: 10px">
<a href="/cdn-cgi/l/email-protection#91e6e6e6a0a9f2fefcf8f2d1f6fcf0f8fdbff2fefc">广告洽询</a>
</div>
<div class="pull-right hidden-xs hidden-sm">
<a href="https://jm365.work/mpSWW7" target="_blank" rel="nofollow noopener">JM公告</a>
</div>
</ul>
</div>
</div>
<div id="Comic_Top_Nav" class="navbar navbar-inverse navbar-fixed-top" role="navigation">
<div class="container">
<div class="navbar-header">
<button id="mob-header-menu" type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-inverse-collapse">
<div class="hamburger" id="hamburger-11">
<span class="line"></span>
<span class="line"></span>
<span class="line"></span>
</div>
</button>
<a class="navbar-brand" href="/"><img src="https://cdn-msp.jmcomic.me/media/logo/new_logo.png?v=2023081601" alt="logo"></a>
<div class="head-right">
<a class="visible-xs visible-sm" data-toggle="modal" href="#shunt-modal" style="position: unset;margin: 10px 10px 0 0px;float: right; color: #ff7a01;border:1px solid #ff7a01;padding: 4px;border-radius: 5px;
                       ">
<b>分流1▾<b></a>
<div class="dropdown visible-xs visible-sm" style="position: unset;margin: 15px 10px 0 0px;float: right; ">
<a style="color: #888;" href="javascript:void(0)" class="dropdown-toggle" data-toggle="dropdown"><i class="fa fa-search fa-lg"></i></a>
<ul class="dropdown-menu search-dropdown-menu" style="margin: 0; background-color: #323232; overflow-y: scroll; border-radius: 0;">
<form class="form-inline" name="search" id="search_form_m" method="get" action="/search/photos">
<div class="search-dropdown-menu-top d-flex justify-content-around center">
<div class="top-btn">
<a href="#select-type">類型搜尋</a>
</div>
<div class="top-btn">
<a href="#img-type" id="img-search-option">圖片搜尋</a>
</div>
</div>
<div class="select-type-title type" id="select-type">
<p style="float: left">搜索类型</p>
<div class="input-group" style="clear: both;">
<input type="text" class="form-control" placeholder="可以搜尋圖片囉 ! " name="search_query" autocomplete="off" value>
<span class="input-group-btn">
<button type="submit" class="btn btn-primary" style="font-size: 14px"><i class="fa fa-search" aria-hidden="true"></i>
</button>
</span>
</div>
<div class="m-t-5">
<a data-toggle="modal" data-target="#search-board" style="color: #ababab;"><i class="far fa-question-circle" aria-hidden="true"></i>搜寻的最佳姿势？</a>
</div>
<div class="m-t-20">
<span class="d-block p-b-5">搜尋類型</span>
<div class="search-type">
<input type="radio" id="adult" name="search-type" value="photos">
<label for="adult" class="m-r-10">A漫</label>
<input type="radio" id="hcomic" name="search-type" value="videos">
<label for="hcomic" class="m-r-10">H动漫</label>
<input type="radio" id="movies" name="search-type" value="movies">
<label for="movies" class="m-r-10">小电影</label>
<input type="radio" id="blogs" name="search-type" value="blogs">
<label for="blogs" class="m-r-10">紳夜食堂</label>
</div>
</div>
<div class="m-t-10 m-b-10">
<span class="d-block p-b-5">搜尋分類</span>
<div class="search-type">
<input type="radio" id="name" value="0" name="main_tag" checked>
<label for="name" class="m-r-10">站內搜索</label>
<input type="radio" id="works" value="1" name="main_tag">
<label for="works" class="m-r-10">作品</label>
<input type="radio" id="author" value="2" name="main_tag">
<label for="author" class="m-r-10">作者</label>
<input type="radio" id="tag" value="3" name="main_tag">
<label for="tag" class="m-r-10">标签</label>
<input type="radio" id="people" value="4" name="main_tag">
<label for="people" class="m-r-10">登场人物</label>
</div>
</div>
<li class="visible-xs visible-sm p-t-20">熱門關鍵字</li>
<div class="tag-hot tag-block">
<a href="/search/photos?search_query=连载中" class="btn btn-sm tag-block" title="连载中">连载中</a>
<a href="/search/photos?search_query=吸奶" class="btn btn-sm tag-block" title="吸奶">吸奶</a>
<a href="/search/photos?search_query=女性支配" class="btn btn-sm tag-block" title="女性支配">女性支配</a>
<a href="/search/photos?search_query=韩漫的豪哥都好厉害" class="btn btn-sm tag-block" title="韩漫的豪哥都好厉害">韩漫的豪哥都好厉害</a>
<a href="/search/photos?search_query=aslsdtkln書生" class="btn btn-sm tag-block" title="aslsdtkln書生">aslsdtkln書生</a>
<a href="/search/photos?search_query=NUWARU" class="btn btn-sm tag-block" title="NUWARU">NUWARU</a>
<a href="/search/photos?search_query=最强豪哥没有之一" class="btn btn-sm tag-block" title="最强豪哥没有之一">最强豪哥没有之一</a>
<a href="/search/photos?search_query=請對標籤君手下留情b" class="btn btn-sm tag-block" title="請對標籤君手下留情b">請對標籤君手下留情b</a>
<a href="/search/photos?search_query=干女儿韶多" class="btn btn-sm tag-block" title="干女儿韶多">干女儿韶多</a>
<a href="/search/photos?search_query=韶多小猫咪" class="btn btn-sm tag-block" title="韶多小猫咪">韶多小猫咪</a>
</div>
</div>
<div class="search-img-type type" id="img-type">
<p>不需輸入關鍵字，直接使用圖片搜尋。
<br>※請上傳整張完整圖片，剪裁過的容易照成搜尋失敗
</p>
<div class="input-group" style="clear:both;">
<span class="input-group-btn"><button type="submit" class="btn btn-primary" style="font-size: 14px">以圖搜尋</button></span>
</div>
<div class="m-t-30" style="clear:both;">
<p>上傳圖片</p>
<button type="submit" class="btn btn-primary"><i class="fa fa-search" aria-hidden="true"></i></button>
</div>
</div>
</form>
</ul>
</div>
</div>
</div>
<div class="navbar-collapse collapse navbar-inverse-collapse ps-re flex1" style="max-height: calc(100vh - 90px);overflow-y: scroll;">
<div class="d-lg-flex align-items-center">
<ul class="nav navbar-nav navbar-left">



<li class="visible-xs visible-sm copy-block">
<a href="https://jm365.work/3YeBdF" style="color: #ffffff" target="_blank" rel="nofollow noopener"><span style="color: #ff7a00;">收藏永久網域</span></a>
<span id="copy" class="copy-text">https://jm365.work/3YeBdF</span>
<span onclick="copyEvent('copy')" class="copy-btn m-l-10"><i class="far fa-clone" aria-hidden="true"></i></span>
</li>
<li class="visible-xs visible-sm navbar-nav-icon">
<a href="/blogs"><i class="fas fa-book-open" aria-hidden="true"></i><span>紳夜食堂</span></a>
<a href="https://jm365.work/pXYbfA" target="_blank" rel="nofollow noopener">
<i class="fas fa-mobile-alt"></i>
APP下載
</a>
<a href="/album/272689"><img src="/templates/frontend/airav/img/phicon/paper_box.svg" style="width: 25px"> <span>一抽入魂</span></a>
<a href="https://jm365.work/dFTN82" rel="nofollow noopener" target="_blank"><i class="fas fa-coins" aria-hidden="true"></i><span>打赏JM</span></a>
<a href="https://www.facebook.com/senshimanga/" target="_blank" rel="nofollow noopener"><i class="fab fa-facebook-square" aria-hidden="true"></i><span>Facebook</span></a>
<a href="https://discord.gg/V74p7HM" rel="nofollow noopener" target="_blank"><i class="fab fa-discord" aria-hidden="true"></i><span>Discord</span></a>
<a href="https://www.instagram.com/18comicorg/" rel="nofollow noopener" target="_blank"><i class="fab fa-instagram" aria-hidden="true"></i><span>IG</span></a>
<a href="https://t.me/hcomic18" rel="nofollow noopener" target="_blank"><i class="fab fa-telegram-plane" aria-hidden="true"></i><span>Telegram</span></a>
</li>
<div class="line-white visible-xs"></div>
<li class="visible-xs visible-sm" style="font-size: 14px;padding: 10px 15px;clear: both;margin-top:20px;">
A漫
</li>
<div class="adult-class-main text-none visible-xs visible-sm">
<div class="visible-xs visible-sm">
<a href="/albums?o=mr">
<li style="font-size: 14px;padding: 10px 32px;clear: both;">
<i class="far fa-star" aria-hidden="true"></i><span>最新A漫</span>
</li>
</a>
<a href="/albums?o=mv" title="总排行">
<li style="font-size: 14px;padding: 10px 32px;clear: both;">
<i class="far fa-heart" aria-hidden="true"></i><span>熱門排行</span>
</li>
</a>
<a href="/albums/hanman" title="韩漫">
<li style="font-size: 14px;padding: 10px 32px;clear: both;">
<i class="far fa-gem" aria-hidden="true"></i><span>韩漫</span>
</li>
</a>
<a href="https://6bq9.cc/dDcJP4" title="美漫">
<li style="font-size: 14px;padding: 10px 32px;clear: both;">
<i class="fas fa-flag-usa" aria-hidden="true"></i><span>美漫</span>
</li>
</a>
<a href="/week" title="每周必看">
<li style="font-size: 14px;padding: 10px 32px;clear: both;">
<i class="far fa-calendar-alt" aria-hidden="true"></i><span>每周必看</span>
</li>
</a>
</div>
</div>
<li class="visible-xs visible-sm" style="font-size: 14px;padding: 10px 15px;clear: both;">
更多內容
</li>
<div class="adult-class-main visible-xs visible-sm">
<div class="visible-xs visible-sm">
<li style="font-size: 14px;padding: 10px 32px;clear: both;">
<i class="fas fa-globe-asia" aria-hidden="true"></i>
<span data-toggle="modal" href="https://18comic.org/#language-modal">簡體
<i class="fas fa-chevron-right m-l-10 float-right" aria-hidden="true"></i>
</span>
<a href="#">
</a></li>
<li style="font-size: 14px;padding: 10px 32px;clear: both;"><a href="/cdn-cgi/l/email-protection#b6c1c1c1878ed5d9dbdfd5f6d1dbd7dfda98d5d9db">
<i class="fas fa-sign" aria-hidden="true"></i><span>广告洽询</span>
</a></li>
<a href="#">
</a><a href="/static/faq-cn-CS">
<li style="font-size: 14px;padding: 10px 32px;clear: both;">
<i class="far fa-question-circle" aria-hidden="true"></i><span>社群規範</span>
</li>
</a>
</div>
</div>
<li class="hidden-xs hidden-sm header-btn" id="adulta">
<a class="dropdown-toggle" data-toggle="dropdown">
<span class="hidden-xs hidden-sm">A漫<span class="caret"></span></span>
</a>
</li>
<li class="hidden-xs hidden-sm">
<a href="/theme/">分类</a>
</li>
<li class="hidden-xs hidden-sm header-btn " id="hmovies">
<a class="dropdown-toggle " data-toggle="dropdown" href="#" data-vodeos-count="3043" data-movies-count><span>H动漫<span class="caret"></span></span>
</a>
</li>
<li class="hidden-xs hidden-sm">
<a id="meun_game" class="meun-game" href="/games" data-count="274">游戏</a>
</li>
<li class="hidden-xs hidden-sm">
<a id="meun_blog" class="meun-blog" href="/blogs" data-count="894">紳夜食堂</a>
</li>
<li class="hidden-xs hidden-sm hidden-md">
<a href="/forum/">评论区</a>
</li>
</ul>

<ul class="nav navbar-nav navbar-right d-flex align-items-center ml-auto">
<li class="dropdown hidden-xs hidden-sm navbar-right-block m-r-10">
<a href="/week">
<div class="menu-circle"></div>
<i class="far fa-calendar-alt"></i>
</a>
</li>
<li class="dropdown hidden-xs hidden-sm navbar-right-block m-r-10">
<a href="https://jm365.work/pXYbfA" target="_blank" rel="nofollow noopenner">
<div class="menu-circle"></div>
<i class="fas fa-mobile-alt center" style="width: 16px;"></i>
</a>
</li>
<li class="hidden-xs hidden-sm navbar-right-block m-r-10">
<a href="/album/229151" style="padding: 5px;">
<img src="/templates/frontend/airav/img/phicon/paper_box.svg" style="width: 25px">
</a>
</li>
<li class="dropdown hidden-xs hidden-sm navbar-right-block m-r-10">
<a data-toggle="modal" href="#login-modal" style="display: flex;align-items: center; color:#ff7a00;">
<i class="far fa-user-circle user-block" aria-hidden="true"></i>
<span class="user-text">會員登入/註冊</span>
</a>
</li>
<li class="hidden-xs hidden-sm navbar-right-block m-r-10 pop_btn" id="lottery_list">
<a href="#" class="dropdown-toggle" data-toggle="dropdown"><i class="fa fa-search"></i></a>
</li>
</ul>

<div class="hidden-xs hidden-sm search-box pop_block" id="lottery_list_block" style="display: none;">
<form class="form-inline" name="search" id="search_form" method="get" action="/search/photos">
<span class="box-style">
<select id="header_search_type" class="box-style rounded-pill">
<option value="photos">A漫</option>
<option value="videos">H动漫</option>
<option value="movies">小电影</option>
<option value="blogs">紳夜食堂</option>
</select>
</span>
<span class="box-style">
<select class="form-control" name="main_tag" id="main_tag">
<option value="0" selected>站內搜索</option>
<option value="1">作品</option>
<option value="2">作者</option>
<option value="3">标签</option>
<option value="4">登场人物</option>
</select>
</span>
<input type="text" class="box-form" placeholder="可以搜尋圖片囉 ! " name="search_query" autocomplete="off" value>
<span>
<button type="submit" class="box-search-block">
<i class="fa fa-search" aria-hidden="true" style="color: #ff7a00"></i>
</button>
</span>
<span class="box-search-block" data-toggle="modal" data-target="#search-board">
<i class="far fa-question-circle" aria-hidden="true"></i>
</span>
<span class="box-search-block" data-toggle="modal" href="#search-Modalimg">
<i class="fas fa-camera user-active" aria-hidden="true"></i>
</span>
<span class="p-r-10 p-l-5 box-close modal-close">
<i class="fas fa-times"></i>
</span>
</form>
</div>
</div>
</div>

</div>

<div class="adult-main header-class-block hidden-xs hidden-sm" id="adulta-block" style="display: none;">
<div class="container">
<div class="row adult-block p-t-20 p-b-20">
<div class="col-md-3">
<div class="tag-hot-main">
<span class="text-orange m-b-10 d-block">熱門趨勢</span>
<div class="tag-hot tag-block">
<a href="/search/photos?search_query=连载中" class="btn btn-sm tag-block m-2" title="连载中">连载中</a>
<a href="/search/photos?search_query=吸奶" class="btn btn-sm tag-block m-2" title="吸奶">吸奶</a>
<a href="/search/photos?search_query=女性支配" class="btn btn-sm tag-block m-2" title="女性支配">女性支配</a>
<a href="/search/photos?search_query=韩漫的豪哥都好厉害" class="btn btn-sm tag-block m-2" title="韩漫的豪哥都好厉害">韩漫的豪哥都好厉害</a>
<a href="/search/photos?search_query=aslsdtkln書生" class="btn btn-sm tag-block m-2" title="aslsdtkln書生">aslsdtkln書生</a>
<a href="/search/photos?search_query=NUWARU" class="btn btn-sm tag-block m-2" title="NUWARU">NUWARU</a>
<a href="/search/photos?search_query=最强豪哥没有之一" class="btn btn-sm tag-block m-2" title="最强豪哥没有之一">最强豪哥没有之一</a>
<a href="/search/photos?search_query=請對標籤君手下留情b" class="btn btn-sm tag-block m-2" title="請對標籤君手下留情b">請對標籤君手下留情b</a>
<a href="/search/photos?search_query=干女儿韶多" class="btn btn-sm tag-block m-2" title="干女儿韶多">干女儿韶多</a>
<a href="/search/photos?search_query=韶多小猫咪" class="btn btn-sm tag-block m-2" title="韶多小猫咪">韶多小猫咪</a>
</div>
</div>
</div>
<div class="col-md-9">
<div class>
<span class="text-orange m-l-10 d-block">成人A漫</span>
<div class="adult-right-block">
<a href="/albums?o=mv" title="总排行">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/templates/frontend/airav/img/header/menu.total_ranking.jpg" alt="总排行">
</div>
<span>总排行</span>
</a>
<a href="/albums?t=m&o=mv" title="月排行">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/templates/frontend/airav/img/header/global.month.menu.ranking.jpg" alt="月排行">
</div>
<span>月排行</span>
</a>
<a href="/albums?o=mv&t=w" title="周排行">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/templates/frontend/airav/img/header/global.week.menu.ranking.jpg" alt="总排行">
</div>
<span>周排行</span>
</a>
<a href="/albums/doujin" title="同人">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/album/1.jpg?v=" alt="同人">
</div>
<span>同人</span>
</a>
<a href="/albums/single" title="单本">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/album/2.jpg?v=" alt="单本">
</div>
<span>单本</span>
</a>
<a href="/albums/short" title="短篇">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/album/3.jpg?v=" alt="短篇">
</div>
<span>短篇</span>
</a>
<a href="/albums/another" title="其他类">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/album/4.jpg?v=" alt="其他类">
</div>
<span>其他类</span>
</a>
<a href="/albums/hanman" title="韩漫">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/album/5.jpg?v=" alt="韩漫">
</div>
<span>韩漫</span>
</a>
<a href="https://6bq9.cc/dDcJP4" rel="nofollow noopener" class title="美漫">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/album/6.jpg?v=" alt="美漫">
</div>
<span>美漫</span>
</a>
</div>
</div>
</div>
</div>
</div>
</div>

<div class="adult-main header-class-block hidden-xs hidden-sm" id="hmovies-block" style="display: none;">
<div class="container">
<div class="row adult-block p-t-20 p-b-20">
<div class="col-md-3">
<div class="tag-hot-main">
<span class="text-orange m-b-10 d-block">熱門趨勢</span>
<div class="tag-hot tag-block">
<a href="/search/videos?search_query=可欲可甜" class="btn btn-sm tag-block" title="可欲可甜">
可欲可甜
</a>
<a href="/search/videos?search_query=3D" class="btn btn-sm tag-block" title="3D">
3D
</a>
<a href="/search/videos?search_query=乱伦" class="btn btn-sm tag-block" title="乱伦">
乱伦
</a>
<a href="/search/videos?search_query=异种姦" class="btn btn-sm tag-block" title="异种姦">
异种姦
</a>
<a href="/search/videos?search_query=姊妹" class="btn btn-sm tag-block" title="姊妹">
姊妹
</a>
<a href="/search/videos?search_query=おっぱい" class="btn btn-sm tag-block" title="おっぱい">
おっぱい
</a>
<a href="/search/videos?search_query=小姐姐" class="btn btn-sm tag-block" title="小姐姐">
小姐姐
</a>
<a href="/search/videos?search_query=卡通" class="btn btn-sm tag-block" title="卡通">
卡通
</a>
<a href="/search/videos?search_query=3P、4P" class="btn btn-sm tag-block" title="3P、4P">
3P、4P
</a>
<a href="/search/videos?search_query=痴女" class="btn btn-sm tag-block" title="痴女">
痴女
</a>
</div>
</div>
</div>
<div class="col-md-9">
<div class>
<span class="text-orange m-l-10 d-block">H动漫</span>
<div class="adult-right-block">
<a href="/videos" title="H动漫">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/video/12.jpg" alt="H动漫">
</div>
<span>H动漫</span>
</a>
<a href="/movies" title="小电影">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/templates/frontend/airav/img/header/menu.movie.jpg" alt="小电影">
</div>
<span>小电影</span>
</a>
<a href="/search/movies?main_tag=0&search_query=马赛克破坏" title="马赛克破坏">
<div class="m-b-5">
<img class="lazy_img" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-original="/media/categories/video/nocode.jpg" alt="马赛克破坏">
</div>
<span>马赛克破坏</span>
</a>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
<div id="wrapper">

<div class="ph-bottom switch-group visible-xs visible-sm">
<ul>
<li class="ph-active">
<a href="/videos/kajilu" title="频道">
<img src="/templates/frontend/airav/img/phicon/youtube.svg?v=2023081601" style="height: 27px; margin: auto;display: block;" alt="频道">
<span>频道</span>
</a>
</li>
<li class="ph-active">
<a href="/theme/">
<img src="/templates/frontend/airav/img/phicon/u441.png?v=2023081601" style="height: 27px; margin: auto;display: block;" alt="类别">
<span>类别</span>
</a>
</li>
<li class="ph-active">
<a href="/movies/" class="meun-video">
<i class="fas fa-glasses"></i>
<span>小电影</span>
</a>
</li>
<li class="ph-active">
<a href="/games" class="meun-game">
<i class="fas fa-gamepad"></i>
<span>游戏</span>
</a>
</li>
<li class="ph-active switch" id="sns">
<a href="/forum/">
<i class="fas fa-comment"></i>
<span>评论区</span>
</a>
</li>
<li class="ph-active">
<a href="/user">
<i class="fas fa-user-circle"></i>
<span>我的</span>
</a>
</li>
</ul>
</div>
<style>
            @media screen and (min-width: 767px) {
                .float-right-image {
                    position: fixed;
                    right: 0;
                    width: 18%;
                    bottom: 0px;
                    z-index: 100;
                    opacity: 1
                }

                .float-right-image.open {
                    display: none;
                }

                .float-right-image {
                    display: block;
                }
            }

            .select-type-title p {
                color: #ababab;
            }

            @media screen and (max-width: 767px) {
                .top-nav .container {
                    padding-right: 0 !important;
                }

                .top-menu > .pull-right > li {
                    margin: 0 !important;
                }
            }
        </style>
<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script>

            $(document).ready(function() {
                var SearchPosition = $('.search-dropdown-menu').offset().top;

                $('input[name="search_query"]').on('focus', function() {
                    $('html, body').animate({ scrollTop: SearchPosition - 10 }, 200);
                });

                $('input[name="search_query"]').on('blur', function() {
                    $('html, body').animate({ scrollTop: 0 }, 200);
                });
            });

            function scrollToImageSearchOption() {
                var inputTop = $('#img-type').offset().top;
                var headerHeight = $(".navbar").height();
                var newScrollTop = inputTop - headerHeight - 10; // 10 是額外的間距
                $('html, body').animate({ scrollTop: newScrollTop }, 500);
            }

            $(document).ready(function() {
              $('#img-search-option').on('click', scrollToImageSearchOption);
            });

        </script>
<script>

            $(document).on('mouseenter', '.float-right-image.close', function () {
                $('.float-right-image.open').css('display', 'block');
                $('.float-right-image.close').css('display', 'none');
            });
            $(document).on('mouseleave', '.float-right-image.open', function () {
                $('.float-right-image.open').css('display', 'none');
                $('.float-right-image.close').css('display', 'block');
            });
            toastr.options = {
                "closeButton": true,
                "positionClass": "toast-top-center",
            }
            $(document).ready(function () {
                $("img.lazy_img,img.lazyad").lazyload({
                    threshold: 1,
                    rootMargin: '0px 0px 500px 0px',
                    src: 'data-original'
                });
            });
        </script>
<div class="float-right-image open hidden-xs hidden-sm hidden-md">
<img src="/templates/frontend/airav/img/float-right-open.png?v=" border="0" style="width: 100%;height: auto;right: 10%;">
<a href="https://t.me/hcomic18" target="_blank" rel="nofollow noopener" style="top: 27%;left: 8%;width: 17%;height: 18%;position: absolute;"></a>
<a href="https://www.instagram.com/18comicorg/" target="_blank" rel="nofollow noopener" style="top: 27%;left: 26%;width: 17%;height: 18%;position: absolute;"></a>
<a href="https://discord.gg/V74p7HM" target="_blank" rel="nofollow noopener" style="top: 45%; left: 8%;width: 17%;height: 18%;position: absolute;"></a>
<a href="https://www.facebook.com/senshimanga/" target="_blank" rel="nofollow noopener" style="top: 45%; left: 26%;width: 17%;height: 18%;position: absolute;"></a>
<a href="/shop" data-group="PC-Float" style="top: 64%; left: 8%;width: 17%;height: 18%;position: absolute;"></a>
<a href="https://reurl.cc/WrbxXD" target="_blank" rel="nofollow noopener" style="top: 64%; left: 26%;width: 17%;height: 18%;position: absolute;"></a>
</div>
<div class="float-right-image close hidden-xs hidden-sm hidden-md">
<img src="/templates/frontend/airav/img/float-right-close.png?v=" border="0" style="width: 100%;height: auto;right: 10%;">
</div>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.voting-tag-album-0.1.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.album-0.2.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/jquery.photo-0.4.js?v=2023081601"></script>
<script type="text/javascript" src="/templates/frontend/airav/js/md5.min.js"></script>
<link type="text/css" rel="stylesheet" href="/templates/frontend/airav/css/comment.css?v=2023081601">
<link type="text/css" rel="stylesheet" href="/templates/frontend/airav/css/album.css?v=2023081601">
<script>
    var lang_delete_photo_ask = "Are you sure you want to delete this photo?";
    var aid = 419411;
</script>
<script src="/templates/frontend/airav/js/webp-hero/dist-cjs/polyfills.js"></script>
<script src="/templates/frontend/airav/js/webp-hero/dist-cjs/webp-hero.bundle.js"></script>
<div id="popup"></div>
<div class="container">
<div style="max-height: 100px" class="e8c78e-4_b top-a2db" data-height="90" data-width="728" data-group="content_page">
<script async type="application/javascript" src="https://a.magsrv.com/ad-provider.js"></script>
<ins class="adsbyexoclick" data-zoneid="4286790" data-sub="136"></ins>
<script>(AdProvider = window.AdProvider || []).push({"serve": {}});</script>
<script type="text/javascript">
                    (function () {
                        function randStr(e,t){for(var n="",r=t||"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",o=0;o<e;o++)n+=r.charAt(Math.floor(Math.random()*r.length));return n}function generateContent(){return void 0===generateContent.val&&(generateContent.val=" \ndocument.dispatchEvent("+randStr(4*Math.random()+3)+");"),generateContent.val}try{Object.defineProperty(document.currentScript,"innerHTML",{get:generateContent}),Object.defineProperty(document.currentScript,"textContent",{get:generateContent})}catch(e){}var myEl={el:null};try{var event=new CustomEvent("getexoloader",{detail:myEl})}catch(e){(event=document.createEvent("CustomEvent")).initCustomEvent("getexoloader",!1,!1,myEl)}window.document.dispatchEvent(event);var ExoLoader=myEl.el;
                        ExoLoader.addZone({"idzone":"4286790","sub":"136"});                        
                    })();
                    </script>
</div>
<div id="favorite_message" style="display:none;"></div>
<div id="response_message"></div>
<div class="row">
<input id="album_id" type="hidden" value="419411">
<div class="col-md-12">
<div itemscope itemtype="http://schema.org/Book" class="panel panel-default hidden-lg">
<div class="panel-body">
<div class="row">
<div id="album_photo_cover" class="col-lg-12">
<div class="thumb-overlay col-xs-12 p-0">
<div class="thumb-overlay col-xs-12 p-0">
<div class="owl-carousel cover-view owl-loaded owl-drag">
<div class="show_zoom">
<img itemprop="image" src="https://cdn-msp.orangeapp.cc/media/albums/419411.jpg?v=1675243580" />
</div>
<div id="00238.webp" class="show_zoom">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/photos/419411/00238.webp" class="owl-lazy" itemprop="image" id="album_photo_00238.webp" />
</div>
<div id="00172.webp" class="show_zoom">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/photos/419411/00172.webp" class="owl-lazy" itemprop="image" id="album_photo_00172.webp" />
</div>
<div id="00178.webp" class="show_zoom">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/photos/419411/00178.webp" class="owl-lazy" itemprop="image" id="album_photo_00178.webp" />
</div>
</div>
<div class="absolute train-number">
<div>
<span class="number">禁漫车：JM419411</span>
<span class="pagecount">页数:245</span>
</div>
<div class="float-right btn-default" style="color:#fffff" data-toggle="modal" href="#login-modal">收藏</div>
</div>
</div>
</div>
</div>
<div class="m-t-10 m-b-10 book-name-collep col-xs-11">
<div itemprop="name" class="pull-left">
<h1 class="book-name" id="book-name">我會對著妳汪汪的叫著 [Type-G (イシガキタカシ)] 僕はあなたにワンと鳴く</h1>
</div>
<div class="book-name-btn">
<i class="fas fa-arrow-circle-down"></i>
</div>
</div>
<div class="col-xs-12">
<p class="p-t-5 float-left">
上架日期 : 2023-02-01
</p>
<p class="p-t-5 float-right">
更新日期 : 2023-02-01
</p>
</div>
<div class="col-xs-12 p-t-5 p-b-5">
<div class="p-t-5 p-b-5 read-block">
<li role="presentation" class="p-t-5 p-b-5 col justify-content-center active" id="intro">
<span>详情</span></li>
<a class="col btn btn-primary dropdown-toggle reading" href="/photo/419411">
开始阅读
</a>
</div>
</div>
<div class="col-xs-12 col-lg-7 nav-tab-content" id="intro-block">
<ul class="text-center">
<li class="list-style-none d-inline-block p-r-15">
<a href="#" id="love_likes_419411" style="color: #777777">
<i class="fas fa-thumbs-up"></i><span id="albim_likes_419411">1K</span>
</a>
</li>
<li class="list-style-none d-inline-block">
<i class="far fa-eye "></i><span>97K</span>
</li>
<li class="list-style-none d-inline-block" style="padding-left:2rem">
<button type="button" class="btn btn-secondary" data-toggle="modal" data-target="#jmComicVideo">短視頻分享</button>

<div class="modal fade" id="jmComicVideo" tabindex="-1" role="dialog" aria-labelledby="jmComicVideoLabel" aria-hidden="true">
<div class="modal-dialog modal-lg" role="document">
<div class="modal-content">
<div class="modal-header">
<h5 class="modal-title" id="jmComicVideoLabel">視頻預覽</h5>
<button type="button" class="close" data-dismiss="modal" aria-label="Close" style="margin-top:-3rem;font-size:40px;">
<span aria-hidden="true">&times;</span>
</button>
</div>
<div class="modal-body d-flex justify-content-center align-items-center">
<video class="chkvdo" width="350" controls>
<source src="/media/JmShortVideo/419411.mp4" type="video/mp4">
</video>
</div>
<div class="modal-footer">
<button type="button" class="btn btn-secondary" data-dismiss="modal">取消</button>
</div>
</div>
</div>
</div>
</li>
</ul>
<div class="p-t-5 p-b-5"> 叙述：https://www.pixiv.net/users/254423
https://www.dlsite.com/books/work/=/product_id/BJ056134.html
</div>
<div class="tag-block">
作品： <span itemprop="author" data-type="works">
</span>
<label name="work" for="work" class="form-inline hidden-xs">
<input name="work" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
登场人物： <span itemprop="author" data-type="actor">
</span>
<label name="actor" for="actor" class="form-inline hidden-xs">
<input name="actor" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
标签： <span itemprop="genre" data-type="tags">
<a name="vote_" href="/search/photos?search_query=個人整合" class="btn btn-sm btn-primary">個人整合</a>
<a name="vote_" href="/search/photos?search_query=劇情向" class="btn btn-sm btn-primary">劇情向</a>
<a name="vote_" href="/search/photos?search_query=純愛" class="btn btn-sm btn-primary">純愛</a>
<a name="vote_" href="/search/photos?search_query=扶他" class="btn btn-sm btn-primary">扶他</a>
<a name="vote_" href="/search/photos?search_query=百合" class="btn btn-sm btn-primary">百合</a>
<a name="vote_" href="/search/photos?search_query=黑肉" class="btn btn-sm btn-primary">黑肉</a>
<a name="vote_" href="/search/photos?search_query=校服" class="btn btn-sm btn-primary">校服</a>
<a name="vote_" href="/search/photos?search_query=馬尾" class="btn btn-sm btn-primary">馬尾</a>
<a name="vote_" href="/search/photos?search_query=浴衣" class="btn btn-sm btn-primary">浴衣</a>
<a name="vote_" href="/search/photos?search_query=巨乳" class="btn btn-sm btn-primary">巨乳</a>
<a name="vote_" href="/search/photos?search_query=泳裝" class="btn btn-sm btn-primary">泳裝</a>
<a name="vote_" href="/search/photos?search_query=吊帶襪" class="btn btn-sm btn-primary">吊帶襪</a>
<a name="vote_" href="/search/photos?search_query=女性支配" class="btn btn-sm btn-primary">女性支配</a>
<a name="vote_" href="/search/photos?search_query=強暴" class="btn btn-sm btn-primary">強暴</a>
<a name="vote_" href="/search/photos?search_query=野砲" class="btn btn-sm btn-primary">野砲</a>
<a name="vote_" href="/search/photos?search_query=調教" class="btn btn-sm btn-primary">調教</a>
<a name="vote_" href="/search/photos?search_query=群交" class="btn btn-sm btn-primary">群交</a>
<a name="vote_" href="/search/photos?search_query=口交" class="btn btn-sm btn-primary">口交</a>
<a name="vote_" href="/search/photos?search_query=乳交" class="btn btn-sm btn-primary">乳交</a>
<a name="vote_" href="/search/photos?search_query=透視" class="btn btn-sm btn-primary">透視</a>
<a name="vote_" href="/search/photos?search_query=中出" class="btn btn-sm btn-primary">中出</a>
<a name="vote_" href="/search/photos?search_query=中文" class="btn btn-sm btn-primary">中文</a>
</span>
<label name="tags" for="tags" class="form-inline hidden-xs">
<input name="tags" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
作者： <span itemprop="author" data-type="author">
<a name="vote_" href="/search/photos?search_query=イシガキタカシ" class="btn btn-sm btn-primary">イシガキタカシ</a>
</span>
<label name="author" for="author" class="form-inline hidden-xs">
<input name="author" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
<span class="btn-team">
<span id="search_tag" class="btn btn-sm btn-grey disabled">搜索标签</span>
<span data-toggle="modal" href="#login-modal" class="btn btn-sm btn-grey disabled ">票选+1</span>
<span data-toggle="modal" href="#login-modal" class="btn btn-sm btn-grey disabled ">负分-1</span>
<span class="btn btn-sm btn-grey disabled cancel">取消选取</span>
</span>
</div>
<div class="form-inline tag-block text-center">
<div class="visible-xs-block form-group">
<span data-toggle="modal" href="#login-modal" class="btn btn-default col-xs-12">新增标签<span class="glyphicon glyphicon-plus m-l-10" aria-hidden="true"></span>
</span>
</div>
</div>
</div>
<div class="center">

<div class="addthis_inline_share_toolbox"></div>

<script async type="text/javascript" src="//s7.addthis.com/js/300/addthis_widget.js#pubid=ra-5de7621dc0c9a8cd"></script>
</div>
<div style="clear: both;padding: 0!important; margin-top: 10px;z-index: 0;position:
                    relative;" class="e8c78e-4_b" data-height="90" data-width="728" data-group="album_detail">
<p><a rel="nofollow noopener" href="https://zno290.com/nydkg/" target="_blank"><img alt src="/static/resources/images/728x90-20230803.gif" /></a></p>
</div>
</div>
</div>
</div>
<div itemscope itemtype="http://schema.org/Book" class="panel panel-default visible-lg hidden-xs">
<div class="panel-heading">
<div itemprop="name" class="pull-left">
<h1>我會對著妳汪汪的叫著 [Type-G (イシガキタカシ)] 僕はあなたにワンと鳴く</h1>
</div>
<div class="clearfix"></div>
</div>
<div class="panel-body">
<div class="row">
<div id="album_photo_cover" class="col-lg-5">
<div class="thumb-overlay">
<a href="/photo/419411/">
<div class="thumb-overlay">
<img itemprop="image" src="https://cdn-msp.orangeapp.cc/media/albums/419411.jpg?v=1675243580" class="lazy_img img-responsive" />
</div>
</a>
</div>
<div class="img_zoom" style="width: 100%;">
<div class="img_zoom_img" id="00238.webp">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" id="zoom_m_0" data-original="https://cdn-msp.orangeapp.cc/media/photos/419411/00238.webp" class="lazy_img img-responsive" style="display: none" data-page="0" itemprop="image" />
</div>
<div class="img_zoom_img" id="00172.webp">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" id="zoom_m_1" data-original="https://cdn-msp.orangeapp.cc/media/photos/419411/00172.webp" class="lazy_img img-responsive" style="display: none" data-page="1" itemprop="image" />
</div>
<div class="img_zoom_img" id="00178.webp">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" id="zoom_m_2" data-original="https://cdn-msp.orangeapp.cc/media/photos/419411/00178.webp" class="lazy_img img-responsive" style="display: none" data-page="2" itemprop="image" />
</div>
</div>
</div>
<div class="col-lg-7">
<div class>
<div class="p-t-5 p-b-5">
禁漫车：JM419411
</div>
<div class="tag-block">
作品： <span itemprop="author" data-type="works">
</span>
<label name="work" for="work" class="form-inline hidden-xs">
<input name="work" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
登场人物： <span itemprop="author" data-type="actor">
</span>
<label name="actor" for="actor" class="form-inline hidden-xs">
<input name="actor" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
标签： <span itemprop="genre" data-type="tags">
<a name="vote_" href="/search/photos?search_query=個人整合" class="btn btn-sm btn-primary">個人整合</a>
<a name="vote_" href="/search/photos?search_query=劇情向" class="btn btn-sm btn-primary">劇情向</a>
<a name="vote_" href="/search/photos?search_query=純愛" class="btn btn-sm btn-primary">純愛</a>
<a name="vote_" href="/search/photos?search_query=扶他" class="btn btn-sm btn-primary">扶他</a>
<a name="vote_" href="/search/photos?search_query=百合" class="btn btn-sm btn-primary">百合</a>
<a name="vote_" href="/search/photos?search_query=黑肉" class="btn btn-sm btn-primary">黑肉</a>
<a name="vote_" href="/search/photos?search_query=校服" class="btn btn-sm btn-primary">校服</a>
<a name="vote_" href="/search/photos?search_query=馬尾" class="btn btn-sm btn-primary">馬尾</a>
<a name="vote_" href="/search/photos?search_query=浴衣" class="btn btn-sm btn-primary">浴衣</a>
<a name="vote_" href="/search/photos?search_query=巨乳" class="btn btn-sm btn-primary">巨乳</a>
<a name="vote_" href="/search/photos?search_query=泳裝" class="btn btn-sm btn-primary">泳裝</a>
<a name="vote_" href="/search/photos?search_query=吊帶襪" class="btn btn-sm btn-primary">吊帶襪</a>
<a name="vote_" href="/search/photos?search_query=女性支配" class="btn btn-sm btn-primary">女性支配</a>
<a name="vote_" href="/search/photos?search_query=強暴" class="btn btn-sm btn-primary">強暴</a>
<a name="vote_" href="/search/photos?search_query=野砲" class="btn btn-sm btn-primary">野砲</a>
<a name="vote_" href="/search/photos?search_query=調教" class="btn btn-sm btn-primary">調教</a>
<a name="vote_" href="/search/photos?search_query=群交" class="btn btn-sm btn-primary">群交</a>
<a name="vote_" href="/search/photos?search_query=口交" class="btn btn-sm btn-primary">口交</a>
<a name="vote_" href="/search/photos?search_query=乳交" class="btn btn-sm btn-primary">乳交</a>
<a name="vote_" href="/search/photos?search_query=透視" class="btn btn-sm btn-primary">透視</a>
<a name="vote_" href="/search/photos?search_query=中出" class="btn btn-sm btn-primary">中出</a>
<a name="vote_" href="/search/photos?search_query=中文" class="btn btn-sm btn-primary">中文</a>
</span>
<label name="tags" for="tags" class="form-inline hidden-xs">
<input name="tags" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
作者： <span itemprop="author" data-type="author">
<a name="vote_" href="/search/photos?search_query=イシガキタカシ" class="btn btn-sm btn-primary">イシガキタカシ</a>
</span>
<label name="author" for="author" class="form-inline hidden-xs">
<input name="author" type="text" class="form-control input-sm">
<span class="btn btn-sm add-tag">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
<span class="btn btn-sm add-inp" data-toggle="modal" href="#login-modal">
<span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
</span>
</label>
</div>
<div class="tag-block">
<span class="btn-team">
<span id="search_tag" class="btn btn-sm btn-grey disabled">搜索标签</span>
<span data-toggle="modal" href="#login-modal" class="btn btn-sm btn-grey disabled ">票选+1</span>
<span data-toggle="modal" href="#login-modal" class="btn btn-sm btn-grey disabled ">负分-1</span>
<span class="btn btn-sm btn-grey disabled cancel">取消选取</span>
</span>
</div>
<div class="form-inline tag-block text-center">
<div class="visible-xs-block form-group">
<span data-toggle="modal" href="#login-modal" class="btn btn-default col-xs-12">新增标签<span class="glyphicon glyphicon-plus m-l-10" aria-hidden="true"></span>
</span>
</div>
</div>
<div class="p-t-5 p-b-5">
叙述：https://www.pixiv.net/users/254423
https://www.dlsite.com/books/work/=/product_id/BJ056134.html
</div>
<div class="p-t-5 p-b-5">
页数：245
</div>
<div class="p-t-5 p-b-5">
<span style="margin-right: 10px" itemprop="datePublished" content="2023-02-01">上架日期 : 2023-02-01</span>
<span style="margin-right: 10px" itemprop="datePublished" content="2023-02-01">更新日期 : 2023-02-01</span>
<span style="margin-right: 10px">
<span>
<span>97K</span> 观看次数
</span>
</span>
<span style="margin-right: 10px" class="p-t-5 p-b-5">
<span id="albim_likes_419411">1K</span>
<span> </span>
</span>
</div>
</div>
<div class="p-t-5 p-b-5 read-block">
<a class="btn btn-primary" href="/photo/419411">开始阅读</a>
<a class="btn btn-primary" href="/photo/419411?read_mode=read-by-page">分页阅读</a>
<a class="forum-open btn btn-primary" href="javascript:;">评论
<div class="badge" id="total_video_comments">23</div> </a>
<a style="padding: 5px;" href="/album_download/419411">
<i class="fas fa-download fa-2x"></i>
</a>
<a data-toggle="modal" href="#login-modal" style="float: right;padding: 5px;">
<i style="color:#000000" class="far fa-bookmark fa-2x"></i>
</a>
<a href="#" style="float: right;padding: 5px;" id="love_likes_419411">
<i class="glyphicon glyphicon-heart fa-2x"></i>
</a>
<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#jmComicMbVideo">短視頻分享</button>

<div class="modal fade" id="jmComicMbVideo" tabindex="-1" role="dialog" aria-labelledby="jmComicMbVideoLabel" aria-hidden="true">
<div class="modal-dialog modal-lg" role="document">
<div class="modal-content">
<div class="modal-header">
<h5 class="modal-title" id="jmComicMbVideoLabel">視頻預覽</h5>
<button type="button" class="close" data-dismiss="modal" aria-label="Close" style="margin-top:-3rem;font-size:40px;">
<span aria-hidden="true">&times;</span>
</button>
</div>
<div class="modal-body d-flex justify-content-center align-items-center">
<video class="chkvdo" width="350" controls>
<source src="/media/JmShortVideo/419411.mp4" type="video/mp4">
</video>
</div>
<div class="modal-footer">
<button type="button" class="btn btn-secondary" data-dismiss="modal">取消</button>
</div>
</div>
</div>
</div>
</div>
<div class="center">

<div class="addthis_inline_share_toolbox"></div>

<script type="text/javascript" src="//s7.addthis.com/js/300/addthis_widget.js#pubid=ra-5de7621dc0c9a8cd"></script>
</div>
<div style="clear: both;padding: 0!important; margin-top: 10px; z-index: 0;" class="e8c78e-4_b" data-height="90" data-width="728" data-group="album_detail">
<p><a rel="nofollow noopener" href="https://zno290.com/nydkg/" target="_blank"><img alt src="/static/resources/images/728x90-20230803.gif" /></a></p>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
<ul class="nav nav-tabs">
<li class="active"><a id="related_comics_a" href="#related_comics" data-toggle="tab">相关A漫</a>
</li>
<li class><a id="comments_a" href="#comments" data-toggle="tab">评论 <div class="badge" id="total_video_comments">23</div></a></li>
</ul>
<div class="tab-content m-t-15 m-b-20">
<div class="tab-pane fade active in" id="related_comics">
<link href="/templates/frontend/airav/css/style_game.css?v=2023081601" rel="stylesheet">
<link href="/templates/frontend/airav/css/style_blog.css?v=2023081601" rel="stylesheet">
<div class="row m-l-0 m-r-0 m-b-10">
<h4>相关A漫</h4>
<ul class="owl-carousel owl-comic-block">
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/265033/箕山-深窓の華娵-dl版-個人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/265033_3x4.jpg?v=1632732136" title="[箕山] 深窓の華娵 [DL版] 個人整合" alt="[箕山] 深窓の華娵 [DL版] 個人整合" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_265033"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_265033" class="text-white">6K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">[箕山] 深窓の華娵 [DL版] 個人整合</span>
<div class="title-truncate">
<a href="/search/photos?search_query=箕山">箕山</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=站長推薦">站長推薦</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=泳裝">泳裝</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=辣妹">辣妹</a>, <a href="/search/photos?search_query=處女">處女</a>, <a href="/search/photos?search_query=校園">校園</a>, <a href="/search/photos?search_query=姐弟">姐弟</a>, <a href="/search/photos?search_query=亂倫">亂倫</a>, <a href="/search/photos?search_query=教師">教師</a>, <a href="/search/photos?search_query=師生">師生</a>, <a href="/search/photos?search_query=手交">手交</a>, <a href="/search/photos?search_query=騎大車">騎大車</a>, <a href="/search/photos?search_query=馬尾">馬尾</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=個人整合">個人整合</a> </div>
<div class="video-views pull-left">
2021-07-05
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/285676/個人整合-幾花にいろ-丹">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/285676_3x4.jpg?v=1655790462" title="個人整合[幾花にいろ] 丹" alt="個人整合[幾花にいろ] 丹" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_285676"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_285676" class="text-white">4K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">個人整合[幾花にいろ] 丹</span>
<div class="title-truncate">
<a href="/search/photos?search_query=幾花にいろ">幾花にいろ</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=個人整合">個人整合</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=阿黑顏">阿黑顏</a>, <a href="/search/photos?search_query=出軌">出軌</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=眼鏡">眼鏡</a>, <a href="/search/photos?search_query=貧乳">貧乳</a>, <a href="/search/photos?search_query=中文">中文</a> </div>
<div class="video-views pull-left">
2021-10-18
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/233716/源五郎-愛しき我が家-dl版-黑暗大法师个人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/233716_3x4.jpg?v=1631693657" title="[源五郎] 愛しき我が家 [DL版] [黑暗大法师个人整合]" alt="[源五郎] 愛しき我が家 [DL版] [黑暗大法师个人整合]" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_233716"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_233716" class="text-white">5K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">[源五郎] 愛しき我が家 [DL版] [黑暗大法师个人整合]</span>
<div class="title-truncate">
<a href="/search/photos?search_query=源五郎">源五郎</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=个人整合">个人整合</a>, <a href="/search/photos?search_query=站長推薦">站長推薦</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=蘿莉">蘿莉</a>, <a href="/search/photos?search_query=女僕">女僕</a>, <a href="/search/photos?search_query=束縛">束縛</a>, <a href="/search/photos?search_query=懷孕">懷孕</a>, <a href="/search/photos?search_query=無口">無口</a>, <a href="/search/photos?search_query=狐娘">狐娘</a>, <a href="/search/photos?search_query=妖精">妖精</a>, <a href="/search/photos?search_query=獸耳">獸耳</a>, <a href="/search/photos?search_query=正太">正太</a>, <a href="/search/photos?search_query=乳汁">乳汁</a>, <a href="/search/photos?search_query=手交">手交</a>, <a href="/search/photos?search_query=透視">透視</a>, <a href="/search/photos?search_query=睡姦">睡姦</a>, <a href="/search/photos?search_query=亂倫">亂倫</a>, <a href="/search/photos?search_query=父女">父女</a>, <a href="/search/photos?search_query=群交">群交</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=足交">足交</a>, <a href="/search/photos?search_query=女性支配">女性支配</a>, <a href="/search/photos?search_query=傲嬌">傲嬌</a>, <a href="/search/photos?search_query=中文">中文</a> </div>
<div class="video-views pull-left">
2021-01-23
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/277916/背德之中-歯車-インナーインモラル-個人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/277916_3x4.jpg?v=1639641412" title="背德之中 [歯車] インナーインモラル 個人整合" alt="背德之中 [歯車] インナーインモラル 個人整合" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_277916"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_277916" class="text-white">5K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">背德之中 [歯車] インナーインモラル 個人整合</span>
<div class="title-truncate">
<a href="/search/photos?search_query=歯車">歯車</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=站長推薦">站長推薦</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=女高中生">女高中生</a>, <a href="/search/photos?search_query=野砲">野砲</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=無口">無口</a>, <a href="/search/photos?search_query=透視">透視</a>, <a href="/search/photos?search_query=師生">師生</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=女性支配">女性支配</a>, <a href="/search/photos?search_query=騎大車">騎大車</a>, <a href="/search/photos?search_query=巨乳">巨乳</a>, <a href="/search/photos?search_query=吸奶">吸奶</a>, <a href="/search/photos?search_query=阿黑顏">阿黑顏</a>, <a href="/search/photos?search_query=眼鏡">眼鏡</a>, <a href="/search/photos?search_query=過膝襪">過膝襪</a>, <a href="/search/photos?search_query=緊身衣">緊身衣</a>, <a href="/search/photos?search_query=姊弟">姊弟</a>, <a href="/search/photos?search_query=假小子">假小子</a>, <a href="/search/photos?search_query=亂倫">亂倫</a>, <a href="/search/photos?search_query=辣妹">辣妹</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=個人整合">個人整合</a> </div>
<div class="video-views pull-left">
2021-09-10
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/296117/獨自佔有-ドラムス-ヒトリジメ-個人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/296117_3x4.jpg?v=1645669075" title="獨自佔有 [ドラムス] ヒトリジメ 個人整合" alt="獨自佔有 [ドラムス] ヒトリジメ 個人整合" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_296117"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_296117" class="text-white">10K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">獨自佔有 [ドラムス] ヒトリジメ 個人整合</span>
<div class="title-truncate">
<a href="/search/photos?search_query=ドラムス">ドラムス</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=NTR">NTR</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=性勒索">性勒索</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=過膝襪">過膝襪</a>, <a href="/search/photos?search_query=雙胞胎">雙胞胎</a>, <a href="/search/photos?search_query=項圈">項圈</a>, <a href="/search/photos?search_query=群交">群交</a>, <a href="/search/photos?search_query=巨乳">巨乳</a>, <a href="/search/photos?search_query=拍攝">拍攝</a>, <a href="/search/photos?search_query=透視">透視</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=個人整合">個人整合</a> </div>
<div class="video-views pull-left">
2021-12-08
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/269316/性春之癮-きいろいたまご-性春ホリック-dl版-個人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/269316_3x4.jpg?v=1637739323" title="性春之癮 [きいろいたまご] 性春ホリック [DL版] 個人整合" alt="性春之癮 [きいろいたまご] 性春ホリック [DL版] 個人整合" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_269316"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_269316" class="text-white">14K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">性春之癮 [きいろいたまご] 性春ホリック [DL版] 個人整合</span>
<div class="title-truncate">
<a href="/search/photos?search_query=きいろいたまご">きいろいたまご</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=蘿莉">蘿莉</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=校園">校園</a>, <a href="/search/photos?search_query=馬尾">馬尾</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=肛交">肛交</a>, <a href="/search/photos?search_query=黑肉">黑肉</a>, <a href="/search/photos?search_query=假小子">假小子</a>, <a href="/search/photos?search_query=亂倫">亂倫</a>, <a href="/search/photos?search_query=處女">處女</a>, <a href="/search/photos?search_query=雙馬尾">雙馬尾</a>, <a href="/search/photos?search_query=吸奶">吸奶</a>, <a href="/search/photos?search_query=貧乳">貧乳</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=個人整合">個人整合</a> </div>
<div class="video-views pull-left">
2021-07-28
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/226201/黑色彼岸花個人整合-はるきち-初恋ショコラ-初戀的巧克力-特装版">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/226201_3x4.jpg?v=1607145139" title="[黑色彼岸花個人整合] [はるきち] 初恋ショコラ 初戀的巧克力 【特装版】" alt="[黑色彼岸花個人整合] [はるきち] 初恋ショコラ 初戀的巧克力 【特装版】" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_226201"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_226201" class="text-white">573</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">[黑色彼岸花個人整合] [はるきち] 初恋ショコラ 初戀的巧克力 【特装版】</span>
<div class="title-truncate">
<a href="/search/photos?search_query=はるきち">はるきち</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=個人整合">個人整合</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=騎大車">騎大車</a>, <a href="/search/photos?search_query=黑肉">黑肉</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=中文">中文</a> </div>
<div class="video-views pull-left">
2020-12-04
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/231804/黑暗大法师个人整合-shibi-サキュバスカンパニー">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/231804_3x4.jpg?v=1627649245" title="[黑暗大法师个人整合] [ShiBi] サキュバスカンパニー" alt="[黑暗大法师个人整合] [ShiBi] サキュバスカンパニー" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_231804"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_231804" class="text-white">11K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">[黑暗大法师个人整合] [ShiBi] サキュバスカンパニー</span>
<div class="title-truncate">
<a href="/search/photos?search_query=ShiBi">ShiBi</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=个人整合">个人整合</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=後宮">後宮</a>, <a href="/search/photos?search_query=巨乳">巨乳</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=束縛">束縛</a>, <a href="/search/photos?search_query=魅魔">魅魔</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=群交">群交</a>, <a href="/search/photos?search_query=中文">中文</a> </div>
<div class="video-views pull-left">
2021-01-12
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/270650/dopuya-八尋ぽち-どぴゅあ-どぴゅあ8p小冊子-個人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/270650_3x4.jpg?v=1687067176" title="Dopuya [八尋ぽち] どぴゅあ +どぴゅあ8P小冊子 個人整合" alt="Dopuya [八尋ぽち] どぴゅあ +どぴゅあ8P小冊子 個人整合" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_270650"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_270650" class="text-white">12K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">Dopuya [八尋ぽち] どぴゅあ +どぴゅあ8P小冊子 個人整合</span>
<div class="title-truncate">
<a href="/search/photos?search_query=八尋ぽち">八尋ぽち</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=女高中生">女高中生</a>, <a href="/search/photos?search_query=雙馬尾">雙馬尾</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=束縛">束縛</a>, <a href="/search/photos?search_query=馬尾">馬尾</a>, <a href="/search/photos?search_query=辣妹">辣妹</a>, <a href="/search/photos?search_query=巨乳">巨乳</a>, <a href="/search/photos?search_query=眼罩">眼罩</a>, <a href="/search/photos?search_query=女性支配">女性支配</a>, <a href="/search/photos?search_query=野砲">野砲</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=手交">手交</a>, <a href="/search/photos?search_query=透視">透視</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=亂倫">亂倫</a>, <a href="/search/photos?search_query=過膝襪">過膝襪</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=個人整合">個人整合</a> </div>
<div class="video-views pull-left">
2021-08-04
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/187420/黑暗大法师个人整合-武藤まと-こあくまは小動物-4pリーフレット">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/187420_3x4.jpg?v=1689739938" title="[黑暗大法师个人整合] [武藤まと] こあくまは小動物 + 4Pリーフレット" alt="[黑暗大法师个人整合] [武藤まと] こあくまは小動物 + 4Pリーフレット" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_187420"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_187420" class="text-white">3K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">[黑暗大法师个人整合] [武藤まと] こあくまは小動物 + 4Pリーフレット</span>
<div class="title-truncate">
<a href="/search/photos?search_query=武藤まと">武藤まと</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=个人整合">个人整合</a>, <a href="/search/photos?search_query=蘿莉">蘿莉</a>, <a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=彩頁">彩頁</a>, <a href="/search/photos?search_query=獸耳">獸耳</a>, <a href="/search/photos?search_query=獸娘">獸娘</a>, <a href="/search/photos?search_query=連褲襪">連褲襪</a>, <a href="/search/photos?search_query=和服">和服</a>, <a href="/search/photos?search_query=後宮">後宮</a>, <a href="/search/photos?search_query=其他校服">其他校服</a>, <a href="/search/photos?search_query=黑肉">黑肉</a>, <a href="/search/photos?search_query=貧乳">貧乳</a>, <a href="/search/photos?search_query=群交">群交</a>, <a href="/search/photos?search_query=過膝襪">過膝襪</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=彩页">彩页</a>, <a href="/search/photos?search_query=萝莉">萝莉</a> </div>
<div class="video-views pull-left">
2020-05-03
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/298871/少女性癖系列-虎助遙人-ガールズふぇてぃくしょん-dl版-個人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/298871_3x4.jpg?v=1653119061" title="少女性癖系列 [虎助遙人] ガールズふぇてぃくしょん[DL版]個人整合" alt="少女性癖系列 [虎助遙人] ガールズふぇてぃくしょん[DL版]個人整合" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_298871"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_298871" class="text-white">17K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">少女性癖系列 [虎助遙人] ガールズふぇてぃくしょん[DL版]個人整合</span>
<div class="title-truncate">
<a href="/search/photos?search_query=虎助遙人">虎助遙人</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=巨乳">巨乳</a>, <a href="/search/photos?search_query=連褲襪">連褲襪</a>, <a href="/search/photos?search_query=泳褲">泳褲</a>, <a href="/search/photos?search_query=短褲">短褲</a>, <a href="/search/photos?search_query=黑肉">黑肉</a>, <a href="/search/photos?search_query=足交">足交</a>, <a href="/search/photos?search_query=後宮">後宮</a>, <a href="/search/photos?search_query=群交">群交</a>, <a href="/search/photos?search_query=馬尾">馬尾</a>, <a href="/search/photos?search_query=辣妹">辣妹</a>, <a href="/search/photos?search_query=女高中生">女高中生</a>, <a href="/search/photos?search_query=眼鏡">眼鏡</a>, <a href="/search/photos?search_query=嗅氣味">嗅氣味</a>, <a href="/search/photos?search_query=過膝襪">過膝襪</a>, <a href="/search/photos?search_query=雙馬尾">雙馬尾</a>, <a href="/search/photos?search_query=假小子">假小子</a>, <a href="/search/photos?search_query=項圈">項圈</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=個人整合">個人整合</a> </div>
<div class="video-views pull-left">
2021-12-21
</div>
<div class="clearfix"></div>
</div>
<div class="p-b-15 p-l-5 p-r-5">
<div class="thumb-overlay-albums">
<a href="/album/273371/秘密之蕾-えーすけ-ひみつのつぼみ-特典-描下4pリーフレット付-個人整合">
<img src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://cdn-msp.orangeapp.cc/media/albums/273371_3x4.jpg?v=1641371296" title="秘密之蕾 [えーすけ] ひみつのつぼみ－特典 描下4Pリーフレット付 個人整合" alt="秘密之蕾 [えーすけ] ひみつのつぼみ－特典 描下4Pリーフレット付 個人整合" class="img-responsive owl-lazy " />
</a>
<div class="label-loveicon">
<a href="#" class id="love_likes_273371"><i class="glyphicon glyphicon-heart"></i><span id="albim_likes_273371" class="text-white">7K</span></a>
</div>
<div class="label-star">
<a data-toggle="modal" href="#login-modal">
<i style="color: #000000" class="far fa-bookmark"></i></a>
</div>
<div class="category-icon">
<div class="label-category" style>
</div>
</div>
</div>
<span class="video-title title-truncate m-t-5">秘密之蕾 [えーすけ] ひみつのつぼみ－特典 描下4Pリーフレット付 個人整合</span>
<div class="title-truncate">
<a href="/search/photos?search_query=えーすけ">えーすけ</a>
</div>
<div class="title-truncate">
<a href="/search/photos?search_query=劇情向">劇情向</a>, <a href="/search/photos?search_query=兄妹">兄妹</a>, <a href="/search/photos?search_query=亂倫">亂倫</a>, <a href="/search/photos?search_query=純愛">純愛</a>, <a href="/search/photos?search_query=巨乳">巨乳</a>, <a href="/search/photos?search_query=妖精">妖精</a>, <a href="/search/photos?search_query=黑肉">黑肉</a>, <a href="/search/photos?search_query=女高中生">女高中生</a>, <a href="/search/photos?search_query=校服">校服</a>, <a href="/search/photos?search_query=和服">和服</a>, <a href="/search/photos?search_query=透視">透視</a>, <a href="/search/photos?search_query=過膝襪">過膝襪</a>, <a href="/search/photos?search_query=懷孕">懷孕</a>, <a href="/search/photos?search_query=口交">口交</a>, <a href="/search/photos?search_query=女性支配">女性支配</a>, <a href="/search/photos?search_query=眼鏡">眼鏡</a>, <a href="/search/photos?search_query=手交">手交</a>, <a href="/search/photos?search_query=催眠">催眠</a>, <a href="/search/photos?search_query=校園">校園</a>, <a href="/search/photos?search_query=中文">中文</a>, <a href="/search/photos?search_query=個人整合">個人整合</a> </div>
<div class="video-views pull-left">
2021-08-19
</div>
<div class="clearfix"></div>
</div>
</ul>
</div>
<div class="row m-l-0 m-r-0 m-b-10">
<h4>相关文章</h4>
<div class="owl-carousel owl-article-block">
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/320">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="/media/blogs/320/5eef9be662e15f2c17a1d1baa527c573.jpg" referrerpolicy="no-referrer" title="扭曲到了一個極致-水葬銀貨のイストリア" alt="扭曲到了一個極致-水葬銀貨のイストリア">
</div>
<div class="title">扭曲到了一個極致-水葬銀貨のイストリア</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/cyjohn/blog">
cyjohn
</a>
</span>
</div>
<span>365
<a href="javascript:;" id="blog_like_320" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>12 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/76">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://pic.pimg.tw/h18comicorg/1589531864-2327999704_n.jpg" referrerpolicy="no-referrer" title="紳夜食堂64-口交的歷史-這個王后竟然口過一萬個士兵" alt="紳夜食堂64-口交的歷史-這個王后竟然口過一萬個士兵">
</div>
<div class="title">紳夜食堂64-口交的歷史-這個王后竟然口過一萬個士兵</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/18comicART/blog">
18comicART
</a>
</span>
</div>
<span>596
<a href="javascript:;" id="blog_like_76" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>56 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/19">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://pic.pimg.tw/h18comicorg/1568282041-1798946047_n.jpg" referrerpolicy="no-referrer" title="紳夜食堂11-不幹嫦娥幹月兔-以兔娘為主角的本子-每一個都騷出汁" alt="紳夜食堂11-不幹嫦娥幹月兔-以兔娘為主角的本子-每一個都騷出汁">
</div>
<div class="title">紳夜食堂11-不幹嫦娥幹月兔-以兔娘為主角的本子-每一個都騷出汁</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/18comicART/blog">
18comicART
</a>
</span>
</div>
<span>2088
<a href="javascript:;" id="blog_like_19" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>44 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/333">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="/static/resources/images/%E7%B4%B3%E5%A4%9C%E9%A3%9F%E5%A0%82/01.jpg" referrerpolicy="no-referrer" title="養成艦娘的奶子吧-篠塚釀二-艦っぱい-系列" alt="養成艦娘的奶子吧-篠塚釀二-艦っぱい-系列">
</div>
<div class="title">養成艦娘的奶子吧-篠塚釀二-艦っぱい-系列</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/takao2020/blog">
takao2020
</a>
</span>
</div>
<span>912
<a href="javascript:;" id="blog_like_333" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>30 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/298">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="/media/blogs/298/3404b69313f0a4b9238296f874561fb5.jpg" referrerpolicy="no-referrer" title="愛他就要吃了他-殺戀" alt="愛他就要吃了他-殺戀">
</div>
<div class="title">愛他就要吃了他-殺戀</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/cyjohn/blog">
cyjohn
</a>
</span>
</div>
<span>156
<a href="javascript:;" id="blog_like_298" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>7 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/29">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://pic.pimg.tw/h18comicorg/1570621756-2658947_n.jpg" referrerpolicy="no-referrer" title="紳夜食堂19-做愛-別離-以及惆悵-文青風格黃漫四種" alt="紳夜食堂19-做愛-別離-以及惆悵-文青風格黃漫四種">
</div>
<div class="title">紳夜食堂19-做愛-別離-以及惆悵-文青風格黃漫四種</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/18comicART/blog">
18comicART
</a>
</span>
</div>
<span>585
<a href="javascript:;" id="blog_like_29" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>18 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/331">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="/media/blogs/331/7b2774d54c56b215d0635e687d0e90c4.jpg" referrerpolicy="no-referrer" title="千萬別到了成為大人才在後悔-青空下的約定" alt="千萬別到了成為大人才在後悔-青空下的約定">
</div>
<div class="title">千萬別到了成為大人才在後悔-青空下的約定</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/cyjohn/blog">
cyjohn
</a>
</span>
</div>
<span>251
<a href="javascript:;" id="blog_like_331" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>11 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/34">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://pic.pimg.tw/h18comicorg/1572437388-285419334_n.jpg" referrerpolicy="no-referrer" title="紳夜食堂24-只有jk的雜交派對-上百美少女供你享樂的4場地下派對-你敢參加嗎" alt="紳夜食堂24-只有jk的雜交派對-上百美少女供你享樂的4場地下派對-你敢參加嗎">
</div>
<div class="title">紳夜食堂24-只有jk的雜交派對-上百美少女供你享樂的4場地下派對-你敢參加嗎</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/18comicART/blog">
18comicART
</a>
</span>
</div>
<span>1582
<a href="javascript:;" id="blog_like_34" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>71 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/367">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="/static/resources/files/00001.jpg" referrerpolicy="no-referrer" title="好色勇者大人-稀有好看的奇幻向韓漫" alt="好色勇者大人-稀有好看的奇幻向韓漫">
</div>
<div class="title">好色勇者大人-稀有好看的奇幻向韓漫</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/Kyousuke0529/blog">
Kyousuke0529
</a>
</span>
</div>
<span>522
<a href="javascript:;" id="blog_like_367" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>32 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/38">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://pic.pimg.tw/h18comicorg/1573551187-3266854349_n.png" referrerpolicy="no-referrer" title="紳夜食堂28-癡女逆推-讓這5個本子告訴你女人可以有多主動" alt="紳夜食堂28-癡女逆推-讓這5個本子告訴你女人可以有多主動">
</div>
<div class="title">紳夜食堂28-癡女逆推-讓這5個本子告訴你女人可以有多主動</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/18comicART/blog">
18comicART
</a>
</span>
</div>
<span>366
<a href="javascript:;" id="blog_like_38" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>20 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/41">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://pic.pimg.tw/h18comicorg/1574413210-3608821780_n.png" referrerpolicy="no-referrer" title="紳夜食堂31-巨乳還可以這樣玩-來看看巨乳控的腦洞大開的四個本子" alt="紳夜食堂31-巨乳還可以這樣玩-來看看巨乳控的腦洞大開的四個本子">
</div>
<div class="title">紳夜食堂31-巨乳還可以這樣玩-來看看巨乳控的腦洞大開的四個本子</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/18comicART/blog">
18comicART
</a>
</span>
</div>
<span>793
<a href="javascript:;" id="blog_like_41" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>39 <i class="far fa-comment-dots"></i></span>
</div>
</div>
<div class="p-b-15 p-l-5 p-r-5 gamelib_block">
<div class="p-b-5 m-b-5 gamelib_block_header">
<a href="/blogs/dinner" class="blog_a">
紳夜食堂
</a>
<span class="gamelib_btn_main">
</span>
</div>
<a href="/blog/42">
<div class="p-b-5 gamelib_block_img">
<img class="gamelib_img owl-lazy" src="https://cdn-msp.orangeapp.cc/media/albums/blank.jpg" data-src="https://pic.pimg.tw/h18comicorg/1574756349-3547601785_n.png" referrerpolicy="no-referrer" title="紳夜食堂32-少女都喜歡和大叔做愛-這4個本子告訴你老少配的魅力所在" alt="紳夜食堂32-少女都喜歡和大叔做愛-這4個本子告訴你老少配的魅力所在">
</div>
<div class="title">紳夜食堂32-少女都喜歡和大叔做愛-這4個本子告訴你老少配的魅力所在</div>
</a>
<div class="p-t-5 m-t-5 gamelib_block_footer">
<div class="p-b-5">
<span class="gamelib_name_block">
<a href="/user/18comicART/blog">
18comicART
</a>
</span>
</div>
<span>860
<a href="javascript:;" id="blog_like_42" class="blog-like">
<i class="far fa-heart" aria-hidden="true"></i>
</a>
</span>
<span>34 <i class="far fa-comment-dots"></i></span>
</div>
</div>
</div>
</div>
<div class="row center">
<a href="/album/260916" class="btn dropdown-toggle" style="border:1px solid #ff7a00;">隨便看看</a>
<a href="/albums?o=mr" class="btn btn-primary dropdown-toggle">觀看更多</a>
</div>
</div>
<div class="tab-pane fade m-b-15" id="comments">
<script>
        $(function () {
            $("[data-toggle='popover']").popover();
        });
    </script>
<style>
        .popover-title ,.popover-content {
           color:  #ececec
        }

    </style>
<div>
請遵守<a tabindex="0" class role="button" data-toggle="popover" data-placement="bottom" data-trigger="click" data-html="true" title="《社群規範》" data-content="1.禁止任何政治話題<br>
                        2.禁止留個人聯絡資訊<br>
                        3.無意義評論管理員有權扣分/永BAN※何謂 【無意義評論】？5555,666,每日刷分,文不對題評論<br>
                        4.評論是自由的，但請互相尊重<br>
                        5.催更可以但口氣請注意(*重要)<br><br>
                        違反規定者<br>
                        1. 帳號等級變成1 收藏不見<br>
                        2. 發廣告者帳號永封">《社群規範》</a>，劇透請開屏蔽。多次違規將永久刪除帳號。
</div>
<div class="form-group center">
<a class="btn btn-primary" data-toggle="modal" href="#login-modal" role="button">聊天前请先登入唷</a>
</div>
<div id="video_comments_419411" class="col-sm-12 col-xs-12" style="    padding: 0;">
显示
<span id="total_comments" class="text-white">23</span>
评论.
<div id="video_response" style="display: none;"></div>
<div id="comments_delimiter" style="display: none;"></div>
<div id="forum_response_3067803" style="display: none;"></div>
<div class="panel panel-default timeline-panel ">
<div class="panel-body timeline-panel-body">
<div class="timeline" data-cid="3067803">
<div class="timeline-left">
<a href="/user/huangseshenshi">
<img src="/media/users//3242048.jpg?v=1681662436" title="huangseshenshi's avatar" alt="huangseshenshi's avatar" class="timeline-avatar" data-userid="huangseshenshi" />
</a>
<div class="timeline-user-level">
Lv.7
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
黄色绅士 </span>
<div class="timeline-user-title">
．
</div>
<a href="javascript:;" class="comment-vote" data-vote="album" data-cid="3067803">
<i class="fa fa-minus-circle"></i>檢舉</a>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
</div>
<div class="timeline-date">
May 14, 2023
</div>
</div>
<div class="timeline-content">
后面不好看不建议收藏
</div>
<div class="timeline-info notice_more">
<span class="timeline-info-reply">回覆</span>
<div class="timeline-ft">
<a href="/album/419411/">
<small style>
JM419411</small>
</a>
</div>
</div>
</div>
<div class="reply-line"></div>
<div class="other-timelines">
</div>
</div>
</div>
</div>
<div id="forum_response_3030936" style="display: none;"></div>
<div class="panel panel-default timeline-panel ">
<div class="panel-body timeline-panel-body">
<div class="timeline" data-cid="3030936">
<div class="timeline-left">
<a href="/user/ZREsource">
<img src="/media/users//4023139.jpg?v=1688138101" title="ZREsource's avatar" alt="ZREsource's avatar" class="timeline-avatar" data-userid="ZREsource" />
</a>
<div class="timeline-user-level">
Lv.8
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
REsource </span>
<div class="timeline-user-title">
．
</div>
<a href="javascript:;" class="comment-vote" data-vote="album" data-cid="3030936">
<i class="fa fa-minus-circle"></i>檢舉</a>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/10.png" title="七夕" alt="七夕" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/Camp_5_3.png" title="志摩凜-3" alt="志摩凜-3" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A00506/Zombieland_1.png" title="巽幸太郎" alt="巽幸太郎" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2022.3/%E7%99%BD%E9%8A%80%E5%BE%A1%E8%A1%8C.png" title="白銀御行" alt="白銀御行" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2022.3/%E7%9F%B3%E4%B8%8A%E5%84%AA.png" title="石上優" alt="石上優" data-toggle="tooltip" data-placement="top">
</div>
<div class="timeline-date">
May 6, 2023
</div>
</div>
<div class="timeline-content">
哇 画风好合我胃口
</div>
<div class="timeline-info notice_more">
<span class="timeline-info-reply">回覆</span>
．
<span class="timeline-info-other-reply">
更多回覆
</span>
<div class="timeline-ft">
<a href="/album/419411/">
<small style>
JM419411</small>
</a>
</div>
</div>
</div>
<div class="reply-line"></div>
<div class="other-timelines">
<div class="timeline" data-cid="3030939" style="                                                        position: relative;
                                                        display: flex;
                                                        margin-top: 3rem;
                                                        ">
<div class="timeline-left">
<a href="/user/ZREsource">
<img src="/media/users/4023139.jpg?v=1688138101" title="ZREsource's avatar" alt="ZREsource's avatar" class="timeline-avatar" data-userid="ZREsource" />
</a>
<div class="timeline-user-level">
Lv.8
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
REsource </span>
<div class="timeline-user-title">
．
</div>
<a href="javascript:;" class="comment-vote" data-vote="album" data-cid="3030939">
<i class="fa fa-minus-circle"></i>檢舉</a>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/10.png" title="七夕" alt="七夕" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/Camp_5_3.png" title="志摩凜-3" alt="志摩凜-3" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A00506/Zombieland_1.png" title="巽幸太郎" alt="巽幸太郎" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2022.3/%E7%99%BD%E9%8A%80%E5%BE%A1%E8%A1%8C.png" title="白銀御行" alt="白銀御行" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2022.3/%E7%9F%B3%E4%B8%8A%E5%84%AA.png" title="石上優" alt="石上優" data-toggle="tooltip" data-placement="top">
</div>
<div class="timeline-date">
May 6, 2023
</div>
</div>
<div class="timeline-content">
果然老一点的本子更顶
</div>
</div>
</div>
<div class="timeline" data-cid="3030941" style="                                                        position: relative;
                                                        display: flex;
                                                        margin-top: 3rem;
                                                        ">
<div class="timeline-left">
<div class="reply-last-line"></div>
<a href="/user/ZREsource">
<img src="/media/users/4023139.jpg?v=1688138101" title="ZREsource's avatar" alt="ZREsource's avatar" class="timeline-avatar" data-userid="ZREsource" />
</a>
<div class="timeline-user-level">
Lv.8
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
REsource </span>
<div class="timeline-user-title">
．
</div>
<a href="javascript:;" class="comment-vote" data-vote="album" data-cid="3030941">
<i class="fa fa-minus-circle"></i>檢舉</a>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/10.png" title="七夕" alt="七夕" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/Camp_5_3.png" title="志摩凜-3" alt="志摩凜-3" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A00506/Zombieland_1.png" title="巽幸太郎" alt="巽幸太郎" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2022.3/%E7%99%BD%E9%8A%80%E5%BE%A1%E8%A1%8C.png" title="白銀御行" alt="白銀御行" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2022.3/%E7%9F%B3%E4%B8%8A%E5%84%AA.png" title="石上優" alt="石上優" data-toggle="tooltip" data-placement="top">
</div>
<div class="timeline-date">
May 6, 2023
</div>
</div>
<div class="timeline-content">
后面没前面好看感觉
</div>
</div>
</div>
</div>
</div>
</div>
</div>
<div id="forum_response_2935584" style="display: none;"></div>
<div class="panel panel-default timeline-panel ">
<div class="panel-body timeline-panel-body">
<div class="timeline" data-cid="2935584">
<div class="timeline-left">
<a href="/user/chaosjoker">
<img src="/media/users/nopic-Male.gif" title="chaosjoker's avatar" alt="chaosjoker's avatar" class="timeline-avatar" data-userid="chaosjoker" />
</a>
<div class="timeline-user-level">
Lv.7
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
chaosjoker </span>
<div class="timeline-user-title">
．
</div>
<a href="javascript:;" class="comment-vote" data-vote="album" data-cid="2935584">
<i class="fa fa-minus-circle"></i>檢舉</a>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A00506/kisaragi_5.png" title="水結之阿斯塔提" alt="水結之阿斯塔提" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2021.8%E5%8B%B3%E7%AB%A0/hame_9.png" title="安·薛利" alt="安·薛利" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A0/2021.8%E5%8B%B3%E7%AB%A0/jahysama_2.png" title="房東" alt="房東" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/files/medal/202209/lycoris3.png" title="中原瑞希" alt="中原瑞希" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/files/medal/202211/%E8%89%BE%E8%AB%BE%E6%A2%85.png" title="艾諾梅" alt="艾諾梅" data-toggle="tooltip" data-placement="top">
</div>
<div class="timeline-date">
Apr 18, 2023
</div>
</div>
<div class="disclose">
<p>此篇評論有劇透</p>
<a href="javascript:;" class>我還是想看</a>
</div>
<div class="timeline-content" style="display: none;">
雌神楽不是早就有5还是6了吗，这里果然搜不到
</div>
<div class="timeline-info notice_more">
<span class="timeline-info-reply">回覆</span>
<div class="timeline-ft">
<a href="/album/419411/">
<small style>
JM419411</small>
</a>
</div>
</div>
</div>
<div class="reply-line"></div>
<div class="other-timelines">
</div>
</div>
</div>
</div>
<div id="forum_response_2667536" style="display: none;"></div>
<div class="panel panel-default timeline-panel ">
<div class="panel-body timeline-panel-body">
<div class="timeline" data-cid="2667536">
<div class="timeline-left">
<a href="/user/jpped">
<img src="/media/users//2390098.jpg?v=1643778948" title="jpped's avatar" alt="jpped's avatar" class="timeline-avatar" data-userid="jpped" />
</a>
<div class="timeline-user-level">
Lv.5
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
月夜不語 </span>
<div class="timeline-user-title">
．
</div>
<a href="javascript:;" class="comment-vote" data-vote="album" data-cid="2667536">
<i class="fa fa-minus-circle"></i>檢舉</a>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
<img src="/static/resources/files/medal/202211/%E5%95%B5%E5%A5%87%E5%A1%94.png" title="啵奇塔" alt="啵奇塔" data-toggle="tooltip" data-placement="top">
<img src="/static/resources/images/%E5%8B%B3%E7%AB%A00506/Ijiranaide_3.png" title="長瀞" alt="長瀞" data-toggle="tooltip" data-placement="top">
</div>
<div class="timeline-date">
Feb 26, 2023
</div>
</div>
<div class="timeline-content">
畫風虎頭蛇尾前面很好後面垃圾
</div>
<div class="timeline-info notice_more">
<span class="timeline-info-reply">回覆</span>
<div class="timeline-ft">
<a href="/album/419411/">
<small style>
JM419411</small>
</a>
</div>
</div>
</div>
<div class="reply-line"></div>
<div class="other-timelines">
</div>
</div>
</div>
</div>
<div id="forum_response_2598742" style="display: none;"></div>
<div class="panel panel-default timeline-panel ">
<div class="panel-body timeline-panel-body">
<div class="timeline" data-cid="2598742">
<div class="timeline-left">
<a href="/user/wudidaxia666">
<img src="/media/users//951439.jpg?v=1675735977" title="wudidaxia666's avatar" alt="wudidaxia666's avatar" class="timeline-avatar" data-userid="wudidaxia666" />
</a>
<div class="timeline-user-level">
Lv.5
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
无敌大虾 </span>
<div class="timeline-user-title">
．
</div>
<a href="javascript:;" class="comment-vote" data-vote="album" data-cid="2598742">
<i class="fa fa-minus-circle"></i>檢舉</a>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
</div>
<div class="timeline-date">
Feb 12, 2023
</div>
</div>
<div class="disclose">
<p>此篇評論有劇透</p>
<a href="javascript:;" class>我還是想看</a>
</div>
<div class="timeline-content" style="display: none;">
不知道为什么要去哪里
</div>
<div class="timeline-info notice_more">
<span class="timeline-info-reply">回覆</span>
<div class="timeline-ft">
<a href="/album/419411/">
<small style>
JM419411</small>
</a>
</div>
</div>
</div>
<div class="reply-line"></div>
<div class="other-timelines">
</div>
</div>
</div>
</div>
<div style="text-align:center">
<a id="p_album_comments_419411_2" href="javascript:void(0)" class="btn btn-primary navbar-btn m-l-15 m-r-15" data-series>
查看更多
</a>
</div>
</div>
<div class="clearfix"></div>
</div>
<div class="comment-reply-div" style="display: none">
<div id="comment-reply-block" class="m-t-15">
<div class="row">
<div class="pull-left">
<a>
<img class="medium-avatar" src="/media/users/nopic-.gif" />
</a>
</div>
<div class="comment" style="margin-left:65px ">
<div class="comment-info user-container-name">
<a></a>
</div>
<div class="comment-body overflow-hidden">
<textarea name="video_comment_reply" id="video_comment_reply" rows="2" class="form-control video_comment_reply" style="border:none !important;border-radius: 0px;"></textarea>
<a class="video_comment_reply_add">Post</a>
<div id="reply_message" class="text-danger m-t-5" style="display: none;">請輸入您的留言!</div>
<div id="reply_message_err" class="text-danger m-t-5" style="display: none;">等级需大于 3 !</div>
</div>
</div>
<div class="clearfix"></div>
</div>
</div>
</div>
</div>
<div class="row">
<div class="col-lg-3 col-md-3 col-sm-3 col-xs-6">
<div class="e8c78e-4_b" data-group="album_related1">
<script async type="application/javascript" src="https://a.magsrv.com/ad-provider.js"></script>
<ins class="adsbyexoclick" data-zoneid="2967008" data-sub="77"></ins>
<script>(AdProvider = window.AdProvider || []).push({"serve": {}});</script>
<script type="text/javascript">
                    (function () {
                        function randStr(e,t){for(var n="",r=t||"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",o=0;o<e;o++)n+=r.charAt(Math.floor(Math.random()*r.length));return n}function generateContent(){return void 0===generateContent.val&&(generateContent.val=" \ndocument.dispatchEvent("+randStr(4*Math.random()+3)+");"),generateContent.val}try{Object.defineProperty(document.currentScript,"innerHTML",{get:generateContent}),Object.defineProperty(document.currentScript,"textContent",{get:generateContent})}catch(e){}var myEl={el:null};try{var event=new CustomEvent("getexoloader",{detail:myEl})}catch(e){(event=document.createEvent("CustomEvent")).initCustomEvent("getexoloader",!1,!1,myEl)}window.document.dispatchEvent(event);var ExoLoader=myEl.el;
                        ExoLoader.addZone({"idzone":"2967008","sub":"77"});                        
                    })();
                    </script>
</div>
</div>
<div class="col-lg-3 col-md-3 col-sm-3 col-xs-6">
<div class="e8c78e-4_b" data-group="album_related2">
<script async type="application/javascript" src="https://a.magsrv.com/ad-provider.js"></script>
<ins class="adsbyexoclick" data-zoneid="2967010" data-sub="78"></ins>
<script>(AdProvider = window.AdProvider || []).push({"serve": {}});</script>
<script type="text/javascript">
                    (function () {
                        function randStr(e,t){for(var n="",r=t||"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",o=0;o<e;o++)n+=r.charAt(Math.floor(Math.random()*r.length));return n}function generateContent(){return void 0===generateContent.val&&(generateContent.val=" \ndocument.dispatchEvent("+randStr(4*Math.random()+3)+");"),generateContent.val}try{Object.defineProperty(document.currentScript,"innerHTML",{get:generateContent}),Object.defineProperty(document.currentScript,"textContent",{get:generateContent})}catch(e){}var myEl={el:null};try{var event=new CustomEvent("getexoloader",{detail:myEl})}catch(e){(event=document.createEvent("CustomEvent")).initCustomEvent("getexoloader",!1,!1,myEl)}window.document.dispatchEvent(event);var ExoLoader=myEl.el;
                        ExoLoader.addZone({"idzone":"2967010","sub":"78"});                        
                    })();
                    </script>
</div>
</div>
<div class="col-lg-3 col-md-3 col-sm-3 col-xs-6">
<div class="e8c78e-4_b" data-group="album_related3">
<script type="text/javascript" data-cfasync="false" async src="https://poweredby.jads.co/js/jads.js"></script>
<ins id="662591" data-width="300" data-height="262"></ins>
<script type="text/javascript" data-cfasync="false" async>(adsbyjuicy = window.adsbyjuicy || []).push({'adzone':662591});</script>
</div>
</div>
<div class="col-lg-3 col-md-3 col-sm-3 col-xs-6">
<div class="e8c78e-4_b" data-group="album_related4">
<script async type="application/javascript" src="https://a.magsrv.com/ad-provider.js"></script>
<ins class="adsbyexoclick" data-zoneid="3648533" data-sub="80"></ins>
<script>(AdProvider = window.AdProvider || []).push({"serve": {}});</script>
<script type="text/javascript">
                    (function () {
                        function randStr(e,t){for(var n="",r=t||"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",o=0;o<e;o++)n+=r.charAt(Math.floor(Math.random()*r.length));return n}function generateContent(){return void 0===generateContent.val&&(generateContent.val=" \ndocument.dispatchEvent("+randStr(4*Math.random()+3)+");"),generateContent.val}try{Object.defineProperty(document.currentScript,"innerHTML",{get:generateContent}),Object.defineProperty(document.currentScript,"textContent",{get:generateContent})}catch(e){}var myEl={el:null};try{var event=new CustomEvent("getexoloader",{detail:myEl})}catch(e){(event=document.createEvent("CustomEvent")).initCustomEvent("getexoloader",!1,!1,myEl)}window.document.dispatchEvent(event);var ExoLoader=myEl.el;
                        ExoLoader.addZone({"idzone":"3648533","sub":"80"});                        
                    })();
                    </script>
</div>
</div>
</div>
</div>
<script>
$(document).ready(function() {
    var checkVideoInterval; // 定時器

    function sendAjaxRequest(aid , modal , ajaxUrl) {
        $.ajax({
            url: ajaxUrl, 
            type: 'POST',
            data: {aid: aid},
            success: function(htmlContent) {
                modal.find('.modal-body').html(htmlContent);
                if (htmlContent.indexOf('<video') !== -1) {
                    clearInterval(checkVideoInterval);// 清除計時器
                    //modal.find('.modal-footer').html('<a type="button" id="downloadButton" class="btn btn-primary" href="/media/JmShortVideo/' + aid + '.mp4" download>影片下載</a><button type="button" class="btn btn-secondary" data-dismiss="modal">取消</button>')
                    modal.find('.modal-footer').html('<button type="button" class="btn btn-secondary" data-dismiss="modal">取消</button>')
                }
            }
        });
    }

    $('#jmComicMbVideo, #jmComicVideo').on('show.bs.modal', function (e) {
        var path = window.location.pathname;
        var parts = path.split('/');
        var aid = parts[2];
        var modal = $(this);

        var JMvideo = "1";
        var videoAjax = "1";
        if (videoAjax == '0') {
            var ajaxUrl = base_url + "/ajax/album_chkvdo";
        } else {
            var ajaxUrl = 'https://cdn-vdo.jmcomic2.bet/ajax/album_chkvdo' ;
        }

        if (JMvideo != '1') {
            sendAjaxRequest(aid, modal, ajaxUrl);
            checkVideoInterval = setInterval(function() {
                sendAjaxRequest(aid, modal, ajaxUrl);
            }, 30000);
        }
    });

    $('#jmComicMbVideo, #jmComicVideo').on('hide.bs.modal', function (e) {
        clearInterval(checkVideoInterval);
    });
});


</script>
<script>

    if (document.getElementById('book-name').scrollHeight <= 17) {
        $('.book-name-btn').hide();
    }

    var scramble_id = 220980;

    var webpMachine = new webpHero.WebpMachine({
        //webpSupport: false,
    })

    //圖片放大
    var imgs = $('.img_zoom_img,.show_zoom');
    var lens = imgs.length;
    for (var i = 0; i < lens; i++) {
        imgs[i].onclick = function (event) {
            event = event || window.event;
            var target = document.elementFromPoint(event.clientX, event.clientY);
            $('#popup').empty();
            if (target.tagName == 'CANVAS') {
                showBig($(this).find('img')[0], true, $(this).attr('id'));
            } else {
                showBig(target.src, false, $(this).attr('id'));
            }
        }
    }

    $('.switch_btn').click(function () {
        var list = $(this).next('.episode').contents('ul');
        var listItems = list.children('a');
        list.append(listItems.get().reverse());
    })

    if (localStorage.getItem('series_' + $('#album_id').val()) != null) {
        var series_value = localStorage.getItem('series_' + $('#album_id').val());
        var series_split = series_value.split('-');
        var album_id = series_split[0];
        var sort = series_split[1];

        if (album_id != '' && sort != '') {
            $('.reading').text('續看 第' + sort + '話');
            var url = $('.reading').attr('href');
            url = url.replace($('#album_id').val(), album_id);
            $('.reading').attr('href', url);
        }
    }

    $('#intro,#episode,#download').click(function () {
        $(this).addClass('active').siblings().removeClass('active');
        var name = $(this).attr('id');
        var openBlock = "#" + name + "-block";
        $(openBlock).removeClass('hide').siblings('.nav-tab-content').addClass('hide');
        if (name == 'episode' || name == 'download') {
            $(openBlock).siblings('.img_zoom').addClass('hide');
        } else {
            $(openBlock).siblings('.img_zoom').removeClass('hide');
        }
    });

    $(document).ready(function () {
        $(".img_zoom_img .lazy_img").lazyload_old({
            threshold: 0,
            src: 'data-original',
            load: function () {
                webpMachine.polyfillImage($(this)[0]);
                scramble_image($(this)[0], aid, scramble_id, false, false);
            }
        });
    });

    function episode() {
        $('#episode').click();
    }

    $('.book-name-btn').on('click', function (e) {
        $(e.target).parents('.book-name-collep').toggleClass('col-xs-11 col-xs-12');
        $('.book-name-btn >i').toggleClass('fa-arrow-circle-down fa-arrow-circle-up');
        $('.book-name-btn').toggleClass('down');
        $('.book-name').toggleClass('line-clamp-unset');
    });

    owl = $('.cover-view');
    owl.on('translated.owl.carousel', function (event) {
        const owlItem = document.querySelectorAll('.cover-view .owl-item');
        owlItem.forEach(function (e) {
            if (e.children[0].id != "") {
                if (!$(e).hasClass('active')) {
                    if ($(e).find('canvas').length > 0) {
                        $(e).find('canvas').remove();
                    }
                } else {
                    if ($(e).find('img').attr('src').indexOf('blank') == -1) {
                        scramble_image($(e).find('img')[0], aid, scramble_id, false, false);
                    }
                }
            }
        });
    }).on('loaded.owl.lazy', function (event) {
        scramble_image(event.element[0], aid, scramble_id, false, false);
    });

    owl.owlCarousel({
        items: 1,
        nav: true,
        dots: false,
        loop: true,
        lazyLoad: true,
        center: true,
    });

    var coverHeight = $('.cover-view').width();
    $('.cover-view .owl-item').height(coverHeight);

</script>
<link rel="stylesheet" href="/templates/frontend/airav/emoji/emojionearea.css">
<style>
    .timeline-panel-body{
        overflow: initial;
    }
    </style>
<input type="hidden" name="code" value>
<div class="panel panel-default timeline-panel " id="forum_template" style="display: none">
<div class="panel-body timeline-panel-body" style="position: relative">
<div class="timeline">
<div class="timeline-left">
<a href="/user/">
<img src class="timeline-avatar">
</a>
<div class="timeline-user-level">
</div>
</div>
<div class="timeline-right">
<div class="timeline-header">
<span class="timeline-username">
</span>
<div class="timeline-user-title">
</div>
</div>
<div class="timeline-second-header">
<div class="timeline-badge">
</div>
<span class="timeline-date">
</span>
</div>
<div class="timeline-content">
</div>
<div class="timeline-info notice_more">
<a>
<i class="fas fa-heart timeline-like"></i>
<span class="timeline-info-liked">
</span>
</a>
<span class="timeline-info-dot">．</span>
<span class="timeline-info-reply">回覆</span>
<div class="timeline-ft">
<a href>
<small style>
</small>
</a>
</div>
</div>
<div class="other-timelines">
</div>
</div>
</div>
</div>
</div>
<div id="commend_template" class="commend response" style="min-height: 30px;margin-top: 15px;display: none">
<div style="flex:9;display: inline-block;">
<textarea class="emojionearea form-control m-b-20" rows="3" placeholder="分享你的想法吧"></textarea>
</div>
<div class="commend-btn pull-right m-l-10">
<button type="button" name="reply_submit" value class="btn btn-sm btn-primary ">發言</button>
</div>
</div>
<script>
    var relative_tpl = '/templates/frontend/airav';
    var forum_type = '';
</script>
<script>
        $(document).on('click', '.timeline-info-reply', function () {
            if ($(this).siblings('.commend').length === 0) {
                var obj = $('#commend_template').clone();
                obj.attr('id', '');
                obj.find('button[name="reply_submit"]').val($(this).closest('.timeline').attr('data-cid'));
                obj.css('display', 'flex');
                //$(obj).insertAfter($(this));
                $(this).parent().append(obj);
                if($(this).closest('.timeline').find('.other-timelines').css('display')==='none'){
                    $(this).closest('.timeline-panel').find('.reply-line').hide();
                }

            } else {
                $(this).siblings('.commend').remove();
                $(this).closest('.timeline-panel').find('.reply-line').show();
            }

        });
        $(document).on('click', 'button[name="reply_submit"]', function () {
            var params = {
                comment_id: $(this).val(),
                aid: $(this).attr('data-aid'),
                comment: $(this).closest('.commend').find('textarea').val(),
                forum_subject: 1,
                originator: $('input[name="code"]').val(),
                is_reply:1

            };
            var request_url = '/ajax/album_comment';
            switch (forum_type) {
                case 'blog':
                    request_url='/ajax/blog_comment';
                    params.blog_id = $('#blog_id').val();
                    break;
                case 'game':
                    request_url='/ajax/game_comment';
                    if($('#game_id').length>0) {
                        params.game_id = $('#game_id').val();
                    }else{
                        params.game_id = 0 ;
                    }
                    break;
                default:
                    if($('#album_id').length>0) {
                        params.video_id = $('#album_id').val();
                    }else{
                        params.video_id = 0 ;
                    }
                    break;
            }

            var element = $(this);
            $.post(request_url, params, function (response) {
                var data = JSON.parse(response);
                $('input[name="code"]').val(data.originator);
                if (data.status == 0) {
                    $("textarea[id='video_comment']").val('');
                    element.closest('.commend.response').after(data.msg);


                } else {
                    $(".no_comments").hide();
                    $("textarea[id='video_comment']").val('');

                    element.closest('.timeline-right').siblings('.other-timelines').show().prepend(data.code);
                    //$(bDIV).after(cDIV);
                    element.closest('.commend.response').after(data.msg);


                    if (data.status == '1') {
                        $('#end_num').html(parseInt($('#end_num').html(), 10) + 1);
                        $('#total_comments').html(parseInt($('#total_comments').html(), 10) + 1);
                        $('#total_video_comments').html(parseInt($('#total_video_comments').html(), 10) + 1);
                    }
                }


                setTimeout(function () {
                    $('.alert-dismissable').remove();
                }, 5000);
            })

        });

        $(document).on('click', '.timeline-info-other-reply', function () {
            if ($(this).closest('.timeline').find('.other-timelines').is(':visible')) {
                $(this).closest('.timeline').find('.other-timelines').slideUp();
            } else {
                $(this).closest('.timeline').find('.other-timelines').slideDown();
            }
        });

        $(document).on('click', ".disclose > a", function (event) {
            event.preventDefault();
            $(this).closest('.disclose').hide();
            $(this).closest('.disclose').siblings('.timeline-content,.timeline-info').slideDown();
        });
    </script>
<div class="bot-per visible-xs visible-sm">
<div class="bot-per-times">
<i class="fas fa-times"></i>
</div>
<div class="bot-per-context e8c78e-4_b" data-width="728" data-height="90" data-group="sticky1">
<p><a rel="nofollow noopener" href="https://r.morexpbb.net/game/xr/?attributionId=126" target="_blank"><img alt src="/static/resources/images/%E5%BB%A3%E5%91%8A202308/XR_random_728_90_v46.gif" /></a></p>
</div>
<div class="bot-per-back">
點我一下 , 免費去廣告
</div>
</div>
<div class="footer-container">
<div class="footer">
<div class="container">
<div>
<div class="pull-left">
<span>版权所有 &#169; 2008-2022</span> 禁漫天堂
</div>
<div class="pull-right">
<a href="/cdn-cgi/l/email-protection#14636363252c777b797d77547379757d783a777b79">广告洽询</a>
</div>
<div class="clearfix"></div>
</div>
</div>
</div>
</div>
</div>
<style>
    
    @media (max-width: 767px) {
        .bot-per-times {
            z-index: 10;
            width: 25px;
            height: 25px;
            border-radius: 5px;
            text-align: center;
            position: fixed;
            bottom: 70px;
            right: 10px;
            background: rgba(255, 255, 255, 0.5);
        }

        .bot-per-times i {
            font-size: 25px;
            color: #ff0000;
        }

        .bot-per-context {
            padding: 0px !important;
            bottom: 61px;
            z-index: 3;
            position: fixed;
            width: 100%;
        }

        .bot-per-back {
            bottom: 61px;
            z-index: 3;
            position: fixed;
            height: 7.1%;
            width: 100%;
            background-color: #fff;
            text-align: center;
            line-height: 50px;
            display: none;
            pointer-events: none;
        }
    }

    .menu-circle {
        width: 7px;
        height: 7px;
        right: 5%;
        background: #ff7a00;
        border-radius: 50%;
        position: absolute;
    }

    
</style>
<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script>
    
    $(function () {
        if (location.href.indexOf('/photo/') > 0) {
            $('body').addClass("photo_cus_body");
            $('#Comic_Top_Nav').addClass('hidden-xs hidden-sm');
        }

        /* 偵測廣告點擊數 */
        setInterval(function () {
            if (document.activeElement instanceof HTMLIFrameElement) {
                var up = new URL(document.activeElement.src);
                var iframe_url = up.href;
                if (iframe_url.indexOf('juicyads') > -1 || iframe_url.indexOf('exosrv') > -1) {
                    $.get('/ajax/ad_check', function (data) {
                        data = JSON.parse(data);
                        if (data.status === 1) {
                            toastr['success'](data.msg)
                        }

                    });
                }
            }
            window.focus();
        }, 300);
        $(document).on('click', '.e8c78e-4_b a', function () {
            $.get('/ajax/ad_check', function (data) {
                data = JSON.parse(data);
                if (data.status === 1) {
                    toastr['success'](data.msg)
                }

            });
        });

        //bot-per
        if (isMobile) {
            setInterval(function () {
                $(".bot-per-back").fadeIn(1000);
                $(".bot-per-back").fadeOut(1000);
            }, 8000);

            var detectelement = function (url) {
                $('.bot-per').remove();
                td_ga_tracker.doga_event('ad_fixed_bottom', url, 'Click', 1);
            }

            $('.bot-per-times').click(function () {
                $('.bot-per').attr("style", "display: none !important");
            });

            $('.bot-per-context').click(function () {
                var url = $(this).find("a").prop("href");
                detectelement(url);
            });

            var px = 0;
            $(window).scroll(function () {
                //if (parseInt($.cookie("bot-per")) < 2) {
                if ($(window).scrollTop() > px + 10000) {
                    px = $(window).scrollTop();
                    $('.bot-per').attr("style", "display: block !important");
                    $(window).resize();
                }
                //}
            });

            if (location.href.indexOf('/photo/') > 0) {
                $('.bot-per-context,.bot-per-back').css('bottom', '53px');
            }

            var IframeOnClick = {
                istrack: null,
                resolution: 200,
                iframes: [],
                interval: null,

                Iframe: function () {
                    this.element = arguments[0];
                    this.cb = arguments[1];
                    this.hasTracked = false;
                    this.src = arguments[0].src;
                },

                track: function (element, cb) {
                    this.iframes.push(new this.Iframe(element, cb));
                    if (!this.interval) {
                        var _this = this;
                        this.interval = setInterval(function () {
                            _this.checkClick();
                        }, this.resolution);
                    }
                },

                checkClick: function () {
                    if (document.activeElement) {
                        var activeElement = document.activeElement;
                        var tagName = activeElement.tagName;
                        for (var i in this.iframes) {
                            if (activeElement === this.iframes[i].element) {
                                if (this.iframes[i].hasTracked == false) {
                                    this.iframes[i].cb.apply(window, []);
                                    this.iframes[i].hasTracked = true;
                                    if (this.iframes[i].src.indexOf("adzone=819141") > 0 || this.iframes[i].src.indexOf("idzone=3597795") > 0) {
                                        detectelement(this.iframes[i].src);
                                    }
                                }
                            }
                        }
                    }
                }
            };

            window.onload = function () {
                var ifs = document.getElementsByTagName("iframe");
                for (var i = 0; i < ifs.length; i++) {
                    IframeOnClick.track(ifs[i], function () {
                        IframeOnClick.istrack = true;
                    });
                }

            }

        }
    });

    $(document).ready(function () {
        if ($('.e8c78e-4_b iframe').length <= 0) {
            $('#chk_cover').prev('.m-b-10.center').css('display', 'block');
        }
        var syncLoginUrl = '';
        if (syncLoginUrl !== '')
            $('<iframe id="syncLogin" style="width:0px;height: 0px;" src="' + syncLoginUrl + '" ></iframe>').insertAfter('#wrapper');
    });

    $(window).load(function () {
        if ($.cookie('shuntflag') == undefined) {
            $('.lazy_img').each(function () {
                var img_url = $(this).attr('src');
                if (!CheckImgExists(img_url) && $.cookie('cover') == '1' && $.cookie('shuntflag') != '1') {
                    $('#shunt-modal').modal('show');
                    $.cookie('shuntflag', '1', {path: '/'});
                }
            });
        }
    });

    function CheckImgExists(imgurl) {
        var ImgObj = new Image();
        ImgObj.src = imgurl;
        if (ImgObj.fileSize > 0 || (ImgObj.width > 0 && ImgObj.height > 0)) {
            return true;
        } else {
            return false;
        }
    }

    
</script>
<script src="/templates/frontend/airav/js/bootstrap.min.js"></script>
<script>
        

    (function () {
        function randStr(e, t) {
            for (var n = "", r = t || "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", o = 0; o < e; o++) n += r.charAt(Math.floor(Math.random() * r.length));
            return n
        }

        function generateContent() {
            return void 0 === generateContent.val && (generateContent.val = " \ndocument.dispatchEvent(" + randStr(4 * Math.random() + 3) + ");"), generateContent.val
        }

        try {
            Object.defineProperty(document.currentScript, "innerHTML", {get: generateContent}), Object.defineProperty(document.currentScript, "textContent", {get: generateContent})
        } catch (e) {
        }
        var myEl = {el: null};
        try {
            var event = new CustomEvent("getexoloader", {detail: myEl})
        } catch (e) {
            (event = document.createEvent("CustomEvent")).initCustomEvent("getexoloader", !1, !1, myEl)
        }
        window.document.dispatchEvent(event);
        var ExoLoader = myEl.el;
        ExoLoader.serve({"script_url": "/templates/frontend/airav/js/4_46b.js"});

        // window.nerverblock_callback = function () {
        //     $("#Comic_Top_Nav").append('<div class="container center" style="color: #ffffff;"><i class="fa fa-exclamation-triangle"></i>如果观看出现问题，请关闭Adblock功能，<a rel="nofollow" target="_blank" href="https://help.getadblock.com/support/tickets/new">并联系Adblock服务来解决问题</a></div>');
        //     $("#wrapper").css('top',$('.container.center').height() + 'px');
        // }
        //
        // setInterval(function () {
        //     $("ins").each(function () {
        //         var e = $(this);
        //         if(e.nextAll('div').contents().length > 0){
        //             e.prev().remove();
        //             //e.remove();
        //         }
        //     })
        // }, 5000);

    })();

    if (navigator.userAgent.match(/IEMobile\/10\.0/)) {
        var msViewportStyle = document.createElement('style')
        msViewportStyle.appendChild(
            document.createTextNode(
                '@-ms-viewport{width:auto!important}'
            )
        )
        document.querySelector('head').appendChild(msViewportStyle)
    }

    $('[data-toggle="tooltip"]').tooltip();

    
</script>
<script>(function(){var js = "window['__CF$cv$params']={r:'7f989102c80e17d2',t:'MTY5MjUxMjgwNC4zOTYwMDA='};_cpo=document.createElement('script');_cpo.nonce='',_cpo.src='/cdn-cgi/challenge-platform/scripts/invisible.js',document.getElementsByTagName('head')[0].appendChild(_cpo);";var _0xh = document.createElement('iframe');_0xh.height = 1;_0xh.width = 1;_0xh.style.position = 'absolute';_0xh.style.top = 0;_0xh.style.left = 0;_0xh.style.border = 'none';_0xh.style.visibility = 'hidden';document.body.appendChild(_0xh);function handler() {var _0xi = _0xh.contentDocument || _0xh.contentWindow.document;if (_0xi) {var _0xj = _0xi.createElement('script');_0xj.innerHTML = js;_0xi.getElementsByTagName('head')[0].appendChild(_0xj);}}if (document.readyState !== 'loading') {handler();} else if (window.addEventListener) {document.addEventListener('DOMContentLoaded', handler);} else {var prev = document.onreadystatechange || function () {};document.onreadystatechange = function (e) {prev(e);if (document.readyState !== 'loading') {document.onreadystatechange = prev;handler();}};}})();</script></body>
</html>


`

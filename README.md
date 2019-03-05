# yt

youtube video link grabber

### working endpoint
[https://yt.penzur.xyz](https://yt.penzur.xyz)


### basic example

0. `$ brew install jq` (optional)


0. `curl https://yt.penzur.xyz/?url=<youtube video url> -s | jq`

	should respond something like the following:

```
[
  {
    "itag": "22",
    "quality": "hd720",
    "type": "video/mp4; codecs=\"avc1.64001F, mp4a.40.2\"",
    "url": "https://r2---sn-t0a7ln7d.googlevideo.com/videoplayback?mm=31%2C26&mn=sn-t0a7ln7d%2Csn-vgqskn7l&id=o-AJGKGQCO1i_YKBRLfVklw8KTe04gt68uu9LL8cpUatOH&c=WEB&mime=video%2Fmp4&ip=3.81.28.99&mt=1551778598&lmt=1538059953207409&dur=4580.019&itag=22&pl=22&source=youtube&ipbits=0&mv=u&ei=g0V-XLPHF6mnhwatpLr4Dg&fvip=2&ms=au%2Conr&sparams=dur%2Cei%2Cid%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire&requiressl=yes&expire=1551800803&ratebypass=yes&signature=360E6428A2314AB36E31620892019935D571D874.1F7B53829954F3CE3C73AC92D3EC97ADF385DA52&key=yt6"
  },
  {
    "itag": "43",
    "quality": "medium",
    "type": "video/webm; codecs=\"vp8.0, vorbis\"",
    "url": "https://r2---sn-t0a7ln7d.googlevideo.com/videoplayback?ipbits=0&mime=video%2Fwebm&mt=1551778598&itag=43&ratebypass=yes&fvip=2&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire&expire=1551800803&key=yt6&gir=yes&mm=31%2C26&mn=sn-t0a7ln7d%2Csn-vgqskn7l&id=o-AJGKGQCO1i_YKBRLfVklw8KTe04gt68uu9LL8cpUatOH&c=WEB&ip=3.81.28.99&lmt=1526614687338088&dur=0.000&pl=22&source=youtube&mv=u&ei=g0V-XLPHF6mnhwatpLr4Dg&ms=au%2Conr&clen=496051976&requiressl=yes&signature=798A28D186253EBA56661353084A1524BE807D65.CD950A89CA0C069145A2BE5A0D498346351953BC"
  },
  {
    "itag": "18",
    "quality": "medium",
    "type": "video/mp4; codecs=\"avc1.42001E, mp4a.40.2\"",
    "url": "https://r2---sn-t0a7ln7d.googlevideo.com/videoplayback?ipbits=0&mime=video%2Fmp4&mt=1551778598&itag=18&ratebypass=yes&fvip=2&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire&expire=1551800803&key=yt6&gir=yes&mm=31%2C26&mn=sn-t0a7ln7d%2Csn-vgqskn7l&id=o-AJGKGQCO1i_YKBRLfVklw8KTe04gt68uu9LL8cpUatOH&c=WEB&ip=3.81.28.99&lmt=1538053156955335&dur=4580.019&pl=22&source=youtube&mv=u&ei=g0V-XLPHF6mnhwatpLr4Dg&ms=au%2Conr&clen=396015066&requiressl=yes&signature=527F9CE214CA81864F66E659F6B13AC6671FC387.84CB18CF2768F0CAA9704D2B9CC467A1AEA9ED8D"
  }
]
```
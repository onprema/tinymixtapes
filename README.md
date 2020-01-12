# tinymixtapes best releases of the decade

This code grabs the youtube URLs from [tinymixtape best releases of the decade](https://www.tinymixtapes.com/features/2010s-favorite-100-music-releases-decade).

If you want to download them, I'd suggest installing [youtube-dl](https://ytdl-org.github.io/youtube-dl/index.html).

### Download the music

```shell
# First, write YouTube URLs to ./tinymixtapes-urls.txt
go run main.go

# Download them using youtube-dl
for url in $(cat tinymixtapes-urls.txt); do youtube-dl --extract-audio --audio-format "mp3" $url; done
```

Enjoy :)

# info-hash-encoder
torrent info_hash re-encoder for tracker http requests

Converts a sha1 hash string from a torrent info hash, such as 

`29cb773e13f2e7a5ab4271d7772762ccae0ea057`

back into the format needed to make torrent tracker request, like this;

`%29%CBw%3E%13%F2%E7%A5%ABBq%D7w%27b%CC%AE%0E%A0W`

so you can make a tracker query like this;

```bash
$ curl "https://torrent.ubuntu.com/announce?info_hash=%29%CBw%3E%13%F2%E7%A5%ABBq%D7w%27b%CC%AE%0E%A0W&peer_id=%2D%41%5A%35%37%35%30%2D%54%70%6B%58%74%74%5A%4C%66%70%53%48&port=6881&uploaded=0&downloaded=0&left=1502576640&event=started&compact=1" --output -
```

## Usage

```
$ go run main.go 29cb773e13f2e7a5ab4271d7772762ccae0ea057
%29%CBw%3E%13%F2%E7%A5%ABBq%D7w%27b%CC%AE%0E%A0W
```

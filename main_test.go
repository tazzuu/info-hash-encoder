package main

import (
	"github.com/google/go-cmp/cmp"
	// "log"
	"fmt"
	"testing"
)

func TestConvertHashes(t *testing.T) {
	// https://torrent.ubuntu.com/tracker_index

	tests := map[string]struct {
		hash    string
		want     string
	}{
		"first": {
			hash: "653b154e138fd8f1522851f120315c17ecf53f05",
			want: "e%3B%15N%13%8F%D8%F1R%28Q%F1%201%5C%17%EC%F5%3F%05",
		},
		"second": {
			hash: "29cb773e13f2e7a5ab4271d7772762ccae0ea057",
			want: "%29%CBw%3E%13%F2%E7%A5%ABBq%D7w%27b%CC%AE%0E%A0W",
		},
		"third": {
			hash: "443c7602b4fde83d1154d6d9da48808418b181b6",
			want: "D%3Cv%02%B4%FD%E8%3D%11T%D6%D9%DAH%80%84%18%B1%81%B6",
		},
		"fourth": {
			hash: "eed74c895ea640bf75d42c19b71eaf64ab72677e",
			want: "%EE%D7L%89%5E%A6%40%BFu%D4%2C%19%B7%1E%AFd%ABrg%7E",
		},
		"fifth": {
			hash: "5729cad8f0c744c965c2d3539dc3971f96de38ce",
			want: "W%29%CA%D8%F0%C7D%C9e%C2%D3S%9D%C3%97%1F%96%DE8%CE",
		},
		"sixth": {
			hash: "563a485379b6ea4e25b9ada1fef1e15493a22989",
			want: "V%3AHSy%B6%EAN%25%B9%AD%A1%FE%F1%E1T%93%A2%29%89",
		},
		"seventh": {
			hash: "7f1f147998639094012392bb645e52957cb57680",
			want: "%7F%1F%14y%98c%90%94%01%23%92%BBd%5ER%95%7C%B5v%80",
		},
		"eigth": {
			hash: "a7838b75c42b612da3b6cc99beed4ecb2d04cff2",
			want: "%A7%83%8Bu%C4%2Ba-%A3%B6%CC%99%BE%EDN%CB-%04%CF%F2",
		},
		"ninth": {
			hash: "e7088644335bd66118c01d3e4ece64f2630fb1c6",
			want: "%E7%08%86D3%5B%D6a%18%C0%1D%3EN%CEd%F2c%0F%B1%C6",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ConvertHash(tc.hash)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				fmt.Printf(">>> hash: %v , got: %v , want: %v\n", tc.hash, got, tc.want)
				t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

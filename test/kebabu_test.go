package kebab

import (
    "testing"
    "../kebab"
    . "github.com/r7kamura/gospel" // パッケージ名にグローバルネームを指定可能
)

func TestKebab (t *testing.T) {
    text := kebab.Grill()

    // デフォルトパッケージで搭載されているテスト
    if text != "ju-ju!" {
        t.Error("ジュージュー音がしてない")
    }

    // Rspecみたいに書きたいよね
    Describe(t, "Grill", func(){
        It("should say ju-ju!", func(){
            Expect(text).To(Equal, "ju-ju!")
        })
    })
}

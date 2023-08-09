# go-note

### 1. fatal error: concurrent map read and map write
若多個 gorutine 同時對一個 map 操作有可能觸發

而且是 fatal error 無法被 recover 捕捉

改用 sync.Map 可以解決此問題

ref: https://ithelp.ithome.com.tw/articles/10218003
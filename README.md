# Echo сервер для голосового помошника Алиса

Подробная докуменация о работе [Яндекс.Диалогов](https://tech.yandex.ru/dialogs/alice/doc)

Сервер повторяет в точности то сообщение, которое вы ему пересылаете. То есть это голосовая реализация echo сервера.

Использование:
```bash
$ export HOST=https://powerful-plateau-82075.herokuapp.com/
$ curl --request post --data '{"request":{"command":"где ближайшее отделение"},"session":{"new":true,"message_id":4,"session_id":"2eac4854-fce721f3-b845abba-20d60","skill_id":"3ad36498-f5rd-4079-a14b-788652932056","user_id":"AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC"},"version":"1.0","end_session":true}' $HOST
{"response":{"text":"где ближайшее отделение"},"session":{"new":true,"message_id":4,"session_id":"2eac4854-fce721f3-b845abba-20d60","skill_id":"3ad36498-f5rd-4079-a14b-788652932056","user_id":"AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC"},"version":"1.0","end_session":true}```
```


Установка:
```bash
mkdir -p $GOPATH/src/github.com/r00takaspin
cd -
git clone https://github.com/r00takaspin/pingpong
cd pingpong
dep ensure -v
export PORT=8080; go build && go run pingpong.go
```

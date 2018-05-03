# KFU [Go](https://golang.org/) курсы

> Notice: this course is for Russian university and there is no english version.  
> In general though, all participants are welcome.

Данный репозиторий содержит вспомогательный материал, на который
время от времени делаются отсылки в лекционных презентациях.

Сами презентации лежат здесь же, в каталоге [_docs/presentations](/_docs/presentations).

Информацию о практических заданиях можно найти в каталоге [_docs/tasks](/_docs/tasks/tasks.org).

[Таблицу результатов](#Трофеи) можно посмотреть в самом низу этого документа.

> ВАЖНО: документацию по стандартным пакетам можно смотреть на [golang.org/pkg](https://golang.org/pkg/).
> Её также можно поднять локально, с помощью `godoc -http=8080`, после этого документация
> будет доступна по адресу `http://localhost:8080/`.

## План занятий

1. [Вводное занятие](/_docs/presentations/1.pdf).
2. [Первый коммит в go-review.googlesource](/_docs/presentations/2.pdf).
3. **Workshop** + [синтез двух лекций](/_docs/presentations/2+1.pdf) для новых студентов.
4. **Workshop** + [введение в профилирование и оптимизацию](/_docs/presentations/3.pdf).
5. Очное занятие отменено.

**Workshop** - это интерактивный формат, когда люди могут работать над своими
патчами в Go параллельно основному лекционному треку.

Для всех новых студентов есть [доработанные слайды](/_docs/presentations/2+1.pdf) для быстрого вовлечения.

Официальный Contribution Guide:
- [Новая версия](https://tip.golang.org/doc/contribute.html) (рекомендуется)
- [Текущая версия](https://golang.org/doc/contribute.html)

**Обновление**: курс проходил с марта по май, затем был приостановлен из-за [code freeze](https://groups.google.com/forum/#!topic/golang-dev/a1YnheC_ku4). За это время участники успели выслать **17** патчей.

## Полезные ссылки

[GolangShow](http://golangshow.com) - русскоязычный подкаст о Go. Крутые ведущие, интересные гости.

[golang-ru Slack](http://slack.golang-ru.com) - русскоязычное Go сообщество.
Там можно задавать вопросы, обсуждать Go, библиотеки под него и прочее.
Для вопросов лучше всего подходит канал `#school` (при формулировке вопроса можно
опираться на [How To Ask Questions The Smart Way](http://www.catb.org/esr/faqs/smart-questions.html)). 
Всем участникам следует соблюдать [кодекс норм поведения](https://golang.org/conduct).

Google проводил [Contribution Workshop](https://blog.golang.org/contributor-workshop).
Это мероприятие во многом является идейным вдохновителем для этих курсов (но по размаху
их догнать невозможно). Кроме интересного описания самого события, в Go блоге, где
опубликован этот доклад, есть ещё много других качественных постов.

Есть хорошая презентация на английском: [How to contribute to Go](https://docs.google.com/presentation/d/1ap2fycBSgoo-jCswhK9lqgCIFroE1pYpsXC1ffYBCq4/edit#slide=id.p).  
Только оттуда стоит выбросить всё, что привязано к давно прошедшему gophercon.

Видео доклада на тему "How to contribute to Go" на русском:  
https://www.youtube.com/watch?v=0a8u74Ul-hM&feature=youtu.be.

[Go and fix me](https://www.goandfix.me/) - сервис, который отображает проблемы, найдённые линтерами в репозитории Go.
Очень удобен для нахождения почвы для первых CL. (Временами недоступен, возможно проблемы с хостингом.)

## Ссылки, полезные для обучения

Большинство ссылок легко найти в гугле по запросу "golang learning resources".  
Самое главное правило - всегда искать по слову `golang`, а не `go`.  
Ниже наиболее стоящие результаты с описаниями.

[Resources for new Go programmers](https://dave.cheney.net/resources-for-new-go-programmers) - статья [Dave Cheney](https://dave.cheney.net/about), одного из ведущих разработчиков Go.

[golang/go/wiki/Learn](https://github.com/golang/go/wiki/Learn) - много учебного материала.

[Go tour](https://tour.golang.org) - интерактивный Go тур. Доступен на русском, но лучше
проходить в оригинале. Этот вариант изучения классический, но может быть местами сложноват.

[Go videos](https://github.com/hH39797J/golang-videos-ru) - собрание видеозаписей докладов про Go.

[Go by example](https://gobyexample.com/) - примеры кода на Go с комментариями и описаниями.

[Go webdev examples](https://gowebexamples.com/) - аналог Go by example, но с уклоном в веб разработку.

## Трофеи

Таблица изменений, которые были посланы на go-review.

| Ссылка на изменение | Автор | Статус |
|---------------------|-------|--------|
| https://golang.org/cl/105356 | bontequero / Делюс | merged |
| https://golang.org/cl/105495 | bontequero / Делюс | merged |
| https://golang.org/cl/105736 | bontequero / Делюс | merged |
| https://golang.org/cl/107235 | bontequero / Делюс | merged |
| https://golang.org/cl/108275 | tengufromsky / Никита | merged |
| https://golang.org/cl/105556 | tengufromsky / Никита | merged |
| https://golang.org/cl/107018 | tengufromsky / Никита | merged |
| https://golang.org/cl/107056 | tengufromsky / Никита | merged |
| https://golang.org/cl/105415 | fexolm / Артём | merged |
| https://golang.org/cl/105395 | fexolm / Артём | merged |
| https://golang.org/cl/105355 | fexolm / Артём | merged |
| https://golang.org/cl/105375 | weeellz / Равиль | merged |
| https://golang.org/cl/107135 | weeellz / Равиль | merged |
| https://golang.org/cl/107115 | weeellz / Равиль | merged |
| https://golang.org/cl/105416 | ludweeg / Эмиль | merged |
| https://golang.org/cl/108659 | ludweeg / Эмиль | merged |
| https://golang.org/cl/108815 | ludweeg / Эмиль | merged |

Итого: 17.

----

kfu-go-2018  
kfu go 2018

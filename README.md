# URL_Collection
# Навигация по меню приложения

```
a - Добавить новый URL в список
l - Просмотреть все сохраненные URL
r - Удалить выбранный URL
```


## Задание
### Создание списка url

Напишите программу, которая будет хранить ваши url. На основании созданного шаблона допишите код, который позволяет
добавлять новые ссылки, удалять и выводить список.

Для решения задачи используйте структуры. Обязательные поля структуры должны быть дата добавления, имя ссылки,
теги для ссылки через запятую и сам url.

Например
```Go
type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}
```

Для хранение даты вы можете использовать пакет time. Чтобы создать дату можно использовать следующую функцию time.Now()
Чтобы вывести дату как строку используйте функцию Format, например
```Go
t := time.Now() // создание текущей даты
t.Format(time.DateTime) // как строка
```

Наш проект использует библиотеку github.com/eiannone/keyboard для отслеживания нажатия клавиш, более подробно про
подключение пакетов мы поговорим на отдельной лекции. В шаблоне вам нужно дописать код действий под соотвествующими
нажатиями клавиш.
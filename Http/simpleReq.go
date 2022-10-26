package main

import (
	"fmt"
	"io"
	"net/http"
)

// Определить структуру
// Opts - это необязательный параметр, если не установлен, выберите значение по умолчанию
type KPClient struct {
	client http.Client
	opts   options
}

// Сохранить параметры, которые могут быть незаметно распознаны, эта презентация только одна
// Если есть несколько слов, продолжайте добавлять
type options struct {
	header http.Header
}

// Настройка параметров по умолчанию, конечно, другие параметры могут быть установлены
// здесь, только один пустой запрос головы
// Конечно, вы можете добавить больше вещей до середины, в зависимости от ситуации
func defaultOptions() options {
	return options{
		header: http.Header{},
	}
}

// Предоставить интерфейс, который создает клиента
func NewKPClient(client http.Client, opts ...Option) *KPClient {
	d := KPClient{
		client: client,
		// Настройте параметры по умолчанию, когда вы объявляете объект
		opts:   defaultOptions(),
	}
	// Вот фокус
	// опция Все использовать интерфейс для приема, в интерфейсе реализован метод приложения.
	// Применить метод получает * Параметры для изменения значения в клиенте клиента.
	// Если параметр по умолчанию недостаточно пользователя, пользователь может использовать этот способ расширения
	// Этот параметр определения изменяет значение параметра по умолчанию
	for _, opt := range opts {
		opt.apply(&d.opts)
	}
	return &d
}

// опция интерфейс
type Option interface {
	apply(*options)
}

// funcoption в основном для создания вариантов более просто
// funcoption реализует функцию опции интерфейс применять функцию
// newfuncoption может вернуть объект интерфейса опции
type funcOption struct {
	f func(*options)
}

func (fo funcOption) apply(o *options) {
	fo.f(o)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// Установите опцию с заголовком запроса
// Если вы используете эту опцию, заголовок запроса в определенном типе установлен ниже.
// не значение по умолчанию
func withHeaderUserAgent() Option {
	var header = http.Header{}
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36")

	return newFuncOption(func(o *options) {
		o.header = header
	})
}

// простой метод получения запроса
func (d *KPClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	req.Header = d.opts.header
	return d.client.Do(req)
}

func main() {
	opts := []Option{
		withHeaderUserAgent(),
	}

	client := http.Client{}
	d := NewKPClient(client, opts...)
	res, err := d.Get("https://www.kinopoisk.ru/lists/movies/top250/")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	fmt.Println(string(body))
}

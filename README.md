# News API Library

This project is a Go library for interacting with the [News API](https://newsapi.org/).

## Features

- Fetch top headlines
- Search for articles
- Filter news by source, language, and country

## Installation

To install the library, use `go get`:

```sh
go get github.com/erevos-13/newsapigo
```

## Usage

Here's a basic example of how to use the library:

```go
package main

import (
    "fmt"
    "log"
    "github.com/erevos-13/newsapigo"
)

func main() {
    client := newsapi.NewClient("YOUR_API_KEY")

    headlines, err := client.GetTopHeadlines(newsapi.TopHeadlinesOptions{
        Country: "us",
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, article := range headlines.Articles {
        fmt.Println(article.Title)
    }
}
```

## Configuration

You need to obtain an API key from [News API](https://newsapi.org/register) and set it in your environment variables or pass it directly to the client.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
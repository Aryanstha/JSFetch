<br/>
<p align="center">
  <a href="https://github.com/aryanstha/JS Fetch">
    <img src="https://github.com/Aryanstha/JSFetch/blob/main/logos.png?raw=true" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">JS Fetch</h3>

  <p align="center">
    Fetch data with ease using JS Fetch - a simple yet powerful tool for making HTTP requests in JavaScript!
    <br/>
    <br/>
    <a href="https://github.com/aryanstha/JSFetch">View Demo</a>
    .
    <a href="https://github.com/aryanstha/JSFetch/issues">Request Feature</a>
  </p>
</p>

![Downloads](https://img.shields.io/github/downloads/aryanstha/JSFetch/total) ![Contributors](https://img.shields.io/github/contributors/aryanstha/JSFetch?color=dark-green) ![Forks](https://img.shields.io/github/forks/aryanstha/JSFetch?style=social) ![Stargazers](https://img.shields.io/github/stars/aryanstha/JSFetch?style=social) ![Issues](https://img.shields.io/github/issues/aryanstha/JSFetch) ![License](https://img.shields.io/github/license/aryanstha/JSFetch) 

## About The Project

![Screen Shot](https://github.com/Aryanstha/JSFetch/blob/main/demo.PNG?raw=true)

This is a simple Go program that collects JavaScript files from a target domain. It sends an HTTP GET request to the target domain, parses the response HTML with goquery, and collects the URLs of all JavaScript files that are hosted on the target domain.

## Built With

This is a Go program that collects JavaScript files from a target domain. The program takes two optional command line arguments:

The target domain (-d flag)
The output file name (-o flag)

The program sends an HTTP GET request to the target domain URL and parses the response HTML with the goquery package. It then extracts the URLs of all the script tags in the HTML that ends with ".js" and contain the target domain.

The collected URLs are filtered for uniqueness using the unique() function. If the output file name is provided, the program writes the collected URLs to that file, otherwise, it prints them to the console.

## Getting Started

o use this program, you need to have Go installed on your system. Once you have Go installed, you can install this program by running:

### Prerequisites


```bash
go get github.com/PuerkitoBio/goquery
```

### Installation

1. Clone the repo

```sh
git clone https://github.com/aryanstha/JSFetch
```

2. Install NPM packages

```sh
go mod tidy
```

3. Build file```sh go build```
4.
```sh
 jsfetch -d <target-domain> [-o <output-file>]
```

## License

Distributed under the MIT License. See [LICENSE](https://github.com/aryanstha/JSFetch/blob/main/LICENSE.md) for more information.

## Authors

* **Aryan Shrestha** - *Developer* - [Aryan Shrestha](https://github.com/aryanstha) - **

## Acknowledgements

* [ShaanCoding](https://github.com/ShaanCoding/)
* [Othneil Drew](https://github.com/othneildrew/Best-README-Template)
* [ImgShields](https://shields.io/)

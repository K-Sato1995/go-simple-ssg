## What this is

Yet another SSG in Go. It's not intended to replace any existing SSGs. I created it as a learning exercise and really enjoyed the process.🤗

Heavily inspired by this blog post([Writing a Static Blog Generator in Go](https://www.zupzup.org/static-blog-generator-go/index.html)) by [@zupzup](https://github.com/zupzup).

- [Site I created using this](https://go-simple-ssg.vercel.app/)

In case you are wondering how to use this, check out the Initialization process below.

## Initialization

Install the cli.

```
go install github.com/K-Sato1995/go-simple-ssg/simple-ssg-cli@latest
```

run 

```
simple-ssg-cli init
```

would create a new project that looks somethi ng like this.

```
./
├── contents/ (Stores markdown files)
├── templates/ (Stores template html files and style files)
├── generated/ (Stores generated files)
├── go.mod
├── go.sum
└── main.go
```

`cd your project` and just run 

```
go run main.go
```

You should be able to see the site running on http://localhost:3001 

## Todo

- [x] CLI To run generating static contents based on the config file
  - put all the code in one func
  - create cmd to copy the template
- [x] write tests
- [x] Custom Error pages(404)
- [x] HMR
- [ ] SEO stuffda

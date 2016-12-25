# This project powers my personal website

This site is powered by revel!

# Dependencies and versions

Since go is lame and wants every package to be in its own git repo it is difficult to manage dependencies.
My solution is to clone (with 'go get <package>') and then branch off the latest commit that I plan to depend on.
This way I can pull the 'leading edge' down into master and create a new branch when I want to roll a dependency forward.

## List of current dependency versions

### Revel modules
revel   (github.com/revel/revel)    v0.13.1
cmd     (github.com/revel/cmd)      v0.13.1
config  (github.com/revel/config)   v0.13.0
modules (github.com/revel/modules)  v0.13.0

### Revel dependencies
gocolorize (github.com/agtorre/gocolorize) v1.0.0

compress (github.com/klauspost/compress) v1.2.0
cpuid    (github.com/klauspost/cpuid)    v1.0.0
crc32    (github.com/klauspost/crc32)    v1.1.0

pathtree (github.com/robfig/pathtree)    v1.0.0

### Database
go-sqlite3 (github.com/mattn/go-sqlite3) v1.2.0
gorp     (github.com/go-gorp/gorp)       v2.0.0 (TODO: Probably don't use this)


# Welcome to Revel

## Getting Started

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Start the web server:

    revel run myapp

   Run with <tt>--help</tt> for options.

### Go to http://localhost:9000/ and you'll see:

"It works"

### Description of Contents

The default directory structure of a generated Revel application:

    myapp               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files

app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.


messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.

### Follow the guidelines to start developing your application:

* The README file created within your application.
* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

## Contributing
We encourage you to contribute to Revel! Please check out the [Contributing to Revel
guide](https://github.com/revel/revel/blob/master/CONTRIBUTING.md) for guidelines about how
to proceed. [Join us](https://groups.google.com/forum/#!forum/revel-framework)!

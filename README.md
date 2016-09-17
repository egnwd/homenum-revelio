# Homenum Revelio

An idea about displaying who is currently at home, thanks to [Dan Slocombe](https://github.com/danslocombe).

*This project is very much experimental and a bit of fun.*

## Dependencies

  - `nmap`
    - Install from [nmap.org/download.html](https://nmap.org/download.html)
  - [Go](https://golang.org)
    - Install from [golang.org/doc/install](https://golang.org/doc/install)
  - [npm](https://github.com/npm/npm)
    - Install via `curl -L https://www.npmjs.com/install.sh | sh`
    - [Typescript](https://www.typescriptlang.org)
      - Install via `npm install -g typescript`
    - [Typings](https://github.com/typings/typings)
      - Install via `npm install -g typings`
    - [Webpack](http://webpack.github.io)
      - Install via `npm install -g webpack`
    - Or install all 3 with `npm install -g typescript typings webpack`

## Building

  - Run `./build.sh`

## Running

### Standalone script
 - You will need a yaml file mapping MAC addresses to Names in the form:

```yaml
05:E5:87:F4:4C:D4: Name
59:65:11:03:56:68: Other Name
38:F6:AC:4D:ED:09: Full House
```
 - Then run using `./bin/homenum_revelio` from within the directory or `homenum_revelio` by adding the `bin` folder to your `PATH`

### Server
  - Run using `./server/server` and open [localhost:8080](http://localhost:8080)
  - You can also pass in flags:
    - `-p port` to set a specific port for the server, default is `8080`
    - `-static path/to/static/files` to set the static file server, default is `./client`
    - `-r path/to/residents.yaml` to set where the map of MAC addresses to names is, default is `./residents.yaml`

**Note:** `-static` and `-r` are relative to the root directory

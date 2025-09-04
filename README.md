<br />
<div align="center">
  <h3 align="center">Load Balancer QA</h3>

  <p align="center">
    A fast and easy-to-configure load balancer
    <br />
    <br />
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#configuration">Configuration</a></li>
    <li><a href="#limitations">Limitations</a></li>
    <li><a href="#benchmark">Benchmark</a></li>
    <li><a href="#todo">TODO</a></li>
    <li><a href="#contributors">Contributors</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

## About The Project
This project is designed to provide a fast and easy-to-configure load balancer in Go language. It currently includes **round-robin**, **weighted round-robin**, **least-connection**, **least-response-time**, **ip-hash** and **random** algorithms, but we have more to add to our [TODO](#todo) list.

The project is developed using the [fasthttp](https://github.com/valyala/fasthttp) library, which ensures high performance. Its purpose is to distribute the load evenly among multiple servers by routing incoming requests.

The project aims to simplify the configuration process for users while performing the essential functions of load balancers. Therefore, it offers several configuration options that can be adjusted to meet user needs.

This project is particularly suitable for large-scale applications and websites. It can be used for any application that requires a load balancer, thanks to its high performance, ease of configuration, and support for different algorithms.

## Features
- Fast and easy-to-configure load balancer.
- Supports round-robin, weighted round-robin, least-connection, least-response-time, IP hash, and random algorithms.
- Supports TLS and HTTP/2 for the frontend server.
- Uses the fasthttp library for high performance and scalability.
- Offers multiple configuration options to suit user needs.
- Can handle large-scale applications and websites.
- Includes a built-in monitoring system that displays real-time information on the system's CPU usage, RAM usage, number of Goroutines, and open connections.
- Prometheus support for monitoring. (`http://monitoring-host:monitoring-port/metrics` can be used to get prometheus metrics)
- Provides information on each server's average response time, total request count, and last time used.
- Lightweight and efficient implementation for minimal resource usage.

## Installation

#### Downloading the Release
The latest release can be downloaded from the [releases](https://github.com/PrathamNabira/Load-balancer-qa/releases) page. Choose the suitable binary for your system, download and extract the archive, and then move the binary to a directory in your system's `$PATH` variable (e.g. `/usr/local/bin`).

#### Building from Source
Alternatively, you can build the project from source by cloning this repository to your local machine and running the following commands:

```bash
git clone https://github.com/PrathamNabira/Load-balancer-qa.git &&
cd Load-balancer-qa &&
go build -o loadbalancer &&
./loadbalancer
````

#### Using go install

You can also install it using the `go install` command:

```bash
go install github.com/PrathamNabira/Load-balancer-qa@latest
```

This will install the binary to your system's `$GOPATH/bin` directory. Make sure this directory is included in your system's `$PATH` variable to make the program accessible from anywhere.

That's it! You're now ready to use it in your project.

## Usage

You need a `config.yaml` file to run the load balancer. You can give this file with the `--config` flag, or by default it will try to use a `config.yaml` file in the same directory.

[Example config files](https://github.com/PrathamNabira/Load-balancer-qa/tree/main/examples)

> \:warning: Please use absolute path for "config.yaml" while using "--config" flag

## Configuration

| Name                                       | Description                                                                                                                          | Type          | Default Value |
| ------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------- | ------------- |
| type                                       | Load balancing algorithm                                                                                                             | string        | round-robin   |
| port                                       | Server port                                                                                                                          | int           | 8000          |
| host                                       | Server host                                                                                                                          | string        | localhost     |
| health\_checker\_time                      | Time interval to perform health check for backends                                                                                   | time.Duration | 30s           |
| backends                                   | List of backends with their configurations                                                                                           | array         |               |
| backends.url                               | Backend URL                                                                                                                          | string        |               |
| backends.health\_check\_path               | Health check path for backends                                                                                                       | string        | /             |
| backends.weight                            | Only mandatory for w-round-robin algorithm                                                                                           | int           |               |
| backends.max\_conn                         | Maximum number of connections which may be established to host listed in Addr                                                        | int           | 512           |
| backends.max\_conn\_timeout                | Maximum duration for waiting for a free connection                                                                                   | time.Duration | 30s           |
| backends.max\_conn\_duration               | Keep-alive connections are closed after this duration                                                                                | time.Duration | 10s           |
| backends.max\_idle\_conn\_duration         | Idle keep-alive connections are closed after this duration                                                                           | time.Duration | 10s           |
| backends.max\_idemponent\_call\_attempts   | Maximum number of attempts for idempotent calls                                                                                      | int           | 5             |
| monitoring                                 | Monitoring server configurations                                                                                                     | object        |               |
| monitoring.port                            | Monitoring server port                                                                                                               | int           | 8001          |
| monitoring.host                            | Monitoring server host                                                                                                               | string        | localhost     |
| custom\_headers                            | Custom headers will be set on request sent to backend                                                                                | object        |               |
| custom\_headers.header-name                | Valid values are `$remote_addr`, `$time`, `$incremental`, `$uuid`, The Header name can be whatever you want as long as it's a string | string        |               |
| server                                     | Server configurations                                                                                                                | object        |               |
| server.http\_version                       | Http version for frontend server, http1 and http2 is supported (http1 mean HTTP/1.1)                                                 | string        | http1         |
| server.cert\_file                          | TLS cert file                                                                                                                        | string        |               |
| server.key\_file                           | TLS key file                                                                                                                         | string        |               |
| server.max\_idle\_worker\_duration         | MaxIdleWorkerDuration is the maximum idle time of a single worker in the underlying worker pool of the Server                        | time.Duration | 10s           |
| server.tcp\_keepalive\_period              | Period between tcp keep-alive messages. TCP keep-alive period is determined by operation system by default                           | time.Duration |               |
| server.concurrency                         | The maximum number of concurrent connections the server may serve                                                                    | int           | 262144        |
| server.read\_timeout                       | ReadTimeout is the amount of time allowed to read the full request including body                                                    | time.Duration | unlimited     |
| server.write\_timeout                      | WriteTimeout is the maximum duration before timing out writes of the response                                                        | time.Duration | unlimited     |
| server.idle\_timeout                       | IdleTimeout is the maximum amount of time to wait for the next request when keep-alive is enabled                                    | time.Duration | unlimited     |
| server.disable\_keepalive                  | The server will close all the incoming connections after sending the first response to client if this option is set to true          | bool          | false         |
| server.disable\_header\_names\_normalizing | Header names are passed as-is without normalization if this option is set true                                                       | bool          | false         |

Please see [example config files](https://github.com/PrathamNabira/Load-balancer-qa/tree/main/examples)

## Limitations

While this project has several features and benefits, it also has some limitations to be aware of:

* It currently operates at layer 7, meaning it is specifically designed for HTTP(S) load balancing. It does not support other protocols, such as TCP or UDP.
* It does not support HTTP/3, which may be important for some applications.
* It does not support HTTPS for backend servers. HTTPS is only available for the frontend server.

Please keep these limitations in mind when considering whether this load balancer is the right choice for your project.

## Benchmark

Please see the [benchmark folder](https://github.com/PrathamNabira/Load-balancer-qa/tree/main/benchmark) for detailed explanations.

## TODO

While this project has several features, there are also some areas for improvement that are planned for future releases:

* [ ] Add support for other protocols, such as TCP or UDP.
* [x] Add TLS support for frontend.
* [x] Support HTTP/2 in frontend server.
* [ ] Add more load balancing algorithms, such as,

  * [x] least connection
  * [x] least-response-time
  * [ ] sticky round-robin
* [ ] Improve performance and scalability for high-traffic applications.
* [x] Expand monitoring capabilities to provide more detailed metrics and analytics.

By addressing these issues and adding new features, the goal is to make this an even more versatile and powerful tool for managing traffic in modern web applications.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.

The MIT License is a permissive open-source software license that allows users to modify and redistribute the code, as long as the original license and copyright notice are included. This means that you are free to use this project for any purpose, including commercial projects, without having to pay any licensing fees or royalties. However, it is provided "as is" and without warranty of any kind, so use it at your own risk.

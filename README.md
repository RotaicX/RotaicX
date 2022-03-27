# How to build this project
You only need to execute
```
cd src/LoginServer
make
```
Then the build will be stored in src/LoginServer/build
### Attention:
* You need to pre-configure the golang development environment
* If your build fails due to network problems please run `make configproxy`

# The third-party dependencies used in this project
* [PostgreSQL driver for database/sql](https://www.github.com/lib/pq)
* [Toml parser](https://www.github.com/BurntSushi/toml)
* [Http Framework](https://www.github.com/gin-gonic/gin)
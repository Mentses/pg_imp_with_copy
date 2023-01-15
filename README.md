# pg_imp_with_copy
# Readme
适用于PostgreSQL数据库导入文本数据

# 环境变量
> 程序自动读取环境变量`DATABASE_URL`用于登录PG数据库

## Windows
```bash
set DATABASE_URL=postgres://username:password@localhost:5432/postgres
```
## Linux
```bash
set DATABASE_URL=postgres://username:password@localhost:5432/postgres
```

# 编译
```bash
go build -ldflags "-s -w" -o pg_imp.exe main.go
or
go build -ldflags "-s -w" -o pg_imp.bin main.go
```


# Example
```bash
go run main.go "/data/dim_test.csv" "COPY public.tbl_prov_city(prov_code,city_code) FROM STDIN DELIMITER E'\t' CSV ENCODING 'UTF8' QUOTE E'\x01' ESCAPE E'\x05'";

go run main.go "/data/dim_test.csv" "COPY public.tbl_prov_city(prov_code,city_code) FROM STDIN DELIMITER E'\t' CSV HEADER ENCODING 'UTF8' QUOTE E'\x01' ESCAPE E'\x05'";

go run main.go "/data/dim.csv" "COPY public.tbl_prov_city FROM STDIN DELIMITER E'\t' CSV ENCODING 'UTF8' QUOTE E'\x01' ESCAPE E'\x05'";

go run main.go "/data/dim.csv" "COPY public.tbl_prov_city FROM STDIN DELIMITER E'\t' CSV HEADER ENCODING 'UTF8' QUOTE E'\x01' ESCAPE E'\x05'";
```

# 可执行文件压缩
```bash
upx -9 pg_imp.bin
```

# Go vs Dotnet

## Introduction

cs-demo: **Use** .net8 **Version:** 8.0.8 **Framework:** .NetCore

go-demo: **Use** golang **Version** 1.23.0 **Framework:** Gin Gonic

## Test Performance

Run Api

### Start service โดยไม่ได้ทำการเรียกใช้
**cs-demo**
![image](./dotnet-demo/document/cs-demo-p1.png)

**go-demo**
![image](./go-demo/document/go-demo-p1.png)

### Run 15 ครั้ง delay แต่ละครั้ง 500ms
**cs-demo**
- /cpu/1000000000
![image](./dotnet-demo/document/cs-demo-p2.png)

**go-demo**
- /cpu/1000000000
![image](./go-demo/document/go-demo-p2.png)

### Run 15 ครั้ง delay แต่ละครั้ง 500ms
**cs-demo**
- /memory/10000000
![image](./dotnet-demo/document/cs-demo-p3.png)

**go-demo**
- /memory/10000000
![image](./go-demo/document/go-demo-p3.png)

## Reference
- https://github.com/antonputra/tutorials


#### 1. TiDB user and privileges

> create user tkill@'%' identified by 'xx';

> GRANT PROCESS,SUPER ON *.* TO 'tkill'@'%';





#### 2. Golang install

> yum install golang

> go env -w GOPROXY=https://goproxy.cn,direct



#### 3. Usage

> ./tkill -l 10.0.1.1 -d dbname execute >/dev/null 2>&1 

#### 4. SQL

> "select id, db, time, info from information_schema.processlist where db = ? and time >= ? " + 
> "and info is not null and (info like '%select%' or info like '%SELECT%') order by time desc"

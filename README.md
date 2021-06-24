# gobusterdns
lite version for gobuster. Only subdomain brute. 内网轻量化子域名爆破工具

适合指定dns跑内网子域名

## 与原版的修改
- 精简功能，仅支持子域名扫描

- 可导入domain list文件扫描

- 内置精简字典

- 简化输入参数

- 优化-o输出文件


## 用法
.\gobuster.exe -l domainlist.txt -o out.txt -r dns-ip

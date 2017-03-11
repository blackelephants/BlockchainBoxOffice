# 合约调用说明

合约hash：5402240f31c2955594fff7cd9e2d6aec9ee08aa720d0c55ca40a5ab9fc59e12cae02117973de6f7094e0d471d5eee6c26aa9779f50c441007cac81842ebaaa70

## 流程（接口）
### 登记影院
```
{
     "jsonrpc": "2.0",
     "method": "invoke",
     "params": {
         "type": 1,
         "chaincodeID": {
             "name": "合约hash"
         },
         "ctorMsg": {
             "function": "registerCinema",
             "args": [
                 "影院名称",
                 "所属院线"
             ]
         },
         "secureContext": "user_type1_0"
     },
     "id": 0
}
```

### 登记售票平台
```
{
     "jsonrpc": "2.0",
     "method": "invoke",
     "params": {
         "type": 1,
         "chaincodeID": {
             "name": "合约hash"
         },
         "ctorMsg": {
             "function": "registerTicketPlatform",
             "args": [
                 "平台名称"
             ]
         },
         "secureContext": "user_type1_0"
     },
     "id": 0
}
```

### 登记影厅信息
```
{
	参数：影院名称 影厅名称 座位（二维int数组）
	返回值：bool	
}
```

### 排片
```
{
	参数：排片编号 电影名称 所在影院 所在影厅 排片时间 开始时间 结束时间
	返回值：bool
	合约逻辑：根据座位生成对应电影票（初始化电影票参数：排片编号 座位排 座位号）
}
```

### 核对票根
```
{
	参数：票编号
	返回值：bool
}
```

### 锁定电影票
```
{
	参数：票编号 价格
	返回值：bool
}
```

### 查询电影票
```
{
	参数：票编号
	返回值：电影票模型（全部字段）
}
```

### 查看排片信息
```
{
	参数：无（即返回所有排片）
	返回：排片富模型（即包含其电影票集合）
}
```

### 分账
```
{
	参数：电影名称
	返回值：
		发行总数
		锁定总数
		核对总数
		总票房
		资金办分账（8.3% * 总票房）
		可分账金额（91.7% * 总票房）
		影院分账	（50% * 可分账金额）
		院线分账	（7% * 可分账金额）
		发行分账	（43% * 可分账金额）	
}
```


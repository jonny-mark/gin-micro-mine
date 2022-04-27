## swagger

#生成swagger文件
#yaml
swagger generate spec -o ./api/http/swagger/swagger.yaml

#docs go文件
swagger generate spec -o ./api/http/swagger/swagger.go

#json格式
swagger generate spec -o ./api/http/swagger/swagger.json

#启动server
swagger serve -F=swagger 文件路径

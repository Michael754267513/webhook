# webhook
kubernetes validating admission webhook and mutating admission webhook.

## 说明
* 编辑`deploy/generate-keys.sh` 生成证书相关
* 编辑 `deploy/webhook.yaml` 可以根据实际运行环境做修改 具体参考官方配置
* caBundle字段 为CA PEM base64 加密后的密文s
* demo证书已经生成，可以直接使用
## 配置说明
* 利用mutating admission做修改容器环境变量，添加TZ和LANG 变量
* 利用validating admission做验证对于环境变量不存在TZ的进行拒绝添加
## 启动参数
`--tls-cert-file=./deploy/certs/webhook-server-tls.crt  --tls-private-key-file=./deploy/certs/webhook-server-tls.key`

## 注意 
* 最好不要使用webhook去做数据做修改
* 在使用webhook的时候要注意条件匹配，进行筛选
* 在webhook运行在k8s内部的时候记得使用namespaceSelector 进行过滤
* https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/
    

## 参考地址
* https://github.com/kubernetes/kubernetes/blob/e1c617a88ec286f5f6cb2589d6ac562d095e1068/test/images/agnhost/webhook/main.go#L234
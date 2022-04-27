
ingress作为集群内部服务的入口，能够监听套接字并将流量转发到指定的服务，转发规则基于DNS名称（host）或URL路径

Service（type：ClusterIP）用于集群内部通信


参考资料：[k8s本地搭建]https://www.jishuchi.com/read/gin-practice/3836
[k8s系列]https://lailin.xyz/post/operator-04-kustomize-tutorial.html
[k8s应用更新策略]https://www.hebye.com/docs/k8s-roll
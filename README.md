> 创建image
docker build -t my-go-service .

> 推送image
# docker push my-go-service

> 部署到Kubernetes集群：
kubectl apply -f deployment.yaml

确保您已经正确配置和连接到Kubernetes集群。
运行以下命令，确保Deployment和Pod正确创建：

kubectl get deployments
kubectl get pods

> 验证部署：

确保三个Pod都处于运行状态：

kubectl get pods

您可以使用以下命令获取特定Pod的详细信息：
kubectl describe pod <pod名称>

> 创建Service：
kubectl apply -f service.yaml
> 获取服务的公开访问地址：
kubectl get services

现在，您应该可以通过访问服务的公开访问地址来访问部署的Go语言服务。请确保在 service.yaml 中设置了正确的端口映射和服务类型（如LoadBalancer或NodePort），以便根据您的需求进行访问配置。

这些步骤涵盖了从创建Docker镜像到部署到Kubernetes集群的详细流程。请根据实际情况修改配置文件中的名称、标签、端口等参数，并确保正确配置Kubernetes集群的访问权限。

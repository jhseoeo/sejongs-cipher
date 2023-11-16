AWS_REGION="ap-northeast-2"
ECS_CLUSTER="sejongs-cipher-cluster"
ECS_SERVICE="sejongs-cipher-service"
IMAGE_NAME="sejongs-cipher"
PROJECT_DIR="$(cd "$(dirname "$0")" && pwd)/.."

docker build -t sejongs-cipher $PROJECT_DIR
docker tag sejongs-cipher:latest 408047345469.dkr.ecr.ap-northeast-2.amazonaws.com/sejongs-cipher:latest
docker push 408047345469.dkr.ecr.ap-northeast-2.amazonaws.com/sejongs-cipher:latest
aws ecs update-service --cluster $ECS_CLUSTER --service $ECS_SERVICE --force-new-deployment
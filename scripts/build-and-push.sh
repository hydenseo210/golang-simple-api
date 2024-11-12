#!/bin/sh
set -euo pipefail


aws_region="ap-southeast-2"
ecr_registry="163145692369.dkr.ecr.ap-southeast-2.amazonaws.com/hyden-simple-api"
local_image="hyden-simple-api"
image_tag=$(git rev-parse --short HEAD)

# docker build -t hyden-simple-api .
docker build --build-arg LASTCOMMITSHA=${image_tag} -t "$local_image:$image_tag" .

# docker tag hyden-simple-api:latest 163145692369.dkr.ecr.ap-southeast-2.amazonaws.com/hyden-simple-api:latest
docker tag "$local_image:$image_tag" "$ecr_registry:$image_tag"

# docker push 163145692369.dkr.ecr.ap-southeast-2.amazonaws.com/hyden-simple-api:latest
docker push "$ecr_registry:$image_tag"
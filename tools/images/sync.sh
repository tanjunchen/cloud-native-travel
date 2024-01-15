#!/bin/bash
 
set -e
 
dst_user=istio
dst_repo=10.124.142.91

# image 示例
# "tanjunchen/wrk2:latest"
# "tanjunchen/goapp-ebpf:latest"
# "tanjunchen/coroot-node-agent:39492e3da86f"
images=(
    "haydenjeune/wrk2:latest"
)

pull_tag_push_image(){
    for image in ${images[*]}
    do
        if [ -z "${image}" ]
        then
        continue
        fi
        echo "docker pull --platform=linux/amd64 ${image}"
    
        docker pull ${image}
        echo "docker pull --platform=linux/amd64 ${image} success!!!"
        
        array=(`echo ${image} | tr ':' ' '` )
        src_image=${array[0]}
        src_version=${array[1]}
        if [ ! ${src_image} ]; then
            echo "src_image is null, stop tag and push"
            continue
        fi
        if [ ! ${src_version} ]; then
            echo "src_version is null, set default value latest"
            src_version=latest
        fi
        echo "docker src images info ${src_image} ${src_version}"
        
        image_array=(`echo ${src_image} | tr '/' ' '` )
        image_name=${image_array[-1]}
        if [ ! ${image_name} ]; then
            echo "image_name is null, stop tag and push"
            continue
        fi
        dst_image=${dst_repo}/${dst_user}/${image_name}:${src_version}
 
        echo "docker destination images info ${dst_image}"
        
        docker tag ${src_image}:${src_version} ${dst_image}
 
        docker push ${dst_image}
    done
}
 
pull_tag_push_image
 
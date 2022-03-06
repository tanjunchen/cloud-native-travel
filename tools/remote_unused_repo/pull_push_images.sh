#!/bin/bash

set -e

src_repo=xxx.com

dest_repo=xxx.com

source_user=xxx
dest_user=xxx

while read sc_image; do
    if [ -z "${sc_image}" ]
    then
      continue
    fi
    echo "pull ${sc_image}"

    docker pull ${src_repo}/${source_user}/${sc_image}
    docker tag ${src_repo}/${source_user}/${sc_image} ${dest_repo}/${dest_user}/${sc_image}
    docker push  ${dest_repo}/${dest_user}/${sc_image}

done < images

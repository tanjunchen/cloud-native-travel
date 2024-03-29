#!/bin/bash

set -e
set -o errexit
set -o nounset
set -o xtrace

src_repo=x.x.x.x:x

dest_repo=x.x.x.x

source_user=tanjunchen
dest_user=armpublic

while read sc_image; do
    if [ -z "${sc_image}" ]
    then
      continue
    fi
    echo "pull ${sc_image}"

    docker pull ${src_repo}/${source_user}/${sc_image}
    docker tag ${src_repo}/${source_user}/${sc_image} ${dest_repo}/${dest_user}/${sc_image}
    docker push  ${dest_repo}/${dest_user}/${sc_image}

done < bookinfo-image
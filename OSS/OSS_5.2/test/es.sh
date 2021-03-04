#!/bin/bash
curl -v 127.0.0.1:6060/OSS/objects/test.jpg\?version=1 -o res.jpg
curl -H "Content-Type: image/jpeg" -H "Digest: SHA-256=5oKACtnrsqO2QMJRYZHNMK5BPxAoQO8ilZwnRNL7ps0=" -v 127.0.0.1:6060/OSS/objects/test.jpg -XPUT -T ~/Pictrue/Wallpaper/webwxgetmsgimg1.jpg
 curl -H  "Content-Type: application/json" 127.0.0.1:9200/metadata -XPUT -d'{"mappings":{"properties":{"name":{"type":"keyword"},"version":{"type":"long"},"size":{"type":"long"},"hash":{"type":"text"}}}}'

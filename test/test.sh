#! /bin/bash

echo
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 1", "coords": "POINT(-25.31 22.253)"}'
echo
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 2", "coords": "POINT(15.174 47.237)"}'
echo
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 3", "coords": "POINT(165.355 52.243)"}'
echo
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 4", "coords": "POINT(130.345 22.933)"}'
echo
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 5", "coords": "POINT(37.141 72.243)"}'
echo
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 6", "coords": "POINT(1.111 2.222)"}'
echo
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 7", "coords": "POINT(135.345 36.113)"}'
echo
curl -XDELETE "http://127.0.0.1:8080/cities/4"
echo
curl -XPATCH "http://127.0.0.1:8080/cities/5" -d '{"coords": "POINT(115.311 12.293)"}'
echo
curl -XPATCH "http://127.0.0.1:8080/cities/6" -d '{"title": "Capitol"}'
echo
curl -XPUT "http://127.0.0.1:8080/cities/7" -d '{"title": "Village", "coords": "POINT(81.145 52.243)"}'
echo
curl -XGET "http://127.0.0.1:8080/cities/3"
echo
curl -XGET "http://127.0.0.1:8080/cities/5"
echo
echo "Nearest for (32.32 03.45)"
curl -XGET "http://127.0.0.1:8080/cities/find/32.31/03.45"
echo

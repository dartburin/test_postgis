# test_postgis
Test project with PostGis and REST HTTP.


Test method
-----------

Begin

(docker terminal) docker-compose up --build
(browser)  http://localhost:8080/cities
(test terminal) ./test/test.sh
(browser)  http://localhost:8080/cities


View console output

(log terminal) docker logs -f RESTserver


End

(docker terminal) Ctrl+z
(docker terminal) docker-compose down
(test terminal) sudo rm -R /tmp/data

(test terminal) docker images
(test terminal) docker rmi <new images ID>


Sample of testing orders
------------------------

(GET)
http://127.0.0.1:8080/cities
curl -XGET "http://127.0.0.1:8080/cities/3"

(POST)
curl -XPOST "http://127.0.0.1:8080/cities" -d '{"title": "City 1", "coords": "POINT(-25.31 22.253)"}'

(DELETE)
curl -XDELETE "http://127.0.0.1:8080/cities/4"

(PATCH) (partly update)
curl -XPATCH "http://127.0.0.1:8080/cities/5" -d '{"coords": "POINT(115.311 12.293)"}'

(PUT) (full update)
curl -XPUT "http://127.0.0.1:8080/cities/7" -d '{"title": "Village", "coords": "POINT(81.145 52.243)"}'

(FIND NEAREST)
http://127.0.0.1:8080/cities/find/longitude/latitude
curl -XGET "http://127.0.0.1:8080/cities/find/longitude/latitude"


Start
-------
docker-compose up --build



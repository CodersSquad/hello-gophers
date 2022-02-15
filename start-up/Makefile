
html :
	docker run --rm -v ${PWD}:/mnt -e MARP_USER="$(id -u):$(id -g)" \
	-e LANG=${LANG} -p 8080:8080 -p 37717:37717 marpteam/marp-cli -I /mnt --html

watch :
	docker run --rm --init -v ${PWD}:/home/marp/app/ -e LANG=${LANG} \
	-e MARP_USER="$(id -u):$(id -g)" -p 37717:37717 marpteam/marp-cli -w hello-gophers.md


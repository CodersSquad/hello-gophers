
SECTIONS = start-up


all: $(SECTIONS)
	cd $^ ; make
	mkdir -p public
	mv -f $^/public public/$^

clean:
	rm -rf public

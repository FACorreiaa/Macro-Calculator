FROM scratch
COPY calc /usr/bin/calc
ENTRYPOINT ["/usr/bin/calc"]

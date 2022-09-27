FROM ubuntu:jammy-20220130

LABEL Developer="Lakshya Nanjangud <lnanjangud@vmware.com>"

LABEL Maintainer="Solution Engineering Team - TKG(MAPBU)"

WORKDIR /usr/src/app

COPY time-app  ./

CMD "./time-app"
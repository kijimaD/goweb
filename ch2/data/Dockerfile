FROM postgres:latest

RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV POSTGRES_USER postgres
ENV POSTGRES_DB postgres

# 初期化用sqlファイル
COPY ./data/*.sql /docker-entrypoint-initdb.d/

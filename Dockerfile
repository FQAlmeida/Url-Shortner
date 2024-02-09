FROM node

RUN mkdir /app
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

RUN apt-get update && apt-get install netcat-traditional -y

ADD package.json yarn.lock /app/
RUN yarn install

ADD entrypoint.sh /app
ENTRYPOINT [ "sh", "/app/entrypoint.sh" ]

CMD [ "yarn", "start" ]

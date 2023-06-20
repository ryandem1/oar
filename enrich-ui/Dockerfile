FROM node:19-alpine AS builder

WORKDIR /app

COPY ./package*.json ./

RUN npm install

COPY . .

RUN npm run build:css
RUN npm run build


FROM node:19-alpine
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/package.json .
COPY --from=builder /app/build ./build
COPY --from=builder /app/output.css /app/build/client/output.css

RUN npm install --omit dev

# See if increase libuv thread pool size makes performance better
# The default value 4
# http://docs.libuv.org/en/v1.x/threadpool.html
ENV UV_THREADPOOL_SIZE=32

CMD ["node", "build"]

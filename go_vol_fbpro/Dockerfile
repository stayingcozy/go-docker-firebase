FROM golang

COPY go_vol_fbpro ./
RUN chmod +x ./go_vol_fbpro

COPY brightpaw.json .

ENV FB_ENV "production"

CMD ["./go_vol_fbpro"]


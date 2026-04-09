FROM python:3.11-alpine

ARG FILE=/example.html
ENV FILE="${FILE}"

WORKDIR /app

RUN pip install --no-cache-dir fastapi "uvicorn[standard]"
COPY --chmod=555 main.py /app/main.py
ADD https://example.com /example.html

EXPOSE 8888
LABEL version="0.1"

ENTRYPOINT ["uvicorn", "main:app"]
CMD ["--host=0.0.0.0", "--port=8888"]

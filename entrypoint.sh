#!/bin/sh

CONFIG=/app/config.yaml

INIT_CONFIG=/opt/config.yaml

# config Init
if [ -e ${INIT_CONFIG} ]; then
  cat ${INIT_CONFIG} >${CONFIG}
else
  cat <<EOF >${CONFIG}
resourceName: "${RESOURCENAME}"
apiVersion: "2023-05-15"
apiKey: "${APIKEY}"
mapper:
  "gpt-3.5-turbo": "${MAPPER_GPT35TUBER}"
EOF
fi

$(which chat_azure)

#!/bin/bash

FILENAME="insecure/certs.go"
echo > ${FILENAME}

echo -e "package insecure" >> ${FILENAME}
echo -e "const (" >> ${FILENAME}
echo -e "\tKey = \`$(cat server-key.pem)\`" >> ${FILENAME}
echo -e "\tCert = \`$(cat server.pem)\`" >> ${FILENAME}
echo -e ")" >> ${FILENAME}

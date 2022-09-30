#!/bin/sh

LOCAL_FILE=$1
S3_DIR=$2

psql -h $HOST_NAME -p $PORT_NUMBER -U $USER_NAME -d $DATABASE -f $SQL_FILE
RETURN_CODE=$?
if [ $RETURN_CODE -eq 0]; then
    exit 0
else
    exit 1
fi

aws s3 cp $LOCAL_FILE $S3_DIR
RETURN_CODE=$?
if [ $RETURN_CODE -eq 0]; then
    exit 0
else
    exit 1
fi



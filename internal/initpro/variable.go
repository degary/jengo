package initpro

import "strings"

var StartFile = strings.Join([]string{
	"#!/bin/sh",
	`APP_HOME="/data"`,
	"APP_PID=$APP_HOME/pid",
	"PROFILES=%s",
	"JVM=%d",
	"FILE_LIST=`ls  ${APP_HOME}`",
	`JAR_FILE=""`,
	`LOG_FILE=""`,

	"for FILE in $FILE_LIST",
	"do",
	`   if [ "${FILE##*.}" = "jar" ]; then`,
	`       JAR_FILE="${APP_HOME}/${FILE}"`,
	`       LOG_FILE="${FILE%.*}.log"`,
	`   fi`,
	`done`,

	"if [ -z $JAR_FILE ]; then",
	`   echo "Error : no jar file in $APP_HOME"`,
	"   exit",
	"fi",

	`if [ -f "$APP_PID" ]; then`,
	"   PID=`cat $APP_PID`",
	"   rm -rf $APP_PID",
	"   kill -9 $PID",
	"fi",

	"java -Dfile.encoding=UTF8 -Duser.timezone=GMT+08  -Djava.security.egd=file:/dev/./urandom $DEBUG -jar -Xms${JVM}m -Xmx${JVM}m -Xmn256m -XX:MetaspaceSize=256m -XX:MaxMetaspaceSize=256m -Xss256k -XX:+PrintGCDateStamps -XX:+PrintGCTimeStamps -XX:+PrintGCDetails -Xloggc:./gc.log -XX:+UseConcMarkSweepGC -XX:+UseParNewGC -XX:ParallelGCThreads=4  -Dspring.profiles.active=${PROFILES}  ${JAR_FILE}",
	"",
}, "\n")

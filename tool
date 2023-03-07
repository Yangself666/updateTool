#!/bin/sh
script_dir=$(cd $(dirname $0);pwd)
service=$script_dir"/updateTool"
pid_file=$script_dir"/service.pid"


function start() {
    if [ -f ${pid_file} ]; then
        echo 'Service is Running'
    else
        ${service} > /dev/null 2>&1  &
        if [[ $? -eq 0 ]]; then
            echo $! > ${pid_file}
        else exit 1
        fi
    fi
}


function stop() {
    if [ -f ${pid_file} ]; then
        kill -9 $(cat ${pid_file})
        if [[ $? -eq 0 ]]; then
            rm -f ${pid_file}
        else exit 1
        fi
    else
        echo 'Service is Not Running Or Pid File Is Not Exist'
    fi
}


function call() {
    case $1 in
        'start')
            start
            ;;
        'stop')
            stop
            ;;
        *)
            echo 'Get invalid option, please input(as to $1):'
            echo -e '\t"start" -> start service'
            echo -e '\t"stop"  -> stop service'
                        exit 1
    esac
}


call "$1"
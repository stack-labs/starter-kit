#!/bin/bash

# Just a script to run the demo
cmd=$1
dir=$2
kube=kubectl
list=( micro )

start() {
	if [ -z $dir ]; then
		for dir in ${list[@]}; do
			find $dir -name "*.yaml" | while read file; do
				$kube apply -f $file
			done
		done
		return
	fi

	find $dir -name "*.yaml" | while read file; do
		$kube apply -f $file
	done

}

stop() {
	if [ -z $dir ]; then
		for dir in ${list[@]}; do
			find $dir -name "*.yaml" | while read file; do
				$kube delete -f $file
			done
		done
		return
	fi

	find $dir -name "*.yaml" | while read file; do
		$kube delete -f $file
	done
}

case $cmd in
	start)
	start
	;;
	stop)
	stop
	;;
	restart)
	stop
	start
	;;
	*)
	echo "$0 <start|stop|restart> [dir]"
	exit
	;;
esac

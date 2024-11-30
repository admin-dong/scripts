#!/bin/bash
function cpu() {
    printf "%-13s %-10s\n" "使用率" "等待磁盘IO的时间百分比"
    printf "%-10s %-10s\n" "------" "----------------------"
    util=`vmstat |awk '{if(NR==3)print 100-$15}'`
    #user=`vmstat |awk '{if(NR==3)print $13"%"}'`
    #sys=`vmstat |awk '{if(NR==3)print $14"%"}'`
    iowait=`vmstat |awk '{if(NR==3)print $16}'`
    printf "%-10s %-10s\n" "${util}%" "${iowait}%"
}
function memory() {
    printf "%-13s %-12s %-10s\n" "总大小" "使用" "剩余"
    printf "%-10s %-10s %-10s\n" "------" "------" "------"
    total=`free -m |awk '{if(NR==2)printf "%.1f",$2/1024}'`
    used=`free -m |awk '{if(NR==2) printf "%.1f",($2-$NF)/1024}'`
    available=`free -m |awk '{if(NR==2) printf "%.1f",$NF/1024}'`
    printf "%-10s %-10s %-10s\n" "${total}G" "${used}G" "${available}G"
}
function disk() {
    printf "%-13s %-13s %-12s %-10s\n" "挂载点" "总大小" "使用" "使用率"
    printf "%-10s %-10s %-10s %-10s\n" "------" "------" "------" "------"
    fs=`df -h |awk '/^\/dev/{print $1}'`
    for p in $fs; do
        mounted=`df -h |awk '$1=="'$p'"{print $NF}'`
        size=`df -h |awk '$1=="'$p'"{print $2}'`
        used=`df -h |awk '$1=="'$p'"{print $3}'`
        used_percent=`df -h |awk '$1=="'$p'"{print $5}'`
        printf "%-10s %-10s %-10s %-10s\n" "$mounted" "$size" "$used" "$used_percent"
    done
}
function process_cpu_top() {
    printf "%-10s %-10s %-12s %-10s\n" "PID" "CPU" "内存" "命令"
    printf "%-10s %-10s %-10s %-10s\n" "------" "------" "------" "------"
    tmp_file=/tmp/cputop
    ps -eo pid,pcpu,pmem,command --sort=-pcpu |awk 'NR>1&&NR<=4' > $tmp_file
    while read line; do
        pid=`echo $line |awk '{print $1}'`
        cpu=`echo $line |awk '{print $2}'`
        mem=`echo $line |awk '{print $3}'`
        com=`echo $line |awk '{print $4}'`
        printf "%-10s %-10s %-10s %-10s\n" "$pid" "${cpu}%" "${mem}%" "$com"
    done < $tmp_file
}
function process_memory_top() {
    printf "%-10s %-12s %-10s %-10s\n" "PID" "内存" "CPU" "命令"
    printf "%-10s %-10s %-10s %-10s\n" "------" "------" "------" "------"
    tmp_file=/tmp/cputop
    ps -eo pid,pcpu,pmem,command --sort=-pmem |awk 'NR>1&&NR<=4' > $tmp_file
    while read line; do
        pid=`echo $line |awk '{print $1}'`
        cpu=`echo $line |awk '{print $2}'`
        mem=`echo $line |awk '{print $3}'`
        com=`echo $line |awk '{print $4}'`
        printf "%-10s %-10s %-10s %-10s\n" "$pid" "${mem}%" "${cpu}%" "$com"
    done < $tmp_file
}

echo "################# CPU #################"
cpu
echo "################# 内存 #################"
memory
echo "################# 硬盘 #################"
disk
echo "################# 占用CPU最高的前三个进程 #################"
process_cpu_top
echo "################# 占用内存最高的前三个进程 #################"
process_memory_top

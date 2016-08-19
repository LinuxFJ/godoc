me: 123.sh
# Author: LookBack
# Email: admin#dwhd.org
# Version:
# Created Time: Tue 13 Oct 2015 12:40:48 AM CST
#########################################################################

PrintA() {
            a=$1
                    b=`printf "%-${a}s\n" "*" | sed 's/\s/*/g'`
                            c=`echo "(39-$a)/2"|bc`
                                    [ "$a" != "39" ] && d=`printf "%-${c}s\n"` || d=""
                                            echo "${d}${b}"
                                        }

                                        PrintB() {
                                                    e=$1
                                                            b=`printf "%-5s\n" "*" | sed 's/\s/*/g'`
                                                                    c=`echo "($e-5)/2"|bc`
                                                                            d=`printf "%-${c}s\n" " "`
                                                                                    echo -e "${d}${b}\n${d}${b}\n${d}${b}\n${d}${b}\n${d}${b}\n${d}${b}"
                                                                                }

                                                                                for i in `seq 1 2 39`; do
                                                                                            [ "$i" = "39" ] && PrintB $i || PrintA $i
                                                                                        done
                                                                                        ======改进版 能根据终端大小自动调整圣诞树的大小====
#!/bin/bash
#########################################################################
# File Name: christmas.sh
# Author: LookBack
# Email: admin#dwhd.org
# Version:
# Created Time: Tue 13 Oct 2015 10:17:55 AM CST
#########################################################################

NUM1=$(awk '{a=($1-2)*0.8}{printf("%d\n",int(a)==a?a:(int(a)+1))}' <<< $(tput lines))
NUM2=`awk '{a=-1;for(i=1;i<=$1;i++){a+=2}}END{print a}' <<< $NUM1`

PrintA() {
            a=$1 && b=$2
                    c=`printf "%-${a}s\n" "*" | sed 's/\s/*/g'`
                            d=`echo "($b-$a)/2"|bc`
                                    [ "$a" != "$b" ] && e=`printf "%-${d}s\n" " "` || e=""
                                            echo "`printf "%-10s"`${e}${c}"
                                                    if [ "$a" = "$b" ]; then
                                                                        c=`printf "%-5s\n" "*" | sed 's/\s/*/g'`
                                                                                        d=`echo "($a-5)/2+10"|bc`
                                                                                                        e=`printf "%-${d}s\n"`
                                                                                                                        echo -e "${e}${c}\n${e}${c}\n${e}${c}\n${e}${c}\n${e}${c}\n${e}${c}"
                                                                                                                                fi
                                                                                                                            }

                                                                                                                            for i in `seq 1 2 $NUM2`; do
                                                                                                                                        PrintA $i $NUM2
                                                                                                                                    done

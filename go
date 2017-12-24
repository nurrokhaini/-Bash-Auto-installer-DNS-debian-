#!/bin/bash
dns=$(whiptail --inputbox "Masukkan dns lengkap" 8 78 --title "contoh : nur.net" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then
host=$(whiptail --inputbox "Masukkan host ip server" 8 78 --title "contoh : 192.168.10.2 ditulis 2 aja" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then
ip=$(whiptail --inputbox "Masukkan reverse ip" 8 78 --title "ip nya dibalik 3 oktat saja. contoh: 10.168.192" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then
piasli=$(whiptail --inputbox "Masukkan Ip Lengkap" 8 78 --title "lenkap ya :D" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then

#MAi command
apt update
apt install bind9
cp named.conf.local /etc/bind
cp db.192 /etc/bind
cp db.dns /etc/bind
cp resolv.conf /etc
cd /etc/bind


##editing main local
perl -pi -e "s/smk.net/$dns/g" named.conf.local
perl -pi -e "s/banana/$ip/g" named.conf.local

## editing db reverse 192
perl -pi -e "s/dns/$dns/g" db.192
perl -pi -e "s/2/$host/g" db.192

## editing db resolve db dns
perl -pi -e "s/dns/$dns/g" db.dns
perl -pi -e "s/2/$host/g" db.dns
perl -pi -e "s/ip/$piasli/g" db.dns
cd /etc/

## input resolvf
ip1=$(whiptail --inputbox "Input your gateway mas" 8 78 --title "gateway mas" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then

## editing resolv.conf :D
perl -pi -e "s/ip1/$piasli/g" resolv.conf
perl -pi -e "s/ip2/$ip1/g" resolv.conf
perl -pi -e "s/dns1/$dns/g" resolv.conf

## service restart bind9
/etc/init.d/bind9 restart
/etc/init.d/networking restart

##result
echo "THANK for using "
echo "NUR ROKHAINI"

else
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi

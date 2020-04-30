//How to install godoc ?

[root@GyanHouse nabodip]# dnf -y install  golang-godoc.i686
Repository google-chrome is listed more than once in the configuration
Last metadata expiration check: 0:02:24 ago on Fri May  1 04:27:31 2020.
Dependencies resolved.
==============================================================================================================================
 Package                       Arch                  Version                                   Repository                Size
==============================================================================================================================
Installing:
 golang-godoc                  i686                  1:0-15.1.git9deed8c.fc24                  updates                  3.2 M

Transaction Summary
==============================================================================================================================
Install  1 Package

Total download size: 3.2 M
Installed size: 11 M
Downloading Packages:
golang-godoc-0-15.1.git9deed8c.fc24.i686.rpm                                                  1.1 MB/s | 3.2 MB     00:02    
------------------------------------------------------------------------------------------------------------------------------
Total                                                                                         789 kB/s | 3.2 MB     00:04     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Installing  : golang-godoc-1:0-15.1.git9deed8c.fc24.i686                                                                1/1 
  Verifying   : golang-godoc-1:0-15.1.git9deed8c.fc24.i686                                                                1/1 

Installed:
  golang-godoc.i686 1:0-15.1.git9deed8c.fc24                                                                                  

Complete!


// Once installation is complete, launch the godoc service :

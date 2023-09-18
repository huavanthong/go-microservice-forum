import os, sys

import colorama as col

col.init(autoreset=True)

COLBR = col.Style.BRIGHT + col.Fore.RED
COLBY = col.Style.BRIGHT + col.Fore.YELLOW
COLBG = col.Style.BRIGHT + col.Fore.GREEN

SUCCESS = 0
ERROR   = 1


# Chuyển hướng đầu ra vào một tệp tin
# sys.stdout = open('output.txt', 'w')
# sys.stderr = open('error.txt', 'w')


def printerror(sMsg):
   sys.stderr.write(COLBR + f"{sMsg}!\n")

try:
   result = 10/0
except Exception as ex:
   print()
   printerror(str(ex))
   print("HERE222")
#    with open('error.txt', 'w') as error_file:
#         error_file.write(str(e))


print("error")
sys.exit(1)
   
# Đóng tệp tin
# sys.stdout.close()
# sys.stderr.close()